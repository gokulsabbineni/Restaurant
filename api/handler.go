package api

import (
	"database/sql"
	"net/http"
	//"Restaurant/user_accounts"
)

// CreateUserHandler is an HTTP handler for creating a new user.
func CreateHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check if the request method is POST
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Call the CreateUser function to handle user creation
		if err := CreateUser(db, w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
