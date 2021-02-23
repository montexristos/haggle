package parsers

import (
	"github.com/gocolly/colly"
	"haggle/models"
	"log"
	"time"

	"github.com/jinzhu/gorm"
)

type Parser interface {
	Initialize()
	Scrape() (bool, error)
	ScrapeHome() (bool, error)
	ScrapeLive() (bool, error)
	ScrapeToday() (bool, error)
	ScrapeTournament(tournamentId string) (bool, error)
	SetDB(db *gorm.DB)
	SetConfig(config *models.SiteConfig)
}

func GetCollector() *colly.Collector {
	c := colly.NewCollector(
		//colly.CacheDir("./_instagram_cache/"),
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
	)

	// Limit the number of threads started by colly to one
	// slow but make sure we don't get banned.
	// can be limited to domain
	_ = c.Limit(&colly.LimitRule{
		Parallelism: 1,
		Delay:       3 * time.Second,
	})
	// Before making a request put the URL with
	// the key of "url" into the context of the request
	c.OnRequest(func(r *colly.Request) {
		r.Ctx.Put("path", r.URL.String())
		log.Println("Visiting", r.URL)
	})
	c.OnError(func(r *colly.Response, e error) {
		log.Println("error:", e, r.Request.URL, string(r.Body))
	})
	return c
}
