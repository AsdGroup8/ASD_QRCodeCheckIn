package model

import "gorm.io/gorm"

// Staff ...
type Staff struct {
	*gorm.Model
	FirstName  string
	LastName   string
	Email      string `gorm:"UNIQUE"`
	Password   string
	Department string
	Privilege  int8
}
