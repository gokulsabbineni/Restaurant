package api

import (
	"errors"
	"time"

	"github.com/Anirudhvunnam/Restaurant/models"
)

func ValidateReservation(reservation models.Reservation) error {
	// Basic input validations
	if reservation.CustomerID <= 0 || reservation.TableID <= 0 || reservation.PartySize <= 0 {
		return errors.New("Invalid reservation details")
	}

	// Additional validations as needed

	return nil
}

func CreateReservation(reservation models.Reservation) error {
	// Implement logic to add reservation to the database
	// Set default status to "Pending"
	reservation.Status = "Pending"
	// Set current date and time
	reservation.ReservationDate = time.Now().Format("2006-01-02")
	reservation.ReservationTime = time.Now().Format("15:04:05")

	// Add reservation to the database
	// ...

	return nil
}

func UpdateReservation(reservationID string, updatedReservation models.Reservation) error {
	// Implement logic to update reservation in the database
	// ...

	return nil
}

func CancelReservation(reservationID string) error {
	// Implement logic to cancel reservation in the database
	// ...

	return nil
}
