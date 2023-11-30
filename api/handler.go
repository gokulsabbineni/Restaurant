package api

import (
	"Restaurant/models"
	"database/sql"
	"encoding/json"
	"net/http"
)

func AddReservation(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var reservation models.Reservation
		if err := json.NewDecoder(r.Body).Decode(&reservation); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := CreateReservationLogic(db, w, reservation); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func UpdateReservation(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var reservation models.Reservation
		if err := json.NewDecoder(r.Body).Decode(&reservation); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := UpdateReservationLogic(db, w, reservation); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func CancelReservation(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var reservationID models.Reservation
		if err := json.NewDecoder(r.Body).Decode(&reservationID); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := CancelReservationLogic(db, w, reservationID.ReservationID); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func GetAllReservations(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		if err := GetAllReservationsLogic(db, w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

// Additional reservation-related handlers can be added as needed
