package api

import (
	"github.com/gorilla/mux"
)

func SetReservationRoutes(router *mux.Router) {
	router.HandleFunc("/api/reservation", CreateReservationHandler).Methods("POST")
	router.HandleFunc("/api/reservation/{reservationID}", UpdateReservationHandler).Methods("PUT")
	router.HandleFunc("/api/reservation/{reservationID}", CancelReservationHandler).Methods("DELETE")
	// Add more routes as needed
}
