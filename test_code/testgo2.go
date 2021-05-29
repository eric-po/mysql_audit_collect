package main

import (
	"encoding/json"
	"fmt"
)

//jsonStr := '{"@timestamp":"2021-05-26T09:34:46.768Z","@metadata":{"beat":"filebeat","type":"_doc","version":"7.3.1","topic":"mysql_audit1"},"message":"20210526 17:34:44,zhc-dba-mysql-dev-01,dba_grb,localhost,652,1251985,QUERY,hehe,'select * from stu',0","db_host":"10.200.11.25","db_port":3301,"host":{"ip":["10.200.11.25","fe80::bc70:13ff:feb9:7067","172.17.0.1","10.200.0.0","fe80::6859:f0ff:fe6e:7037","10.200.0.1","fe80::84ef:b1ff:fee3:4a1e"]}}'
//var jsonStr = `{"message":"20210526 17:34:44,zhc-dba-mysql-dev-01,dba_grb,localhost,652,1251985,QUERY,hehe,select * from stu',0","db_host":"10.200.11.25","db_port":3301,"host":{"ip":["10.200.11.25","fe80::bc70:13ff:feb9:7067","172.17.0.1","10.200.0.0","fe80::6859:f0ff:fe6e:7037","10.200.0.1","fe80::84ef:b1ff:fee3:4a1e"]}}`
var jsonStr = `{"@timestamp":"2021-05-26T09:34:49.769Z","@metadata":{"beat":"filebeat","type":"_doc","version":"7.3.1","topic":"mysql_audit1"},"message":"20210526 17:34:49,zhc-dba-mysql-dev-01,dba_grb,localhost,652,1251991,QUERY,hehe,'select * from stu',0","db_host":"10.200.11.25","db_port":3301,"host":{"ip":["10.200.11.25","fe80::bc70:13ff:feb9:7067","172.17.0.1","10.200.0.0","fe80::6859:f0ff:fe6e:7037","10.200.0.1","fe80::84ef:b1ff:fee3:4a1e"]}}`

type sqlStruct struct {
	Message string `message`
	dbHost  string `db_host`
	dbPort  int    `db_port`
}

func main() {
	var config sqlStruct
	if err := json.Unmarshal([]byte(jsonStr), &config); err == nil {
		fmt.Println("================json str 转struct==")
		fmt.Println(config)
		fmt.Printf("haha %s\n", config.dbHost)
		fmt.Printf("hehe %s\n", config.Message)
	} else {
		fmt.Println(err)
	}

	//var err = json.Unmarshal(sqlStruct, &basket)
	//if err != nil {
	//	log.Println(err)
	//}
	//json.Unmarshal(bs, &t)

}
