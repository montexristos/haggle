package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"haggle/models"
	"haggle/parsers"
	"io/ioutil"
	"testing"
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
func Test_PameStoixima(t *testing.T) {
	db := GetDb()
	app := Application{
		db,
	}
	if _, err := app.ScrapeSite("pamestoixima"); err != nil {
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

	db := GetDb()
	app := Application{
		db: db,
	}
	if _, err := app.ScrapeSite("pokerstars"); err != nil {
		t.Error(err.Error())
	}
}
func Test_betsson(t *testing.T) {

	db := GetDb()
	app := Application{
		db: db,
	}
	if _, err := app.ScrapeSite("betsson"); err != nil {
		t.Error(err.Error())
	}
}

func TestBetsson_MatchEventMarkets(t *testing.T) {
	p := &parsers.Betsson{}
	json, _ := ioutil.ReadFile("test_input/betsson/tournament.json")
	markets, err := p.MatchEventMarkets(json)
	if err != nil {
		t.Error(err.Error())
		t.Fail()
	}
	if len(markets) == 0 {

		t.Fail()
	}
}

func TestClearDB(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			"test1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ClearDB()
			rows, _ := GetDb().Raw(`SELECT * FROM haggle.events;`).Rows()
			defer rows.Close()
			if rows.Next() {
				t.Errorf("db should be empty")
				t.Fail()
			}
		})
	}
}

func Test_getTournamentUpcomingData(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "test1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getTournamentUpcomingData()
		})
	}
}
