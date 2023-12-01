package api

import (
	"Restaurant/dataservice"
	"Restaurant/model"
	"database/sql"
	"errors"
	"net/http"
)

func CreateUser(db *sql.DB, w http.ResponseWriter, user model.User) error {

	if exists, err := dataservice.UserExists(db, user.Id); err != nil {
		return err
	} else if exists {
		http.Error(w, "user already exists", http.StatusBadRequest)
		return errors.New("user already exists")
	}
	return dataservice.AddQuery(db, w, user)

}

func UpdateUser(db *sql.DB, w http.ResponseWriter, r *http.Request, id string, updatedUser model.User) error {

	if exists, err := dataservice.UserExists(db, updatedUser.Id); err != nil {
		return err
	} else if exists {
		return dataservice.UpadateQuery(db, w, r, id, updatedUser)

	}
	return nil

}

// DeleteUser deletes a user from the database.
func DeleteUser(db *sql.DB, w http.ResponseWriter, r *http.Request, id string) error {
	// if exists, err := dataservice.UserLookup(db, id); err != nil {
	// 	return err
	// } else if exists {
	// 	http.Error(w, "user doesn't exist", http.StatusBadRequest)
	// 	return errors.New("user doesn't exist")
	// }

	return dataservice.DeleteQuery(db, w, r, id)

}
