package api

import (
	model "Restaurant/model"
	"database/sql"
	"encoding/json"
	"net/http"
)

func AddTable(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method you have declared is not allowed", http.StatusMethodNotAllowed)
			return
		}
		var table model.Table
		if err := json.NewDecoder(r.Body).Decode(&table); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := CreateTableLogic(db, w, table); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func GetAllTables(db *sql.DB) http.HandlerFunc {
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
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method you have declared is not allowed", http.StatusMethodNotAllowed)
			return
		}
		var table model.Table_num
		if err := json.NewDecoder(r.Body).Decode(&table); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := DeleteTableLogic(db, w, table.Table_number); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
