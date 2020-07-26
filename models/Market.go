package models

type Market struct {
	ID         string `gorm:"primary_key;auto_increment:false"`
	EventID    int
	Name       string
	MarketId   string
	Type       string
	Selections []Selection
}

// Set User's table name to be `profiles`
func (Market) TableName() string {
	return "markets"
}
