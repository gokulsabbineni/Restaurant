package main

import (
	"Restaurant/api"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Set up routes
	api.SetReservationRoutes(r)

	http.Handle("/", r)
	http.ListenAndServe(":3306", nil)
}
