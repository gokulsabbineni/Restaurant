package api

import (
	"database/sql"
	"fmt"
	"net/http"
)

func RegisterRoutes(db *sql.DB) {
	fmt.Println("here")

	// Reservation routes
	http.HandleFunc("/add_reservation", AddReservation(db))
	http.HandleFunc("/update_reservation", UpdateReservation(db))
	http.HandleFunc("/cancel_reservation/{id}", CancelReservation(db))
	http.HandleFunc("/reservations", GetAllReservations(db))

	// Additional reservation routes can be added as needed
}
