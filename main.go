package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"haggle/fixtureModels"
	"haggle/models"
	"haggle/parsers"
	"haggle/tools"
	"io"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"path"
	"strconv"
	"sync"
	"time"
)

/*
docker run \
--detach \
--name=mysql \
--env="MYSQL_ROOT_PASSWORD=123" \
--publish 6602:3306 \
--volume=data:/var/lib/mysql \
mysql
*/

func main() {
	db := GetDb()
	app := Application{
		db,
	}
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/data", app.dataLink)
	router.HandleFunc("/scrape", app.scrapeLink)
	router.HandleFunc("/home", homeLink)
	router.HandleFunc("/cache", cacheLink)
	// router.HandleFunc("/ml", mlLink)
	router.HandleFunc("/week/{week}", weekLink)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./build/")))

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"http://localhost:3000", "http://142.93.163.59:8088"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	//initiate an initial scrape
	//_, _ =scrapeAll()

	// start server listen
	log.Fatal(http.ListenAndServe(":8088", handlers.CompressHandler(handlers.CORS(originsOk, headersOk, methodsOk)(router))))
	//log.Fatal(http.ListenAndServe(":8088", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}

type Application struct {
	db *gorm.DB
}

func GetDb() *gorm.DB {
	//CREATE SCHEMA `haggle` DEFAULT CHARACTER SET utf8 ;
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:123@(localhost:6602)/haggle?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	err = db.AutoMigrate(&models.Event{}, &models.Market{}, &models.Selection{})
	if err != nil {
		panic(err.Error())
	}
	return db
}

func ClearDB() {
	db := GetDb()
	db.Where("1 = 1").Unscoped().Delete(&models.Selection{})
	db.Where("1 = 1").Unscoped().Delete(&models.Market{})
	db.Where("1 = 1").Unscoped().Delete(&models.Event{})
	if db.Error != nil {
		panic(db.Error.Error())
	}
}

func (app *Application) scrapeAll() (map[string]string, error) {
	ClearDB()
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
			res, err := app.ScrapeSite(site)
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

func GetParser(website string, db *gorm.DB) (parsers.Parser, error) {
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
		return parser, nil
	}
	return nil, fmt.Errorf("parser not found")
}

func (app *Application) ReadTournamentList(website string) ([]interface{}, error) {
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

func (app *Application) ScrapeSite(website string) (bool, error) {
	//parser := GetParser(website, app.db)
	parser, err := GetParser(website, app.db)
	tournaments, err := app.ReadTournamentList(website)
	if parser != nil {
		_, err = parser.ScrapeHome()
		_, err = parser.ScrapeLive()
		_, err = parser.ScrapeToday()
		config := parser.GetConfig()
		for t := 0; t < len(tournaments); t++ {
			tourUrl := tournaments[t]
			if tourUrl != nil {
				_, err = parser.ScrapeTournament(tourUrl.(string))
			}
		}
		if err != nil {
			return false, fmt.Errorf("Parser %s initialize error", config.Id)
		}

		return parser.Scrape()
	}
	return false, fmt.Errorf("Parser not found")
}

func getScrapeResults() (map[string]interface{}, error) {
	// find events that appear in more than one site
	//rows, err := GetDb().Raw(`SELECT betradar_id FROM (SELECT count(distinct site_id) as matches, betradar_id FROM haggle.events group by betradar_id) as tab WHERE matches > 1;`).Rows()
	rows, err := GetDb().Raw(`SELECT canonical_name FROM (SELECT count(distinct site_id) as matches, canonical_name FROM haggle.events group by canonical_name) as tab WHERE matches > 1;;`).Rows()
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
	GetDb().Preload("Markets").Preload("Markets.Selections").Where("canonical_name in (?)", eventIds).Find(&events)
	matches := make(map[string][]models.Event)
	for _, event := range events {
		//var matchResultMarket *models.Market
		//var overUnderMarket *models.Market
		//var bttsMarket *models.Market
		if _, found := matches[event.CanonicalName]; !found {
			matches[event.CanonicalName] = make([]models.Event, 0)
		}

		matches[event.CanonicalName] = append(matches[event.CanonicalName], event)
	}
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
	arbs := FindArbs(matches)
	sites := make(map[int]string)
	for _, site := range siteList {
		parser, _ := GetParser(site, GetDb())
		if parser != nil {
			sites[parser.GetConfig().SiteID] = site
		}
	}

	return map[string]interface{}{
		"events": matches,
		"sites":  sites,
		"arbs":   arbs,
	}, nil
}

func FindArbs(matches map[string][]models.Event) map[string]string {
	arbs := make(map[string]string)
	temp := make(map[string]interface{})
	for name, siteEvent := range matches {
		for _, event := range siteEvent {
			for _, market := range event.Markets {
				for index, selection := range market.Selections {
					found := checkMarket(event.CanonicalName, selection, market.MarketType, index, temp)
					if found != "" {
						arbs[name] = found
					}
				}
			}
		}
	}
	return arbs
}

func checkMarket(eventName string, selection models.Selection, marketType string, index int, temp map[string]interface{}) string {
	key := fmt.Sprintf("%s-%d", marketType, index)
	if selection.Line > 0 {
		key = fmt.Sprintf("%s-%f", key, selection.Line)
	}
	if sel, found := temp[key]; found {
		if testArbs(sel.(models.Selection).Price, selection.Price) {
			return fmt.Sprintf("%s - %s - %s (%.2f - %.2f)", eventName, marketType, key, selection.Price, sel.(models.Selection).Price)
		}
	} else {
		temp[key] = selection
	}
	return ""
}

func testArbs(odd1, odd2 float64) bool {
	diff := math.Abs(odd1 - odd2)
	max := math.Max(odd1, odd2)
	thres := diff * 100 / max
	if thres > 20 {
		return true
	}
	return false
}

func (app *Application) scrapeLink(w http.ResponseWriter, r *http.Request) {
	ClearDB()
	result, err := app.scrapeAll()
	if err != nil {
		_ = json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}
	_ = json.NewEncoder(w).Encode(result)
}

func (app *Application) dataLink(w http.ResponseWriter, r *http.Request) {
	result, err := getScrapeResults()
	if err != nil {
		_ = json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}
	_ = json.NewEncoder(w).Encode(result)
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	week := tools.GetHomeDateRange()
	layout := "02-01-2006"
	tournamentList := getTournamentData(false, fixtureModels.Week{Start: time.Now(), End: time.Now().AddDate(1, 0, 0)})
	stats := fixtureModels.NewStats()
	getQueryParams(r, stats)
	fixtureModels.ProcessFixtures(&tournamentList, stats, week, week.Start)
	_ = json.NewEncoder(w).Encode(struct {
		TournamentList []*fixtureModels.Tournament `json:"tournaments"`
		Weeks          []*fixtureModels.Week       `json:"weeks"`
		End            string                      `json:"end"`
		Start          string                      `json:"start"`
		Stats          *fixtureModels.Stats        `json:"stats"`
	}{
		TournamentList: tournamentList,
		Weeks:          getWeeks(),
		End:            week.End.Format(layout),
		Start:          week.Start.Format(layout),
		Stats:          stats,
	})
}

func cacheLink(w http.ResponseWriter, r *http.Request) {
	week := tools.GetHomeDateRange()
	tournamentList := getTournamentData(true, week)
	_ = json.NewEncoder(w).Encode(tournamentList)
}

func getWeeks() []*fixtureModels.Week {
	year, currentWeek := time.Now().ISOWeek()
	weeks := make([]*fixtureModels.Week, 0)
	for i := currentWeek; i > 0; i-- {
		start, end := tools.WeekRange(year, i)
		weeks = append(weeks, &fixtureModels.Week{
			Start: start,
			End:   end,
		})
	}
	for i := 52; i > 1; i-- {
		start, end := tools.WeekRange(year-1, i)
		weeks = append(weeks, &fixtureModels.Week{
			Start: start,
			End:   end,
		})
	}
	return weeks
}

func weekLink(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	weekEnd := vars["week"]
	layout := "02-01-2006"
	day, _ := time.Parse(layout, weekEnd)
	end := day
	start := day.AddDate(0, 0, -7)
	week := fixtureModels.Week{
		Start: time.Date(start.Year(), start.Month(), start.Day(), 0, 0, 0, 0, time.Now().Location()),
		End:   time.Date(end.Year(), end.Month(), end.Day(), 23, 59, 59, 0, time.Now().Location()),
	}
	stats := fixtureModels.NewStats()
	getQueryParams(r, stats)
	tournamentList, stats := getDataForWeek(week, stats)
	_ = json.NewEncoder(w).Encode(struct {
		TournamentList []*fixtureModels.Tournament `json:"tournaments"`
		Weeks          []*fixtureModels.Week       `json:"weeks"`
		End            string                      `json:"end"`
		Start          string                      `json:"start"`
		Stats          *fixtureModels.Stats        `json:"stats"`
	}{
		TournamentList: tournamentList,
		Weeks:          getWeeks(),
		End:            end.Format(layout),
		Start:          start.Format(layout),
		Stats:          stats,
	})
}

func getQueryParams(r *http.Request, stats *fixtureModels.Stats) {
	v := r.URL.Query()
	if threshold0 := v.Get("threshold0"); threshold0 != "" {
		stats.Threshold0, _ = strconv.ParseFloat(threshold0, 64)
	}
	if threshold1 := v.Get("threshold1"); threshold1 != "" {
		stats.Threshold1, _ = strconv.ParseFloat(threshold1, 64)
	}
	if threshold2 := v.Get("threshold2"); threshold2 != "" {
		stats.Threshold2, _ = strconv.ParseFloat(threshold2, 64)
	}
	if threshold3 := v.Get("threshold3"); threshold3 != "" {
		stats.Threshold3, _ = strconv.ParseFloat(threshold3, 64)
	}
	if threshold4 := v.Get("threshold4"); threshold4 != "" {
		stats.Threshold4, _ = strconv.ParseFloat(threshold4, 64)
	}
}

func allLink(w http.ResponseWriter, r *http.Request) {
	layout := "02-01-2006"
	now := time.Now()
	week := fixtureModels.Week{
		Start: time.Date(now.Year(), 1, 1, 0, 0, 0, 0, time.Now().Location()),
		End:   time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, time.Now().Location()),
	}
	tournamentList := getTournamentData(false, fixtureModels.Week{Start: time.Now(), End: time.Now().AddDate(1, 0, 0)})
	stats := fixtureModels.NewStats()
	fixtureModels.ProcessFixtures(&tournamentList, stats, week, time.Now())

	_ = json.NewEncoder(w).Encode(struct {
		TournamentList []*fixtureModels.Tournament `json:"tournaments"`
		Weeks          []*fixtureModels.Week       `json:"weeks"`
		End            string                      `json:"end"`
		Start          string                      `json:"start"`
		Stats          *fixtureModels.Stats        `json:"stats"`
	}{
		TournamentList: tournamentList,
		Weeks:          getWeeks(),
		End:            week.Start.Format(layout),
		Start:          week.End.Format(layout),
		Stats:          stats,
	})
}

func getDataForWeek(week fixtureModels.Week, stats *fixtureModels.Stats) ([]*fixtureModels.Tournament, *fixtureModels.Stats) {
	tournamentList := getTournamentData(false, week)
	fixtureModels.ProcessFixtures(&tournamentList, stats, week, week.Start)
	return tournamentList, stats
}

func getTournamentUpcomingFixtures(tournament *fixtureModels.Tournament, cache bool, week fixtureModels.Week) ([]*fixtureModels.Fixture, error) {
	var fixtures []*fixtureModels.Fixture
	var fixtureDTOs []fixtureModels.FixtureDTO
	var err error
	filePath := fmt.Sprintf("data/%d.json", tournament.Id)

	// Open our jsonFile
	jsonFile, err := os.Open(filePath)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	if err == nil && !cache {
		// defer the closing of our jsonFile so that we can parse it later on
		defer jsonFile.Close()

		byteValue, _ := ioutil.ReadAll(jsonFile)

		err = json.Unmarshal([]byte(byteValue), &fixtures)
		if err == nil || len(fixtures) > 0 {
			return fixtures, nil
		}
	}
	if tournament.Id == 0 {
		return fixtures, fmt.Errorf("No id for tournament %s", tournament.Name)
	}
	fixtureDTOs, err = tournament.GetFixturesNew(week.Start.Format("2006-01-02"), week.End.Format("2006-01-02"))
	//fixtureDTOs, err = fixtureModels.GetFixtures(tournament)

	fixtures = fixtureModels.MapFixtures(fixtureDTOs, tournament)

	if err != nil {
		fmt.Println("Error getting fixtureModels", err, tournament)
		return make([]*fixtureModels.Fixture, 0), err
	}
	if len(fixtures) == 0 {
		fmt.Println("No fixtureModels for tournament ", tournament.Name)
		return make([]*fixtureModels.Fixture, 0), err
	}
	if cache {
		jsonValue, _ := json.Marshal(fixtures)
		err = ioutil.WriteFile(filePath, jsonValue, 0644)
		if err != nil {
			fmt.Println("Error writing json file", err)
		}
	}
	return fixtures, nil
}

func getTournamentData(cache bool, week fixtureModels.Week) []*fixtureModels.Tournament {
	wg := sync.WaitGroup{}
	tournaments := sync.Map{}
	for _, tournament := range parseConfig() {
		wg.Add(1)
		go getCompetitionData(tournament, &wg, &tournaments, cache, week)
	}
	wg.Wait()
	tournamentList := make([]*fixtureModels.Tournament, 0)
	tournaments.Range(func(key interface{}, value interface{}) bool {
		tour := value.(**fixtureModels.Tournament)
		upcomingFixtures, err := getTournamentUpcomingFixtures(*tour, cache, week)
		(*tour).Fixtures = append((*tour).Fixtures, upcomingFixtures...)
		if err != nil {
			fmt.Println(err)
		}
		tournamentList = append(tournamentList, *tour)
		return true
	})

	return tournamentList
}

func getTournamentUpcomingData() {
	wg := sync.WaitGroup{}
	week := tools.GetHomeDateRange()
	tournaments := make(map[int]string)
	for _, tournament := range parseConfig() {
		tournaments[tournament.Id] = tournament.Fd
	}
	wg.Wait()
	tournamentList := make([]*fixtureModels.Tournament, 0)
	for tournamentId, tournamentKey := range tournaments {
		tour := &fixtureModels.Tournament{
			Id: tournamentId,
			Fd: tournamentKey,
		}
		upcomingFixtures, err := getTournamentUpcomingFixtures(tour, true, week)
		(*tour).Fixtures = append((*tour).Fixtures, upcomingFixtures...)
		if err != nil {
			fmt.Println(err)
		}
		tournamentList = append(tournamentList, tour)
	}
}

func parseConfig() []*fixtureModels.Tournament {
	yamlFile, err := ioutil.ReadFile("config/tournaments.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	y := fixtureModels.TournamentConfig{}

	err = yaml.Unmarshal(yamlFile, &y)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return y.Tournaments
}

// Parses csv files to get the competition fixtures and teams
func getCompetitionData(tournament *fixtureModels.Tournament, wg *sync.WaitGroup, tournaments *sync.Map, cache bool, week fixtureModels.Week) {
	csvPath, _ := getCSVPath(tournament, false)
	if cache && csvPath != "" {
		if err := tools.DownloadFile(csvPath, tournament.Csv); err != nil {
			log.Println("error downloading file " + tournament.Csv)
		}
		csvPath, err := getCSVPath(tournament, true)
		if err != nil {
			log.Println("error finding file " + tournament.Old)
		} else {
			if err := tools.DownloadFile(csvPath, tournament.Old); err != nil {
				log.Printf("error downloading file %s error: %s", tournament.Old, err.Error())
			}
		}
	}
	readCSVTournaments(tournament, week, true)
	readCSVTournaments(tournament, week, false)
	wg.Done()
	tournaments.Store(tournament.Name, &tournament)
}

func getCSVPath(tournament *fixtureModels.Tournament, old bool) (string, error) {
	if old {
		if tournament.Old != "" {
			url := tournament.Old
			filename := path.Base(url)
			return fmt.Sprintf("data/old/%s", filename), nil
		}
		return "", fmt.Errorf("file not found")
	}
	url := tournament.Csv
	filename := path.Base(url)
	return fmt.Sprintf("data/%s", filename), nil
}

func readCSVTournaments(tournament *fixtureModels.Tournament, week fixtureModels.Week, old bool) {
	csvPath, err := getCSVPath(tournament, old)
	if err != nil {
		return
	}
	var csvfile *os.File
	// Parse the file
	csvfile, err = os.Open(csvPath)
	if err != nil {
		return
	}
	r := csv.NewReader(csvfile)
	// Iterate through the records
	var event fixtureModels.Event
	var headers []string
	index := 0
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
		if index == 0 {
			index++
			headers = record
			continue
		}

		fixture, err := event.Parse(record, headers, tournament, week)
		if err == nil {
			if fixture.Date.Before(week.End) {
				tournament.Fixtures = append(tournament.Fixtures, &fixture)
			}
		}
	}
}

func GenerateCSV(filename string) {
	weeks := getWeeks()
	//csv := fmt.Sprintf("%s, %s, %s, %s, %s\r\n", "homeOver", "homeGG", "awayOver", "awayGG", "result")
	csv := ""
	for _, week := range weeks {
		stats := fixtureModels.NewStats()
		tournamentList, stats := getDataForWeek(*week, stats)
		fixtureModels.ProcessFixtures(&tournamentList, stats, *week, week.Start)
		for _, tour := range tournamentList {
			for _, fix := range tour.Fixtures {
				if fix.Score == "" {
					continue
				}
				row, err := fix.GenerateRowOver(false)
				if err == nil {
					csv += row
				}
			}
		}
	}
	//write csv to file
	file, err := os.Create(filename)
	if err == nil {
		_, _ = file.Write([]byte(csv))
		_ = file.Close()
	}
}

func GenerateCSVString(filename string) {
	weeks := getWeeks()
	csv := fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s\r\n", "homeOver", "homeGG", "awayOver", "awayGG", "homeGoals", "awayGoals", "result")
	for _, week := range weeks {
		stats := fixtureModels.NewStats()
		tournamentList, stats := getDataForWeek(*week, stats)
		fixtureModels.ProcessFixtures(&tournamentList, stats, *week, week.Start)
		for _, tour := range tournamentList {
			for _, fix := range tour.Fixtures {
				if fix.Score == "" {
					continue
				}
				row, err := fix.GenerateRow(false)
				if err == nil {
					csv += row
				}
			}
		}
	}
	//write csv to file
	file, err := os.Create(filename)
	if err == nil {
		_, _ = file.Write([]byte(csv))
		_ = file.Close()
	}
}
