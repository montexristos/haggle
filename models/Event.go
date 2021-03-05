package models

import (
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	SiteID     int
	BetradarID int
	Name       string
	Markets    []Market
}

// Set User's table name to be `profiles`
func (Event) TableName() string {
	return "events"
}

func GetCreateEvent(db *gorm.DB, eventID int, siteID int, name string) Event {
	var e Event
	db.Where("site_id = ? AND betradar_id = ?", siteID, eventID).Preload("Markets").Preload("Markets.Selections").First(&e)
	if e.ID == 0 {
		e = Event{
			Name:       name,
			BetradarID: eventID,
			SiteID:     siteID,
		}
	}
	return e
}
