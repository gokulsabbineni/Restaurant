package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yourusername/restaurant/api"
)

func main() {
	r := mux.NewRouter()

	// Define your API endpoints
	r.HandleFunc("/api/reservation", api.AddReservationEndpoint).Methods("POST")

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
