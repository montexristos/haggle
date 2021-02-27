package parsers

import (
	"fmt"
	"github.com/gocolly/colly"
	"haggle/models"
	"log"
	"time"

	"github.com/jinzhu/gorm"
)

type Parser interface {
	Initialize()
	GetSiteID() int
	GetDB() *gorm.DB
	Scrape() (bool, error)
	ScrapeHome() (bool, error)
	ScrapeLive() (bool, error)
	ScrapeToday() (bool, error)
	ScrapeTournament(tournamentId string) (bool, error)
	GetEventID(event map[string]interface{}) int
	GetEventName(event map[string]interface{}) string
	GetEventIsAntepost(event map[string]interface{}) bool
	GetEventMarkets(event map[string]interface{})  []interface{}
	ParseMarketType(market map[string]interface{}) string
	ParseMarketId(market map[string]interface{}) string
	ParseMarketName(market map[string]interface{}) string
	ParseSelectionName(selectionData map[string]interface{}) string
	ParseSelectionPrice(selectionData map[string]interface{}) float64
	ParseSelectionLine(selectionData map[string]interface{}) float64
	//ParseMarket(p Parser, market map[string]interface{}, event models.Event) models.Market
	//ParseSelection(eventId int, market models.Market, selection map[string]interface{}) models.Selection
	SetDB(db *gorm.DB)
	SetConfig(config *models.SiteConfig)
}

func ParseEvent(p Parser, event map[string]interface{}) (*models.Event, error) {
	if p.GetEventIsAntepost(event) {
		return &models.Event{}, fmt.Errorf("antepost")
	}
	eventID := p.GetEventID(event)
	eventName := p.GetEventName(event)
	eventMarkets := p.GetEventMarkets(event)
	db := p.GetDB()
	siteID := p.GetSiteID()
	e := models.GetCreateEvent(p.GetDB(), eventID, siteID, eventName)
	markets := make([]models.Market, 0)
	for _, market := range eventMarkets {
		m := ParseMarket(p, market.(map[string]interface{}), e)
		markets = append(markets, m)
	}
	e.Markets = markets
	db.Save(&e)
	return &e, nil
}

func ParseMarket(p Parser, market map[string]interface{}, event models.Event) models.Market {
	marketType := p.ParseMarketType(market)
	marketName := p.ParseMarketName(market)
	marketId := p.ParseMarketId(market)
	m := models.Market{
		Name:     marketName,
		Type:     marketType,
		ID:       marketId,
		SiteID:   p.GetSiteID(),
	}
	selections := ParseMarketSelections(market)
	for _, selection := range selections {
		sel := ParseSelection(p, event.ID, m, selection.(map[string]interface{}))
		m.Selections = append(m.Selections, sel)
	}
	return m
}

func ParseMarketSelections(market map[string]interface{}) []interface{} {
	return market["selections"].([]interface{})
}

func ParseSelection(p Parser, eventId int, market models.Market, selection map[string]interface{}) models.Selection {
	sel := models.Selection{
		ID:    fmt.Sprintf(`%d:%s:%s`, eventId, market.ID, selection["id"].(string)),
		Name:  p.ParseSelectionName(selection),
		Price: p.ParseSelectionPrice(selection),
		SiteID:   p.GetSiteID(),
		MarketID: market.ID,
		Line:     p.ParseSelectionLine(selection),
	}
	return sel
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
