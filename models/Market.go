package models

import "gorm.io/gorm"

type Market struct {
	gorm.Model
	Name       string
	Type       string
	EventID    uint
	Selections []Selection
}

// Set User's table name to be `profiles`
func (Market) TableName() string {
	return "markets"
}
