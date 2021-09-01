package parsers

import (
	"github.com/gocolly/colly"
	"gorm.io/gorm"
	"haggle/models"
)

type Bet struct {
	Parser
	db     *gorm.DB
	config *models.SiteConfig
	c      *colly.Collector
	ID     int
}
