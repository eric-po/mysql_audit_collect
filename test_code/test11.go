package main

import (
	"encoding/json"
	"fmt"
)

type QueryRecord struct {
	Message string `json:"message"`
	DbHost  string `json:"db_host"`
	DbPort  int    `json:"db_port"`
}

func RecordParse(jsonStr string) *QueryRecord {
	var qr QueryRecord
	if err := json.Unmarshal([]byte(jsonStr), &qr); err == nil {
		return &qr
	} else {
		fmt.Println(err)
		return nil
	}
}

func main() {
	//var ss string
	//s = '{"@timestamp":"2021-06-03T03:42:14.981Z","@metadata":{"beat":"filebeat","type":"_doc","version":"7.3.1","topic":"audit_log_for_mysql1"},"db_host":"192.168.11.27","db_port":3301,"message":"aa"}'
	//ss = 's'
	//ss := "20210531 14:25:44,dc-dba-mysql-pro-156,a_gulf_rw,192.168.11.148,1566279,17436347284,QUERY,gulf,'select     t2.*     from t_terminal_vehicle t1 left join t_vehicle t2 on t1.vehicle_id = t2.vehicle_id     where t1.terminal_id = \\'862932043955344\\'',0"
	ss := `{"@timestamp":"2021-06-03T05:48:17.260Z","@metadata":{"beat":"filebeat","type":"_doc","version":"7.3.1","topic":"mysql_audit1"},"message":"20210603 13:48:16,zhc-dba-mysql-dev-01,a_hehe_rw,localhost,225932,0,DISCONNECT,,,0","db_host":"10.200.11.25","db_port":3301,"host":{"ip":["10.200.11.25","fe80::bc70:13ff:feb9:7067","172.17.0.1","10.200.0.0","fe80::6859:f0ff:fe6e:7037","10.200.0.1","fe80::84ef:b1ff:fee3:4a1e"]}}`
	a := RecordParse(ss)
	fmt.Println(a)
	//fmt.Println(ss)
}
