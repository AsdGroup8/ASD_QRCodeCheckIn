package model

import (
	"time"

	"github.com/AsdGroup8/ASD_QRCodeCheckIn/internal/ec"
	"gorm.io/gorm"
)

// Customer ...
type Customer struct {
	*gorm.Model
	Name     string `json:"name"`
	Gender   string `json:"gender"`
	DOB      int64  `json:"dob"`
	Email    string `gorm:"UNIQUE" json:"email"` // login
	Phone    string `gorm:"UNIQUE" json:"phone"`
	Password string `json:"password"`
	Address  string `json:"address"`
}

// BeforeCreate create hook
func (c *Customer) BeforeCreate(tx *gorm.DB) error {
	// check not null fields
	switch "" {
	case c.Name:
		return ec.ErrInvalidFormat
	case c.Email:
		return ec.ErrInvalidFormat
	case c.Phone:
		return ec.ErrInvalidFormat
	}
	// check dob is valid
	if c.DOB >= time.Now().Unix() {
		return ec.ErrInvalidFormat
	}
	return nil
}
