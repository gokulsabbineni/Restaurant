package api

import (
	dataservice "Restaurant/dataservice"
	"Restaurant/model"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
)

func CreateTableLogic(db *sql.DB, w http.ResponseWriter, table model.Table) error {
	// fmt.Println(table.Author)
	if exists, err := dataservice.TableExists(db, table.Table_number, table.Table_id); err != nil {
		fmt.Println(err)
		return err
	} else if exists {
		fmt.Println(exists)
		http.Error(w, "table already exists", http.StatusBadRequest)
		return errors.New("table already exists")
	}
	return dataservice.AddTable(db, w, table)
}

func GetAllTablesLogic(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	return dataservice.GetAllTables(db, w, r)
}

func SearchTablesLogic(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	return dataservice.SearchTables(db, w, r)
}

func DeleteTableLogic(db *sql.DB, w http.ResponseWriter, table_num int) error {
	// fmt.Println(table.Author)
	if exists, err := dataservice.TableLookup(db, table_num); err != nil {
		fmt.Println(err)
		return err
	} else if !exists {
		fmt.Println(exists)
		http.Error(w, "table doesn't exist", http.StatusBadRequest)
		return errors.New("table doesn't exist")
	}
	return dataservice.DeleteTable(db, w, table_num)
}
