package model

import "gorm.io/gorm"

// Person person who made reservation to check in
type Person struct {
	*gorm.Model
	Name    string
	Gender  EGender
	Age     int16
	Email   string `gorm:"UNIQUE"`
	Phone   string `gorm:"UNIQUE"`
	Address string
}
