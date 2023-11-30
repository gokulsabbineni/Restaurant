package api

import (
	dataservice "Restaurant/dataservice"
	"Restaurant/models"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
)

func CreateReservationLogic(db *sql.DB, w http.ResponseWriter, reservation models.Reservation) error {
	if exists, err := dataservice.ReservationExists(db, reservation.ReservationID); err != nil {
		fmt.Println(err)
		return err
	} else if exists {
		fmt.Println(exists)
		http.Error(w, "reservation already exists", http.StatusBadRequest)
		return errors.New("reservation already exists")
	}

	return dataservice.AddReservation(db, w, reservation)
}

func UpdateReservationLogic(db *sql.DB, w http.ResponseWriter, reservation models.Reservation) error {
	if exists, err := dataservice.ReservationExists(db, reservation.ReservationID); err != nil {
		fmt.Println(err)
		return err
	} else if !exists {
		fmt.Println(exists)
		http.Error(w, "reservation doesn't exist", http.StatusBadRequest)
		return errors.New("reservation doesn't exist")
	}

	return dataservice.UpdateReservation(db, w, reservation)
}

func CancelReservationLogic(db *sql.DB, w http.ResponseWriter, reservationID int) error {
	if exists, err := dataservice.ReservationExists(db, reservationID); err != nil {
		fmt.Println(err)
		return err
	} else if !exists {
		fmt.Println(exists)
		http.Error(w, "reservation doesn't exist", http.StatusBadRequest)
		return errors.New("reservation doesn't exist")
	}

	return dataservice.CancelReservation(db, w, reservationID)
}

func GetAllReservationsLogic(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	return dataservice.GetAllReservations(db, w, r)
}

// Additional reservation-related business logic functions can be added as needed
