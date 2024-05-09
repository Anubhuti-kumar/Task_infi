package Config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB() *sql.DB {
	var err error
	db, err = sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/cetec")
	if err != nil {
		panic(err)
	}
	return db
}
