package models

import (
	"github.com/jinzhu/gorm"
)

type Event struct {
	ID         int `gorm:"primary_key"`
	Name       string
	Markets    []Market
	SiteID     int
	BetradarID int
}

// Set User's table name to be `profiles`
func (Event) TableName() string {
	return "events"
}

func GetCreateEvent(db *gorm.DB, eventID int, siteID int, name string) Event {
	var e Event
	db.Where("site_id = ? AND id = ?", siteID, eventID).First(&e)
	if e.ID == 0 {
		e = Event{
			Name:   name,
			ID:     eventID,
			SiteID: siteID,
		}
		db.Create(&e)
	}
	return e
}
