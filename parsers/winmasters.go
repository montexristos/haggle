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
	"github.com/Jeffail/gabs"
	"github.com/gocolly/colly"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"haggle/models"
	"strings"
)

type Winmasters struct {
	Parser
	db     *gorm.DB
	config *models.SiteConfig
	c      *colly.Collector
	ID     int
}

func (w Winmasters) Initialize() {
	w.c = GetCollector()
	w.c.OnHTML("body", func(e *colly.HTMLElement) {
		text := e.Text
		fmt.Println(text)
		if strings.HasPrefix(e.Text, `window["initial_state"]=`) {
			jsonParsed, err := gabs.ParseJSON([]byte(e.Text[24:]))
			if err != nil {
				fmt.Println(err.Error())
			}
			topEvents := jsonParsed.Path("data.topEvents").Data()
			if topEvents != nil {
				w.parseTopEvents(topEvents.([]interface{}))
			}
			//TODO parse other items (fmt.Println(jsonParsed))
		}
	})
}

func (w *Winmasters) SetConfig(c *models.SiteConfig) {
	w.config = c
	w.ID = c.SiteID
}
func (w *Winmasters) GetConfig() *models.SiteConfig {
	return w.config
}

func (w *Winmasters) Scrape() (bool, error) {

	return true, nil
}
func (w *Winmasters) ScrapeHome() (bool, error) {
	return true, nil
}

func (w *Winmasters) ScrapeLive() (bool, error) {
	_ = w.c.Visit(fmt.Sprintf("%s/%s", w.config.BaseUrl, w.config.Urls["live"]))
	return true, nil
}

func (w *Winmasters) ScrapeToday() (bool, error) {
	_ = w.c.Visit(fmt.Sprintf("%s/%s", w.config.BaseUrl, w.config.Urls["day"]))
	return true, nil
}

func (w *Winmasters) ScrapeTournament(tournamentId string) (bool, error) {
	return true, nil
}

func (w *Winmasters) SetDB(db *gorm.DB) {
	w.db = db
}

func (w *Winmasters) parseTopEvents(sports []interface{}) {
	for i := 0; i < len(sports); i++ {
		sport := sports[0].(map[string]interface{})
		events := sport["events"].([]interface{})
		if len(events) > 0 {
			for j := 0; j < len(events); j++ {
				_, _ = ParseEvent(w, events[j].(map[string]interface{}))
			}
		}
	}
}

func (w *Winmasters) ParseEvents(events []interface{}, markets []interface{}) {
	for i := 0; i < len(events); i++ {
		event := events[i].(map[string]interface{})
		eventObject, err := ParseEvent(w, event)
		if err != nil {
			fmt.Println("error parsing event")
			continue
		}
		for j := 0; j < len(markets); j++ {
			if markets[j].(map[string]interface{})["eventId"] == event["id"] {
				market, parseError := ParseMarket(w, markets[j].(map[string]interface{}), *eventObject)
				if parseError == nil {
					eventObject.Markets = append(eventObject.Markets, market)
				}

			}
		}
		w.db.Save(&eventObject)
	}
}

func (w *Winmasters) GetDB() *gorm.DB {
	return w.db
}

func (w *Winmasters) GetEventIsAntepost(event map[string]interface{}) bool {
	return event["type"] == "Outright"
}

func (w *Winmasters) GetEventID(event map[string]interface{}) string {
	return cast.ToString(event["id"])
}

func (w *Winmasters) GetEventName(event map[string]interface{}) string {
	return cast.ToString(event["name"])
}

func (w *Winmasters) GetEventMarkets(event map[string]interface{}) []interface{} {
	return make([]interface{}, 0)
}

func (w *Winmasters) ParseMarketId(market map[string]interface{}) string {
	return market["id"].(string)
}

func (w *Winmasters) ParseMarketType(market map[string]interface{}) string {
	return cast.ToString(market["marketType"].(map[string]interface{})["id"])
}

func (w *Winmasters) ParseMarketName(market map[string]interface{}) string {
	return cast.ToString(market["marketType"].(map[string]interface{})["name"])
}

func (w *Winmasters) ParseSelectionName(selectionData map[string]interface{}) string {
	return selectionData["name"].(string)
}

func (w *Winmasters) ParseSelectionPrice(selectionData map[string]interface{}) float64 {
	return selectionData["trueOdds"].(float64)
}

func (w *Winmasters) ParseSelectionLine(selectionData map[string]interface{}, marketData map[string]interface{}) float64 {
	line := 0.0
	if selectionData["points"] != nil {
		line = selectionData["points"].(float64)
	}
	return line
}
