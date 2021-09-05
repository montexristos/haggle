package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"haggle/models"
	"haggle/parsers"
	"log"
	"net/http"
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

func (app *Application) scrapeAll() (map[string]string, error) {
	result := make(map[string]string)
	if res, err := app.ScrapeSite("stoiximan"); res {
		result["stoiximan"] = "ok"
	} else {
		result["stoiximan"] = err.Error()
	}
	if res, err := app.ScrapeSite("novibet"); res {
		result["novibet"] = "ok"
	} else {
		result["novibet"] = err.Error()
	}
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
	}
	if parser != nil {
		parser.SetDB(db)
		parser.SetConfig(config)
		parser.Initialize()
		return parser, nil
	}
	return nil, fmt.Errorf("parser not found")
}

func (app *Application) ScrapeSite(website string) (bool, error) {
	//parser := GetParser(website, app.db)
	parser, err := GetParser(website, app.db)
	if parser != nil {
		_, err = parser.ScrapeHome()
		_, err = parser.ScrapeLive()
		_, err = parser.ScrapeToday()
		config := parser.GetConfig()
		for t := 0; t < len(parser.GetConfig().Tournaments); t++ {
			_, err = parser.ScrapeTournament(config.Tournaments[t])
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
	rows, err := GetDb().Raw(`SELECT betradar_id FROM (SELECT count(distinct site_id) as matches, betradar_id FROM haggle.events group by betradar_id) as tab WHERE matches > 1;`).Rows()
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
	GetDb().Preload("Markets").Preload("Markets.Selections").Where("betradar_id in (?)", eventIds).Find(&events)
	matches := make(map[int][]models.Event)
	for _, event := range events {
		var matchResultMarket models.Market
		var overUnderMarket models.Market
		var bttsMarket models.Market
		if _, found := matches[event.BetradarID]; !found {
			matches[event.BetradarID] = make([]models.Event, 0)
		}
		//iterate event markets and order them
		for _, v := range event.Markets {
			if v.MarketType == models.NewMatchResult().Name {
				matchResultMarket = v
			}
			if v.MarketType == models.NewOverUnder().Name {
				overUnderMarket = v
			}
			if v.MarketType == models.NewBtts().Name {
				bttsMarket = v
			}
		}
		event.Markets = make([]models.Market, 0)
		event.Markets = append(event.Markets, matchResultMarket)
		event.Markets = append(event.Markets, overUnderMarket)
		event.Markets = append(event.Markets, bttsMarket)
		matches[event.BetradarID] = append(matches[event.BetradarID], event)
	}

	return map[string]interface{}{
		"events": matches,
	}, nil
}

func (app *Application) scrapeLink(w http.ResponseWriter, r *http.Request) {
	GetDb().Raw(`SET FOREIGN_KEY_CHECKS=0;TRUNCATE TABLE haggle.selections;TRUNCATE TABLE haggle.markets;TRUNCATE TABLE haggle.events;SET FOREIGN_KEY_CHECKS=1;`)
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
