package parsers

import (
	"flag"
	"github.com/gocolly/colly"
	"github.com/gorilla/websocket"
	"gorm.io/gorm"
	"haggle/models"
	"log"
	"net/url"
	"os"
	"os/signal"
)

type PameStoixima struct {
	Parser
	db     *gorm.DB
	config *models.SiteConfig
	c      *colly.Collector
	ID     int
	Conn   *websocket.Conn
}

var addr = flag.String("addr", "wss://www.pamestoixima.gr/api/864/dl3p24b0/websocket", "http service address")

func (p *PameStoixima) Initialize() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "wss", Host: "www.pamestoixima.gr", Path: "/api/864/dl3p24b0/websocket"}
	log.Printf("connecting to %s", u.String())
	var err error
	p.Conn, _, err = websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer p.Conn.Close()

	done := make(chan struct{})

	//go func() {
	go func() {
		defer close(done)
		for {
			_, message, err := p.Conn.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()
	//send connect
	msg := `["CONNECT\nprotocol-version:1.5\naccept-version:1.1,1.0\nheart-beat:10000,10000\n\n\u0000"]`
	p.Conn.WriteMessage(1, []byte(msg))

}

func (p *PameStoixima) SetConfig(c *models.SiteConfig) {
	p.config = c
	p.ID = c.SiteID
}

func (p *PameStoixima) GetConfig() *models.SiteConfig {
	return p.config
}

func (p *PameStoixima) Scrape() (bool, error) {

	return true, nil
}
func (p *PameStoixima) ScrapeHome() (bool, error) {
	//err := p.c.Visit(fmt.Sprintf("%s/%s", p.config.BaseUrl, p.config.Urls["home"]))
	//if err != nil {
	//	return false, err
	//}

	return true, nil
}

func (p *PameStoixima) ScrapeLive() (bool, error) {
	//err := p.c.Visit(fmt.Sprintf("%s/%s", p.config.BaseUrl, p.config.Urls["live"]))
	//if err != nil {
	//	return false, err
	//}
	return true, nil
}

func (p *PameStoixima) ScrapeToday() (bool, error) {
	//&from= #2021-09-02T21:00:00.000Z&to=2021-09-03T21:00:00.000Z
	//now := time.Now().Format(time.RFC3339)
	//tomorrow := time.Now().Add(time.Hour * 24).Format(time.RFC3339)
	//url := fmt.Sprintf("%s/%s&from=%s&to=%s", p.config.BaseUrl, p.config.Urls["day"], now, tomorrow)
	//err := p.c.Visit(url)
	//if err != nil {
	//	return false, err
	//}
	return true, nil
}

func (p *PameStoixima) ScrapeTournament(tournamentUrl string) (bool, error) {
	// first get tournaments
	msg := `["SUBSCRIBE\nid:/api/eventgroups/soccer-uk-sb_type_10247-all-match-events-grouped-by-type\ndestination:/api/eventgroups/soccer-uk-sb_type_10247-all-match-events-grouped-by-type\nlocale:el\n\n\u0000"]`
	p.Conn.WriteMessage(1, []byte(msg))

	return true, nil
}

func (p *PameStoixima) SetDB(db *gorm.DB) {
	p.db = db
}

func (p *PameStoixima) GetDB() *gorm.DB {
	return p.db
}

func (p *PameStoixima) GetEventID(event map[string]interface{}) int {
	if live, found := event["live"]; found {
		if live != nil {
			if matchId, found := live.(map[string]interface{})["match_id"]; found {
				return int(matchId.(float64))
			}
		}
	}

	return int(event["id"].(float64))
}

func (p *PameStoixima) GetEventName(event map[string]interface{}) string {
	if name, exist := event["label"]; exist {
		return name.(string)
	}
	return ""
}
func (p *PameStoixima) GetEventCanonicalName(event map[string]interface{}) string {
	if name, exist := event["label"]; exist {
		return name.(string)
	}
	return ""
}

func (p *PameStoixima) GetEventMarkets(event map[string]interface{}) []interface{} {
	markets := event["bets"].(map[string]interface{})
	result := make([]interface{}, 0)
	for _, market := range markets {
		result = append(result, market)
	}
	return result
}

func (p *PameStoixima) GetEventDate(event map[string]interface{}) string {
	return event["start"].(string)
}

func (p *PameStoixima) ParseMarketName(market map[string]interface{}) string {
	if name, exist := market["question"]; exist {
		captions := name.(map[string]interface{})
		return captions["label"].(string)
	}
	return ""
}

func (p *PameStoixima) ParseSelectionName(selectionData map[string]interface{}) string {
	if name, exist := selectionData["actor"]; exist {
		captions := name.(map[string]interface{})
		return captions["label"].(string)
	}
	return ""
}

func (p *PameStoixima) ParseSelectionPrice(selectionData map[string]interface{}) float64 {
	return selectionData["odd"].(float64)
}

func (p *PameStoixima) GetEventIsAntepost(event map[string]interface{}) bool {
	return false
}

func (p *PameStoixima) ParseSelectionLine(selectionData map[string]interface{}, marketData map[string]interface{}) float64 {
	line := 0.0

	return line
}

func (p *PameStoixima) ParseMarketType(market map[string]interface{}) string {
	categoryType := market["type"].(map[string]interface{})
	categoryId := categoryType["id"].(float64)
	switch categoryId {
	case 1:
		return "SOCCER_MATCH_RESULT"
	}
	return ""
}

func (p *PameStoixima) MatchMarketType(market map[string]interface{}, marketType string) (models.MarketType, error) {

	return models.MarketType{}, nil
}

func (p *PameStoixima) ParseMarketId(market map[string]interface{}) string {
	return market["id"].(string)
}
func (p *PameStoixima) GetMarketSelections(market map[string]interface{}) []interface{} {
	return market["choices"].([]interface{})
}
