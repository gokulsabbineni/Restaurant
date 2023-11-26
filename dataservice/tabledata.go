package dataservice

import (
	model "Restaurant/model"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

func AddTable(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	var table model.Table
	if err := json.NewDecoder(r.Body).Decode(&table); err != nil {
		return err
	}

	fmt.Println(table)

	query := "INSERT INTO tables (table_id, table_number, capacity, location, description) VALUES (?,?,?)"
	_, err := db.Exec(query, table.Table_id, table.Table_number, table.Capacity, table.Location, table.Description)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(table)
	return nil
}

func TableExists(db *sql.DB, table_number int, table_id int) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM tables WHERE table_number=? or table_id=?)"
	err := db.QueryRow(query, table_number, table_id).Scan(&exists)
	return exists, err
}

func GetAllTables(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	query := "SELECT * FROM tables"
	resp, err := db.Exec(query)

}
