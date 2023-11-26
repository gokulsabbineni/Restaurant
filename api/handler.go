package api

import (
	"database/sql"
	"fmt"
	"net/http"
)

func AddTable(db *sql.DB) http.HandlerFunc {
	fmt.Println("here")
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method you have declared is not allowed", http.StatusMethodNotAllowed)
			return
		}

		if err := CreateTableLogic(db, w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func GetAllTables(db *sql.DB) http.HandlerFunc {
	fmt.Println("here")
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method you have declared is not allowed", http.StatusMethodNotAllowed)
			return
		}

		if err := GetAllTablesLogic(db, w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func FindTable(db *sql.DB) http.HandlerFunc {
	fmt.Println("here")
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method you have declared is not allowed", http.StatusMethodNotAllowed)
			return
		}

		if err := SearchTablesLogic(db, w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func RemoveTable(db *sql.DB) http.HandlerFunc {
	fmt.Println("here")
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method you have declared is not allowed", http.StatusMethodNotAllowed)
			return
		}

		if err := DeleteTableLogic(db, w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
