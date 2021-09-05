package parsers

import (
	"fmt"
	"github.com/Jeffail/gabs"
	"github.com/gocolly/colly"
	"gorm.io/gorm"
	"haggle/models"
	"strings"
	"time"
)

type Stoiximan struct {
	Parser
	db     *gorm.DB
	config *models.SiteConfig
	c      *colly.Collector
	ID     int
}

func (s *Stoiximan) Initialize() {
	s.c = GetCollector()
	s.c.OnHTML("script", func(e *colly.HTMLElement) {
		if strings.HasPrefix(e.Text, `window["initial_state"]=`) {
			jsonParsed, err := gabs.ParseJSON([]byte(e.Text[24:]))
			if err != nil {
				fmt.Println(err.Error())
			}
			topEvents := jsonParsed.Path("data.topEvents").Data()
			if topEvents != nil {
				s.parseTopEvents(topEvents.([]interface{}))
			}
			//TODO parse other items (fmt.Println(jsonParsed))
			block := jsonParsed.Path("data.blocks").Data()
			if block != nil {
				s.parseTopEvents(block.([]interface{}))
			}
		}
	})
}

func (s *Stoiximan) Scrape() (bool, error) {

	return true, nil
}

func (s *Stoiximan) ScrapeHome() (bool, error) {
	_ = s.c.Visit(fmt.Sprintf("%s", s.config.BaseUrl))
	return true, nil
}

func (s *Stoiximan) ScrapeLive() (bool, error) {
	_ = s.c.Visit(fmt.Sprintf("%s/%s", s.config.BaseUrl, s.config.Urls["live"]))
	return true, nil
}

func (s *Stoiximan) ScrapeToday() (bool, error) {
	_ = s.c.Visit(fmt.Sprintf("%s/%s", s.config.BaseUrl, s.config.Urls["day"]))
	return true, nil
}

func (s *Stoiximan) ScrapeTournament(tournamentUrl string) (bool, error) {
	_ = s.c.Visit(fmt.Sprintf("%s/%s", s.config.BaseUrl, tournamentUrl))
	return true, nil
}

func (s *Stoiximan) SetDB(db *gorm.DB) {
	s.db = db
}

func (s *Stoiximan) GetDB() *gorm.DB {
	return s.db
}

func (s *Stoiximan) SetConfig(c *models.SiteConfig) {
	s.config = c
	s.ID = c.SiteID
}
func (s *Stoiximan) GetConfig() *models.SiteConfig {
	return s.config
}

func (s *Stoiximan) GetEventID(event map[string]interface{}) int {
	return int(event["betRadarId"].(float64))
}

func (s *Stoiximan) GetEventName(event map[string]interface{}) string {
	return event["name"].(string)
}

func (s *Stoiximan) GetEventMarkets(event map[string]interface{}) []interface{} {
	return event["markets"].([]interface{})
}

func (s *Stoiximan) GetEventDate(event map[string]interface{}) string {
	tm := time.Unix(int64(event["startTime"].(float64)/1000), 0)
	return tm.Format("2006-01-02 15:04:05")
}
func (s *Stoiximan) parseTopEvents(sports []interface{}) {
	for i := 0; i < len(sports); i++ {
		sport := sports[0].(map[string]interface{})
		events := sport["events"].([]interface{})
		if len(events) > 0 {
			for j := 0; j < len(events); j++ {
				_, _ = ParseEvent(s, events[j].(map[string]interface{}))
			}
		}
	}
}

func (s *Stoiximan) ParseMarketName(market map[string]interface{}) string {
	return market["name"].(string)
}

func (s *Stoiximan) ParseSelectionName(selectionData map[string]interface{}) string {
	return selectionData["name"].(string)
}

func (s *Stoiximan) ParseSelectionPrice(selectionData map[string]interface{}) float64 {
	return selectionData["price"].(float64)
}

func (s *Stoiximan) ParseSelectionLine(selectionData map[string]interface{}, marketData map[string]interface{}) float64 {
	line := 0.0
	//TODO get line
	return line
}

func (s *Stoiximan) GetEventIsAntepost(event map[string]interface{}) bool {
	return false
}

func (s *Stoiximan) ParseMarketType(market map[string]interface{}) string {
	return market["type"].(string)
}

func (s *Stoiximan) MatchMarketType(market map[string]interface{}, marketType string) models.MarketType {
	switch marketType {
	case "MRES":
		return models.NewMatchResult().MarketType
	case "HCTG":
		if market["handicap"] == 2.5 {
			return models.NewOverUnder().MarketType
		}
		return models.MarketType{}
	case "BTSC":
		return models.NewBtts().MarketType
	}
	return models.MarketType{}
}

func (s *Stoiximan) ParseMarketId(market map[string]interface{}) string {
	return market["id"].(string)
}

func (s *Stoiximan) GetMarketSelections(market map[string]interface{}) []interface{} {
	return market["selections"].([]interface{})
}
