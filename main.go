package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	api "Restaurant/api"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "root:password@tcp(127.0.0.1:3306)/restaurant_db?parseTime=true"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	fmt.Println("DB drivers established")

	api.RegisterRoutes(db)

	fmt.Println("Sent Request")

	log.Println("Listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
