package parsers

//
//curl 'https://sbapi.sbtech.com/winmasters/sportscontent/sportsbook/v1/Events/GetByLeagueId' \
//-H 'authority: sbapi.sbtech.com' \
//-H 'block-id: EventsWrapper_Center_LeagueViewResponsiveBlock_22345' \
//-H 'authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJTaXRlSWQiOjExNywiU2Vzc2lvbklkIjoiNjQyODI0OWYtY2NhNS00ZWMzLWJmYWMtNjJjMTZkYWFkNjBlIiwibmJmIjoxNjA3MDY1MzQ5LCJleHAiOjE2MDc2NzAxNzksImlhdCI6MTYwNzA2NTM3OX0.Hi6r2ZAeSfL_0LBxn9BPO4DUzsucjjMc2mCDf9TUdYE' \
//-H 'locale: gr' \
//-H 'user-agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 11_0_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36' \
//-H 'content-type: application/json-patch+json' \
//-H 'accept: */*' \
//-H 'origin: https://www.winmasters.gr' \
//-H 'sec-fetch-site: cross-site' \
//-H 'sec-fetch-mode: cors' \
//-H 'sec-fetch-dest: empty' \
//-H 'referer: https://www.winmasters.gr/' \
//-H 'accept-language: en-US,en;q=0.9,el;q=0.8,it;q=0.7,fr;q=0.6,es;q=0.5' \
//--data-binary '{"eventState":"Mixed","eventTypes":["Fixture","AggregateFixture"],"ids":["40253"],"regionIds":["86"],"marketTypeRequests":[{"sportIds":["1"],"marketTypeIds":["1_39","2_39","3_39","1_169","1707","1_0","2_0","3_0"],"statement":"Include"}]}' \
//--compressed

import (
	"fmt"
	"strings"

	"github.com/Jeffail/gabs"
	"github.com/gocolly/colly"
	"github.com/montexristos/haggle/models"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

type Interwetten struct {
	Parser
	db     *gorm.DB
	config *models.SiteConfig
	c      *colly.Collector
	ID     int
}

func (p Interwetten) Initialize() {
	p.c = GetCollector()
	p.c.OnHTML("body", func(e *colly.HTMLElement) {
		text := e.Text
		fmt.Println(text)
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
}

func (p *Interwetten) SetConfig(c *models.SiteConfig) {
	p.config = c
	p.ID = c.SiteID
}
func (p *Interwetten) GetConfig() *models.SiteConfig {
	return p.config
}

func (p *Interwetten) Scrape() (bool, error) {

	return true, nil
}
func (p *Interwetten) ScrapeHome() (bool, error) {
	return true, nil
}

func (p *Interwetten) ScrapeLive() (bool, error) {
	_ = p.c.Visit(fmt.Sprintf("%s/%s", p.config.BaseUrl, p.config.Urls["live"]))
	return true, nil
}

func (p *Interwetten) ScrapeToday() (bool, error) {
	_ = p.c.Visit(fmt.Sprintf("%s/%s", p.config.BaseUrl, p.config.Urls["day"]))
	return true, nil
}

func (p *Interwetten) ScrapeTournament(tournamentId string) (bool, error) {
	return true, nil
}

func (p *Interwetten) SetDB(db *gorm.DB) {
	p.db = db
}

func (p *Interwetten) parseTopEvents(sports []interface{}) {
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

func (p *Interwetten) ParseEvents(events []interface{}, markets []interface{}) {
	for i := 0; i < len(events); i++ {
		event := events[i].(map[string]interface{})
		eventObject, err := ParseEvent(p, event)
		if err != nil {
			fmt.Println("error parsing event")
			continue
		}
		for j := 0; j < len(markets); j++ {
			if markets[j].(map[string]interface{})["eventId"] == event["id"] {
				market, parseError := ParseMarket(p, markets[j].(map[string]interface{}), *eventObject)
				if parseError == nil {
					eventObject.Markets = append(eventObject.Markets, market)
				}
				eventObject.Markets = append(eventObject.Markets, market)
				p.db.Save(&eventObject)
			}
		}
	}
}

func (p *Interwetten) GetDB() *gorm.DB {
	return p.db
}

func (p *Interwetten) GetEventIsAntepost(event map[string]interface{}) bool {
	return event["type"] == "Outright"
}

func (p *Interwetten) GetEventIsLive(event map[string]interface{}) bool {
	return event["live"] == "true"
}

func (p *Interwetten) GetEventID(event map[string]interface{}) string {
	return cast.ToString(event["id"])
}

func (p *Interwetten) GetEventName(event map[string]interface{}) string {
	return cast.ToString(event["name"])
}

func (p *Interwetten) GetEventMarkets(event map[string]interface{}) []interface{} {
	return make([]interface{}, 0)
}

func (p *Interwetten) ParseMarketId(market map[string]interface{}) string {
	return market["id"].(string)
}

func (p *Interwetten) ParseMarketType(market map[string]interface{}) string {
	return cast.ToString(market["marketType"].(map[string]interface{})["id"])
}

func (p *Interwetten) ParseMarketName(market map[string]interface{}) string {
	return cast.ToString(market["marketType"].(map[string]interface{})["name"])
}

func (p *Interwetten) ParseSelectionName(selectionData map[string]interface{}) string {
	return selectionData["name"].(string)
}

func (p *Interwetten) ParseSelectionPrice(selectionData map[string]interface{}) float64 {
	return selectionData["trueOdds"].(float64)
}

func (p *Interwetten) ParseSelectionLine(selectionData map[string]interface{}, marketData map[string]interface{}) float64 {
	line := 0.0
	if selectionData["points"] != nil {
		line = selectionData["points"].(float64)
	}
	return line
}
