package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/montexristos/haggle/database"
	"github.com/montexristos/haggle/models"
	"github.com/montexristos/haggle/parsers"
	"github.com/montexristos/haggle/scrape"
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
	if _, err := scrape.ScrapeSite("stoiximan"); err != nil {
		t.Error(err.Error())
	}
}
func Test_PameStoixima(t *testing.T) {
	if _, err := scrape.ScrapeSite("pamestoixima"); err != nil {
		t.Error(err.Error())
	}
}
func Test_winmastersParse(t *testing.T) {
	if _, err := scrape.ScrapeSite("winmasters"); err != nil {
		t.Error(err.Error())
	}
}

func Test_novibet(t *testing.T) {
	if _, err := scrape.ScrapeSite("novibet"); err != nil {
		t.Error(err.Error())
	}
}
func Test_netbet(t *testing.T) {
	if _, err := scrape.ScrapeSite("netbet"); err != nil {
		t.Error(err.Error())
	}
}
func Test_all_active(t *testing.T) {
	if _, err := scrape.ScrapeAll(); err != nil {
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
	if _, err := scrape.ScrapeSite("bwin"); err != nil {
		t.Error(err.Error())
	}
}

func Test_pokerstars(t *testing.T) {
	if _, err := scrape.ScrapeSite("pokerstars"); err != nil {
		t.Error(err.Error())
	}
}
func Test_betsson(t *testing.T) {
	if _, err := scrape.ScrapeSite("betsson"); err != nil {
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
			database.ClearDB()
			rows, _ := database.GetDb().Raw(`SELECT * FROM haggle.events;`).Rows()
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
	resp, _ := scrape.GetScrapeResults()
	if len(resp["arbs"].(map[string]string)) > 0 {
		t.Log("found")
	}

}
func Test_getAllResults(t *testing.T) {
	resp := scrape.AllResults()
	if len(resp["arbs"].(map[string]string)) > 0 {
		t.Log("found")
	}

}
func Test_getArbDetect(t *testing.T) {
	tests := []struct {
		odd1   float64
		odd2   float64
		result float64
	}{
		{1.2, 1.3, 0.1},
		{1.2, 1.5, 0.1},
		{2, 2.6, 0.1},
		{2.2, 2.3, 0.1},
		{1.01, 1.1, 0.1},
		{12, 13, 0.1},
	}
	for _, test := range tests {
		if res := scrape.TestArbs(test.odd1, test.odd2); res != test.result {
			t.Fail()
		}
	}
}

func Test_CheckMarket_differentIndex(t *testing.T) {
	temp := make(map[string][]scrape.SiteOdd)
	tests := []struct {
		eventName  string
		selection  models.Selection
		marketType string
		index      int
		temp       map[string][]scrape.SiteOdd
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

		result := scrape.CheckMarket(test.eventName, test.selection, test.marketType, test.index, temp, test.site)
		if result != "" {
			t.Error("should not find arb")
			t.Fail()
		}
	}
}
func Test_CheckMarket_sameOdds(t *testing.T) {
	temp := make(map[string][]scrape.SiteOdd)
	tests := []struct {
		eventName  string
		selection  models.Selection
		marketType string
		index      int
		temp       map[string][]scrape.SiteOdd
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

		result := scrape.CheckMarket(test.eventName, test.selection, test.marketType, test.index, temp, test.site)
		if result != "" {
			t.Error("should not find arb")
			t.Fail()
		}
	}
}
func Test_CheckMarket_diffOdds(t *testing.T) {
	temp := make(map[string][]scrape.SiteOdd)
	tests := []struct {
		eventName  string
		selection  models.Selection
		marketType string
		index      int
		temp       map[string][]scrape.SiteOdd
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

		result := scrape.CheckMarket(test.eventName, test.selection, test.marketType, test.index, temp, test.site)
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
