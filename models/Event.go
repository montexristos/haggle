package models

import (
	"gorm.io/gorm"
)

type Event struct {
	ID            uint `gorm:"primarykey"`
	SiteID        int
	BetradarID    string
	Name          string
	Date          string
	CanonicalName string
	Markets       []Market
	Url           string
	Tournament    string
	Live          bool
	Time          float64
}

// Set User's table name to be `profiles`
func (Event) TableName() string {
	return "events"
}

func GetCreateEvent(db *gorm.DB, eventID string, siteID int, name string, live bool) Event {
	if live {
		return Event{
			Name:       name,
			BetradarID: eventID,
			SiteID:     siteID,
			Live:       live,
		}
	}
	var e Event
	db.Where("site_id = ? AND betradar_id = ?", siteID, eventID).Preload("Markets").Preload("Markets.Selections").First(&e)
	if e.ID == 0 {
		e = Event{
			Name:       name,
			BetradarID: eventID,
			SiteID:     siteID,
			Live:       live,
		}
	}
	return e
}
