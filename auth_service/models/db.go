package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func getDbCon() *sql.DB {
	db, err := sql.Open("mysql", "root:123qwe@tcp(127.0.0.1:3306)/sys01")
	if err != nil {
		panic(err)
	}
	return db
}
