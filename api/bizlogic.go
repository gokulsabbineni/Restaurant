package api

import (
	dataservice "Restaurant/dataservice"
	"Restaurant/model"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func CreateTableLogic(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	var table model.Table
	if err := json.NewDecoder(r.Body).Decode(&table); err != nil {
		return err
	}
	// fmt.Println(table.Author)
	if exists, err := dataservice.TableExists(db, table.Table_number, table.Table_id); err != nil {
		fmt.Println(err)
		return err
	} else if exists {
		fmt.Println(exists)
		http.Error(w, "table already exists", http.StatusBadRequest)
		return errors.New("table already exists")
	}
	return dataservice.AddTable(db, w, r)
}

func GetAllTablesLogic(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	return dataservice.GetAllTables(db, w, r)
}
