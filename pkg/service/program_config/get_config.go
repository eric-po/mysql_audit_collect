package program_config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

func (c *KafkaInfo) getConf(confFile string) *KafkaInfo {
	yamlFile, err := ioutil.ReadFile(confFile)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}

type KafkaInfo struct {
	InstanceList           string `yaml:"InstanceList"`
	Topic                  string `yaml:"Topic"`
	GroupId                string `yaml:"GroupId"`
	ClickHouseInstanceHost string `yaml:"ClickHouseInstanceHost"`
	ClickHouseInstancePort int    `yaml:"ClickHouseInstancePort"`
	ClickHouseDatabase     string `yaml:"ClickHouseDatabase"`
}

func GetKafkaConfig(confFile string) KafkaInfo {
	//kafkaConfig := KafkaInfo{
	//	instanceList: "10.1.1.232:9095",
	//	topic:        "mysql_audit_park",
	//}
	var k KafkaInfo
	k.getConf(confFile)
	return k
}
