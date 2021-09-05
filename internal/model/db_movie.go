package model

import "gorm.io/gorm"

// Movie ...
type Movie struct {
	*gorm.Model
	Name        string
	Desc        string
	OriginPrice float32
	Discount    float32
}

// Price movie ticket price
func (m *Movie) Price() float32 {
	return m.OriginPrice * (1 - m.Discount)
}
