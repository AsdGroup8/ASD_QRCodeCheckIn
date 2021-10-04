package service

import (
	"errors"

	"github.com/AsdGroup8/ASD_QRCodeCheckIn/internal/db"
	"github.com/AsdGroup8/ASD_QRCodeCheckIn/internal/log"
	"gorm.io/gorm"
)

var dbmgr *gorm.DB

// InitMovies ...
func InitMovies() {
	if _, err := GetInTheatersMovies(); err != nil {
		log.Errorf("fail to get in theaters movies. %v", err)
	}
}

// Init initialize service
func Init() error {
	dbmgr = db.GetDB()
	if dbmgr == nil {
		return errors.New("database manager not initialized")
	}
	return nil
}
