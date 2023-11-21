package api

import (
	"Restaurant/model"
	"database/sql"
	"encoding/json"
	"net/http"
)

func CreateUser(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	var user model.User

	// Decode the JSON request body into the user struct
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return err
	}

	// Perform the database query to insert the new user
	query := `INSERT INTO user(id, username, firstname, lastname, roles, email, contact) VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err := db.Exec(query, user.Id, user.Username, user.Firstname, user.Lastname, user.Roles, user.Email, user.Contact)
	if err != nil {
		return err
	}

	// Set the HTTP status code to 201 (Created)
	w.WriteHeader(http.StatusCreated)

	// Encode the user as JSON and write it to the response
	json.NewEncoder(w).Encode(user)

	return nil
}
