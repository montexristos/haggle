package parsers

import (
	"fmt"
	"github.com/Jeffail/gabs"
	"github.com/gocolly/colly"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"haggle/models"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strconv"
	"time"
)

type PokerStars struct {
	Parser
	db     *gorm.DB
	config *models.SiteConfig
	c      *colly.Collector
	ID     int
}

func (p *PokerStars) SetConfig(c *models.SiteConfig) {
	p.config = c
	p.ID = c.SiteID
}

func (p *PokerStars) GetConfig() *models.SiteConfig {
	return p.config
}

func (p *PokerStars) Initialize() {
	p.c = GetCollector()
	p.c.OnResponse(func(response *colly.Response) {
		jsonParsed, err := gabs.ParseJSON(response.Body)
		if err != nil {
			fmt.Println(err.Error())
		}

		events, err := jsonParsed.Search("event").Children()
		if err != nil {
			print(err.Error())
		}
		if events != nil {
			for _, eventList := range events {
				evt := eventList.Data()
				ParseEvent(p, evt.(map[string]interface{}))
			}
		}
	})
}

func (p *PokerStars) GetDB() *gorm.DB {
	return p.db
}

func (p *PokerStars) Scrape() (bool, error) {
	return true, nil
}

func (p *PokerStars) ScrapeHome() (bool, error) {
	return true, nil
}

func (p *PokerStars) ScrapeLive() (bool, error) {
	return true, nil
}

func (p *PokerStars) ScrapeToday() (bool, error) {
	return true, nil
}

func (p *PokerStars) ScrapeTournament(tournamentId string) (bool, error) {
	// first get tournaments
	tourUrl := fmt.Sprintf("%s/%s", p.config.BaseUrl, tournamentId)

	err := p.c.Visit(tourUrl)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (p *PokerStars) SetDB(db *gorm.DB) {
	p.db = db
}

func (p *PokerStars) GetEventID(event map[string]interface{}) string {
	return strconv.Itoa(int(event["id"].(float64)))
}

func (p *PokerStars) GetEventName(event map[string]interface{}) string {
	return event["name"].(string)
}

func (p *PokerStars) GetEventIsAntepost(event map[string]interface{}) bool {
	return false
}

func (p *PokerStars) GetEventIsLive(event map[string]interface{}) bool {
	return cast.ToBool(event["isInplay"])
}

func (p *PokerStars) GetEventMarkets(event map[string]interface{}) []interface{} {
	return event["markets"].([]interface{})
}

func (p *PokerStars) ParseMarketType(market map[string]interface{}) string {
	return market["type"].(string)
}

func (p *PokerStars) ParseMarketId(market map[string]interface{}) string {
	return market["id"].(string)
}

func (p *PokerStars) ParseMarketName(market map[string]interface{}) string {
	return market["name"].(string)
}

func (p *PokerStars) ParseSelectionName(selectionData map[string]interface{}) string {
	return selectionData["name"].(string)
}

func (p *PokerStars) ParseSelectionPrice(selectionData map[string]interface{}) float64 {
	if odds, found := selectionData["odds"]; found {
		if dec, found := odds.(map[string]interface{})["dec"]; found {
			return cast.ToFloat64(dec)
		}
	}
	return 0.0
}

func (p *PokerStars) ParseSelectionLine(selectionData map[string]interface{}, marketData map[string]interface{}) float64 {
	if marketLine, found := marketData["line"]; found {
		return cast.ToFloat64(marketLine)
	}
	return 0
}

func (p *PokerStars) ParseMarketLine(market map[string]interface{}) float64 {
	return cast.ToFloat64(market["line"])
}

func (p *PokerStars) GetEventDate(event map[string]interface{}) string {
	return time.Unix(cast.ToInt64(event["eventTime"])/1000, 0).Format("2006-01-02 15:04:05")
}

func (p *PokerStars) GetMarketSelections(marketData map[string]interface{}) []interface{} {
	//sort odds
	selections := marketData["selection"].([]interface{})

	sort.Slice(selections, func(i, j int) bool {
		if selections[i] != nil && selections[j] != nil {
			pos1 := selections[i].(map[string]interface{})["pos"]
			pos2 := selections[j].(map[string]interface{})["pos"]
			if pos1 != nil && pos2 != nil {
				row1 := pos1.(map[string]interface{})["col"].(float64)
				row2 := pos2.(map[string]interface{})["col"].(float64)
				return row1 < row2
			}
		}
		return true
	})
	return selections
}

func (p *PokerStars) FetchEvent(e *models.Event) error {
	eventUrl := fmt.Sprintf(`sportsbook/v1/api/getEvent?eventId=%s&channelId=19&locale=en-gr`, e.BetradarID)
	url := fmt.Sprintf("%s/%s", p.config.BaseUrl, eventUrl)

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
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	jsonParsed, err := gabs.ParseJSON(body)
	if err != nil {
		fmt.Println(err.Error())
	}

	marks := jsonParsed.Path("markets")
	if marks != nil {
		results, _ := marks.Children()
		for _, markets := range results {
			market := markets.Data()
			parsedMarket, parseError := ParseMarket(p, market.(map[string]interface{}), *e)
			if parseError == nil {
				e.Markets = append(e.Markets, parsedMarket)
			}
		}
	}

	return fmt.Errorf("could not fetch details")
}

func (p *PokerStars) GetEventUrl(event map[string]interface{}) string {
	return ""
}
func (p *PokerStars) MatchMarketType(market map[string]interface{}, marketType string) (models.MarketType, error) {
	switch marketType {
	case "SOCCER:FT:AXB":
		return models.NewMatchResult().MarketType, nil
	case "SOCCER:FT:DNB":
		return models.NewDrawNoBet().MarketType, nil
	case "SOCCER:FT:OU":
		handicap := cast.ToFloat64(market["line"])
		return models.NewOverUnderHandicap(handicap).MarketType, nil
	case "SOCCER:P:OU":
		handicap := cast.ToFloat64(market["line"])
		if period, found := market["period"]; found && period == "P1" {
			return models.NewUnderOverHalf(handicap).MarketType, nil
		}
		return models.MarketType{}, fmt.Errorf("could not match market type")
	case "SOCCER:FT:OU1.5":
		return models.NewOverUnderHandicap(1.5).MarketType, nil
	case "SOCCER:FT:OU-A":
		handicap := cast.ToFloat64(market["line"])
		return models.NewUnderOverHome(handicap).MarketType, nil
	case "SOCCER:FT:OU-B":
		handicap := cast.ToFloat64(market["line"])
		return models.NewUnderOverAway(handicap).MarketType, nil
	case "SOCCER:FT:BTS":
		return models.NewBtts().MarketType, nil
	case "SOCCER:FT:DC":
		return models.NewDoubleChance().MarketType, nil
	}
	return models.MarketType{}, fmt.Errorf("could not match market type")

}
func (p *PokerStars) GetEventCanonicalName(event map[string]interface{}) string {
	if name, exist := event["name"]; exist {
		return name.(string)
	}
	return ""
}
