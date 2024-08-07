package dataservice

import (
	"Restaurant/model"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

func UserExists(db *sql.DB, id string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM user WHERE Id=?)"
	err := db.QueryRow(query, id).Scan(&exists)
	return exists, err
}

func AddQuery(db *sql.DB, w http.ResponseWriter, user model.User) error {

	query := `INSERT INTO user(id, username, firstname, lastname, password1, roles, email, contact) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := db.Exec(query, user.Id, user.Username, user.Firstname, user.Lastname, user.Password1, user.Roles, user.Email, user.Contact)
	if err != nil {
		return err
	}
	// Set the HTTP status code to 201 (Created)
	w.WriteHeader(http.StatusCreated)

	// Encode the user as JSON and write it to the response
	json.NewEncoder(w).Encode(user)

	fmt.Println("Hello")

	return nil
}

func UpadateQuery(db *sql.DB, w http.ResponseWriter, r *http.Request, id string, updatedUser model.User) error {

	query := "UPDATE user SET username = ?, firstname = ?, password1 = ?, lastname = ?, roles = ?, email = ?, contact = ? WHERE id = ?"

	_, err := db.Exec(query,
		updatedUser.Username,
		updatedUser.Firstname,
		updatedUser.Password1,
		updatedUser.Lastname,
		updatedUser.Roles,
		updatedUser.Email,
		updatedUser.Contact,
		id,
	)
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User updated successfully"))

	return nil

}
func DeleteQuery(db *sql.DB, w http.ResponseWriter, r *http.Request, id string) error {

	query := "DELETE FROM user WHERE id = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User deleted successfully"))
	return nil
}
