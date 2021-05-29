package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func AuditContentParse() *[]string {
	s := "20210527 11:04:43,zhc-dba-mysql-dev-01,dba_grb,localhost,653,1252216,QUERY,,'select \\'a\\' , b from test.hehe ;',0"
	reg1 := regexp.MustCompile(`(.*),(.*),(.*),(.*),(.*),(.*),(.*),(.*),'(.*)',(.*)`)
	result0 := reg1.FindAllStringSubmatch(s, -1)
	fmt.Println("haha : ", result0)

	for _, param := range result0[0] {
		fmt.Printf("diwn : %s\n", param)
	}
	return &result0[0]
}

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
	reg1 := regexp.MustCompile(`(.*),(.*),(.*),(.*),(.*),(.*),(.*),(.*),[']?(.*)?,(.*)`)
	result0 := reg1.FindAllStringSubmatch(message, -1)
	fmt.Println("haha : ", result0)

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
	s := "20210527 11:04:43,zhc-dba-mysql-dev-01,dba_grb,localhost,653,1252216,QUERY,,'select \\'a\\' , b from test.hehe ;',0"

	msg := MessageParse(s)
	fmt.Println(msg)
}
