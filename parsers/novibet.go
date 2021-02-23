package parsers

import (
	"fmt"
	"github.com/Jeffail/gabs"
	"github.com/gocolly/colly"
	"github.com/jinzhu/gorm"
	"github.com/spf13/cast"
	"haggle/models"
	"strconv"
	"strings"
)

type Novibet struct {
	Parser
	db     *gorm.DB
	config *models.SiteConfig
	c      *colly.Collector
	ID     int
}

func (n *Novibet) Initialize() {
	n.c = GetCollector()
	n.c.OnHTML("body", func(e *colly.HTMLElement) {
		text := e.Text
		fmt.Println(text)
		if strings.HasPrefix(e.Text, `window["initial_state"]=`) {
			jsonParsed, err := gabs.ParseJSON([]byte(e.Text[24:]))
			if err != nil {
				fmt.Println(err.Error())
			}
			topEvents := jsonParsed.Path("data.topEvents").Data()
			if topEvents != nil {
				n.parseTopEvents(topEvents.([]interface{}))
			}
			//TODO parse other items (fmt.Println(jsonParsed))
		}
	})
}

func (n *Novibet) SetConfig(c *models.SiteConfig) {
	n.config = c
	n.ID = c.SiteID
}

func (n *Novibet) Scrape() (bool, error) {

	return true, nil
}
func (n *Novibet) ScrapeHome() (bool, error) {
	return true, nil
}

func (n *Novibet) ScrapeLive() (bool, error) {
	_ = n.c.Visit(fmt.Sprintf("%s/%s", n.config.BaseUrl, n.config.Urls["live"]))
	return true, nil
}

func (n *Novibet) ScrapeToday() (bool, error) {
	_ = n.c.Visit(fmt.Sprintf("%s/%s", n.config.BaseUrl, n.config.Urls["day"]))
	return true, nil
}

func (n *Novibet) ScrapeTournament(tournamentId string) (bool, error) {
	return true, nil
}

func (n *Novibet) SetDB(db *gorm.DB) {
	n.db = db
}

func (n *Novibet) parseTopEvents(sports []interface{}) {
	for i := 0; i < len(sports); i++ {
		sport := sports[0].(map[string]interface{})
		events := sport["events"].([]interface{})
		if len(events) > 0 {
			for j := 0; j < len(events); j++ {
				n.ParseEvent(events[j].(map[string]interface{}))
			}
		}
	}
}

func (n *Novibet) ParseEvent(event map[string]interface{}) {
	eventID := cast.ToInt(event["betRadarId"])
	name := event["name"].(string)
	e := models.GetCreateEvent(n.db, eventID, n.ID, name)
	markets := make([]models.Market, 0)
	for _, market := range event["markets"].([]interface{}) {
		m := n.parseMarket(market.(map[string]interface{}), e)
		markets = append(markets, m)
	}
	e.Markets = markets
	n.db.Save(&e)
}

func (n *Novibet) parseMarket(market map[string]interface{}, event models.Event) models.Market {
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
		sel := n.parseSelection(event.ID, m, selection.(map[string]interface{}))
		m.Selections = append(m.Selections, sel)
	}
	return m
}

func (n *Novibet) parseSelection(eventId int, market models.Market, selection map[string]interface{}) models.Selection {
	sel := models.Selection{
		ID:    fmt.Sprintf(`%d:%s:%s`, eventId, market.ID, selection["id"].(string)),
		Name:  selection["name"].(string),
		Price: selection["price"].(float64),
	}
	return sel
}
