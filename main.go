package main

import (
	"net/http"

	"github.com/Anirudhvunnam/Restaurant/api" // Adjusted import path
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Set up routes
	api.SetReservationRoutes(r)

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
