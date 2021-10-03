package service

import (
	"github.com/AsdGroup8/ASD_QRCodeCheckIn/internal/model"
)

// GetAllReservation ...
func GetAllReservation(cusID uint) ([]model.Reservation, error) {
	res := make([]model.Reservation, 0, 4)
	if err := dbmgr.Where("customer_id = ?", cusID).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// CreateReservation ...
func CreateReservation(r *model.Reservation) error {
	// TODO: Check movie id valid
	return dbmgr.Create(r).Error
}
