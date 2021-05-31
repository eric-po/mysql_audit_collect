package clickhouse_utils

import (
	"database/sql"
	"fmt"
	_ "github.com/ClickHouse/clickhouse-go"
	"github.com/Shopify/sarama"
	"github.com/ssesse/mysql_audit_collect/pkg/service/data_parse"
	"log"
	"os"
	"strconv"
)

var sqlFormat = "insert into sql_record(db_host  ,db_port  ,query_timestamp  ,serverhost  ,username  ,host  ,connectionid  ,queryid  ,operation  ,database  ,object  ,retcode) values(?,?,?,?,?,?,?,?,?,?,?,?) ;"

func QueryRecordHandle(message *sarama.ConsumerMessage) {
	var qr *data_parse.QueryRecord
	qr = data_parse.RecordParse(string(message.Value))
	msg := data_parse.MessageParse(qr.Message)
	fmt.Println(*msg)
	chCon := GetChCon()
	var (
		tx, _   = chCon.Begin()
		stmt, _ = tx.Prepare(sqlFormat)
	)
	//stmt.Exec(qr.DbHost,qr.DbPort,*msg.queryTimestamp,*msg.serverhost)
	if _, err := stmt.Exec(
		qr.DbHost,
		qr.DbPort,
		msg.QueryTimestamp,
		msg.Serverhost,
		msg.Username,
		msg.Host,
		msg.Connectionid,
		msg.Queryid,
		msg.Operation,
		msg.Database,
		msg.Object,
		msg.Retcode,
	); err != nil {
		log.Fatal(err)
	}
	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
}

var ChCon *sql.DB

func ChConnect(instanceHost string, instancePort int, database string) *sql.DB {
	DSNConnect := fmt.Sprintf("tcp://%s:%d?database=%s", instanceHost, instancePort, database)
	connect, err := sql.Open("clickhouse", DSNConnect)
	fmt.Println(DSNConnect)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(sqlFormat)
	return connect
}

func GetChCon() *sql.DB {
	if ChCon == nil {
		//instanceHost := os.Getenv("INSTANCE_HOST")
		instanceHost, ex := os.LookupEnv("CH_INSTANCE_HOST")
		if !ex {
			instanceHost = "10.200.11.26"
			log.Printf("The env variable %s is not set.\n", "CH_INSTANCE_HOST")
		}
		//instanceHost := "10.200.11.26"
		//instancePort := 9091
		instancePortStr, ex := os.LookupEnv("CH_INSTANCE_PORT")
		var instancePort int
		var portParseError error
		if ex {
			instancePort = 9091
			log.Printf("The env variable %s is not set.\n", "CH_INSTANCE_PORT")
		} else {
			instancePort, portParseError = strconv.Atoi(instancePortStr)
			if portParseError != nil {
				log.Panic("clickhoue port define error !")
			}
		}
		chDatabase, ex := os.LookupEnv("CH_DATABASE")
		if !ex {
			chDatabase = "hehe"
			log.Printf("The env variable %s is not set.\n", "CH_DATABASE")
		}
		ChCon = ChConnect(instanceHost, instancePort, chDatabase)
		//return ChCon
	}
	return ChCon
}
