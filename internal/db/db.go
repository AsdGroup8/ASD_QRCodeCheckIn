package db

import (
	"github.com/AsdGroup8/ASD_QRCodeCheckIn/internal/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

type M = map[string]interface{}

// Init initialize database
func Init(connStr string) error {
	dbConn, err := gorm.Open(mysql.Open(connStr), &gorm.Config{
		Logger: logger.New(log.GetDBLogger(), logger.Config{
			Colorful: true,
		}),
	})
	if err != nil {
		return err
	}
	db = dbConn
	return nil
}

// GetDB ...
func GetDB() *gorm.DB {
	return db
}
