package models

type Market struct {
	ID         uint `gorm:"primarykey"`
	Name       string
	Type       string
	MarketType string
	EventID    uint
	Selections []Selection
	Line       float64
	Status     string
}

// Set User's table name to be `profiles`
func (Market) TableName() string {
	return "markets"
}
