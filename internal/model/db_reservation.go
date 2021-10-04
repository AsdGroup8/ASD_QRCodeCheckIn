package model

import (
	"time"

	"gorm.io/gorm"
)

// DBReservation ...
type DBReservation struct {
	*gorm.Model
	CustomerID uint      `json:"customer_id"`
	MovieID    string    `json:"movie_id"`
	Customer   *Customer `gorm:"foreignKey:CustomerID"`
	StartTime  int64     `json:"starttime"`
}

// Packed pack DBReservation without endtime and movie name
func (dr *DBReservation) Packed() *Reservation {
	return &Reservation{
		ID:        dr.ID,
		CreatedAt: dr.CreatedAt.Format("2006-01-02 15:04"),
		StartAt:   time.Unix(dr.StartTime/1000, 0).Format("2006-01-02 15:04"),
	}
}
