package fixtureModels

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
)

// Fixture the event model
type Fixture struct {
	Date              time.Time `json:"date"`
	HomeTeam          *Team     `json:"homeTeam"`
	AwayTeam          *Team     `json:"awayTeam"`
	HomeTeamName      string    `json:"homeTeamName"`
	AwayTeamName      string    `json:"awayTeamName"`
	Score             string    `json:"score"`
	Background        string    `json:"omit"`
	HomeScored        int       `json:"-"`
	HomeConceided     int       `json:"-"`
	AwayScored        int       `json:"-"`
	AwayConceided     int       `json:"-"`
	HomeAwayScored    int       `json:"-"`
	HomeAwayConceided int       `json:"-"`
	AwayHomeScored    int       `json:"-"`
	AwayHomeConceided int       `json:"-"`
	HomeCorners       int       `json:"homeCorners"`
	AwayCorners       int       `json:"awayCorners"`
	HomeYellowCards   int       `json:"homeYellowCards"`
	AwayYellowCards   int       `json:"awayYellowCards"`
	HomeRedCards      int       `json:"homeRedCards"`
	AwayRedCards      int       `json:"awayRedCards"`
	HomeIndex         float64   `json:"homeIndex"`
	AwayIndex         float64   `json:"awayIndex"`
	TotalIndex        float64   `json:"totalIndex"`
	PredictedOver0    int       `json:"predictedOver0"`
	PredictedOver1    int       `json:"predictedOver1"`
	PredictedOver2    int       `json:"predictedOver2"`
	PredictedOver3    int       `json:"predictedOver3"`
	PredictedOver4    int       `json:"predictedOver4"`
	PredictedUnder0   int       `json:"predictedUnder0"`
	PredictedUnder1   int       `json:"predictedUnder1"`
	PredictedUnder2   int       `json:"predictedUnder2"`
	PredictedUnder3   int       `json:"predictedUnder3"`
	PredictedUnder4   int       `json:"predictedUnder4"`
	HomeOdd           float32   `json:"homeOdd"`
	DrawOdd           float32   `json:"drawOdd"`
	AwayOdd           float32   `json:"awayOdd"`
	OverOdd           float32   `json:"overOdd"`
	UnderOdd          float32   `json:"underOdd"`
	Over0             int       `json:"over0"`
	Over1             int       `json:"over1"`
	Over2             int       `json:"over2"`
	Over3             int       `json:"over3"`
	Over4             int       `json:"over4"`
	Under0            int       `json:"under0"`
	Under1            int       `json:"under1"`
	Under2            int       `json:"under2"`
	Under3            int       `json:"under3"`
	Under4            int       `json:"under4"`
	CornerClass       string    `json:"cornerClass"`
	CardClass         string    `json:"cardClass"`
}

type FixtureJson struct {
	Date            time.Time `json:"date"`
	HomeTeamName    string    `json:"homeTeamName"`
	AwayTeamName    string    `json:"awayTeamName"`
	Score           string    `json:"score"`
	HomeCorners     int       `json:"homeCorners"`
	AwayCorners     int       `json:"awayCorners"`
	HomeYellowCards int       `json:"homeYellowCards"`
	AwayYellowCards int       `json:"awayYellowCards"`
	HomeRedCards    int       `json:"homeRedCards"`
	AwayRedCards    int       `json:"awayRedCards"`
}

func (f *Fixture) Export() FixtureJson {
	return FixtureJson{
		Date:            f.Date,
		HomeTeamName:    f.HomeTeamName,
		AwayTeamName:    f.AwayTeamName,
		Score:           f.Score,
		HomeCorners:     f.HomeCorners,
		AwayCorners:     f.AwayCorners,
		HomeYellowCards: f.HomeYellowCards,
		AwayYellowCards: f.AwayYellowCards,
		HomeRedCards:    f.HomeRedCards,
		AwayRedCards:    f.AwayRedCards,
	}
}

type EventListDTO struct {
	Events []FixtureDTO `json:"events"`
}

// FixtureDTO the fixture transfer model
type FixtureDTO struct {
	Date     string `json:"dateEvent"`
	HomeTeam string `json:"strHomeTeam"`
	AwayTeam string `json:"strAwayTeam"`
	Time     string `json:"strTime"`
}

type FixtureFD struct {
	Count   int `json:"count"`
	Filters struct {
	} `json:"filters"`
	Competition struct {
		ID   int `json:"id"`
		Area struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"area"`
		Name        string    `json:"name"`
		Code        string    `json:"code"`
		Plan        string    `json:"plan"`
		LastUpdated time.Time `json:"lastUpdated"`
	} `json:"competition"`
	Matches []struct {
		ID     int `json:"id"`
		Season struct {
			ID              int    `json:"id"`
			StartDate       string `json:"startDate"`
			EndDate         string `json:"endDate"`
			CurrentMatchday int    `json:"currentMatchday"`
		} `json:"season"`
		UtcDate     time.Time `json:"utcDate"`
		Status      string    `json:"status"`
		Matchday    int       `json:"matchday"`
		Stage       string    `json:"stage"`
		Group       string    `json:"group"`
		LastUpdated time.Time `json:"lastUpdated"`
		Odds        struct {
			Msg string `json:"msg"`
		} `json:"odds"`
		Score struct {
			Winner   string `json:"winner"`
			Duration string `json:"duration"`
			FullTime struct {
				HomeTeam int `json:"homeTeam"`
				AwayTeam int `json:"awayTeam"`
			} `json:"fullTime"`
			HalfTime struct {
				HomeTeam int `json:"homeTeam"`
				AwayTeam int `json:"awayTeam"`
			} `json:"halfTime"`
			ExtraTime struct {
				HomeTeam interface{} `json:"homeTeam"`
				AwayTeam interface{} `json:"awayTeam"`
			} `json:"extraTime"`
			Penalties struct {
				HomeTeam interface{} `json:"homeTeam"`
				AwayTeam interface{} `json:"awayTeam"`
			} `json:"penalties"`
		} `json:"score"`
		HomeTeam struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"homeTeam"`
		AwayTeam struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"awayTeam"`
		Referees []struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Role        string `json:"role"`
			Nationality string `json:"nationality"`
		} `json:"referees"`
	} `json:"matches"`
}

func (fixture Fixture) Normalize(goals int) float64 {
	//goalsMin := 0
	//goalsMax := 5
	if goals > 4 {
		goals = 5
	}
	return float64(goals)
	//return float64(goals-goalsMin) / float64(goalsMax-goalsMin)
}

func (f *Fixture) GenerateRow(upcoming bool) (string, error) {
	if f.HomeTeam.NoEvents < 10 || f.AwayTeam.NoEvents < 10 {
		return "", fmt.Errorf("few events")
	}
	homeOver := float64(f.HomeTeam.Over2) / float64(f.HomeTeam.NoEvents)
	homeGG := float64(f.HomeTeam.Gg) / float64(f.HomeTeam.NoEvents)
	awayOver := float64(f.AwayTeam.Over2) / float64(f.AwayTeam.NoEvents)
	awayGG := float64(f.AwayTeam.Gg) / float64(f.AwayTeam.NoEvents)

	homeScored := float64(f.HomeTeam.HomeGoalsScored) / float64(f.HomeTeam.HomeEvents)
	awayScored := float64(f.AwayTeam.AwayGoalsScored) / float64(f.AwayTeam.AwayEvents)

	if f.Score == "" && upcoming {
		result := "yes"
		return fmt.Sprintf("%.2f, %.2f, %.2f, %.2f,%.2f,%.2f,%s\r\n", homeOver, homeGG, awayOver, awayGG, homeScored, awayScored, result), nil
	}
	Scores := strings.Split(f.Score, " - ")
	homeScore, _ := strconv.Atoi(Scores[0])
	awayScore, _ := strconv.Atoi(Scores[1])
	score := homeScore + awayScore
	isGGandOver := homeScore > 0 && awayScore > 0 && score > 2
	result := "no"
	if isGGandOver {
		result = "yes"
	}
	return fmt.Sprintf("%.2f,%.2f,%.2f,%.2f,%.2f,%.2f,%s\r\n", homeOver, homeGG, awayOver, awayGG, homeScored, awayScored, result), nil
}

func (f *Fixture) GenerateRowOver(upcoming bool) (string, error) {

	if f.HomeTeam.HomeEvents < 10 || f.AwayTeam.AwayEvents < 10 {
		return "", fmt.Errorf("few events")
	}

	homeOver := float64(f.HomeTeam.Over2) / float64(f.HomeTeam.NoEvents)
	homeGG := float64(f.HomeTeam.Gg) / float64(f.HomeTeam.NoEvents)
	awayOver := float64(f.AwayTeam.Over2) / float64(f.AwayTeam.NoEvents)
	awayGG := float64(f.AwayTeam.Gg) / float64(f.AwayTeam.NoEvents)
	homeScored := float64(f.HomeTeam.HomeGoalsScored) / float64(f.HomeTeam.HomeEvents)
	awayScored := float64(f.AwayTeam.AwayGoalsScored) / float64(f.AwayTeam.AwayEvents)
	if f.Score == "" && upcoming {
		return fmt.Sprintf("%.2f,%.2f,%.2f,%.2f,%.2f,%.2f,%.2f\r\n", homeOver, homeGG, awayOver, awayGG, homeScored, awayScored, 4.00), nil
	}
	Scores := strings.Split(f.Score, " - ")
	homeScore, _ := strconv.Atoi(Scores[0])
	awayScore, _ := strconv.Atoi(Scores[1])
	score := homeScore + awayScore
	return fmt.Sprintf("%.2f,%.2f,%.2f,%.2f,%.2f,%.2f,%.2f\r\n", homeOver, homeGG, awayOver, awayGG, homeScored, awayScored, float64(score)), nil
}

func (f *Fixture) CalculateMetrics(homeOp int, awayOp int) (float64, float64, float64) {
	//generate test suite by saving fixture in file
	//var err error
	//homeIndex, err := f.calculateHomeIndex(f.HomeTeam)
	//awayIndex, err := f.calculateAwayIndex(f.AwayTeam)
	//if err != nil {
	//	return 0, 0, 0
	//}
	//totalIndex := homeIndex*0.40 + awayIndex*0.40 + float64(homeOp)*10.0 + float64(awayOp)*10.0
	//return totalIndex, homeIndex, awayIndex

	if f.HomeTeam.HomeEvents > 0 && f.AwayTeam.AwayEvents > 0 {
		weights := map[string]float64{
			"homeScoreMatch": 20,
			"awayScoreMatch": 20,
			"scoreIndex":     20,
			"conceidedIndex": 20,
			"homeOp":         5,
			"awayOp":         5,
		}

		homeScoreIndex := f.getMatchesScoredIndex(f.HomeTeam.HomeGoalsScored, f.HomeTeam.HomeEvents, f.HomeTeam.HomeMatchesScored)
		awayScoreIndex := f.getMatchesScoredIndex(f.AwayTeam.AwayGoalsScored, f.AwayTeam.AwayEvents, f.AwayTeam.AwayMatchesScored)
		//homeScoreIndex := float64(f.HomeTeam.HomeMatchesScored) / float64(f.HomeTeam.HomeEvents)
		awayConceidedIndex := float64(f.AwayTeam.AwayMatchesConceided) / float64(f.AwayTeam.AwayEvents)
		//awayScoreIndex := float64(f.AwayTeam.AwayMatchesScored) / float64(f.AwayTeam.AwayEvents)
		homeConceidedIndex := float64(f.HomeTeam.HomeMatchesConceided) / float64(f.HomeTeam.HomeEvents)

		return homeScoreIndex*.35 + awayScoreIndex*.35 + awayConceidedIndex*.15 + homeConceidedIndex*.15, homeScoreIndex * .50, awayScoreIndex * .50
		//return (homeScoreIndex + awayScoreIndex) * 100 / 4, homeScoreIndex * 50, awayScoreIndex * 50

		scoreIndex := (f.HomeTeam.HomeGoalsScored/float64(f.HomeTeam.HomeEvents) + f.AwayTeam.AwayGoalsScored/float64(f.AwayTeam.AwayEvents)) / 4
		conceidedIndex := (f.HomeTeam.HomeGoalsConceided/float64(f.HomeTeam.HomeEvents) + f.AwayTeam.AwayGoalsConceided/float64(f.AwayTeam.AwayEvents)) / 4
		return homeScoreIndex*weights["homeScoreMatch"] +
			awayScoreIndex*weights["awayScoreMatch"] +
			scoreIndex*weights["scoreIndex"] +
			conceidedIndex*weights["conceidedIndex"] +
			float64(homeOp)*weights["homeOp"] +
			float64(awayOp)*weights["awayOp"], 10 * scoreIndex, 10 * conceidedIndex
	}
	return 0.0, 0.0, 0.0
}

func (f *Fixture) calculateHomeIndex(t *Team) (float64, error) {
	homeScoredWeight := 1.5
	homeConceidedWeight := 2.0
	awayScoredWeight := .5
	awayConceidedWeight := .5
	return f.calculateIndex(homeScoredWeight, homeConceidedWeight, awayScoredWeight, awayConceidedWeight, t)
}

func (f *Fixture) calculateAwayIndex(t *Team) (float64, error) {
	homeScoredWeight := 0.5
	homeConceidedWeight := 0.5
	awayScoredWeight := 2.0
	awayConceidedWeight := 1.5
	return f.calculateIndex(homeScoredWeight, homeConceidedWeight, awayScoredWeight, awayConceidedWeight, t)
}

func (f *Fixture) calculateIndex(
	homeScoredWeight, homeConceidedWeight, awayScoredWeight, awayConceidedWeight float64,
	t *Team,
) (float64, error) {

	homeGoalsScored, homeGoalsConceided, awayGoalsScored, awayGoalsConceided := 0.0, 0.0, 0.0, 0.0
	if t.HomeEvents > 0 {
		homeGoalsScored = f.getGoalIndex(t.HomeGoalsScored, t.HomeEvents, t.HomeMatchesScored)
		homeGoalsConceided = f.getGoalIndex(t.HomeGoalsConceided, t.HomeEvents, t.HomeMatchesConceided)
	}
	if t.AwayEvents > 0 {
		awayGoalsScored = f.getGoalIndex(t.AwayGoalsScored, t.AwayEvents, t.AwayMatchesScored)
		awayGoalsConceided = f.getGoalIndex(t.AwayGoalsConceided, t.AwayEvents, t.AwayMatchesConceided)
	}
	sumWeights := homeScoredWeight + homeConceidedWeight + awayScoredWeight + awayConceidedWeight
	sum := homeGoalsScored*homeScoredWeight + homeGoalsConceided*homeConceidedWeight + awayGoalsScored*awayScoredWeight + awayGoalsConceided*awayConceidedWeight
	div := sumWeights * homeScoredWeight * homeConceidedWeight * awayScoredWeight * awayConceidedWeight
	if div == 0 || homeGoalsScored == 0 || awayGoalsScored == 0 {
		return 0, fmt.Errorf("no events")
	}
	return sum / div, nil
}

func (f *Fixture) getGoalIndex(goals float64, events int, matches int) float64 {
	//mult := 100.0 / float64(events)
	//abs := math.Abs(float64(3*events)-goals) / float64(events*3)
	abs := math.Abs(float64(matches) / float64(events))
	return math.Round(100 - (100 * abs))
}

func (f *Fixture) getMatchesScoredIndex(goals float64, events int, matches int) float64 {
	matchesScored := 100 * float64(matches) / float64(events)
	index := 0.0
	if matchesScored > 90 {
		index = 90.0
	} else if matchesScored > 70 {
		index = 70.0
	} else if matchesScored > 50 {
		index = 50.0
	}

	return index + math.Round(goals/float64(events))
}

func (f *Fixture) getOp(t1 *Team, t2 *Team) int {
	pointSpread1 := float64(t1.Points) / float64(t1.HomeEvents+t1.AwayEvents)
	pointSpread2 := float64(t2.Points) / float64(t2.HomeEvents+t2.AwayEvents)
	if t1.Points > 10 && t2.Points > 10 && pointSpread1-pointSpread2 > 0.7 && math.Abs(float64(t1.NoEvents-t2.NoEvents)) < 3 {
		return 1
	}
	return 0
}

// ProcessFixtures calculates fixture metadata
func ProcessFixtures(tournamentList *[]*Tournament, stats *Stats, week Week, calculateUntil time.Time) {
	var teamNames map[string]interface{}
	yamlFile, err := ioutil.ReadFile("config/names.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &teamNames)
	if err != nil {
		log.Printf("yamlFile.parse err   #%v ", err)
	}
	for _, v := range *tournamentList {
		fixtures := make([]*Fixture, 0)
		for _, fixture := range v.Fixtures {
			//match teams from tournament
			fixture.HomeTeam = v.GetTeam(fixture.HomeTeamName, &teamNames)
			fixture.AwayTeam = v.GetTeam(fixture.AwayTeamName, &teamNames)
			if fixture.Date.Before(calculateUntil) && fixture.HomeTeam != nil && fixture.AwayTeam != nil {
				fixture.ProcessFixture(fixture.HomeTeam, fixture.AwayTeam, stats)
			} else {
				//fmt.Println("event skipped")
			}

			if fixture.Date.After(week.Start) && fixture.Date.Before(week.End) {
				fixture.calculate(stats)
				fixtures = append(fixtures, fixture)
			}
		}
		v.Fixtures = fixtures
		//for _, t := range v.Teams {
		//	t.Fixtures = make([]FixtureJson, 0)
		//}
	}
	//write team names back to file
	out, err := yaml.Marshal(teamNames)
	_ = ioutil.WriteFile("config/names.yaml", out, os.FileMode(755))
}

func (fixture *Fixture) ProcessFixture(homeTeam, awayTeam *Team, stats *Stats) {
	homeTeam.NoEvents++
	awayTeam.NoEvents++
	homeTeam.HomeEvents++
	awayTeam.AwayEvents++
	homeTeam.Fixtures = append(homeTeam.Fixtures, fixture.Export())
	awayTeam.Fixtures = append(awayTeam.Fixtures, fixture.Export())
	homeTeam.GoalsScored += fixture.Normalize(fixture.HomeScored)
	homeTeam.GoalsConceided += fixture.Normalize(fixture.AwayScored)
	homeTeam.HomeGoalsScored += fixture.Normalize(fixture.HomeScored)
	homeTeam.HomeGoalsConceided += fixture.Normalize(fixture.AwayScored)
	homeTeam.HomeCorners += fixture.HomeCorners
	homeTeam.HomeAwayCorners += fixture.AwayCorners
	homeTeam.Corners += fixture.HomeCorners + fixture.AwayCorners
	homeTeam.YellowCards += fixture.HomeYellowCards
	homeTeam.RedCards += fixture.HomeRedCards
	homeTeam.TotalCards += (fixture.HomeYellowCards + 2*fixture.HomeRedCards + fixture.AwayYellowCards + 2*fixture.AwayRedCards)

	awayTeam.GoalsScored += fixture.Normalize(fixture.AwayScored)
	awayTeam.GoalsConceided += fixture.Normalize(fixture.AwayConceided)
	awayTeam.AwayGoalsScored += fixture.Normalize(fixture.AwayScored)
	awayTeam.AwayGoalsConceided += fixture.Normalize(fixture.AwayConceided)
	awayTeam.AwayCorners += fixture.AwayCorners
	awayTeam.AwayHomeCorners += fixture.HomeCorners
	awayTeam.Corners += fixture.HomeCorners + fixture.AwayCorners
	awayTeam.YellowCards += fixture.AwayYellowCards
	awayTeam.RedCards += fixture.AwayRedCards
	awayTeam.TotalCards += (fixture.HomeYellowCards + 2*fixture.HomeRedCards + fixture.AwayYellowCards + 2*fixture.AwayRedCards)

	if fixture.HomeScored > fixture.AwayScored {
		homeTeam.Points += 3
	} else if fixture.AwayScored == fixture.HomeScored {
		awayTeam.Points++
		homeTeam.Points++
	} else if fixture.HomeScored < fixture.AwayScored {
		awayTeam.Points += 3
	}
	if fixture.HomeScored+fixture.AwayScored > 0 {
		homeTeam.Over0++
		awayTeam.Over0++
		stats.Over0++
	}
	if fixture.HomeScored+fixture.AwayScored > 1 {
		homeTeam.Over1++
		awayTeam.Over1++
		stats.Over1++
	} else {
		homeTeam.Under1++
		awayTeam.Under1++
		stats.Under1++
	}
	if fixture.HomeScored+fixture.AwayScored > 2 {
		homeTeam.Over2++
		awayTeam.Over2++
		stats.Over2++
	} else {
		homeTeam.Under2++
		awayTeam.Under2++
		stats.Under2++
	}
	if fixture.HomeScored+fixture.AwayScored > 3 {
		homeTeam.Over3++
		awayTeam.Over3++
		stats.Over3++
	} else {
		homeTeam.Under3++
		awayTeam.Under3++
		stats.Under3++
	}
	if fixture.HomeScored+fixture.AwayScored > 4 {
		homeTeam.Over4++
		awayTeam.Over4++
		stats.Over4++
	} else {
		homeTeam.Under4++
		awayTeam.Under4++
		stats.Under4++
	}
	if fixture.HomeScored > 0 && fixture.AwayScored > 0 {
		homeTeam.Gg++
		awayTeam.Gg++
		homeTeam.HomeGg++
		awayTeam.AwayGg++
	}
	if fixture.HomeScored > 0 {
		homeTeam.HomeMatchesScored++
		awayTeam.AwayMatchesConceided++
	}
	if fixture.AwayScored > 0 {
		homeTeam.HomeMatchesConceided++
		awayTeam.AwayMatchesScored++
	}
}

func (f *Fixture) calculate(stats *Stats) {
	cardIndex := 0
	cornerIndex := 0
	f.PredictedOver0 = -1
	f.PredictedOver1 = -1
	f.PredictedOver2 = -1
	f.PredictedOver3 = -1
	f.PredictedOver4 = -1
	f.PredictedUnder0 = -1
	f.PredictedUnder1 = -1
	f.PredictedUnder2 = -1
	f.PredictedUnder3 = -1
	f.PredictedUnder4 = -1
	homeNoEvents := 0
	awayNoEvents := 0
	cornerThreshold := 12
	cardThreshold := 5

	homeTeam := f.HomeTeam
	awayTeam := f.AwayTeam
	if homeTeam != nil && awayTeam != nil && homeTeam.HomeEvents > 0 && awayTeam.AwayEvents > 0 {
		homeNoEvents = homeTeam.HomeEvents + homeTeam.AwayEvents
		awayNoEvents = awayTeam.HomeEvents + awayTeam.AwayEvents
		//extra bonus for away goals* mult
		f.HomeCorners = (homeTeam.HomeCorners + homeTeam.HomeAwayCorners) / homeNoEvents
		f.AwayCorners = (awayTeam.AwayCorners + awayTeam.AwayHomeCorners) / awayNoEvents

		homeOp := f.getOp(homeTeam, awayTeam)
		awayOp := f.getOp(awayTeam, homeTeam)
		f.TotalIndex, f.HomeIndex, f.AwayIndex = f.CalculateMetrics(homeOp, awayOp)
		cornerIndex = int(math.Round(float64(cornerIndex)))
		if f.TotalIndex > stats.Threshold4 {
			f.PredictedOver4 = 1
			f.PredictedUnder4 = 0
		} else if f.TotalIndex <= stats.Threshold4 {
			f.PredictedOver4 = 0
			f.PredictedUnder4 = 1
		}
		if f.TotalIndex > stats.Threshold3 {
			f.PredictedOver3 = 1
			f.PredictedUnder3 = 0
		} else if f.TotalIndex <= stats.Threshold3 {
			f.PredictedOver3 = 0
			f.PredictedUnder3 = 1
		}
		if f.TotalIndex > stats.Threshold2 {
			f.PredictedOver2 = 1
			f.PredictedUnder2 = 0
		} else if f.TotalIndex <= stats.Threshold2 {
			f.PredictedOver2 = 0
			f.PredictedUnder2 = 1
		}
		if f.TotalIndex > stats.Threshold1 {
			f.PredictedOver1 = 1
			f.PredictedUnder1 = 0
		} else if f.TotalIndex <= stats.Threshold1 {
			f.PredictedOver1 = 0
			f.PredictedUnder1 = 1
		}
		if f.TotalIndex > stats.Threshold0 {
			f.PredictedOver0 = 1
			f.PredictedUnder0 = 0
		} else if f.TotalIndex <= stats.Threshold0 {
			f.PredictedOver0 = 0
			f.PredictedUnder0 = 1
		}
		if f.Score != "" {
			Scores := strings.Split(f.Score, " - ")
			homeScore, _ := strconv.Atoi(Scores[0])
			awayScore, _ := strconv.Atoi(Scores[1])
			score := homeScore + awayScore
			if score > 4 {
				f.Over4 = 1
				f.Under4 = -1
			} else {
				f.Over4 = -1
				f.Under4 = 1
			}
			if score > 3 {
				f.Over3 = 1
				f.Under3 = -1
			} else {
				f.Over3 = -1
				f.Under3 = 1
			}
			if score > 2 {
				f.Over2 = 1
				f.Under2 = -1
			} else {
				f.Over2 = -1
				f.Under2 = 1
			}
			if score > 1 {
				f.Over1 = 1
				f.Under1 = -1
			} else {
				f.Over1 = -1
				f.Under1 = 1
			}
			if score > 0 {
				f.Over0 = 1
			} else {
				f.Over0 = -1
			}
			if score > 4 && f.PredictedOver4 == 1 {
				stats.OverCorrect4++
			}
			if score > 4 && f.PredictedOver4 == 0 {
				stats.UnderWrong4++
			}
			if score <= 4 && f.PredictedOver4 == 1 {
				stats.OverWrong4++
			}
			if score <= 4 && f.PredictedOver4 == 0 {
				stats.UnderCorrect4++
			}
			if score > 3 && f.PredictedOver3 == 1 {
				stats.OverCorrect3++
			}
			if score > 3 && f.PredictedOver3 == 0 {
				stats.UnderWrong3++
			}
			if score <= 3 && f.PredictedOver3 == 1 {
				stats.OverWrong3++
			}
			if score <= 3 && f.PredictedOver3 == 0 {
				stats.UnderCorrect3++
			}
			if score > 2 && f.PredictedOver2 == 1 {
				stats.OverCorrect2++
			}
			if score > 2 && f.PredictedOver2 == 0 {
				stats.UnderWrong2++
			}
			if score <= 2 && f.PredictedOver2 == 1 {
				stats.OverWrong2++
			}
			if score <= 2 && f.PredictedOver2 == 0 {
				stats.UnderCorrect2++
			}
			if score > 1 && f.PredictedOver1 == 1 {
				stats.OverCorrect1++
			}
			if score > 1 && f.PredictedOver1 == 0 {
				stats.UnderWrong1++
			}
			if score <= 1 && f.PredictedOver1 == 1 {
				stats.OverWrong1++
			}
			if score <= 1 && f.PredictedOver1 == 0 {
				stats.UnderCorrect1++
			}
			if score > 0 && f.PredictedOver0 == 1 {
				stats.OverCorrect0++
			}
			if score <= 0 && f.PredictedOver0 == 1 {
				stats.OverWrong0++
			}
		}
		if cornerIndex > cornerThreshold {
			f.CornerClass = "cornerOver"
		} else {
			f.CornerClass = "cornerUnder"
		}
		fixtureCorners := f.HomeTeam.Corners + f.AwayTeam.Corners
		if fixtureCorners > 0 {
			if fixtureCorners > cornerThreshold && cornerIndex > cornerThreshold {
				f.CornerClass += " cornerCorrect"
				stats.CorrectCorner++
			}
			if fixtureCorners > cornerThreshold && cornerIndex <= cornerThreshold {
				f.CornerClass += " cornerWrong"
			}
			if fixtureCorners <= cornerThreshold && cornerIndex > cornerThreshold {
				f.CornerClass += " cornerWrong"
				stats.WrongCorner++
			}
			if fixtureCorners <= cornerThreshold && cornerIndex <= cornerThreshold {
				f.CornerClass += " cornerCorrect"
			}
		}
		cardIndex = int(math.Round(float64(cardIndex)))
		if cardIndex > cardThreshold {
			f.CardClass = "cardOver"
		} else {
			f.CardClass = "cardUnder"
		}
		fixtureCards := f.HomeTeam.YellowCards + f.HomeTeam.RedCards + f.AwayTeam.YellowCards + f.AwayTeam.RedCards
		if fixtureCards > 0 {
			if cardIndex > cardThreshold && fixtureCards > cardThreshold {
				f.CardClass += " cardCorrect"
				stats.CorrectCards++
			}
			if cardIndex > cardThreshold && fixtureCards <= cardThreshold {
				f.CardClass += " cardWrong"
				stats.WrongCards++
			}
			if cardIndex <= cardThreshold && fixtureCards > cardThreshold {
				f.CardClass += " cardWrong"
			}
			if cardIndex <= cardThreshold && fixtureCards <= cardThreshold {
				f.CardClass += " cardCorrect"
			}
		}
	}
}

// GetFixtures returns upcoming fixtureModels for a tournament
func GetFixtures(tournament *Tournament) ([]FixtureDTO, error) {
	url := fmt.Sprintf("https://www.thesportsdb.com/api/v1/json/1/eventsnextleague.php?id=%d", tournament.Id)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var fixtureList EventListDTO
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	respByte := buf.Bytes()
	if err := json.Unmarshal(respByte, &fixtureList); err != nil {
		return nil, err
	}
	return fixtureList.Events, nil
}

func MapFixtures(fixtureList []FixtureDTO, tournament *Tournament) []*Fixture {
	var fixtures []*Fixture
	today := time.Now()
	twoDaysAfter := today.AddDate(0, 0, 15)
	for _, fixtureItem := range fixtureList {
		dateObj, _ := time.Parse(`2006-01-02  15:04:05`, fmt.Sprintf("%s %s", fixtureItem.Date, fixtureItem.Time))
		if dateObj.After(twoDaysAfter) {
			continue
		}
		fixture := &Fixture{
			Date:         dateObj,
			HomeTeamName: fixtureItem.HomeTeam,
			AwayTeamName: fixtureItem.AwayTeam,
			Score:        "",
		}
		fixtures = append(fixtures, fixture)
	}

	return fixtures
}
