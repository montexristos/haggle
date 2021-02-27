package models

type Selection struct {
	ID       string `gorm:"primary_key;auto_increment:false"`
	SiteID     int `gorm:"primaryKey;autoIncrement:false"`
	MarketID string `gorm:"primaryKey;autoIncrement:false"`
	Name     string
	Price    float64
	Line     float64
}

// Set User's table name to be `profiles`
func (Selection) TableName() string {
	return "selections"
}
