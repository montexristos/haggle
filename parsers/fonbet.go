package parsers

import (
	"bytes"
	"fmt"
	"github.com/Jeffail/gabs"
	"github.com/gocolly/colly"
	"gorm.io/gorm"
	"haggle/models"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Fonbet struct {
	Parser
	db     *gorm.DB
	config *models.SiteConfig
	c      *colly.Collector
	ID     int
}

/**
 first visit https://www.Netbet.gr/api/marketviews/findtimestamps?lang=el-GR&oddsR=1&timeZ=GTB%20Standard%20Time&usrGrp=G
then make request for each content key
*/
func (p *Fonbet) Initialize() {

	p.c = GetCollector()
	p.c.OnResponse(func(response *colly.Response) {
		jsonParsed, err := gabs.ParseJSON(response.Body)
		if err != nil {
			fmt.Println(err.Error())
		}

		events, err := jsonParsed.Search("components", "data", "events").Children()
		if err != nil {
			print(err.Error())
		}
		if events != nil {
			for _, eventList := range events {
				evt := eventList.Data()
				for _, event := range evt.([]interface{}) {
					ParseEvent(p, event.(map[string]interface{}))
				}
			}
		}
	})
}

func (p *Fonbet) SetConfig(c *models.SiteConfig) {
	p.config = c
	p.ID = c.SiteID
}

func (p *Fonbet) GetConfig() *models.SiteConfig {
	return p.config
}

func (p *Fonbet) Scrape() (bool, error) {

	return true, nil
}
func (p *Fonbet) ScrapeHome() (bool, error) {
	//err := p.c.Visit(fmt.Sprintf("%s/%s", p.config.BaseUrl, p.config.Urls["home"]))
	//if err != nil {
	//	return false, err
	//}

	return true, nil
}

func (p *Fonbet) ScrapeLive() (bool, error) {
	//err := p.c.Visit(fmt.Sprintf("%s/%s", p.config.BaseUrl, p.config.Urls["live"]))
	//if err != nil {
	//	return false, err
	//}
	return true, nil
}

func (p *Fonbet) ScrapeToday() (bool, error) {
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

func (p *Fonbet) ScrapeTournament(tournamentUrl string) (bool, error) {
	// first get tournaments
	tourUrl := fmt.Sprintf("%s", p.config.BaseUrl)
	var formData = `{
    "context":
    {
        "url_key": "%s",
        "version": "1.0.1",
        "device": "web_vuejs_desktop",
        "lang": "en",
        "timezone": "-1",
        "url_params":
        {}
    },
    "components":
    [
        {
            "tree_compo_key": "prematch_event_list",
            "params":
            {}
        }
    ]
}`
	formData = fmt.Sprintf(formData, tournamentUrl)
	err := p.c.PostRaw(tourUrl, []byte(formData))
	if err != nil {
		return false, err
	}
	return true, nil
}

func (p *Fonbet) SetDB(db *gorm.DB) {
	p.db = db
}

func (p *Fonbet) GetDB() *gorm.DB {
	return p.db
}

func (p *Fonbet) GetEventID(event map[string]interface{}) string {
	if live, found := event["live"]; found {
		if live != nil {
			if matchId, found := live.(map[string]interface{})["match_id"]; found {
				return strconv.Itoa(int(matchId.(float64)))
			}
		}
	}

	return strconv.Itoa(int(event["id"].(float64)))
}

func (p *Fonbet) GetEventName(event map[string]interface{}) string {
	if name, exist := event["label"]; exist {
		return name.(string)
	}
	return ""
}
func (p *Fonbet) GetEventCanonicalName(event map[string]interface{}) string {
	if name, exist := event["label"]; exist {
		return name.(string)
	}
	return ""
}

func (p *Fonbet) GetEventMarkets(event map[string]interface{}) []interface{} {
	markets := event["bets"].(map[string]interface{})
	result := make([]interface{}, 0)
	for _, market := range markets {
		result = append(result, market)
	}
	return result
}

func (p *Fonbet) GetEventDate(event map[string]interface{}) string {
	return event["start"].(string)
}

func (p *Fonbet) ParseMarketName(market map[string]interface{}) string {
	if name, exist := market["question"]; exist {
		captions := name.(map[string]interface{})
		return captions["label"].(string)
	}
	return ""
}

func (p *Fonbet) ParseSelectionName(selectionData map[string]interface{}) string {
	if name, exist := selectionData["actor"]; exist {
		captions := name.(map[string]interface{})
		return captions["label"].(string)
	}
	return ""
}

func (p *Fonbet) ParseSelectionPrice(selectionData map[string]interface{}) float64 {
	return selectionData["odd"].(float64)
}

func (p *Fonbet) GetEventIsAntepost(event map[string]interface{}) bool {
	return false
}

func (p *Fonbet) GetEventIsLive(event map[string]interface{}) bool {
	return false
}

func (p *Fonbet) ParseSelectionLine(selectionData map[string]interface{}, marketData map[string]interface{}) float64 {
	line := 0.0

	return line
}

func (p *Fonbet) ParseMarketType(market map[string]interface{}) string {
	categoryType := market["type"].(map[string]interface{})
	categoryId := categoryType["id"].(float64)
	switch categoryId {
	case 1:
		return "SOCCER_MATCH_RESULT"
	case 10:
		return "SOCCER_DOUBLE_CHANCE"
	case 24:
		return "SOCCER_BOTH_TEAMS_TO_SCORE"
	case 39:
		return "SOCCER_AWAY_UNDER_OVER"
	case 9:
		return "SOCCER_UNDER_OVER"
	case 38:
		return "SOCCER_HOME_UNDER_OVER"
	}
	return ""
}

func (p *Fonbet) MatchMarketType(market map[string]interface{}, marketType string) (models.MarketType, error) {
	switch marketType {
	case "SOCCER_MATCH_RESULT":
		return models.NewMatchResult().MarketType, nil
	case "SOCCER_UNDER_OVER":
		return models.NewOverUnder().MarketType, nil
	case "SOCCER_BOTH_TEAMS_TO_SCORE":
		return models.NewBtts().MarketType, nil
	}
	return models.MarketType{}, fmt.Errorf("could not match market type")
}

func (p *Fonbet) ParseMarketId(market map[string]interface{}) string {
	return market["id"].(string)
}
func (p *Fonbet) GetMarketSelections(market map[string]interface{}) []interface{} {
	return market["choices"].([]interface{})
}

func (p *Fonbet) FetchEvent(e *models.Event) error {

	///ekdílosi/4067630-λάτσιο-inter-milan/
	formData := `{
    "context":
    {
        "url_key": "%s",
        "version": "1.0.1",
        "device": "web_vuejs_desktop",
        "lang": "en",
        "timezone": "-1",
        "url_params":
        {}
    },
    "components":
    [
        {
            "tree_compo_key": "event_market_list",
            "params": null
        }
    ]
}`
	formData = fmt.Sprintf(formData, e.Url)
	url := fmt.Sprintf("%s", p.config.BaseUrl)

	client := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte(formData)))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	jsonParsed, err := gabs.ParseJSON(body)
	if err != nil {
		fmt.Println(err.Error())
	}

	marks := jsonParsed.Path("components.data.event.bets")
	if marks != nil {
		results, _ := marks.Children()
		for _, markets := range results {

			marketMap, err := markets.ChildrenMap()
			if err != nil {
				return fmt.Errorf("could not fetch markets")
			}
			for _, mark := range marketMap {
				market := mark.Data()
				parsedMarket, parseError := ParseMarket(p, market.(map[string]interface{}), *e)
				if parseError == nil {
					e.Markets = append(e.Markets, parsedMarket)
				}
			}
		}
	}

	return fmt.Errorf("could not fetch details")
}

func (p *Fonbet) GetEventUrl(event map[string]interface{}) string {
	if url, found := event["url"]; found {
		return url.(string)
	}
	return ""
}
