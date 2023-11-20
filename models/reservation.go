package models

import "time"

type Reservation struct {
	ID              int
	UserID          int
	ReservationTime time.Time
	Status          string
}
