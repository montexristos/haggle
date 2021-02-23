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
	"github.com/jinzhu/gorm"
	"github.com/spf13/cast"
	"haggle/models"
	"strconv"
	"strings"
)

type Winmasters struct {
	Parser
	db     *gorm.DB
	config *models.SiteConfig
	c      *colly.Collector
	ID     int
}

func (w *Winmasters) Initialize() {
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
				w.ParseEvent(events[j].(map[string]interface{}))
			}
		}
	}
}

func (w *Winmasters) ParseEvents(events []interface{}, markets []interface{}) {
	for i := 0; i < len(events); i++ {
		event := events[i].(map[string]interface{})
		eventObject, err := w.ParseEvent(event)
		if err != nil {
			fmt.Println("error parsing event")
			continue
		}
		for j := 0; j < len(markets); j++ {
			if markets[j].(map[string]interface{})["eventId"] == event["id"] {
				market := w.ParseMarket(markets[j].(map[string]interface{}))
				eventObject.Markets = append(eventObject.Markets, market)
				w.db.Save(&eventObject)
			}
		}
	}
}

func (w *Winmasters) ParseEvent(event map[string]interface{}) (*models.Event, error) {
	if event["type"] == "Outright" {
		return &models.Event{}, fmt.Errorf("antepost")
	}
	eventID := cast.ToInt(event["id"])
	name := cast.ToString(event["name"])
	e := models.GetCreateEvent(w.db, eventID, w.ID, name)
	return &e, nil
}

func (w *Winmasters) ParseMarket(market map[string]interface{}) models.Market {
	marketType := market["marketType"].(map[string]interface{})["id"].(string)
	name := market["marketType"].(map[string]interface{})["name"].(string)
	eventID, _ := strconv.Atoi(market["eventId"].(string))
	marketId := market["id"].(string)
	m := models.Market{
		Name:     name,
		MarketId: marketId,
		Type:     marketType,
		ID:       fmt.Sprintf(`%d:%s:%d`, eventID, marketId, w.ID),
	}
	selections := market["selections"].([]interface{})
	for _, selection := range selections {
		sel := w.parseSelection(eventID, m, selection.(map[string]interface{}))
		m.Selections = append(m.Selections, sel)
	}
	return m
}

func (w *Winmasters) parseSelection(eventId int, market models.Market, selection map[string]interface{}) models.Selection {
	line := 0.0
	if selection["points"] != nil {
		line = selection["points"].(float64)
	}
	sel := models.Selection{
		ID:       fmt.Sprintf(`%d:%s:%s`, eventId, market.ID, selection["id"].(string)),
		MarketID: market.ID,
		Name:     selection["name"].(string),
		Price:    selection["trueOdds"].(float64),
		Line:     line,
	}
	return sel
}
