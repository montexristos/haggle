package models

type Selection struct {
	ID       string `gorm:"primary_key;auto_increment:false"`
	MarketID string
	Name     string
	Price    float64
}

// Set User's table name to be `profiles`
func (Selection) TableName() string {
	return "selections"
}