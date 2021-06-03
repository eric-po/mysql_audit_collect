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
	"sync"
	"time"
)

//var sqlFormat = Sprintf("insert into %s(db_host  ,db_port  ,query_timestamp  ,serverhost  ,username  ,host  ,connectionid  ,queryid  ,operation  ,database  ,object  ,retcode) values(?,?,?,?,?,?,?,?,?,?,?,?) ;", table_name)

var sqlFormatStandard = "insert into %s(db_host  ,db_port  ,query_timestamp  ,serverhost  ,username  ,host  ,connectionid  ,queryid  ,operation  ,database  ,object  ,retcode) values(?,?,?,?,?,?,?,?,?,?,?,?) ;"

func timeStringFormat(timeStr string) string {
	t, _ := time.Parse("20060102 15:04:05", timeStr)
	a := t.Format("2006-01-02 15:04:05")
	return a
}

func QueryRecordHandleBatch(cMessage []*sarama.ConsumerMessage) {
	chCon := GetChCon()

	chTable, ex := os.LookupEnv("CH_TABLE")
	if !ex {
		chTable = "sql_record"
		log.Printf("The env variable %s is not set.\n", "CH_TABLE")
	}
	sqlFormat := fmt.Sprintf(sqlFormatStandard, chTable)
	//fmt.Println(sqlFormat)
	var (
		tx, _   = chCon.Begin()
		stmt, _ = tx.Prepare(sqlFormat)
	)

	for _, message := range cMessage {
		qr := data_parse.RecordParse(string(message.Value))
		msg := data_parse.MessageParse(qr.Message)
		if msg.Object == "" {
			continue
		}

		if _, err := stmt.Exec(
			qr.DbHost,
			qr.DbPort,
			//msg.QueryTimestamp,
			timeStringFormat(msg.QueryTimestamp),
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
	}

	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
}

//func QueryRecordHandle(message *sarama.ConsumerMessage) {
//	var qr *data_parse.QueryRecord
//	qr = data_parse.RecordParse(string(message.Value))
//	msg := data_parse.MessageParse(qr.Message)
//	if msg.Object == "" {
//		return
//	}
//	//fmt.Println(*msg)
//	chCon := GetChCon()
//	var (
//		tx, _   = chCon.Begin()
//		stmt, _ = tx.Prepare(sqlFormat)
//	)
//	//stmt.Exec(qr.DbHost,qr.DbPort,*msg.queryTimestamp,*msg.serverhost)
//	if _, err := stmt.Exec(
//		qr.DbHost,
//		qr.DbPort,
//		//msg.QueryTimestamp,
//		timeStringFormat(msg.QueryTimestamp),
//		msg.Serverhost,
//		msg.Username,
//		msg.Host,
//		msg.Connectionid,
//		msg.Queryid,
//		msg.Operation,
//		msg.Database,
//		msg.Object,
//		msg.Retcode,
//	); err != nil {
//		log.Fatal(err)
//	}
//	if err := tx.Commit(); err != nil {
//		log.Fatal(err)
//	}
//}

var ChCon *sql.DB
var dbMutex sync.Mutex

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
	dbMutex.Lock()
	defer dbMutex.Unlock()
	for {
		if ChCon == nil {
			log.Println("Database pointer is nil . Now we create connection to the clickhouse . ")
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
			if !ex {
				instancePort = 9091
				log.Printf("The env variable %s is not set.\n", "CH_INSTANCE_PORT")
			} else {
				instancePort, portParseError = strconv.Atoi(instancePortStr)
				if portParseError != nil {
					log.Panic("clickhouse port define error !")
				}
			}
			chDatabase, ex := os.LookupEnv("CH_DATABASE")
			if !ex {
				chDatabase = "hehe"
				log.Printf("The env variable %s is not set.\n", "CH_DATABASE")
			}
			fmt.Println(instanceHost, instancePort, chDatabase)
			ChCon = ChConnect(instanceHost, instancePort, chDatabase)
			//return ChCon
		}
		return ChCon
	}
}
