package model

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dsn = `tango:tango@unix(/var/run/mysqld/mysqld.sock)/sigo?charset=utf8mb4`
	db  *sql.DB
)

func init() {
	db, _ = sql.Open("mysql", dsn)
	db.SetMaxOpenConns(500)
	db.SetMaxIdleConns(300)
	db.Ping()

	err := db.Ping()
	if err != nil {
		fmt.Println("mysql connect fail\n" + err.Error())
	}
}
