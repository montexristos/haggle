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

func Test_getScrapeResults(t *testing.T) {
	resp, _ := getScrapeResults()
	if len(resp["arbs"].(map[string]string)) > 0 {
		t.Log("found")
	}

}
func Test_getArbDetect(t *testing.T) {
	tests := []struct {
		odd1   float64
		odd2   float64
		result bool
	}{
		{1.2, 1.3, false},
		{1.2, 1.5, true},
		{2, 2.6, true},
		{2.2, 2.3, false},
		{1.01, 1.1, false},
		{12, 13, false},
	}
	for _, test := range tests {
		if res := testArbs(test.odd1, test.odd2); res != test.result {
			t.Fail()
		}
	}
}

func Test_CheckMarket_differentIndex(t *testing.T) {
	temp := make(map[string][]SiteOdd)
	tests := []struct {
		eventName  string
		selection  models.Selection
		marketType string
		index      int
		temp       map[string][]SiteOdd
		site       string
	}{
		{
			eventName: "test1",
			selection: models.Selection{
				Price:    1.3,
				Name:     "kaka",
				Line:     1.0,
				MarketID: 1,
			},
			marketType: "MRES",
			index:      1,
			temp:       temp,
			site:       "stoiximan",
		}, {
			eventName: "test1",
			selection: models.Selection{
				Price:    1.3,
				Name:     "kaka",
				Line:     1.0,
				MarketID: 1,
			},
			marketType: "MRES",
			index:      0,
			temp:       temp,
			site:       "novi",
		},
	}

	for _, test := range tests {

		result := CheckMarket(test.eventName, test.selection, test.marketType, test.index, temp, test.site)
		if result != "" {
			t.Error("should not find arb")
			t.Fail()
		}
	}
}
func Test_CheckMarket_sameOdds(t *testing.T) {
	temp := make(map[string][]SiteOdd)
	tests := []struct {
		eventName  string
		selection  models.Selection
		marketType string
		index      int
		temp       map[string][]SiteOdd
		site       string
	}{
		{
			eventName: "test1",
			selection: models.Selection{
				Price:    1.3,
				Name:     "kaka",
				Line:     1.0,
				MarketID: 1,
			},
			marketType: "MRES",
			index:      1,
			temp:       temp,
			site:       "stoiximan",
		}, {
			eventName: "test1",
			selection: models.Selection{
				Price:    1.3,
				Name:     "kaka",
				Line:     1.0,
				MarketID: 1,
			},
			marketType: "MRES",
			index:      1,
			temp:       temp,
			site:       "novi",
		},
	}

	for _, test := range tests {

		result := CheckMarket(test.eventName, test.selection, test.marketType, test.index, temp, test.site)
		if result != "" {
			t.Error("should not find arb")
			t.Fail()
		}
	}
}
func Test_CheckMarket_diffOdds(t *testing.T) {
	temp := make(map[string][]SiteOdd)
	tests := []struct {
		eventName  string
		selection  models.Selection
		marketType string
		index      int
		temp       map[string][]SiteOdd
		site       string
	}{
		{
			eventName: "test1",
			selection: models.Selection{
				Price:    1.3,
				Name:     "kaka",
				Line:     1.0,
				MarketID: 1,
			},
			marketType: "MRES",
			index:      1,
			temp:       temp,
			site:       "stoiximan",
		}, {
			eventName: "test1",
			selection: models.Selection{
				Price:    4,
				Name:     "kaka",
				Line:     1.0,
				MarketID: 1,
			},
			marketType: "MRES",
			index:      1,
			temp:       temp,
			site:       "novi",
		},
	}

	for index, test := range tests {

		result := CheckMarket(test.eventName, test.selection, test.marketType, test.index, temp, test.site)
		if index == 0 && result != "" {
			t.Error("should not find arb")
			t.Fail()
		}
		if index == 1 && result == "" {
			t.Error("should  find arb")
			t.Fail()
		}
	}
}
