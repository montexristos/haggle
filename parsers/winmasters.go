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
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/mxschmitt/playwright-go"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"haggle/models"
	"log"
	"strings"
	"sync"
	"time"
)

type Winmasters struct {
	Parser
	db      *gorm.DB
	config  *models.SiteConfig
	c       *colly.Collector
	ID      int
	pw      *playwright.Playwright
	browser *playwright.Browser
	context *playwright.BrowserContext
	wg      *sync.WaitGroup
}

func (w *Winmasters) Initialize() {

	w.wg = &sync.WaitGroup{}
}

func (w *Winmasters) Destruct() {
	var err error
	context := *w.context
	err = context.Close()
	if err != nil {
		log.Fatalf("could close context: %v", err)
	}
	browser := *w.browser
	if err = browser.Close(); err != nil {
		log.Fatalf("could not close browser: %v", err)
	}
	if err = w.pw.Stop(); err != nil {
		log.Fatalf("could not stop Playwright: %v", err)
	}
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
	return true, nil
}

func (w *Winmasters) ScrapeToday() (bool, error) {
	return true, nil
}

func (w *Winmasters) ScrapeTournament(tournamentId string) (bool, error) {
	url := fmt.Sprintf("%s/%s", w.config.BaseUrl, tournamentId)
	w.getPage(url)
	return true, nil
}

func (w *Winmasters) getPage(url string) {
	var err error
	w.pw, err = playwright.Run()
	if err != nil {
		panic(err.Error())
	}
	browser, err := w.pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(true),
	})
	if err != nil {
		panic(err.Error())
	}
	context, err := browser.NewContext(playwright.BrowserNewContextOptions{
		UserAgent: playwright.String("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36"),
	})
	page, err := context.NewPage()
	if err != nil {
		log.Fatalf("could create page: %v", err)
	}
	page.Once("websocket", func(ws playwright.WebSocket) {
		ws.On("framereceived", func(payload []byte) {
			jsonStr := string(payload)
			fmt.Println(jsonStr)
			if strings.Contains(jsonStr, "AGGREGATOR") && strings.Contains(jsonStr, "INITIAL_DUMP") {
				defer w.wg.Done()
				var result interface{}
				err = json.Unmarshal(payload, &result)
				if err != nil {
					fmt.Errorf(err.Error())
				}
				tournament := result.([]interface{})[len(result.([]interface{}))-1]
				if records, found := tournament.(map[string]interface{})["records"]; found {
					items := records.([]interface{})
					for _, item := range items {
						if item.(map[string]interface{})["_type"] == "MATCH" {
							ParseEvent(w, item.(map[string]interface{}))
						}
					}
				}
			}
		})
	})
	fmt.Println("Visiting page %s", url)
	page.SetDefaultNavigationTimeout(15*1000)
	page.SetDefaultTimeout(15*1000)
	w.wg.Add(1)
	_, err = page.Goto(url, playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateNetworkidle,
	})
	w.wg.Wait()
}

func (w *Winmasters) SetDB(db *gorm.DB) {
	w.db = db
}

func (w *Winmasters) GetDB() *gorm.DB {
	return w.db
}

func (w *Winmasters) GetEventUrl(event map[string]interface{}) string {
	//sports/i/event/1/podosfairo/ellada/super-league-1/ofi-atromitos/{eventid}
	sport := strings.ToLower(event["shortSportName"].(string))
	category := strings.ToLower(event["categoryName"].(string))
	tournament := strings.ToLower(event["shortParentName"].(string))
	eventName := strings.ToLower(event["name"].(string))
	eventId := event["id"].(string)
	eventName = strings.Replace(eventName, " - ", "-", -1)
	sport = strings.Replace(sport, " ", "-", -1)
	category = strings.Replace(category, " ", "-", -1)
	tournament = strings.Replace(tournament, " ", "-", -1)
	eventName = strings.Replace(eventName, " ", "-", -1)
	return fmt.Sprintf("sports/i/event/1/%s/%s/%s/%s/%s", sport, category, tournament, eventName, eventId)
}

func (w *Winmasters) GetEventIsAntepost(event map[string]interface{}) bool {
	return false
}

func (w *Winmasters) GetEventIsLive(event map[string]interface{}) bool {
	return false
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

func (w *Winmasters) GetEventDate(event map[string]interface{}) string {
	start := time.Unix(cast.ToInt64(event["startTime"])/1000, 0)
	return start.Format("2006-01-02 15:04:05")
}


func (w *Winmasters) FetchEvent(e *models.Event) error {
	url := fmt.Sprintf("%s/%s", w.config.BaseUrl, e.Url)
	w.getPage(url)
	return nil
}