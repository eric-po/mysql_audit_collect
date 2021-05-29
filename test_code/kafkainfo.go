package main

import "fmt"

type kafkaInfo struct {
	instanceList string
	topic        string
}

type KafkaConfig interface {
	getKafkaConfig() kafkaInfo
}

func getKafkaConfig() kafkaInfo {
	kafkaConfig := kafkaInfo{
		instanceList: "10.1.1.232:9095",
		topic:        "mysql_audit_park",
	}
	return kafkaConfig
}

func main() {
	var k KafkaConfig
	a := k.getKafkaConfig()
	//a := k.getKafkaConfig()
	fmt.Println(a)
}
