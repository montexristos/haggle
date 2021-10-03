package main

import (
	"encoding/json"
	"fmt"
	"haggle/models"
	"haggle/parsers"
	"io/ioutil"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func Test_bet(t *testing.T) {

}

/**
SET FOREIGN_KEY_CHECKS=0;
TRUNCATE TABLE haggle.selections;
TRUNCATE TABLE haggle.markets;
TRUNCATE TABLE haggle.events;
SET FOREIGN_KEY_CHECKS=1;
*/

func Test_stoiximan(t *testing.T) {
	db := GetDb()
	app := Application{
		db,
	}
	if _, err := app.ScrapeSite("stoiximan"); err != nil {
		t.Error(err.Error())
	}
}
func Test_winmastersParse(t *testing.T) {
	db := GetDb()
	app := Application{
		db,
	}
	parser, _ := GetParser("winmasters", db)
	//read file and parse
	file := "./test_input/winmasters/premierLeague.json"
	event, _ := ioutil.ReadFile(file)

	if _, err := app.ScrapeSite("winmasters"); err != nil {
		t.Error(err.Error())
	}
	var parsed interface{}
	json.Unmarshal(event, &parsed)
	for key, value := range parsed.(map[string]interface{}) {
		if key == "events" {
			for i := 0; i < len(value.([]interface{})); i++ {
				_, _ = parsers.ParseEvent(parser, value.([]interface{})[i].(map[string]interface{}))
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
	app := Application{
		db: db,
	}
	if _, err := app.ScrapeSite("novibet"); err != nil {
		t.Error(err.Error())
	}
}
func Test_netbet(t *testing.T) {
	db := GetDb()
	app := Application{
		db: db,
	}
	if _, err := app.ScrapeSite("netbet"); err != nil {
		t.Error(err.Error())
	}
}
func Test_all_active(t *testing.T) {
	db := GetDb()
	app := Application{
		db: db,
	}
	if _, err := app.scrapeAll(); err != nil {
		t.Error(err.Error())
	}
}

func Test_transform(t *testing.T) {
	name := parsers.TransformName("Λεστερ / Μπερνλι")
	if name != "Lester - Mpernli" {
		t.Fail()
	}
	name = parsers.TransformName("ΝΠΣ Βόλος - Ατρόμητος")
	if name != "NPS Bolos - Atromhtos" {
		t.Error(fmt.Sprintf("name is not %s", name))
		t.Fail()
	}
}

func Test_bwin(t *testing.T) {
	db := GetDb()
	app := Application{
		db: db,
	}
	if _, err := app.ScrapeSite("bwin"); err != nil {
		t.Error(err.Error())
	}
}

func Test_pokerstars(t *testing.T) {

}
