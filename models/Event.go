package models

type Event struct {
	ID      int `gorm:"primary_key"`
	Name    string
	Markets []Market
}

// Set User's table name to be `profiles`
func (Event) TableName() string {
	return "events"
}