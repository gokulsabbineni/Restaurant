package api

import (
	"database/sql"
	"fmt"
	"net/http"
)

func RegisterRoutes(db *sql.DB) {
	fmt.Println("here")
	http.HandleFunc("/add", AddTable(db))
	http.HandleFunc("/tables", GetAllTables(db))
	http.HandleFunc("/table/{table_num}", FindTable(db))
	http.HandleFunc("/remove_table", RemoveTable(db))
	// http.HandleFunc("/update", UpdateHandler(db))
}
