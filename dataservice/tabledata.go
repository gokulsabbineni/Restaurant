package dataservice

import (
	model "Restaurant/model"
	"database/sql"
	"encoding/json"
	"net/http"
)

func AddTable(db *sql.DB, w http.ResponseWriter, table model.Table) error {
	query := "INSERT INTO tables (table_id, table_number, capacity, location, description) VALUES (?,?,?,?,?)"
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
	query := "SELECT EXISTS(SELECT 1 FROM tables WHERE table_number=? OR table_id=?)"
	err := db.QueryRow(query, table_number, table_id).Scan(&exists)
	return exists, err
}

func TableLookup(db *sql.DB, table_number int) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM tables WHERE table_number=?)"
	err := db.QueryRow(query, table_number).Scan(&exists)
	return exists, err
}

func GetAllTables(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	var results []model.Table
	query := "SELECT * FROM tables"
	resp, err := db.Query(query)

	if err != nil {
		return err
	}

	defer resp.Close()

	for resp.Next() {
		var table_id int
		var table_number int
		var capacity int
		var location string
		var description string

		err := resp.Scan(&table_id, &table_number, &capacity, &location, &description)

		if err != nil {
			return err
		}

		results = append(results, model.Table{Table_id: table_id, Table_number: table_number, Capacity: capacity, Location: location, Description: description})
	}
	jsonBytes, err := json.Marshal(results)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
	return nil
}

func SearchTables(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	var results []model.Table
	param := r.URL.Query().Get("table_num")
	query := "SELECT * FROM tables WHERE table_number = ?"
	resp, err := db.Query(query, param)
	if err != nil {
		return err
	}

	defer resp.Close()

	for resp.Next() {
		var table_id int
		var table_number int
		var capacity int
		var location string
		var description string

		err := resp.Scan(&table_id, &table_number, &capacity, &location, &description)

		if err != nil {
			return err
		}

		results = append(results, model.Table{Table_id: table_id, Table_number: table_number, Capacity: capacity, Location: location, Description: description})
	}
	jsonBytes, err := json.Marshal(results)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
	return nil
}

func DeleteTable(db *sql.DB, w http.ResponseWriter, table_num int) error {

	query := "DELETE FROM tables WHERE table_number = ?"
	_, err := db.Exec(query, table_num)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusAccepted)
	return nil
}
