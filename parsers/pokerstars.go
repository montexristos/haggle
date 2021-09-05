package parsers

import (
	"fmt"
	"haggle/models"
	"strings"

	"github.com/Jeffail/gabs"
	"github.com/gocolly/colly"
	"gorm.io/gorm"
)

type PokerStars struct {
	Parser
	db     *gorm.DB
	config *models.SiteConfig
	c      *colly.Collector
	ID     int
}

func (p *PokerStars) SetConfig(c *models.SiteConfig) {
	p.config = c
	p.ID = c.SiteID
}

func (p *PokerStars) GetConfig() *models.SiteConfig {
	return p.config
}

func (p *PokerStars) Initialize() {
	p.c = GetCollector()
}

func (p *PokerStars) GetDB() *gorm.DB {
	return p.db
}

func (p *PokerStars) Scrape() (bool, error) {
	p.c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Content-Type", "application/json;charset=UTF-8")
	})
	// err := c.PostRaw(URL, payload)
	p.c.OnHTML("script", func(e *colly.HTMLElement) {
		if strings.HasPrefix(e.Text, `window["initial_state"]=`) {
			jsonParsed, err := gabs.ParseJSON([]byte(e.Text[24:]))
			if err != nil {
				fmt.Println(err.Error())
			}
			topEvents := jsonParsed.Path("data.topEvents").Data()
			if topEvents != nil {
				p.parseTopEvents(topEvents.([]interface{}))
			}
			//TODO parse other items (fmt.Println(jsonParsed))
		}
	})

	_ = p.c.Visit(fmt.Sprintf("%s", p.config.BaseUrl))
	return true, nil
}

func (p *PokerStars) ScrapeHome() (bool, error) {
	return true, nil
}

func (p *PokerStars) ScrapeLive() (bool, error) {
	_ = p.c.Visit(fmt.Sprintf("%s/%s", p.config.BaseUrl, p.config.Urls["live"]))
	return true, nil
}

func (p *PokerStars) ScrapeToday() (bool, error) {
	_ = p.c.Visit(fmt.Sprintf("%s/%s", p.config.BaseUrl, p.config.Urls["day"]))
	return true, nil
}

func (p *PokerStars) ScrapeTournament(tournamentId string) (bool, error) {
	return true, nil
}

func (p *PokerStars) SetDB(db *gorm.DB) {
	p.db = db
}

func (p *PokerStars) parseTopEvents(sports []interface{}) {
	for i := 0; i < len(sports); i++ {
		sport := sports[0].(map[string]interface{})
		events := sport["events"].([]interface{})
		if len(events) > 0 {
			for j := 0; j < len(events); j++ {
				_, _ = ParseEvent(p, events[j].(map[string]interface{}))
			}
		}
	}
}

func (p *PokerStars) GetEventID(event map[string]interface{}) int {
	return int(event["betRadarId"].(float64))
}

func (p *PokerStars) GetEventName(event map[string]interface{}) string {
	return event["name"].(string)
}

func (p *PokerStars) GetEventIsAntepost(event map[string]interface{}) bool {
	return false
}

func (p *PokerStars) GetEventMarkets(event map[string]interface{}) []interface{} {
	return event["markets"].([]interface{})
}

func (p *PokerStars) ParseMarketType(market map[string]interface{}) string {
	return market["type"].(string)
}

func (p *PokerStars) ParseMarketId(market map[string]interface{}) string {
	return market["id"].(string)
}

func (p *PokerStars) ParseMarketName(market map[string]interface{}) string {
	return market["name"].(string)
}

func (p *PokerStars) ParseSelectionName(selectionData map[string]interface{}) string {
	return selectionData["name"].(string)
}

func (p *PokerStars) ParseSelectionPrice(selectionData map[string]interface{}) float64 {
	return selectionData["price"].(float64)
}

func (p *PokerStars) ParseSelectionLine(selectionData map[string]interface{}, marketData map[string]interface{}) float64 {
	return 0
}
