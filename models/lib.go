package models

type Reservation struct {
	ReservationID   int
	CustomerID      int
	TableID         int
	ReservationDate string
	ReservationTime string
	PartySize       int
	Status          string
}
