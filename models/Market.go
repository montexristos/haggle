package models

import "gorm.io/gorm"

type Market struct {
	gorm.Model
	Name       string
	Type       string
	MarketType string
	EventID    uint
	Selections []Selection
	Line       float64
}

// Set User's table name to be `profiles`
func (Market) TableName() string {
	return "markets"
}
