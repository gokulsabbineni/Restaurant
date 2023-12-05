package api

import (
	"database/sql"
	"net/http"
)

func RegisterRoutes(db *sql.DB) {
	http.HandleFunc("/create", CreateHandler(db)) // function in handler.go
	http.HandleFunc("/update", UpdateHandler(db))
	http.HandleFunc("/delete", DeleteHandler(db))

}
