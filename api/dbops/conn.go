package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConn *sql.DB
	err error
)

func init() {
	dbConn, err = sql.Open("mysql",
		"root:umpay123@tcp(10.10.55.113:3306)/video_server?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
}