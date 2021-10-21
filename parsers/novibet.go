package parsers

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"haggle/models"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"time"
)

type Novibet struct {
	Parser
	db     *gorm.DB
	config *models.SiteConfig
	c      *colly.Collector
	ID     int
}

/**
 first visit https://www.novibet.gr/api/marketviews/findtimestamps?lang=el-GR&oddsR=1&timeZ=GTB%20Standard%20Time&usrGrp=G
then make request for each content key
*/
func (n *Novibet) Initialize() {

	n.c = GetCollector()
	n.c.OnResponse(func(response *colly.Response) {
		var resp interface{}
		err := json.Unmarshal(response.Body, &resp)
		if err != nil {
			panic(err.Error())
		}
		v := reflect.ValueOf(resp)
		switch v.Kind() {
		case reflect.Bool:
			fmt.Printf("bool: %v\n", v.Bool())
		case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64:
			fmt.Printf("int: %v\n", v.Int())
		case reflect.Uint, reflect.Uint8, reflect.Uint32, reflect.Uint64:
			fmt.Printf("int: %v\n", v.Uint())
		case reflect.Float32, reflect.Float64:
			fmt.Printf("float: %v\n", v.Float())
		case reflect.String:
			fmt.Printf("string: %v\n", v.String())
		case reflect.Slice:
			//check response for tournaments api/marketviews/coupon/16/440817/3464004?lang=el-GR&oddsR=1&timeZ=GTB%20Standard%20Time&usrGrp=GR&_=1630499278281
			if items, exists := resp.([]interface{}); exists {
				for _, v := range items {
					item := v.(map[string]interface{})
					if itemType, found := item["BetViews"]; found {
						if reflect.TypeOf(itemType).Kind() == reflect.Slice {
							for _, betview := range itemType.([]interface{}) {
								n.ParseBetViews(betview.(map[string]interface{}))
							}
						}
					}
				}
			}
			fmt.Printf("slice: len=%d, %v\n", v.Len(), v.Interface())
		case reflect.Map:
			//check response for tournaments api/marketviews/coupon/16/440817/3464004?lang=el-GR&oddsR=1&timeZ=GTB%20Standard%20Time&usrGrp=GR&_=1630499278281
			if items, exists := resp.(map[string]interface{})["Items"]; exists {
				for _, v := range items.([]interface{}) {
					item := v.(map[string]interface{})
					if itemType, found := item["BetViews"]; found {
						if reflect.TypeOf(itemType).Kind() == reflect.Slice {
							for _, betview := range itemType.([]interface{}) {
								n.ParseBetViews(betview.(map[string]interface{}))
							}
						}
					}
				}
				return
			}
			if items, exists := resp.(map[string]interface{})["MarketViews"]; exists {
				for _, v := range items.([]interface{}) {
					item := v.(map[string]interface{})
					if itemType, found := item["BetViews"]; found {
						if reflect.TypeOf(itemType).Kind() == reflect.Slice {
							for _, betview := range itemType.([]interface{}) {
								n.ParseBetViews(betview.(map[string]interface{}))
							}
						}
					}
				}
			}
			fmt.Printf("map: %v\n", v.Interface())
		case reflect.Chan:
			fmt.Printf("chan %v\n", v.Interface())
		default:
			fmt.Println(resp)
		}
	})
}

func (n *Novibet) SetConfig(c *models.SiteConfig) {
	n.config = c
	n.ID = c.SiteID
}

func (n *Novibet) GetConfig() *models.SiteConfig {
	return n.config
}

func (n *Novibet) Scrape() (bool, error) {

	return true, nil
}
func (n *Novibet) ScrapeHome() (bool, error) {
	err := n.c.Visit(fmt.Sprintf("%s/%s%d", n.config.BaseUrl, n.config.Urls["home"], time.Now().Unix()))
	if err != nil {
		return false, err
	}

	return true, nil
}

func (n *Novibet) ScrapeLive() (bool, error) {
	err := n.c.Visit(fmt.Sprintf("%s/%s%d", n.config.BaseUrl, n.config.Urls["live"], time.Now().Unix()))
	if err != nil {
		return false, err
	}
	return true, nil
}

func (n *Novibet) ScrapeToday() (bool, error) {
	err := n.c.Visit(fmt.Sprintf("%s/%s%d", n.config.BaseUrl, n.config.Urls["day"], time.Now().Unix()))
	if err != nil {
		return false, err
	}
	return true, nil
}

func (n *Novibet) ScrapeTournament(tournamentUrl string) (bool, error) {
	// first get tournaments
	tourUrl := fmt.Sprintf("%s/%s%d", n.config.BaseUrl, tournamentUrl, time.Now().Unix())
	err := n.c.Visit(tourUrl)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (n *Novibet) SetDB(db *gorm.DB) {
	n.db = db
}

func (n *Novibet) GetDB() *gorm.DB {
	return n.db
}

func (n *Novibet) ParseBetViews(betview map[string]interface{}) {
	if competitions, exist := betview["Competitions"]; exist {
		for _, competition := range competitions.([]interface{}) {
			n.ParseCompetition(competition.(map[string]interface{}))
		}
		return
	}
	if events, exist := betview["Items"]; exist {
		for _, event := range events.([]interface{}) {
			ParseEvent(n, event.(map[string]interface{}))
		}
		return
	}
}

func (n *Novibet) ParseCompetition(competition map[string]interface{}) {
	if events, exist := competition["Events"]; exist {
		for _, event := range events.([]interface{}) {
			ParseEvent(n, event.(map[string]interface{}))
		}
	}
}

func (n *Novibet) parseTopEvents(sports []interface{}) {
	for i := 0; i < len(sports); i++ {
		sport := sports[0].(map[string]interface{})
		events := sport["events"].([]interface{})
		if len(events) > 0 {
			for j := 0; j < len(events); j++ {
				ParseEvent(n, events[j].(map[string]interface{}))
			}
		}
	}
}

func (n *Novibet) GetEventID(event map[string]interface{}) string {
	if id, found := event["SportradarMatchId"]; found && id != nil {
		return strconv.Itoa(int(id.(float64)))
	}
	if id, found := event["EventBetContextId"]; found && id != nil {
		return strconv.Itoa(int(id.(float64)))
	}
	if id, found := event["BetContextId"]; found && id != nil {
		return strconv.Itoa(int(id.(float64)))
	}
	return ""
}

func (n *Novibet) GetEventName(event map[string]interface{}) string {
	if captionsMap, exist := event["AdditionalCaptions"]; exist {
		captions := captionsMap.(map[string]interface{})
		return fmt.Sprintf("%s - %s", captions["Competitor1"].(string), captions["Competitor2"].(string))
	}
	return event["Path"].(string)
}

func (n *Novibet) GetEventCanonicalName(event map[string]interface{}) string {
	if captionsMap, exist := event["AdditionalCaptions"]; exist {
		captions := captionsMap.(map[string]interface{})
		return fmt.Sprintf("%s - %s", captions["Competitor1"].(string), captions["Competitor2"].(string))
	}
	return event["Path"].(string)
}

func (n *Novibet) GetEventMarkets(event map[string]interface{}) []interface{} {
	return event["Markets"].([]interface{})
}

func (n *Novibet) GetEventDate(event map[string]interface{}) string {
	return ""
}

func (n *Novibet) ParseMarketName(market map[string]interface{}) string {
	if marketType, found := market["Caption"]; found {
		return marketType.(string)
	}
	if marketType, found := market["BetTypeSysname"]; found {
		return marketType.(string)
	}
	if marketType, found := market["MarkettSysname"]; found {
		return marketType.(string)
	}

	return ""
}

func (n *Novibet) ParseSelectionName(selectionData map[string]interface{}) string {
	return selectionData["Caption"].(string)
}

func (n *Novibet) ParseSelectionPrice(selectionData map[string]interface{}) float64 {
	return selectionData["Price"].(float64)
}

func (n *Novibet) GetEventIsAntepost(event map[string]interface{}) bool {
	return false
}

func (n *Novibet) ParseSelectionLine(selectionData map[string]interface{}, marketData map[string]interface{}) float64 {
	line := 0.0
	hc := marketData["BetItems"].([]interface{})[0].(map[string]interface{})["InstanceCaption"]
	if hc == nil {
		return line
	}
	handicap, err := strconv.ParseFloat(hc.(string), 64)
	if err == nil {
		return handicap
	}
	return line
}

func (n *Novibet) ParseMarketType(market map[string]interface{}) string {
	if marketType, found := market["BetTypeSysname"]; found {
		return marketType.(string)
	}
	if marketType, found := market["Sysname"]; found {
		return marketType.(string)
	}
	if marketType, found := market["MarketSysname"]; found {
		return marketType.(string)
	}
	return ""
}

func (n *Novibet) MatchMarketType(market map[string]interface{}, marketType string) (models.MarketType, error) {
	switch marketType {
	case "SOCCER_MATCH_RESULT":
		return models.NewMatchResult().MarketType, nil
	case "SOCCER_UNDER_OVER":
		if len(market["BetItems"].([]interface{})) > 0 {
			if hc, found := market["BetItems"].([]interface{})[0].(map[string]interface{})["InstanceCaption"]; found {
				return models.NewOverUnderHandicap(cast.ToFloat64(hc)).MarketType, nil
			}
		}
		return models.NewOverUnder().MarketType, nil
	case "SOCCER_BOTH_TEAMS_TO_SCORE":
		return models.NewBtts().MarketType, nil
	case "INTS":
		return models.NewNextTeamToScore().MarketType, nil
	case "SOCCER_DOUBLE_CHANCE":
		return models.NewDoubleChance().MarketType, nil
	case "SOCCER_MATCH_RESULT_NODRAW":
		return models.NewDrawNoBet().MarketType, nil
	case "SOCCER_HOME_UNDER_OVER":
		handicap := 2.5
		if len(market["BetItems"].([]interface{})) > 0 {
			if hc, found := market["BetItems"].([]interface{})[0].(map[string]interface{})["InstanceCaption"]; found {
				return models.NewUnderOverHome(cast.ToFloat64(hc)).MarketType, nil
			}
		}
		return models.NewUnderOverHome(handicap).MarketType, nil
	case "SOCCER_AWAY_UNDER_OVER":
		handicap := 2.5
		if len(market["BetItems"].([]interface{})) > 0 {
			if hc, found := market["BetItems"].([]interface{})[0].(map[string]interface{})["InstanceCaption"]; found {
				return models.NewUnderOverAway(cast.ToFloat64(hc)).MarketType, nil
			}
		}
		return models.NewUnderOverAway(handicap).MarketType, nil
	case "OUH1":
		handicap := 2.5
		if len(market["BetItems"].([]interface{})) > 0 {
			if hc, found := market["BetItems"].([]interface{})[0].(map[string]interface{})["InstanceCaption"]; found {
				return models.NewUnderOverAway(cast.ToFloat64(hc)).MarketType, nil
			}
		}
		return models.NewUnderOverAway(handicap).MarketType, nil
	case "FG28":
		return models.NewFirstGoalEarly().MarketType, nil
	case "SOCCER_CORNERS_UNDER_OVER":
		if len(market["BetItems"].([]interface{})) > 0 {
			if hc, found := market["BetItems"].([]interface{})[0].(map[string]interface{})["InstanceCaption"]; found {
				return models.NewUnderOverCorners(cast.ToFloat64(hc)).MarketType, nil
			}
		}
		return models.NewUnderOverCorners(2.5).MarketType, nil
	}
	return models.MarketType{}, fmt.Errorf("could not match market type")
}

func (n *Novibet) ParseMarketId(market map[string]interface{}) string {
	return market["MarketId"].(string)
}
func (n *Novibet) GetMarketSelections(market map[string]interface{}) []interface{} {
	return market["BetItems"].([]interface{})
}

func (n *Novibet) FetchEvent(e *models.Event) error {
	url := fmt.Sprintf("%s/api/marketviews/event/16/%s?lang=en-US", n.config.BaseUrl, e.BetradarID)
	client := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
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
	var eventData interface{}
	err = json.Unmarshal(body, &eventData)
	if marketGroups, found := eventData.(map[string]interface{})["MarketCategories"]; found {
		for _, marketGroup := range marketGroups.([]interface{}) {
			if markets, found := marketGroup.(map[string]interface{})["Items"]; found {
				for _, marketTypes := range markets.([]interface{}) {
					for _, market := range marketTypes.(map[string]interface{})["BetViews"].([]interface{}) {
						parsedMarket, parseError := ParseMarket(n, market.(map[string]interface{}), *e)
						if parseError == nil {
							e.Markets = append(e.Markets, parsedMarket)
						}
					}
				}
			}
		}
		return nil
	}
	return fmt.Errorf("could not fetch details")
}

func (n *Novibet) GetEventUrl(event map[string]interface{}) string {
	return ""
}
