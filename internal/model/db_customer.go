package model

import (
	"time"

	"github.com/AsdGroup8/ASD_QRCodeCheckIn/internal/ec"
	"gorm.io/gorm"
)

// Customer ...
type Customer struct {
	*gorm.Model
	Name     string
	Gender   EGender
	DOB      int64
	Email    string `gorm:"UNIQUE"` // login
	Phone    string `gorm:"UNIQUE"`
	Password string
	Address  string
}

// BeforeCreate create hook
func (c *Customer) BeforeCreate(tx *gorm.DB) error {
	// check not null fields
	switch "" {
	case c.Name:
		return ec.ErrCustomerNameEmpty
	case c.Email:
		return ec.ErrCustomerEmailEmpty
	case c.Phone:
		return ec.ErrCustomerPhoneEmpty
	}
	// check dob is valid
	if c.DOB >= time.Now().Unix() {
		return ec.ErrCustomerDOBInvalid
	}
	return nil
}
