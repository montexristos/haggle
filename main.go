package main

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"haggle/models"
	"haggle/parsers"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)

/*
docker run \
--detach \
--name=mysql \
--env="MYSQL_ROOT_PASSWORD=123" \
--publish 6603:3306 \
--volume=data:/var/lib/mysql \
mysql
*/

func main() {
	db := GetDb()
	defer db.Close()
	app := Application{
		db,
	}
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/data", app.dataLink)
	router.HandleFunc("/scrape", app.scrapeLink)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./build/")))

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"http://localhost:3000"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	//initiate an initial scrape
	//_, _ =scrapeAll()

	// start server listen
	log.Fatal(http.ListenAndServe(":8088", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}

type Application struct {
	db *gorm.DB
}

func GetDb() *gorm.DB {
	//CREATE SCHEMA `haggle` DEFAULT CHARACTER SET utf8 ;
	db, err := gorm.Open("mysql", "root:123@(localhost:6603)/haggle?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(&models.Event{}, &models.Market{}, &models.Selection{})
	return db
}

func (app *Application) scrapeAll() (map[string]string, error) {
	result := make(map[string]string)
	if res, err := app.scrapeSite(models.ParseSiteConfig("stoiximan")); res {
		result["stoiximan"] = "ok"
	} else {
		result["stoiximan"] = err.Error()
	}
	if res, err := app.scrapeSite(models.ParseSiteConfig("pokerstars")); res {
		result["pokerstars"] = "ok"
	} else {
		result["pokerstars"] = err.Error()
	}
	return result, nil
}

func getParser(id string) parsers.Parser {
	switch id {
	case `stoiximan`:
		return &parsers.Stoiximan{}
	//case `bet365`:
	//	return parsers.Bet365{}
	case `novibet`:
		return &parsers.Novibet{}
	case `pokerstars`:
		return &parsers.PokerStars{}
	case `winmasters`:
		return &parsers.Winmasters{}
	}
	return nil
}

func (app *Application) scrapeSite(config *models.SiteConfig) (bool, error) {
	if !config.Active {
		return false, fmt.Errorf("Parser %s disabled", config.Id)
	}

	parser := getParser(config.Id)
	if parser != nil {
		parser.Initialize()
		parser.SetDB(app.db)
		parser.SetConfig(config)
		var err error
		_, err = parser.ScrapeHome()
		_, err = parser.ScrapeLive()
		_, err = parser.ScrapeToday()
		for t := 0; t < len(config.Tournaments); t++ {
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
	events := make([]models.Event, 0)
	GetDb().Preload("Markets").Preload("Markets.Selections").Find(&events)
	return map[string]interface{}{
		"events": events,
	}, nil
}

func (app *Application) scrapeLink(w http.ResponseWriter, r *http.Request) {
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
