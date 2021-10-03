package service

import (
	"github.com/AsdGroup8/ASD_QRCodeCheckIn/internal/log"
	"github.com/AsdGroup8/ASD_QRCodeCheckIn/internal/model"
)

// GetAllReservation ...
func GetAllReservation(cusID uint) []model.Reservation {
	res := make([]model.Reservation, 0, 4)
	if err := dbmgr.Where("customer_id = ?", cusID).Find(&res).Error; err != nil {
		log.Errorf("fail to find reservation for customer %d. %v", cusID, err)
		return nil
	}
	return res
}
func DeleteReservation(reservID uint) error {
	if err := dbmgr.Delete(&model.Reservation{}, reservID).Error; err != nil {
		log.Errorf("fail to delete reservation %d. %v", reservID, err)
		return err
	}
	return nil
}
