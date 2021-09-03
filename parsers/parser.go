package parsers

import (
	"fmt"
	"github.com/gocolly/colly"
	"gorm.io/gorm"
	"haggle/models"
	"log"
	"time"
)

type Parser interface {
	Initialize()
	//GetSiteID() int
	GetDB() *gorm.DB
	Scrape() (bool, error)
	ScrapeHome() (bool, error)
	ScrapeLive() (bool, error)
	ScrapeToday() (bool, error)
	ScrapeTournament(tournamentUrl string) (bool, error)
	GetEventID(event map[string]interface{}) int
	GetEventName(event map[string]interface{}) string
	GetEventDate(event map[string]interface{}) string
	GetEventIsAntepost(event map[string]interface{}) bool
	GetEventMarkets(event map[string]interface{}) []interface{}
	ParseMarketType(market map[string]interface{}) string
	MatchMarketType(market map[string]interface{}, marketType string) models.MarketType
	ParseMarketId(market map[string]interface{}) string
	ParseMarketName(market map[string]interface{}) string
	ParseSelectionName(selectionData map[string]interface{}) string
	ParseSelectionPrice(selectionData map[string]interface{}) float64
	ParseSelectionLine(selectionData map[string]interface{}) float64
	GetMarketSelections(marketData map[string]interface{}) []interface{}
	SetDB(db *gorm.DB)
	SetConfig(config *models.SiteConfig)
	GetConfig() *models.SiteConfig
}

func GetSiteID(p Parser) int {
	return p.GetConfig().SiteID
}

func ParseEvent(p Parser, event map[string]interface{}) (*models.Event, error) {
	if p.GetEventIsAntepost(event) {
		return &models.Event{}, fmt.Errorf("antepost")
	}
	eventID := p.GetEventID(event)
	if eventID < 0 {
		return nil, fmt.Errorf("wrong event id")
	}
	eventName := p.GetEventName(event)
	eventMarkets := p.GetEventMarkets(event)
	date := p.GetEventDate(event)
	db := p.GetDB()
	siteID := GetSiteID(p)
	e := models.GetCreateEvent(p.GetDB(), eventID, siteID, eventName)
	e.Date = date
	if len(e.Markets) == 0 {
		markets := make([]models.Market, 0)
		for _, market := range eventMarkets {
			m := ParseMarket(p, market.(map[string]interface{}), e)
			markets = append(markets, m)
		}
		e.Markets = markets
	} else {
		for _, market := range e.Markets {
			UpdateMarket(p, market, eventMarkets)
		}
	}
	db.Debug().Session(&gorm.Session{FullSaveAssociations: true}).Save(&e)
	return &e, nil
}

func UpdateMarket(p Parser, market models.Market, markets []interface{}) {
	for _, m := range markets {
		if p.ParseMarketType(m.(map[string]interface{})) == market.Type {
			selections := ParseMarketSelections(p, m.(map[string]interface{}))
			UpdateSelections(p, market, selections)
		}
	}
}
func UpdateSelections(p Parser, market models.Market, selections []interface{}) {
	for _, sel := range selections {
		fmt.Println(sel)
	}
}

func ParseMarket(p Parser, market map[string]interface{}, event models.Event) models.Market {
	marketType := p.ParseMarketType(market)
	marketTypeId := p.MatchMarketType(market, marketType)
	marketName := p.ParseMarketName(market)
	//marketId := p.ParseMarketId(market)
	m := models.Market{
		Name:       marketName,
		Type:       marketType,
		MarketType: marketTypeId.Name,
	}
	selections := ParseMarketSelections(p, market)
	for _, selection := range selections {
		sel := ParseSelection(p, m, selection.(map[string]interface{}))
		m.Selections = append(m.Selections, sel)
	}
	p.GetDB().Debug().Session(&gorm.Session{FullSaveAssociations: true}).Updates(&m)
	return m
}

func ParseMarketSelections(p Parser, market map[string]interface{}) []interface{} {
	selections := p.GetMarketSelections(market)
	return selections
}

func ParseSelection(p Parser, market models.Market, selection map[string]interface{}) models.Selection {
	sel := models.Selection{
		Name:  p.ParseSelectionName(selection),
		Price: p.ParseSelectionPrice(selection),
		Line:  p.ParseSelectionLine(selection),
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
