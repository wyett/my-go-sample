package mysql

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

/**
 * @author     : wyettLei
 * @date       : Created in 2021/2/23 10:33
 * @description: TODO
 */

type mysqlconn struct {
	host     string
	port     uint32
	username string
	password string
	dbname   string
	charset  string
}

type tresult struct {
	id       int64
	no       int64
	name     string
	createat time.Time
}

func NewMysqlconn() {

}

func MySQLQuery() {

	var tt tresult
	//username:password@protocol(address)/dbname?param=value
	conn, err := sql.Open("mysql", "wyett_rw:jumpjump@tcp(10.16.17.103:3306)/wyett?charset=utf8mb4")
	if err != nil {
		panic(err)
	}
	conn.SetMaxIdleConns(10)
	conn.SetConnMaxLifetime(10)
	conn.SetConnMaxIdleTime(10)
	conn.SetMaxOpenConns(20)

	rows, err := conn.Query("select id from t limit 1")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		err = rows.Scan(&tt)
		if err != nil {
			panic(err)
		}

		fmt.Printf(tt.name)
	}

	//fmt.Println(row.Columns())
	//for i, val := range row {
	//}

	//fmt.Printf("",row.Columns())

	defer conn.Close()
}
