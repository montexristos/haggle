package parsers

import (
	"github.com/gocolly/colly"
	"github.com/jinzhu/gorm"
	"haggle/models"
)

type Parser interface {
	Scrape(config models.SiteConfig, c *colly.Collector, db *gorm.DB) (bool, error)
}
