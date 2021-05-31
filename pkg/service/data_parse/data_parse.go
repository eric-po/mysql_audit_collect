package data_parse

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
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

type MessageInfo struct {
	dbHost         string
	dbPort         int
	QueryTimestamp string
	Serverhost     string
	Username       string
	Host           string
	Connectionid   int
	Queryid        int
	Operation      string
	Database       string
	Object         string
	Retcode        int
}

func MessageParse(message string) *MessageInfo {
	reg1 := regexp.MustCompile(`(\d{8} \d{2}:\d{2}:\d{2}),(.*),(.*),(.*),(\d*),(\d*),(.*),(.*),['](.*)['],(\d*$)`)
	//reg1 := regexp.MustCompile(`(\d{8} \d{2}:\d{2}:\d{2}),(.*),(.*),(.*),(\d*),(\d*),(CONNECT|QUERY|READ|WRITE|CREATE|ALTER|RENAME|DROP),(.*),['](.*)['],(\d*$)`)

	result0 := reg1.FindAllStringSubmatch(message, -1)
	//fmt.Println("haha : ", result0)
	if len(result0) == 0 {
		reg1 = regexp.MustCompile(`(\d{8} \d{2}:\d{2}:\d{2}),(.*),(.*),(.*),(\d*),(\d*),(.*),(.*),(.*),(\d*$)`)
		//reg1 = regexp.MustCompile(`(\d{8} \d{2}:\d{2}:\d{2}),(.*),(.*),(.*),(\d*),(\d*),(CONNECT|DISCONNECT|QUERY|READ|WRITE|CREATE|ALTER|RENAME|DROP),(.*),(),(\d*$)`)
		result0 = reg1.FindAllStringSubmatch(message, -1)
		//fmt.Println(result0)
	}
	//
	//for _, param := range result0[0] {
	//	fmt.Printf("diwn : %s\n", param)
	//}
	var msgResult MessageInfo
	msgResult.QueryTimestamp = result0[0][1]
	msgResult.Serverhost = result0[0][2]
	msgResult.Username = result0[0][3]
	msgResult.Host = result0[0][4]
	msgResult.Connectionid, _ = strconv.Atoi(result0[0][5])
	msgResult.Queryid, _ = strconv.Atoi(result0[0][6])
	msgResult.Operation = result0[0][7]
	msgResult.Database = result0[0][8]
	msgResult.Object = result0[0][9]
	msgResult.Retcode, _ = strconv.Atoi(result0[0][10])
	//return &result0[0]
	return &msgResult
}
