package parsers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/Jeffail/gabs"
	"github.com/gocolly/colly"
	"github.com/montexristos/haggle/models"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

type Stoiximan struct {
	Parser
	db     *gorm.DB
	config *models.SiteConfig
	c      *colly.Collector
	ID     int
}

func (s *Stoiximan) Initialize() {
	s.c = GetCollector()

	s.c.OnResponse(func(response *colly.Response) {
		body := response.Body
		var jsonString map[string]interface{}
		err := json.Unmarshal(body, &jsonString)
		if err == nil {
			if eventData, found := jsonString["event"]; found {
				ParseEvent(s, eventData.(map[string]interface{}))
			}
			if events, found := jsonString["events"]; found {
				markets := jsonString["markets"]
				selections := jsonString["selections"]
				for _, eventData := range events.(map[string]interface{}) {
					eventMap := eventData.(map[string]interface{})
					eventMap["markets"] = make([]interface{}, 0)
					for _, marketId := range eventData.(map[string]interface{})["marketIdList"].([]interface{}) {
						market := markets.(map[string]interface{})[cast.ToString(marketId)].(map[string]interface{})
						market["selections"] = make([]interface{}, 0)
						for _, selectionId := range market["selectionIdList"].([]interface{}) {
							selection := selections.(map[string]interface{})[cast.ToString(selectionId)].(map[string]interface{})
							market["selections"] = append(market["selections"].([]interface{}), selection)
						}
						eventMap["markets"] = append(eventMap["markets"].([]interface{}), market)
					}
					//override live now
					eventMap["liveNow"] = true
					ParseEvent(s, eventMap)
				}
			}
		}
	})
	s.c.OnHTML("script", func(e *colly.HTMLElement) {
		if strings.HasPrefix(e.Text, `window["initial_state"]=`) {
			jsonParsed, err := gabs.ParseJSON([]byte(e.Text[24:]))
			if err != nil {
				fmt.Println(err.Error())
			}
			topEvents := jsonParsed.Path("data.topEvents").Data()
			if topEvents != nil {
				s.parseTopEvents(topEvents.([]interface{}))
			}
			//TODO parse other items (fmt.Println(jsonParsed))
			block := jsonParsed.Path("data.blocks").Data()
			if block != nil {
				s.parseTopEvents(block.([]interface{}))
			}
		}
	})
}
func (s *Stoiximan) Destruct() {
}

func (s *Stoiximan) Scrape() (bool, error) {

	return true, nil
}

func (s *Stoiximan) ScrapeHome() (bool, error) {
	//_ = s.c.Visit(fmt.Sprintf("%s", s.config.BaseUrl))
	return true, nil
}

func (s *Stoiximan) ScrapeLive() (bool, error) {
	err := s.c.Request("GET", "https://en.stoiximan.gr/danae-webapi/api/live/overview/latest", nil, nil, http.Header{
		"x-language": {"1"},
		"x-operator": {"2"},
	})

	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *Stoiximan) ScrapeToday() (bool, error) {
	//_ = s.c.Visit(fmt.Sprintf("%s/%s", s.config.BaseUrl, s.config.Urls["day"]))
	return true, nil
}

func (s *Stoiximan) ScrapeTournament(tournamentUrl string) (bool, error) {
	_ = s.c.Visit(fmt.Sprintf("%s/%s", s.config.BaseUrl, tournamentUrl))
	return true, nil
}

func (s *Stoiximan) SetDB(db *gorm.DB) {
	s.db = db
}

func (s *Stoiximan) GetDB() *gorm.DB {
	return s.db
}

func (s *Stoiximan) SetConfig(c *models.SiteConfig) {
	s.config = c
	s.ID = c.SiteID
}
func (s *Stoiximan) GetConfig() *models.SiteConfig {
	return s.config
}

func (s *Stoiximan) GetEventID(event map[string]interface{}) string {
	if id, found := event["betradarMatchId"]; found {
		return cast.ToString(id)
	}
	if id, found := event["betRadarId"]; found {
		return cast.ToString(id)
	}

	return ""
}

func (s *Stoiximan) GetEventName(event map[string]interface{}) string {
	if name, found := event["name"]; found {
		return name.(string)
	}
	if participants, found := event["participants"]; found && len(participants.([]interface{})) > 1 {
		return fmt.Sprintf("%s - %s", cast.ToString(participants.([]interface{})[0]), cast.ToString(participants.([]interface{})[1]))
	}
	return ""
}
func (s *Stoiximan) GetEventCanonicalName(event map[string]interface{}) string {
	return event["name"].(string)
}

func (s *Stoiximan) GetEventMarkets(event map[string]interface{}) []interface{} {
	return event["markets"].([]interface{})
}

func (s *Stoiximan) GetEventDate(event map[string]interface{}) string {
	tm := time.Unix(int64(event["startTime"].(float64)/1000), 0)
	return tm.Format("2006-01-02 15:04:05")
}

func (s *Stoiximan) GetEventRunningTime(event map[string]interface{}) float64 {
	if liveData, found := event["liveData"]; found {
		if clock, found := liveData.(map[string]interface{})["clock"]; found {
			if timeData, found := clock.(map[string]interface{})["secondsSinceStart"]; found {
				return cast.ToFloat64(timeData) / 60
			}
		}
	}
	return -1.0
}
func (s *Stoiximan) GetEventScore(event map[string]interface{}) string {
	home := 0
	away := 0
	if liveData, found := event["liveData"]; found {
		if score, found := liveData.(map[string]interface{})["score"]; found {
			if homeScore, found := score.(map[string]interface{})["home"]; found {
				home = cast.ToInt(homeScore)
			} else {
				return ""
			}
			if awayScore, found := score.(map[string]interface{})["away"]; found {
				away = cast.ToInt(awayScore)
			} else {
				return ""
			}
			return fmt.Sprintf("%d-%d", home, away)
		}
	}
	return ""
}

func (s *Stoiximan) parseTopEvents(sports []interface{}) {
	for i := 0; i < len(sports); i++ {
		sport := sports[0].(map[string]interface{})
		events := sport["events"].([]interface{})
		if len(events) > 0 {
			for j := 0; j < len(events); j++ {
				_, _ = ParseEvent(s, events[j].(map[string]interface{}))
			}
		}
	}
}

func (s *Stoiximan) ParseMarketName(market map[string]interface{}) string {
	return market["name"].(string)
}

func (s *Stoiximan) ParseSelectionName(selectionData map[string]interface{}) string {
	return selectionData["name"].(string)
}

func (s *Stoiximan) ParseSelectionPrice(selectionData map[string]interface{}) float64 {
	return selectionData["price"].(float64)
}

func (s *Stoiximan) ParseSelectionLine(selectionData map[string]interface{}, marketData map[string]interface{}) float64 {
	line := 0.0
	if hc, found := marketData["handicap"]; found {
		return hc.(float64)
	}
	//TODO get line
	return line
}
func (s *Stoiximan) ParseSelectionId(selectionData map[string]interface{}) uint {
	if id, found := selectionData["id"]; found {
		return cast.ToUint(id)
	}
	return 0
}

func (s *Stoiximan) GetEventIsAntepost(event map[string]interface{}) bool {
	return false
}

func (s *Stoiximan) GetEventIsLive(event map[string]interface{}) bool {
	return cast.ToBool(event["liveNow"])
}

func (s *Stoiximan) ParseMarketType(market map[string]interface{}) string {
	return market["type"].(string)
}

func (s *Stoiximan) MatchMarketType(market map[string]interface{}, marketType string) (models.MarketType, error) {
	switch marketType {
	case "MR12":
		return models.NewMatchResult().MarketType, nil
	case "MRES":
		return models.NewMatchResult().MarketType, nil
	case "HCTG":
		if len(market["selections"].([]interface{})) < 3 {
			handicap := cast.ToFloat64(market["handicap"])
			return models.NewOverUnderHandicap(handicap).MarketType, nil
		}
	case "BTSC":
		return models.NewBtts().MarketType, nil
	case "INTS":
		return models.NewNextTeamToScore().MarketType, nil
	case "DBLC":
		return models.NewDoubleChance().MarketType, nil
	case "DNOB":
		return models.NewDrawNoBet().MarketType, nil
	case "OUHG":
		if len(market["selections"].([]interface{})) < 3 {
			handicap := cast.ToFloat64(market["handicap"])
			return models.NewUnderOverHome(handicap).MarketType, nil
		}
	case "OUAG":
		if len(market["selections"].([]interface{})) < 3 {
			handicap := cast.ToFloat64(market["handicap"])
			return models.NewUnderOverAway(handicap).MarketType, nil
		}
	case "OUH1":
		if len(market["selections"].([]interface{})) < 3 {
			handicap := cast.ToFloat64(market["handicap"])
			return models.NewUnderOverHalf(handicap).MarketType, nil
		}
	case "FG28":
		return models.NewFirstGoalEarly().MarketType, nil
	case "CNOU":
		handicap := cast.ToFloat64(market["handicap"])
		return models.NewUnderOverCorners(handicap).MarketType, nil
	case "3966":
		handicap := cast.ToFloat64(market["handicap"])
		return models.NewBttsOrOver(handicap).MarketType, nil
	}
	return models.MarketType{}, fmt.Errorf("could not match market type")
}

func (s *Stoiximan) ParseMarketLine(market map[string]interface{}) float64 {
	return cast.ToFloat64(market["handicap"])
}

func (s *Stoiximan) ParseMarketId(market map[string]interface{}) string {
	return market["id"].(string)
}

func (s *Stoiximan) GetMarketSelections(market map[string]interface{}) []interface{} {
	return market["selections"].([]interface{})
}

func (s *Stoiximan) FetchEvent(e *models.Event) error {
	url := fmt.Sprintf("%s/api%s", s.config.BaseUrl, e.Url)
	client := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
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
	var eventData interface{}
	err = json.Unmarshal(body, &eventData)
	if eventData == nil || eventData.(map[string]interface{})["data"] == nil {
		return fmt.Errorf("error fetcing event")
	}
	if event, found := eventData.(map[string]interface{})["data"].(map[string]interface{})["event"]; found {
		//parse markets
		if markets, found := event.(map[string]interface{})["markets"]; found {
			for _, market := range markets.([]interface{}) {
				parsedMarket, parseError := ParseMarket(s, market.(map[string]interface{}), *e)
				if parseError == nil {
					e.Markets = append(e.Markets, parsedMarket)
				}
			}
			return nil
		}
	}
	return fmt.Errorf("could not fetch details")
}

func (s *Stoiximan) GetEventUrl(event map[string]interface{}) string {
	if url, found := event["url"]; found {
		return url.(string)
	}
	return ""
}
func (s *Stoiximan) GetEventTournament(event map[string]interface{}) string {
	if league, found := event["leagueId"]; found {
		return cast.ToString(league)
	}
	return ""
}
func (s *Stoiximan) GetEventSport(event map[string]interface{}) string {
	if sport, found := event["sportId"]; found {
		return cast.ToString(sport)
	}
	return ""
}
