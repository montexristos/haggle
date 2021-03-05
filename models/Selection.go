package models

import "gorm.io/gorm"

type Selection struct {
	gorm.Model
	Name     string
	Price    float64
	Line     float64
	MarketID uint
}

// Set User's table name to be `profiles`
func (Selection) TableName() string {
	return "selections"
}
