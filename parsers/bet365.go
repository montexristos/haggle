package parsers

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gorilla/websocket"
	"gorm.io/gorm"
	"haggle/models"
	"log"
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"time"
)

type Bet struct {
	Parser
	db     *gorm.DB
	config *models.SiteConfig
	c      *colly.Collector
	ID     int
}

func (p *Bet) Initialize() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/echo"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	//go func() {
	func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

}

func (p *Bet) Scrape() (bool, error) {
	return true, nil
}

func (p *Bet) ScrapeHome() (bool, error) {
	_ = p.c.Visit(fmt.Sprintf("%s", p.config.BaseUrl))
	return true, nil
}

func (p *Bet) ScrapeLive() (bool, error) {
	_ = p.c.Visit(fmt.Sprintf("%s/%s", p.config.BaseUrl, p.config.Urls["live"]))
	return true, nil
}

func (p *Bet) ScrapeToday() (bool, error) {
	_ = p.c.Visit(fmt.Sprintf("%s/%s", p.config.BaseUrl, p.config.Urls["day"]))
	return true, nil
}

func (p *Bet) ScrapeTournament(tournamentUrl string) (bool, error) {
	_ = p.c.Visit(fmt.Sprintf("%s/%s", p.config.BaseUrl, tournamentUrl))
	return true, nil
}

func (p *Bet) SetDB(db *gorm.DB) {
	p.db = db
}

func (p *Bet) GetDB() *gorm.DB {
	return p.db
}

func (p *Bet) SetConfig(c *models.SiteConfig) {
	p.config = c
	p.ID = c.SiteID
}
func (p *Bet) GetConfig() *models.SiteConfig {
	return p.config
}

func (p *Bet) GetEventID(event map[string]interface{}) string {
	return strconv.Itoa(int(event["betRadarId"].(float64)))
}

func (p *Bet) GetEventName(event map[string]interface{}) string {
	return event["name"].(string)
}
func (p *Bet) GetEventCanonicalName(event map[string]interface{}) string {
	return event["name"].(string)
}

func (p *Bet) GetEventMarkets(event map[string]interface{}) []interface{} {
	return event["markets"].([]interface{})
}

func (p *Bet) GetEventDate(event map[string]interface{}) string {
	tm := time.Unix(int64(event["startTime"].(float64)/1000), 0)
	return tm.Format("2006-01-02 15:04:05")
}
func (p *Bet) parseTopEvents(sports []interface{}) {
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

func (p *Bet) ParseMarketName(market map[string]interface{}) string {
	return market["name"].(string)
}

func (p *Bet) ParseSelectionName(selectionData map[string]interface{}) string {
	return selectionData["name"].(string)
}

func (p *Bet) ParseSelectionPrice(selectionData map[string]interface{}) float64 {
	return selectionData["price"].(float64)
}

func (p *Bet) ParseSelectionLine(selectionData map[string]interface{}, marketData map[string]interface{}) float64 {
	line := 0.0
	//TODO get line
	return line
}

func (p *Bet) GetEventIsAntepost(event map[string]interface{}) bool {
	return false
}

func (p *Bet) ParseMarketType(market map[string]interface{}) string {
	return market["type"].(string)
}

func (p *Bet) MatchMarketType(market map[string]interface{}, marketType string) (models.MarketType, error) {
	switch marketType {
	case "MRES":
		return models.NewMatchResult().MarketType, nil
	case "HCTG":
		if market["handicap"] == 2.5 {
			return models.NewOverUnder().MarketType, nil
		}
		return models.MarketType{}, nil
	case "BTSC":
		return models.NewBtts().MarketType, nil
	}
	return models.MarketType{}, fmt.Errorf("could not match market type")
}

func (p *Bet) ParseMarketId(market map[string]interface{}) string {
	return market["id"].(string)
}

func (p *Bet) GetMarketSelections(market map[string]interface{}) []interface{} {
	return market["selections"].([]interface{})
}
