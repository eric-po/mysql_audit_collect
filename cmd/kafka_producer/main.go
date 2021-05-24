package main

import (
	"fmt"
	kafka_utils "github.com/ssesse/mysql_audit_collect/pkg/service/kafka_utils"
)

func main() {
	fmt.Println("Here test producer")
	kafka_utils.kafkaproducer.kafkaProducer()

}
