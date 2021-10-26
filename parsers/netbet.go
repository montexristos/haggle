package parsers

import (
	"bytes"
	"fmt"
	"github.com/Jeffail/gabs"
	"github.com/gocolly/colly"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"haggle/models"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Netbet struct {
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
func (n *Netbet) Initialize() {

	n.c = GetCollector()
	n.c.OnResponse(func(response *colly.Response) {
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
					ParseEvent(n, event.(map[string]interface{}))
				}
			}
		}
	})
}

func (n *Netbet) SetConfig(c *models.SiteConfig) {
	n.config = c
	n.ID = c.SiteID
}

func (n *Netbet) GetConfig() *models.SiteConfig {
	return n.config
}

func (n *Netbet) Scrape() (bool, error) {

	return true, nil
}
func (n *Netbet) ScrapeHome() (bool, error) {
	//err := n.c.Visit(fmt.Sprintf("%s/%s", n.config.BaseUrl, n.config.Urls["home"]))
	//if err != nil {
	//	return false, err
	//}

	return true, nil
}

func (n *Netbet) ScrapeLive() (bool, error) {
	//err := n.c.Visit(fmt.Sprintf("%s/%s", n.config.BaseUrl, n.config.Urls["live"]))
	//if err != nil {
	//	return false, err
	//}
	return true, nil
}

func (n *Netbet) ScrapeToday() (bool, error) {
	//&from= #2021-09-02T21:00:00.000Z&to=2021-09-03T21:00:00.000Z
	//now := time.Now().Format(time.RFC3339)
	//tomorrow := time.Now().Add(time.Hour * 24).Format(time.RFC3339)
	//url := fmt.Sprintf("%s/%s&from=%s&to=%s", n.config.BaseUrl, n.config.Urls["day"], now, tomorrow)
	//err := n.c.Visit(url)
	//if err != nil {
	//	return false, err
	//}
	return true, nil
}

func (n *Netbet) ScrapeTournament(tournamentUrl string) (bool, error) {
	// first get tournaments
	tourUrl := fmt.Sprintf("%s", n.config.BaseUrl)
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
	err := n.c.PostRaw(tourUrl, []byte(formData))
	if err != nil {
		return false, err
	}
	return true, nil
}

func (n *Netbet) SetDB(db *gorm.DB) {
	n.db = db
}

func (n *Netbet) GetDB() *gorm.DB {
	return n.db
}

func (n *Netbet) GetEventID(event map[string]interface{}) string {
	if live, found := event["live"]; found {
		if live != nil {
			if matchId, found := live.(map[string]interface{})["match_id"]; found {
				return strconv.Itoa(int(matchId.(float64)))
			}
		}
	}

	return strconv.Itoa(int(event["id"].(float64)))
}

func (n *Netbet) GetEventName(event map[string]interface{}) string {
	if name, exist := event["label"]; exist {
		return name.(string)
	}
	return ""
}
func (n *Netbet) GetEventCanonicalName(event map[string]interface{}) string {
	if name, exist := event["label"]; exist {
		return name.(string)
	}
	return ""
}

func (n *Netbet) GetEventMarkets(event map[string]interface{}) []interface{} {
	markets := event["bets"].(map[string]interface{})
	result := make([]interface{}, 0)
	for _, market := range markets {
		result = append(result, market)
	}
	return result
}

func (n *Netbet) GetEventDate(event map[string]interface{}) string {
	return event["start"].(string)
}

func (n *Netbet) ParseMarketName(market map[string]interface{}) string {
	if name, exist := market["question"]; exist {
		captions := name.(map[string]interface{})
		return captions["label"].(string)
	}
	return ""
}

func (n *Netbet) ParseSelectionName(selectionData map[string]interface{}) string {
	if name, exist := selectionData["actor"]; exist {
		captions := name.(map[string]interface{})
		return captions["label"].(string)
	}
	return ""
}

func (n *Netbet) ParseSelectionPrice(selectionData map[string]interface{}) float64 {
	return selectionData["odd"].(float64)
}

func (n *Netbet) GetEventIsAntepost(event map[string]interface{}) bool {
	return false
}

func (n *Netbet) GetEventIsLive(event map[string]interface{}) bool {
	return false
}

func (n *Netbet) ParseSelectionLine(selectionData map[string]interface{}, marketData map[string]interface{}) float64 {
	line := 0.0

	return line
}

func (n *Netbet) ParseMarketLine(market map[string]interface{}) float64 {
	return cast.ToFloat64(market["handicap"])
}

func (n *Netbet) ParseMarketType(market map[string]interface{}) string {
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

func (n *Netbet) MatchMarketType(market map[string]interface{}, marketType string) (models.MarketType, error) {
	switch marketType {
	case "SOCCER_MATCH_RESULT":
		return models.NewMatchResult().MarketType, nil
	case "SOCCER_UNDER_OVER":
		return models.NewOverUnderHandicap(2.5).MarketType, nil
	case "SOCCER_BOTH_TEAMS_TO_SCORE":
		return models.NewBtts().MarketType, nil
	}
	return models.MarketType{}, fmt.Errorf("could not match market type")
}

func (n *Netbet) ParseMarketId(market map[string]interface{}) string {
	return market["id"].(string)
}

func (n *Netbet) GetMarketSelections(market map[string]interface{}) []interface{} {
	selections := market["choices"].([]interface{})
	if market["type"].(map[string]interface{})["id"].(float64) == 9 {
		selections = make([]interface{}, 0)
		for _, sel := range market["choices"].([]interface{}) {
			if strings.Contains(sel.(map[string]interface{})["actor"].(map[string]interface{})["label"].(string), "2.5") {
				selections = append(selections, sel)
			}
		}
	}
	return selections
}

func (n *Netbet) FetchEvent(e *models.Event) error {

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
	url := fmt.Sprintf("%s", n.config.BaseUrl)

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
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	defer res.Body.Close()
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
				parsedMarket, parseError := ParseMarket(n, market.(map[string]interface{}), *e)
				if parseError == nil {
					e.Markets = append(e.Markets, parsedMarket)
				}
			}
		}
	}

	return fmt.Errorf("could not fetch details")
}

func (n *Netbet) GetEventUrl(event map[string]interface{}) string {
	if url, found := event["url"]; found {
		return url.(string)
	}
	return ""
}
