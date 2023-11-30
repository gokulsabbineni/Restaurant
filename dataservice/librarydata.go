package dataservice

import (
	"Restaurant/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

func AddReservation(db *sql.DB, w http.ResponseWriter, reservation models.Reservation) error {
	fmt.Println(reservation)

	query := "INSERT INTO reservations (reservation_id, customer_id, table_id, reservation_date, reservation_time, party_size, status) VALUES (?,?,?,?,?,?,?)"
	_, err := db.Exec(query, reservation.ReservationID, reservation.CustomerID, reservation.TableID, reservation.ReservationDate, reservation.ReservationTime, reservation.PartySize, reservation.Status)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(reservation)
	return nil
}

func UpdateReservation(db *sql.DB, w http.ResponseWriter, reservation models.Reservation) error {
	fmt.Println(reservation)

	query := "UPDATE reservations SET customer_id=?, table_id=?, reservation_date=?, reservation_time=?, party_size=?, status=? WHERE reservation_id=?"
	_, err := db.Exec(query, reservation.CustomerID, reservation.TableID, reservation.ReservationDate, reservation.ReservationTime, reservation.PartySize, reservation.Status, reservation.ReservationID)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	return nil
}

func CancelReservation(db *sql.DB, w http.ResponseWriter, reservationID int) error {
	query := "UPDATE reservations SET status='Cancelled' WHERE reservation_id=?"
	_, err := db.Exec(query, reservationID)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	return nil
}

func GetAllReservations(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	var results []models.Reservation
	query := "SELECT * FROM reservations"
	resp, err := db.Query(query)

	if err != nil {
		return err
	}

	defer resp.Close()

	for resp.Next() {
		var reservationID int
		var customerID int
		var tableID int
		var reservationDate string
		var reservationTime string
		var partySize int
		var status string

		err := resp.Scan(&reservationID, &customerID, &tableID, &reservationDate, &reservationTime, &partySize, &status)

		if err != nil {
			return err
		}

		results = append(results, models.Reservation{ReservationID: reservationID, CustomerID: customerID, TableID: tableID, ReservationDate: reservationDate, ReservationTime: reservationTime, PartySize: partySize, Status: status})
	}
	jsonBytes, err := json.Marshal(results)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
	return nil
}

// ReservationExists checks if a reservation with the given ID exists in the database.
func ReservationExists(db *sql.DB, reservationID int) (bool, error) {
	query := "SELECT EXISTS(SELECT 1 FROM reservations WHERE reservation_id = ?)"
	var exists bool
	err := db.QueryRow(query, reservationID).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

// Additional reservation-related functions can be added as needed
