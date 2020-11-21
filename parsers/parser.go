package parsers

import (
	"haggle/models"

	"github.com/gocolly/colly"
	"github.com/jinzhu/gorm"
)

type Parser interface {
	Scrape(config models.SiteConfig, c *colly.Collector) (bool, error)
	SetDB(db *gorm.DB)
}
