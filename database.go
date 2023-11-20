package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "username:password@tcp(localhost:3306)/restaurant")
	if err != nil {
		log.Fatal(err)
	}
}
