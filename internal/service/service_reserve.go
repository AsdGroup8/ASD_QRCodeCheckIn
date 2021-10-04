package service

import (
	"github.com/AsdGroup8/ASD_QRCodeCheckIn/internal/log"
	"github.com/AsdGroup8/ASD_QRCodeCheckIn/internal/model"
)

// GetAllReservation ...
func GetAllReservation(cusID uint) ([]model.DBReservation, error) {
	res := make([]model.DBReservation, 0, 4)
	if err := dbmgr.Where("customer_id = ?", cusID).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// CreateReservation ...
func CreateReservation(r *model.DBReservation) error {
	// TODO: Check movie id valid
	return dbmgr.Create(r).Error
}

// DeleteReservation ...
func DeleteReservation(reservID uint) error {
	if err := dbmgr.Delete(&model.DBReservation{}, reservID).Error; err != nil {
		log.Errorf("fail to delete reservation %d. %v", reservID, err)
		return err
	}
	return nil
}
