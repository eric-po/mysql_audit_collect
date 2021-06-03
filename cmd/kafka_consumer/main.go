package main

import (
	"flag"
	"fmt"
	"github.com/ssesse/mysql_audit_collect/pkg/service/kafka_utils"
	"github.com/ssesse/mysql_audit_collect/pkg/service/program_config"
	"log"
)

func main() {
	var configFileName string
	flag.StringVar(&configFileName, "f", "db_source.yaml", "yaml config file. ")
	flag.Parse()
	a := program_config.GetKafkaConfig(configFileName)
	fmt.Println(a)
	//kafka_utils.KafkaConsume(a.InstanceList, a.Topic, a.GroupId)
	//sqlFormat := "insert into sql_record(db_host  ,db_port  ,query_timestamp  ,serverhost  ,username  ,host  ,connectionid  ,queryid  ,operation  ,database  ,object  ,retcode) values(?,?,?,?,?,?,?,?,?,?,?,?) ;"
	//kConfig.
	var kConfig = new(kafka_utils.KafkaConfig)
	kConfig.Oldest = false
	kConfig.Assignor = "roundrobin"
	kConfig.Verbose = false
	//ckConnect := clickhouse_utils.ChConnect(a.ClickHouseInstanceHost, a.ClickHouseInstancePort, a.ClickHouseDatabase, sqlFormat)
	for {
		kafka_utils.KafkaConsumeNew(a.InstanceList, a.Topic, a.GroupId, kConfig)
		log.Println("Now we start consuming again. ")
	}
}
