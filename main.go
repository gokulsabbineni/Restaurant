package main

import (
	"Restaurant/api"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	fmt.Println("Welcome to Restaurant")

	dsn := "root:Yashwanth@2199@tcp(127.0.0.1:3306)/sys?parseTime=true" // syntax dsn(Data source name) := "username:password@tcp(127.0.0.1:3306)/sys?parseTime=true"

	db, err := sql.Open("mysql", dsn) // it does not really opens the sql. It jst logins and here it expects two outputs so here we use db and err as output
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	api.RegisterRoutes(db)

	log.Println("server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
