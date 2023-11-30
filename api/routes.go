package api

import (
	"database/sql"
	"net/http"
)

func RegisterRoutes(db *sql.DB) {
	http.HandleFunc("/add", AddTable(db))
	http.HandleFunc("/tables", GetAllTables(db))
	http.HandleFunc("/table", FindTable(db))
	http.HandleFunc("/remove_table", RemoveTable(db))
	// http.HandleFunc("/update", UpdateHandler(db))
}
