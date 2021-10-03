package service

import "github.com/AsdGroup8/ASD_QRCodeCheckIn/internal/model"

// MigrateModel migrate models
func MigrateModel() error {
	if err := dbmgr.AutoMigrate(&model.Customer{}); err != nil {
		return err
	}
	if err := dbmgr.AutoMigrate(&model.Staff{}); err != nil {
		return err
	}
	return dbmgr.AutoMigrate(&model.Reservation{})
}
