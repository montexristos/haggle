package parsers

import (
	"fmt"
	"github.com/Jeffail/gabs"
	"github.com/gocolly/colly"
	"gorm.io/gorm"
	"haggle/models"
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

func (n *Novibet) GetConfig() *models.SiteConfig {
	return n.config
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

func (n *Novibet) GetDB() *gorm.DB {
	return n.db
}

func (n *Novibet) parseTopEvents(sports []interface{}) {
	for i := 0; i < len(sports); i++ {
		sport := sports[0].(map[string]interface{})
		events := sport["events"].([]interface{})
		if len(events) > 0 {
			for j := 0; j < len(events); j++ {
				ParseEvent(n, events[j].(map[string]interface{}))
			}
		}
	}
}

func (n *Novibet) GetEventID(event map[string]interface{}) int {
	return int(event["betRadarId"].(float64))
}

func (n *Novibet) GetEventName(event map[string]interface{}) string {
	return event["name"].(string)
}

func (n *Novibet) GetEventMarkets(event map[string]interface{}) []interface{} {
	return event["markets"].([]interface{})
}

func (n *Novibet) ParseMarketName(market map[string]interface{}) string {
	return market["name"].(string)
}

func (n *Novibet) ParseSelectionName(selectionData map[string]interface{}) string {
	return selectionData["name"].(string)
}

func (n *Novibet) ParseSelectionPrice(selectionData map[string]interface{}) float64 {
	return selectionData["price"].(float64)
}

func (n *Novibet) GetEventIsAntepost(event map[string]interface{}) bool {
	return false
}

func (n *Novibet) ParseSelectionLine(selectionData map[string]interface{}) float64 {
	line := 0.0
	//TODO get line
	return line
}

func (n *Novibet) ParseMarketType(market map[string]interface{}) string {
	return market["type"].(string)
}
func (n *Novibet) ParseMarketId(market map[string]interface{}) string {
	return market["id"].(string)
}
