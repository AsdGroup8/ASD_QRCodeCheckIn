package db

import (
	"github.com/AsdGroup8/ASD_QRCodeCheckIn/internal/log"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

// M ...
type M = map[string]interface{}

// Init initialize database
func Init(connStr string) error {
	log.Debug("connecting to database", zap.String("connstr", connStr))
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
