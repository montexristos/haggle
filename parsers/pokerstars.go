package parsers

import (
	"fmt"
	"haggle/models"
	"strconv"
	"strings"

	"github.com/Jeffail/gabs"
	"github.com/gocolly/colly"
	"github.com/jinzhu/gorm"
)

type PokerStars struct {
	Parser
	db *gorm.DB
}

func (p PokerStars) Scrape(config models.SiteConfig, c *colly.Collector, db *gorm.DB) (bool, error) {
	p.db = db
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Content-Type", "application/json;charset=UTF-8")
	})
	// err := c.PostRaw(URL, payload)
	c.OnHTML("script", func(e *colly.HTMLElement) {
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

	_ = c.Visit(fmt.Sprintf("%s", config.BaseUrl))
	_ = c.Visit(fmt.Sprintf("%s/%s", config.BaseUrl, config.Urls["live"]))
	_ = c.Visit(fmt.Sprintf("%s/%s", config.BaseUrl, config.Urls["day"]))
	return true, nil
}

func (p PokerStars) parseTopEvents(sports []interface{}) {
	for i := 0; i < len(sports); i++ {
		sport := sports[0].(map[string]interface{})
		events := sport["events"].([]interface{})
		if len(events) > 0 {
			for j := 0; j < len(events); j++ {
				p.parseEvent(events[j].(map[string]interface{}))
			}
		}
	}
}

func (p PokerStars) parseEvent(event map[string]interface{}) {
	eventID := int(event["betRadarId"].(float64))
	var e models.Event
	p.db.First(&e, eventID)
	if e.ID == 0 {
		e = models.Event{
			Name: event["name"].(string),
			ID:   eventID,
		}
		p.db.Create(&e)
	}

	markets := make([]models.Market, 0)
	for _, market := range event["markets"].([]interface{}) {
		m := p.parseMarket(market.(map[string]interface{}), e)
		markets = append(markets, m)
	}
	e.Markets = markets
	p.db.Save(&e)
}

func (p PokerStars) parseMarket(market map[string]interface{}, event models.Event) models.Market {
	var marketId string
	if market["handicap"].(float64) > 0 {
		handicap := strconv.FormatFloat(market["handicap"].(float64), 'f', 2, 64)
		marketId = fmt.Sprintf(`%s:%s`, market["type"].(string), handicap)
	} else {
		marketId = fmt.Sprintf(`%s`, market["type"].(string))
	}

	m := models.Market{
		Name:     market["name"].(string),
		MarketId: market["id"].(string),
		Type:     market["type"].(string),
		ID:       fmt.Sprintf(`%d:%s`, event.ID, marketId),
	}
	selections := market["selections"].([]interface{})
	for _, selection := range selections {
		s := p.parseSelection(event.ID, m, selection.(map[string]interface{}))
		m.Selections = append(m.Selections, s)
	}
	return m
}

func (p PokerStars) parseSelection(eventId int, market models.Market, selection map[string]interface{}) models.Selection {
	s := models.Selection{
		ID:    fmt.Sprintf(`%d:%s:%s`, eventId, market.ID, selection["id"].(string)),
		Name:  selection["name"].(string),
		Price: selection["price"].(float64),
	}
	return s
}
