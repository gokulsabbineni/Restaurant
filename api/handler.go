package api

import (
	"encoding/json"
	"net/http"

	"Restaurant/models"

	"github.com/gorilla/mux"
)

func CreateReservationHandler(w http.ResponseWriter, r *http.Request) {
	var reservation models.Reservation
	err := json.NewDecoder(r.Body).Decode(&reservation)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = ValidateReservation(reservation)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = CreateReservation(reservation)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func UpdateReservationHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	reservationID := vars["reservationID"]

	var updatedReservation models.Reservation
	err := json.NewDecoder(r.Body).Decode(&updatedReservation)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = UpdateReservation(reservationID, updatedReservation)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func CancelReservationHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	reservationID := vars["reservationID"]

	err := CancelReservation(reservationID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
