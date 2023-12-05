package api

import (
	"Restaurant/model"
	"database/sql"
	"encoding/json"
	"net/http"
)

// CreateUserHandler is an HTTP handler for creating a new user.
func CreateHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check if the request method is POST
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var user model.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Call the CreateUser function to handle user creation
		if err := CreateUser(db, w, user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

		}
	}
}

func UpdateHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Create a User object to hold the updated data
		var updatedUser model.User

		// Decode the JSON request body into the User struct
		err := json.NewDecoder(r.Body).Decode(&updatedUser)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if err := UpdateUser(db, w, r, updatedUser.Id, updatedUser); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

// DeleteHandler is an HTTP handler for deleting a user.
func DeleteHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check if the request method is DELETE
		if r.Method != http.MethodDelete {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var user model.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Call the DeleteUser function to handle user deletion

		if err := DeleteUser(db, w, r, user.Id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
