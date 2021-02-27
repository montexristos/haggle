package models

type Market struct {
	ID         string `gorm:"primary_key;auto_increment:false"`
	EventID    int `gorm:"primary_key;auto_increment:false"`
	Name       string
	SiteID int `gorm:"primary_key;auto_increment:false"`
	Type       string
	Selections []Selection
}

// Set User's table name to be `profiles`
func (Market) TableName() string {
	return "markets"
}
