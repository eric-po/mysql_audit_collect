package main

import (
	"fmt"
	"regexp"
	"strconv"
)

type MessageInfo struct {
	dbHost         string
	dbPort         int
	queryTimestamp string
	serverhost     string
	username       string
	host           string
	connectionid   int
	queryid        int
	operation      string
	database       string
	object         string
	retcode        int
}

func MessageParse(message string) *MessageInfo {
	reg1 := regexp.MustCompile(`(\d{8} \d{2}:\d{2}:\d{2}),(.*),(.*),(.*),(\d*),(\d*),(.*),(.*),['](.*)['],(\d*$)`)
	//reg1 := regexp.MustCompile(`(\d{8} \d{2}:\d{2}:\d{2}),(.*),(.*),(.*),(\d*),(\d*),(CONNECT|QUERY|READ|WRITE|CREATE|ALTER|RENAME|DROP),(.*),['](.*)['],(\d*$)`)

	result0 := reg1.FindAllStringSubmatch(message, -1)
	fmt.Println("haha : ", result0)
	if len(result0) == 0 {
		reg1 = regexp.MustCompile(`(\d{8} \d{2}:\d{2}:\d{2}),(.*),(.*),(.*),(\d*),(\d*),(.*),(.*),(.*),(\d*$)`)
		//reg1 = regexp.MustCompile(`(\d{8} \d{2}:\d{2}:\d{2}),(.*),(.*),(.*),(\d*),(\d*),(CONNECT|DISCONNECT|QUERY|READ|WRITE|CREATE|ALTER|RENAME|DROP),(.*),(),(\d*$)`)
		result0 = reg1.FindAllStringSubmatch(message, -1)
		fmt.Println(result0)
	}
	for _, param := range result0[0] {
		fmt.Printf("diwn : %s\n", param)
	}
	var msgResult MessageInfo
	msgResult.queryTimestamp = result0[0][1]
	msgResult.serverhost = result0[0][2]
	msgResult.username = result0[0][3]
	msgResult.host = result0[0][4]
	msgResult.connectionid, _ = strconv.Atoi(result0[0][5])
	msgResult.queryid, _ = strconv.Atoi(result0[0][6])
	msgResult.operation = result0[0][7]
	msgResult.database = result0[0][8]
	msgResult.object = result0[0][9]
	msgResult.retcode, _ = strconv.Atoi(result0[0][10])
	//return &result0[0]
	return &msgResult
}

func main() {
	//s := "20210528 11:12:51,zhc-dba-mysql-dev-01,a_hehe_rw,localhost,906,0,DISCONNECT,,,0"
	//s := "20210527 11:04:43,zhc-dba-mysql-dev-01,dba_grb,localhost,653,1252216,QUERY,,'select \\'a\\' , b from test.hehe ;',0"
	//s := "20210531 12:58:16,dc-dba-mysql-pro-10,a_uis_rw,192.168.2.219,6044145,4958880683,QUERY,,'select userdailyb0_.id as id1_4_, userdailyb0_.app_id as app_id2_4_, userdailyb0_.create_time as create_t3_4_, userdailyb0_.end_gow as end_gow4_4_, userdailyb0_.end_gow_minute as end_gow_5_4_, userdailyb0_.end_gtw as end_gtw6_4_, userdailyb0_.end_gtw_minute as end_gtw_7_4_, userdailyb0_.start_gow as start_go8_4_, userdailyb0_.start_gow_minute as start_go9_4_, userdailyb0_.start_gtw as start_g10_4_, userdailyb0_.start_gtw_minute as start_g11_4_, userdailyb0_.update_time as update_12_4_, userdailyb0_.user_id as user_id13_4_, userdailyb0_.working_day as working14_4_, userdailyb0_.working_day_type as working15_4_ from t_user_daily_behavior userdailyb0_ where userdailyb0_.user_id=897873 and userdailyb0_.app_id=2',0"
	s := "20210531 14:25:44,dc-dba-mysql-pro-156,a_gulf_rw,192.168.11.148,1566279,17436347284,QUERY,gulf,'select     t2.*     from t_terminal_vehicle t1 left join t_vehicle t2 on t1.vehicle_id = t2.vehicle_id     where t1.terminal_id = \\'862932043955344\\'',0"
	//s := "20210531 14:30:28,zhc-dba-mysql-dev-01,a_hehe_rw,localhost,2710,0,CONNECT,,,0"
	fmt.Println("#########################")
	msg := MessageParse(s)
	fmt.Println(msg)
}
