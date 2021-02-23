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
TRUNCATE TABLE haggle.selections;
TRUNCATE TABLE haggle.markets;
TRUNCATE TABLE haggle.events;
*/

func Test_stoiximan(t *testing.T) {
	db := GetDb()
	defer db.Close()
	app := Application{
		db: db,
	}
	config := models.ParseSiteConfig("stoiximan")
	if _, err := app.scrapeSite(config); err != nil {
		t.Error(err.Error())
	}
}
func Test_winmastersParse(t *testing.T) {
	db := GetDb()
	defer db.Close()

	//read file and parse
	file := "./test_input/winmasters/premierLeague.json"
	event, _ := ioutil.ReadFile(file)
	//
	//f, err := os.OpenFile(``, os.O_RDONLY, os.ModePerm)
	//if err != nil {
	//	log.Fatalf("open file error: %v", err)
	//	return
	//}
	//event, _ := ioutil.ReadFile(file)
	//defer f.Close()
	//sc := bufio.NewScanner(f)
	//for sc.Scan() {
	//	//jsonText := sc.Text()
	//	var parsed interface{}
	//	json.Unmarshal(sc.Bytes(), &parsed)
	//	fmt.Println(parsed)
	//}
	parser := parsers.Winmasters{}
	parser.SetDB(db)
	var parsed interface{}
	json.Unmarshal(event, &parsed)
	for key, value := range parsed.(map[string]interface{}) {
		if key == "events" {
			for i := 0; i < len(value.([]interface{})); i++ {
				parser.ParseEvent(value.([]interface{})[i].(map[string]interface{}))
			}
		}
		if key == "markets" {
			for i := 0; i < len(value.([]interface{})); i++ {
				parser.ParseMarket(value.([]interface{})[i].(map[string]interface{}))
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
	if _, err := app.scrapeSite(config); err != nil {
		t.Error(err.Error())
	}
}

func Test_pokerstars(t *testing.T) {

}
