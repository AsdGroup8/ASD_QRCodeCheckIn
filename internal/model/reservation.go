package model

// Reservation ...
type Reservation struct {
	ID        uint   `json:"reservationid"`
	StartAt   string `json:"starttime"`
	CreatedAt string `json:"reservationtime"`
	MovieName string `json:"movietitle"`
}
