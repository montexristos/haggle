package scrape

import (
	"fmt"
	"github.com/montexristos/haggle/database"
	"github.com/montexristos/haggle/models"
	"github.com/montexristos/haggle/parsers"
	"gopkg.in/yaml.v3"
	"gorm.io/gorm"
	"io/ioutil"
	"math"
	"strings"
	"sync"
)

func ScrapeAll() (map[string]string, error) {
	database.ClearDB()
	result := make(map[string]string)
	sites := []string{
		"stoiximan",
		"novibet",
		"netbet",
		//"fonbet",
		"betsson",
		"pokerstars",
	}
	wg := &sync.WaitGroup{}
	for _, site := range sites {
		wg.Add(1)
		site := site
		go func() {
			res, err := ScrapeSite(site)
			if err == nil && res {
				result[site] = "ok"
			} else {
				result[site] = err.Error()
			}
			wg.Done()
		}()
	}
	wg.Wait()

	return result, nil
}

func GetParser(website string, db *gorm.DB) (*parsers.Parser, error) {
	config := models.ParseSiteConfig(website)
	if !config.Active {
		return nil, fmt.Errorf("Parser %s disabled", config.Id)
	}

	var parser parsers.Parser
	switch config.Id {
	case `stoiximan`:
		parser = &parsers.Stoiximan{}
		break
	//case `bet365`:
	//	parser = &parsers.Bet365{}
	//	break
	case `novibet`:
		parser = &parsers.Novibet{}
		break
	case `pokerstars`:
		parser = &parsers.PokerStars{}
		break
	case `winmasters`:
		parser = &parsers.Winmasters{}
		break
	case `bwin`:
		parser = &parsers.Bwin{}
		break
	case `netbet`:
		parser = &parsers.Netbet{}
		break
	case `pamestoixima`:
		parser = &parsers.PameStoixima{}
		break
	case `betsson`:
		parser = &parsers.Betsson{}
		break
	case `fonbet`:
		parser = &parsers.Fonbet{}
		break
	}
	if parser != nil {
		parser.SetDB(db)
		parser.SetConfig(config)
		parser.Initialize()
		return &parser, nil
	}
	return nil, fmt.Errorf("parser not found")
}

func ReadTournamentList(website string) ([]interface{}, error) {
	tournaments := make([]interface{}, 0)
	//read file
	yamlFile, err := ioutil.ReadFile("config/sites/tournaments.yaml")
	if err != nil {
		return tournaments, err
	}
	var siteTournaments map[string]interface{}
	err = yaml.Unmarshal(yamlFile, &siteTournaments)
	if err != nil {
		return tournaments, err
	}
	for _, sites := range siteTournaments {
		tournaments = append(tournaments, sites.(map[interface{}]interface{})[website])
	}
	return tournaments, nil
}

func ScrapeSite(website string) (bool, error) {
	//parser := GetParser(website, app.db)
	parser, err := GetParser(website, database.GetDb())
	if parser == nil {
		return false, fmt.Errorf("no parser for %s", website)
	}
	defer (*parser).Destruct()
	tournaments, err := ReadTournamentList(website)
	if parser != nil {
		_, err = (*parser).ScrapeHome()
		_, err = (*parser).ScrapeLive()
		_, err = (*parser).ScrapeToday()
		config := (*parser).GetConfig()
		for t := 0; t < len(tournaments); t++ {
			tourUrl := tournaments[t]
			if tourUrl != nil {
				_, err = (*parser).ScrapeTournament(tourUrl.(string))
			}
		}
		if err != nil {
			return false, fmt.Errorf("Parser %s initialize error", config.Id)
		}

		return true, nil
	}
	return false, fmt.Errorf("Parser not found")
}

func GetScrapeResults() (map[string]interface{}, error) {
	// find events that appear in more than one site
	//rows, err := GetDb().Raw(`SELECT betradar_id FROM (SELECT count(distinct site_id) as matches, betradar_id FROM haggle.events group by betradar_id) as tab WHERE matches > 1;`).Rows()
	rows, err := database.GetDb().Raw(`SELECT canonical_name FROM (SELECT count(distinct site_id) as matches, canonical_name FROM haggle.events group by canonical_name) as tab WHERE matches > 1;`).Rows()
	if err != nil {
		return map[string]interface{}{
			"error": err,
		}, nil
	}
	defer rows.Close()
	var eventMatch interface{}
	eventIds := make([]string, 0)
	for rows.Next() {
		err = rows.Scan(&eventMatch)
		if err != nil {
			return map[string]interface{}{
				"error": err,
			}, nil
		}
		eventIds = append(eventIds, fmt.Sprintf("%s", eventMatch.([]byte)))
	}
	events := make([]models.Event, 0)
	database.GetDb().Preload("Markets").Preload("Markets.Selections").Where("canonical_name in (?)", eventIds).Find(&events)
	matches := make(map[string][]models.Event)
	for _, event := range events {
		//var matchResultMarket *models.Market
		//var overUnderMarket *models.Market
		//var bttsMarket *models.Market
		canonicalName := parsers.TransformName(event.Name)
		if _, found := matches[canonicalName]; !found {
			matches[canonicalName] = make([]models.Event, 0)
		}

		matches[canonicalName] = append(matches[canonicalName], event)
	}
	sites := GetSites()
	arbs := FindArbs(matches, sites)

	return map[string]interface{}{
		"events": matches,
		"sites":  sites,
		"arbs":   arbs,
	}, nil
}

func FindArbs(matches map[string][]models.Event, sites map[int]string) map[string]string {
	arbs := make(map[string]string)
	for name, siteEvent := range matches {
		temp := make(map[string][]SiteOdd)
		for _, event := range siteEvent {
			site := sites[event.SiteID]
			for _, market := range event.Markets {
				for index, selection := range market.Selections {
					found := CheckMarket(event.CanonicalName, selection, market.MarketType, index, temp, site)
					if found != "" {
						arbs[name] = found
					}
				}
			}
		}
	}
	return arbs
}

type SiteOdd struct {
	Site       string
	Price      float64
	Name       string
	MarketName string
	Line       float64
}

func CheckMarket(eventName string, selection models.Selection, marketType string, index int, temp map[string][]SiteOdd, site string) string {
	results := make([]string, 0)
	key := fmt.Sprintf("%s-%d", marketType, index)
	if selection.Line > 0 {
		key = fmt.Sprintf("%s-%.2f", key, selection.Line)
	}
	arbKey := ""
	switch marketType {
	case "HCTG":
	case "BTSC":
	case "INTS":
	case "DNOB":
	case "OUHG":
	case "OUAG":
	case "OUH1":
	case "OU":
		arbKey = fmt.Sprintf("%s-%d", marketType, 1-index)
		if selection.Line > 0 {
			arbKey = fmt.Sprintf("%s-%.2f", arbKey, selection.Line)
		}
	}

	//skip some markets for now
	if marketType == "DBLC" || selection.Price > 12 {
		return ""
	}
	if _, found := temp[key]; !found {
		temp[key] = make([]SiteOdd, 0)
	}
	if _, found := temp[arbKey]; found && arbKey != "" {
		for _, sel := range temp[arbKey] {
			if arbValue := TestArbs(sel.Price, selection.Price); arbValue > 0 {
				results = append(results, fmt.Sprintf("%.2f%% %s:%s (%s:%.2f - %s:%.2f)", arbValue, eventName, key, site, selection.Price, sel.Site, sel.Price))
			}
		}
	} else if _, found = temp[key]; found {
		for _, sel := range temp[key] {
			if testArbSameIndex(sel.Price, selection.Price) {
				results = append(results, fmt.Sprintf("%s:%s (%s:%.2f - %s:%.2f)", eventName, key, site, selection.Price, sel.Site, sel.Price))
			}
		}
	}
	temp[key] = append(temp[key], SiteOdd{
		Site:       site,
		Price:      selection.Price,
		Name:       selection.Name,
		MarketName: marketType,
		Line:       selection.Line,
	})
	return strings.Join(results, "\r\n")
}

func TestArbs(odd1, odd2 float64) float64 {
	perc1 := 1 / odd1 * 100
	perc2 := 1 / odd2 * 100
	arb := math.Round(100*(perc1+perc2)) / 100
	if arb < 102 {
		return 100 - arb
	}
	return 0.0
}

func testArbSameIndex(odd1, odd2 float64) bool {
	diff := math.Abs(odd1 - odd2)
	max := math.Max(odd1, odd2)
	thres := diff * 100 / max
	if thres > 30 {
		return true
	}
	return false
}

func AllResults() map[string]interface{} {
	rows, err := database.GetDb().Model(&models.Event{}).Select("id").Rows()
	defer rows.Close()
	var eventMatch interface{}
	eventIds := make([]string, 0)
	for rows.Next() {
		err = rows.Scan(&eventMatch)
		if err != nil {
			fmt.Println(err.Error())
		}
		eventIds = append(eventIds, fmt.Sprintf("%s", eventMatch.([]byte)))
	}
	events := make([]models.Event, 0)
	database.GetDb().Preload("Markets").Preload("Markets.Selections").Where("id in (?)", eventIds).Find(&events)
	matches := make(map[string][]models.Event)
	for _, event := range events {
		//var matchResultMarket *models.Market
		//var overUnderMarket *models.Market
		//var bttsMarket *models.Market
		canonicalName := parsers.TransformName(event.Name)
		if _, found := matches[canonicalName]; !found {
			matches[canonicalName] = make([]models.Event, 0)
		}
		event.CanonicalName = canonicalName
		matches[canonicalName] = append(matches[canonicalName], event)
	}

	sites := GetSites()
	arbs := FindArbs(matches, sites)

	return map[string]interface{}{
		"events": matches,
		"sites":  sites,
		"arbs":   arbs,
	}
}

func GetSites() map[int]string {
	siteList := map[int]string{
		1:  `bet`,
		2:  `novibet`,
		3:  `pokerstars`,
		4:  `stoiximan`,
		5:  `winmasters`,
		6:  `bwin`,
		8:  `netbet`,
		12: `betsson`,
	}
	sites := make(map[int]string)
	for _, site := range siteList {
		parser, _ := GetParser(site, database.GetDb())
		if parser != nil {
			sites[(*parser).GetConfig().SiteID] = site
		}
	}
	return sites
}
