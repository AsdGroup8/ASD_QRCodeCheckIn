package model

import (
	"database/sql"

	"gorm.io/gorm"
)

// Reservation ...
type Reservation struct {
	*gorm.Model
	CustomerID uint         `json:"customer_id"`
	MovieID    uint         `json:"movie_id"`
	Customer   *Customer    `gorm:"foreignKey:CustomerID"`
	Movie      *Movie       `gorm:"foreignKey:MovieID"`
	StartTime  sql.NullTime `json:"starttime"`
	EndTime    sql.NullTime `json:"EndTime"`
}
