package kafka_utils

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/ssesse/mysql_audit_collect/pkg/service/clickhouse_utils"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"
)

// Sarama configuration options

type KafkaConfig struct {
	Oldest   bool
	Assignor string
	Verbose  bool
}

func KafkaConsumeNew(kafkaInstanceList string, consumeTopic string, GroupId string, kConfig *KafkaConfig) {
	log.Println("Starting a new Sarama consumer")

	if kConfig.Verbose {
		sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)
	}
	//version, err := sarama.ParseKafkaVersion(version)
	//if err != nil {
	//	log.Panicf("Error parsing Kafka version: %v", err)
	//}

	/**
	 * Construct a new Sarama configuration.
	 * The Kafka cluster version has to be defined before the consumer/producer is initialized.
	 */
	config := sarama.NewConfig()
	switch kConfig.Assignor {
	case "sticky":
		config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategySticky
	case "roundrobin":
		config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	case "range":
		config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
	default:
		log.Panicf("Unrecognized consumer group partition assignor: %s", kConfig.Assignor)
	}

	if kConfig.Oldest {
		config.Consumer.Offsets.Initial = sarama.OffsetOldest
	} else {
		config.Consumer.Offsets.Initial = sarama.OffsetNewest
	}

	/**
	 * Setup a new Sarama consumer group
	 */
	consumer := Consumer{
		ready: make(chan bool),
	}

	ctx, cancel := context.WithCancel(context.Background())
	client, err := sarama.NewConsumerGroup(strings.Split(kafkaInstanceList, ","), GroupId, config)
	if err != nil {
		//log.Panicf("Error creating consumer group client: %v", err)
		log.Printf("Kafka consumer error , brokers(%s) is  not reachable. \n \n", kafkaInstanceList)
		cancel()
		return
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			// `Consume` should be called inside an infinite loop, when a
			// server-side rebalance happens, the consumer session will need to be
			// recreated to get the new claims
			if err := client.Consume(ctx, strings.Split(consumeTopic, ","), &consumer); err != nil {
				log.Panicf("Error from consumer: %v", err)
			}
			// check if context was cancelled, signaling that the consumer should stop
			if ctx.Err() != nil {
				return
			}
			consumer.ready = make(chan bool)
		}
	}()

	<-consumer.ready // Await till the consumer has been set up
	log.Println("Sarama consumer up and running!...")

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-ctx.Done():
		log.Println("terminating: context cancelled")
	case <-sigterm:
		log.Println("terminating: via signal")
	}
	cancel()
	wg.Wait()
	if err = client.Close(); err != nil {
		log.Panicf("Error closing client: %v", err)
	}
}

// Consumer represents a Sarama consumer group consumer
type Consumer struct {
	ready chan bool
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	close(consumer.ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// NOTE:
	// Do not move the code below to a goroutine.
	// The `ConsumeClaim` itself is called within a goroutine, see:
	// https://github.com/Shopify/sarama/blob/master/consumer_group.go#L27-L29
	//listOfNumberMessages := []*sarama.ConsumerMessage{}
	var consumerMessageSlice []*sarama.ConsumerMessage
	var sum int
	sum = 0

	insertBatchSizeStr, ex := os.LookupEnv("CH_BATCH_SIZE")
	var insertBatchSize int
	var portParseError error
	if !ex {
		insertBatchSize = 1000
		log.Printf("The env variable %s is not set.\n", "CH_BATCH_SIZE")
	} else {
		insertBatchSize, portParseError = strconv.Atoi(insertBatchSizeStr)
		if portParseError != nil {
			log.Panic("clickhouse port define error !")
		}
	}

	for message := range claim.Messages() {
		//log.Println(string(message.Value))
		//listOfNumberMessages = append(listOfNumberMessages, message)
		//clickhouse_utils.QueryRecordHandle(message)
		if sum < insertBatchSize {
			consumerMessageSlice = append(consumerMessageSlice, message)
			sum += 1
		} else {
			clickhouse_utils.QueryRecordHandleBatch(consumerMessageSlice)
			sum = 0
			consumerMessageSlice = nil
		}
		//log.Printf("Message claimed: value = %s, timestamp = %v, topic = %s", string(message.Value), message.Timestamp, message.Topic)
		session.MarkMessage(message, "")
	}
	return nil
}
