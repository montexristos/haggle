package parsers

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"gorm.io/gorm"
	"haggle/models"
	"log"
	"strconv"
	"strings"
	"time"
)

type Bwin struct {
	Parser
	db     *gorm.DB
	config *models.SiteConfig
	c      *colly.Collector
	ID     int
}

/**
 first visit https://www.Bwin.gr/api/marketviews/findtimestamps?lang=el-GR&oddsR=1&timeZ=GTB%20Standard%20Time&usrGrp=G
then make request for each content key
*/
func (b *Bwin) Initialize() {

	b.c = GetCollector()
	b.c.OnResponse(func(response *colly.Response) {
		var resp interface{}
		err := json.Unmarshal(response.Body, &resp)
		if err != nil {
			panic(err.Error())
		}

		json := resp.(map[string]interface{})
		if fixtures, found := json["fixtures"]; found {
			for _, event := range fixtures.([]interface{}) {
				ParseEvent(b, event.(map[string]interface{}))
			}
		}
		log.Println(json)
	})

}

func (b *Bwin) SetConfig(c *models.SiteConfig) {
	b.config = c
	b.ID = c.SiteID
}

func (b *Bwin) GetConfig() *models.SiteConfig {
	return b.config
}

func (b *Bwin) Scrape() (bool, error) {

	return true, nil
}
func (b *Bwin) ScrapeHome() (bool, error) {
	err := b.c.Visit(fmt.Sprintf("%s/%s", b.config.BaseUrl, b.config.Urls["home"]))
	if err != nil {
		return false, err
	}

	return true, nil
}

func (b *Bwin) ScrapeLive() (bool, error) {
	err := b.c.Visit(fmt.Sprintf("%s/%s", b.config.BaseUrl, b.config.Urls["live"]))
	if err != nil {
		return false, err
	}
	return true, nil
}

func (b *Bwin) ScrapeToday() (bool, error) {
	//&from= #2021-09-02T21:00:00.000Z&to=2021-09-03T21:00:00.000Z
	now := time.Now().Format(time.RFC3339)
	tomorrow := time.Now().Add(time.Hour * 24).Format(time.RFC3339)
	url := fmt.Sprintf("%s/%s&from=%s&to=%s", b.config.BaseUrl, b.config.Urls["day"], now, tomorrow)
	err := b.c.Visit(url)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (b *Bwin) ScrapeTournament(tournamentUrl string) (bool, error) {
	// first get tournaments
	tourUrl := fmt.Sprintf("%s/%s", b.config.BaseUrl, tournamentUrl)
	err := b.c.Visit(tourUrl)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (b *Bwin) SetDB(db *gorm.DB) {
	b.db = db
}

func (b *Bwin) GetDB() *gorm.DB {
	return b.db
}

func (b *Bwin) GetEventID(event map[string]interface{}) int {
	if addons, found := event["addons"]; found && addons != nil {
		if betradarid, exists := addons.(map[string]interface{})["betRadar"]; exists && betradarid != nil {
			return int(betradarid.(float64))
		}
	}
	return -1
}

func (b *Bwin) GetEventName(event map[string]interface{}) string {
	if name, exist := event["name"]; exist {
		captions := name.(map[string]interface{})
		return captions["value"].(string)
	}
	return ""
}

func (b *Bwin) GetEventMarkets(event map[string]interface{}) []interface{} {
	return event["games"].([]interface{})
}

func (b *Bwin) GetEventDate(event map[string]interface{}) string {
	return ""
}

func (b *Bwin) ParseMarketName(market map[string]interface{}) string {
	if name, exist := market["name"]; exist {
		captions := name.(map[string]interface{})
		return captions["value"].(string)
	}
	return ""
}

func (b *Bwin) ParseSelectionName(selectionData map[string]interface{}) string {
	if name, exist := selectionData["sourceName"]; exist {
		captions := name.(map[string]interface{})
		return captions["value"].(string)
	}
	if name, exist := selectionData["name"]; exist {
		captions := name.(map[string]interface{})
		return captions["value"].(string)
	}
	return ""
}

func (b *Bwin) ParseSelectionPrice(selectionData map[string]interface{}) float64 {
	return selectionData["odds"].(float64)
}

func (b *Bwin) GetEventIsAntepost(event map[string]interface{}) bool {
	return false
}

func (b *Bwin) ParseSelectionLine(selectionData map[string]interface{}, marketData map[string]interface{}) float64 {
	line := 0.0
	//TODO get line
	if marketLine, found := marketData["attr"]; found {
		hc := marketLine.(string)
		hc = normalizeFloat(hc)
		hcfloat, _ := strconv.ParseFloat(hc, 64)
		return hcfloat
	}
	return line
}

func (b *Bwin) ParseMarketType(market map[string]interface{}) string {
	categoryId := market["categoryId"].(float64)
	switch categoryId {
	case 25:
		return "SOCCER_MATCH_RESULT"
	case 31:
		return "SOCCER_UNDER_OVER"
	case 261:
		return "SOCCER_BOTH_TEAMS_TO_SCORE"
	}
	return ""
}

func normalizeFloat(old string) string {
	s := strings.Replace(old, ".", "", 1)
	s = strings.Replace(s, ",", ".", -1)
	return s
}

func (b *Bwin) MatchMarketType(market map[string]interface{}, marketType string) (models.MarketType, error) {
	switch marketType {
	case "SOCCER_MATCH_RESULT":
		return models.NewMatchResult().MarketType, nil
	case "SOCCER_UNDER_OVER":
		hc := market["attr"].(string)
		hc = normalizeFloat(hc)
		hcfloat, _ := strconv.ParseFloat(hc, 64)
		if hcfloat == 2.5 {
			return models.NewOverUnder().MarketType, nil
		}
		return models.MarketType{}, nil
	case "SOCCER_BOTH_TEAMS_TO_SCORE":
		return models.NewBtts().MarketType, nil
	}
	return models.MarketType{}, fmt.Errorf("could not match market type")
}

func (b *Bwin) ParseMarketId(market map[string]interface{}) string {
	return market["id"].(string)
}
func (b *Bwin) GetMarketSelections(market map[string]interface{}) []interface{} {
	return market["results"].([]interface{})
}
