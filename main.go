package main

import (
	"encoding/json"
	"fmt"
	"haggle/models"
	"haggle/parsers"
	"log"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/gocolly/colly"
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
var db *gorm.DB

func main() {
	db := getDb()
	defer db.Close()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/data", dataLink)
	router.HandleFunc("/scrape", scrapeLink)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./build/")))

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"http://localhost:3000"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	//initiate an initial scrape
	//_, _ =scrapeAll()

	// start server listen
	log.Fatal(http.ListenAndServe(":8088", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}

func getDb() *gorm.DB {
	if db != nil {
		return db
	}
	var err error
	db, err = gorm.Open("mysql", "root:123@(localhost:6603)/haggle?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(&models.Event{}, &models.Market{}, &models.Selection{})

	return db
}

func scrapeAll() (map[string]string, error) {
	result := make(map[string]string)
	if res, err := scrapeSite(models.ParseSiteConfig("stoiximan")); res {
		result["stoiximan"] = "ok"
	} else {
		result["stoiximan"] = err.Error()
	}
	if res, err := scrapeSite(models.ParseSiteConfig("pokerstars")); res {
		result["pokerstars"] = "ok"
	} else {
		result["pokerstars"] = err.Error()
	}
	return result, nil
}

func getCollector() *colly.Collector {
	c := colly.NewCollector(
		//colly.CacheDir("./_instagram_cache/"),
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
	)

	// Limit the number of threads started by colly to one
	// slow but make sure we don't get banned.
	// can be limited to domain
	_ = c.Limit(&colly.LimitRule{
		Parallelism: 1,
		Delay:       3 * time.Second,
	})
	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})

	c.OnError(func(r *colly.Response, e error) {
		log.Println("error:", e, r.Request.URL, string(r.Body))
	})
	return c
}

func getParser(id string) parsers.Parser {
	switch id {
	case `stoiximan`:
		return parsers.Stoiximan{}
	//case `bet365`:
	//	return parsers.Bet365{}
	//case `novibet`:
	//	return parsers.Novibet{}
	case `pokerstars`:
		return parsers.PokerStars{}
	}
	return nil
}

func scrapeSite(config models.SiteConfig) (bool, error) {
	if !config.Active {
		return false, fmt.Errorf("Parser disabled")
	}
	c := getCollector()
	parser := getParser(config.Id)
	if parser != nil {
		parser.SetDB(db)

		// Before making a request put the URL with
		// the key of "url" into the context of the request
		c.OnRequest(func(r *colly.Request) {
			r.Ctx.Put("path", r.URL.String())
		})
		parser.ScrapeHome()
		parser.ScrapeLive()
		parser.ScrapeToday()
		for t := 0; t < len(config.Tournaments); t++ {
			parser.ScrapeTournament(t)
		}

		return parser.Scrape(config, c)
	}
	return false, fmt.Errorf("Parser not found")
}

func getScrapeResults() (map[string]interface{}, error) {
	events := make([]models.Event, 0)
	getDb().Preload("Markets").Preload("Markets.Selections").Find(&events)
	return map[string]interface{}{
		"events": events,
	}, nil
}

func scrapeLink(w http.ResponseWriter, r *http.Request) {
	result, err := scrapeAll()
	if err != nil {
		_ = json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}
	_ = json.NewEncoder(w).Encode(result)
}

func dataLink(w http.ResponseWriter, r *http.Request) {
	result, err := getScrapeResults()
	if err != nil {
		_ = json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}
	_ = json.NewEncoder(w).Encode(result)
}
