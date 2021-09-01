package service

import (
	"errors"
	"qr_checkin/internal/db"

	"gorm.io/gorm"
)

var (
	dbmgr *gorm.DB
)

// Init initialize service
func Init() error {
	dbmgr = db.GetDB()
	if dbmgr == nil {
		return errors.New("database manager not initialized")
	}
	return nil
}
