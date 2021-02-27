package main

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"haggle/models"
	"haggle/parsers"
	"io/ioutil"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func setupParser(site string, db *gorm.DB) parsers.Parser {
	config := models.ParseSiteConfig(site)
	parser := getParser(config.Id)
	parser.SetDB(db)
	parser.SetConfig(config)
	parser.Initialize()
}

func Test_bet(t *testing.T) {

}

/**
TRUNCATE TABLE haggle.selections;
TRUNCATE TABLE haggle.markets;
TRUNCATE TABLE haggle.events;
*/

func Test_stoiximan(t *testing.T) {
	db := GetDb()
	defer db.Close()
	parser := setupParser("winmasters", db)
	config := models.ParseSiteConfig("stoiximan")
	if _, err := app.ScrapeSite(config); err != nil {
		t.Error(err.Error())
	}
}
func Test_winmastersParse(t *testing.T) {
	db := GetDb()
	defer db.Close()
	parser := setupParser("winmasters", db)

	//read file and parse
	file := "./test_input/winmasters/premierLeague.json"
	event, _ := ioutil.ReadFile(file)

	var parsed interface{}
	json.Unmarshal(event, &parsed)
	for key, value := range parsed.(map[string]interface{}) {
		if key == "events" {
			for i := 0; i < len(value.([]interface{})); i++ {
				parsers.ParseEvent(parser, value.([]interface{})[i].(map[string]interface{}))
			}
		}
		if key == "markets" {
			e := models.Event{}
			for i := 0; i < len(value.([]interface{})); i++ {
				parsers.ParseMarket(parser, value.([]interface{})[i].(map[string]interface{}), e)
			}
		}
	}
	fmt.Println(parsed)

}

func Test_novibet(t *testing.T) {
	db := GetDb()
	defer db.Close()
	app := Application{
		db: db,
	}
	config := models.ParseSiteConfig("novibet")
	if _, err := app.ScrapeSite(config); err != nil {
		t.Error(err.Error())
	}
}

func Test_pokerstars(t *testing.T) {

}
