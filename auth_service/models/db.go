package models

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func getDbCon() *sql.DB {
	db, err := sql.Open("mysql", "root:123qwe@tcp(mysql01:3306)/mysql01")
	if err != nil {
		panic(err)
	}
	migrateAndSeed(db)
	return db
}

func migrateAndSeed(db *sql.DB) {
	fmt.Println("checkTablesAndMigrate")
	// Check if tables exist
	_, err := db.Query("SELECT * FROM `users` LIMIT 1;")
	if err != nil {
		fmt.Println("MIGRATE")
		// Create tables
		_, _ = db.Exec(`
			CREATE TABLE users (
				id int NOT NULL AUTO_INCREMENT,
				email varchar(255) default '',
				password varchar(255) default '',
				PRIMARY KEY (id)
			);
		`)
		_, _ = db.Exec(`
			CREATE TABLE users_sessions (
				id int NOT NULL AUTO_INCREMENT,
				session_id varchar(255),
				user_id int,
				PRIMARY KEY (id)
			);
		`)
		// Password is 123
		_, _ = db.Exec(`
			INSERT INTO users (email,password) VALUES ('test@test.com','202cb962ac59075b964b07152d234b70');   
		`)
	}
}
