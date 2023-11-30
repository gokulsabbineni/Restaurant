package models

type Reservation struct {
	ReservationID   int    `json:"reservation_id"`
	CustomerID      int    `json:"customer_id"`
	TableID         int    `json:"table_id"`
	ReservationDate string `json:"reservation_date"`
	ReservationTime string `json:"reservation_time"`
	PartySize       int    `json:"party_size"`
	Status          string `json:"status"`
}
