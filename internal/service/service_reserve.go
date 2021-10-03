package service

import (
	"github.com/AsdGroup8/ASD_QRCodeCheckIn/internal/log"
	"github.com/AsdGroup8/ASD_QRCodeCheckIn/internal/model"
)

// GetAllReservation ...
func GetAllReservation(cusID int) []model.Reservation {
	res := make([]model.Reservation, 0, 4)
	if err := dbmgr.Where("customer_id = ?", cusID).Find(&res).Error; err != nil {
		log.Errorf("fail to find reservation for customer %d. %v", cusID, err)
		return nil
	}
	return res
}
