package models

type Selection struct {
	ID       uint `gorm:"primarykey"`
	Name     string
	Price    float64
	Line     float64
	MarketID uint
	Status   string
}

// Set User's table name to be `profiles`
func (Selection) TableName() string {
	return "selections"
}
