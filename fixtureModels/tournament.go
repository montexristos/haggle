package fixtureModels

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

type TournamentConfig struct {
	Tournaments []*Tournament `yaml:"Tournaments"`
}

// Tournament the tournament model
type Tournament struct {
	Teams    []*Team    `json:"teams"`
	Name     string     `json:"name"`
	Fixtures []*Fixture `json:"fixtures"`
	Csv      string     `json:"csv"`
	Old      string     `json:"old"`
	Id       int        `json:"id"`
	Fd       string     `json:"fd"`
}

// MarshalJSON exports the tournament to a custom json object
func (t *Tournament) MarshalJSON() ([]byte, error) {
	fixtures := make([]interface{}, 0)
	for _, value := range t.Fixtures {
		fixtures = append(fixtures, value)
	}
	return json.Marshal(&struct {
		Name     string        `json:"name"`
		Teams    []*Team       `json:"teams"`
		Fixtures []interface{} `json:"fixtures"`
	}{
		Name:     t.Name,
		Teams:    t.Teams,
		Fixtures: fixtures,
	})
}

// Parse returns a new tournament
func (t *Tournament) Parse(tournamentDTO TournamentDTO) (*Tournament, error) {

	return t, nil
}

type TournamentDTO struct {
}

func (t *Tournament) GetTeam(name string, teamNames *map[string]interface{}) *Team {
	name = strings.TrimSpace(name)
	mappedName, found := (*teamNames)[name]
	if !found {
		//fmt.Printf("name not found: %s\r\n", name)
		//(*teamNames)[name] = name
		//mappedName = name
	}
	for _, team := range t.Teams {
		if team.Name == mappedName {
			return team
		}
	}
	if mappedName != nil {
		name = mappedName.(string)
	} else {
		if name != "" {
			fmt.Printf("%s: %s\r\n", name, name)
		}
	}
	newTeam := Team{
		Name:     name,
		Fixtures: make([]FixtureJson, 0),
	}
	//save a new team
	t.Teams = append(t.Teams, &newTeam)
	return &newTeam
}

//
//func() {
//	csvPath, _ := tournament.getCSVPath(tournament, false)
//	if cache && csvPath != "" {
//		if err := tools.DownloadFile(csvPath, tournament.Csv); err != nil {
//			log.Println("error downloading file " + tournament.Csv)
//		}
//		csvPath, err := getCSVPath(tournament, true)
//		if err != nil {
//			log.Println("error finding file " + tournament.Old)
//		} else {
//			if err := tools.DownloadFile(csvPath, tournament.Old); err != nil {
//				log.Printf("error downloading file %s error: %s", tournament.Old, err.Error())
//			}
//		}
//	}
//}

// GetFixtures returns upcoming fixtureModels for a tournament
func (t *Tournament) GetFixturesNew(startDate string, endDate string) ([]FixtureDTO, error) {
	var fixtureList EventListDTO
	upcomingPath, _ := t.getUpcomingPath(endDate)
	//read file and if not exists download from service

	jsonText, err := ioutil.ReadFile(upcomingPath)
	if err != nil && strings.Compare(endDate, time.Now().Format("2006-01-02")) > 0 {
		client := &http.Client{}
		url := fmt.Sprintf("https://api.football-data.org/v2/competitions/%s/matches?dateFrom=%s&dateTo=%s", t.Fd, startDate, endDate)
		r, _ := http.NewRequest("GET", url, nil)
		r.Header.Add("X-Auth-Token", "652dd0613b134e6c80a3cd172e17349d")
		r.Header.Add("Content-Type", "application/json")
		resp, err := client.Do(r)
		if err == nil {
			defer resp.Body.Close()
			buf := new(bytes.Buffer)
			_, _ = buf.ReadFrom(resp.Body)
			respByte := buf.Bytes()
			if strings.Contains(string(respByte), "errorCode") {
				return nil, fmt.Errorf("error in response")
			}
			_ = ioutil.WriteFile(upcomingPath, respByte, os.ModePerm)
			jsonText, err = ioutil.ReadFile(upcomingPath)
		}
		if err != nil {
			return nil, err
		}
	}
	var matches FixtureFD

	if err := json.Unmarshal(jsonText, &matches); err != nil {
		return nil, err
	}
	for _, match := range matches.Matches {
		fixtureList.Events = append(fixtureList.Events, FixtureDTO{
			Date:     match.UtcDate.Format("2006-01-02"),
			HomeTeam: match.HomeTeam.Name,
			AwayTeam: match.AwayTeam.Name,
			Time:     match.UtcDate.Format("15:04:05"),
		})
	}
	return fixtureList.Events, nil
}

func (tournament *Tournament) getUpcomingPath(endDate string) (string, error) {
	return fmt.Sprintf("data/%s-%s.json", tournament.Fd, endDate), nil
}
