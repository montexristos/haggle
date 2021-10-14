package fixtureModels

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestFixture_calculateMetrics(t *testing.T) {

	f := &Fixture{
		HomeTeam: &Team{
			Name:               "Leverkusen",
			GoalsScored:        77,
			GoalsConceided:     53,
			HomeGoalsScored:    40,
			HomeGoalsConceided: 27,
			AwayGoalsScored:    37,
			AwayGoalsConceided: 26,
			NoEvents:           42,
			HomeEvents:         20,
			AwayEvents:         22,
			Points:             81,
			Corners:            418,
			HomeCorners:        150,
			HomeAwayCorners:    56,
			AwayCorners:        112,
			AwayHomeCorners:    100,
			YellowCards:        75,
			RedCards:           5,
			TotalCards:         170,
			Gg:                 16,
			HomeGg:             9,
			AwayGg:             7,
		},
		AwayTeam: &Team{
			Name:               "Hertha",
			GoalsScored:        63,
			GoalsConceided:     77,
			HomeGoalsScored:    27,
			HomeGoalsConceided: 43,
			AwayGoalsScored:    36,
			AwayGoalsConceided: 34,
			NoEvents:           42,
			HomeEvents:         21,
			AwayEvents:         21,
			Points:             48,
			Corners:            383,
			HomeCorners:        83,
			HomeAwayCorners:    95,
			AwayCorners:        78,
			AwayHomeCorners:    127,
			YellowCards:        94,
			RedCards:           5,
			TotalCards:         182,
			Gg:                 16,
			HomeGg:             9,
			AwayGg:             7,
		},
		HomeTeamName: "Leverkusen",
		AwayTeamName: "Hertha",
		//HomeScored:        tt.fields.HomeScored,
		//HomeConceided:     tt.fields.HomeConceided,
		//AwayScored:        tt.fields.AwayScored,
		//AwayConceided:     tt.fields.AwayConceided,
		//HomeAwayScored:    tt.fields.HomeAwayScored,
		//HomeAwayConceided: tt.fields.HomeAwayConceided,
		//AwayHomeScored:    tt.fields.AwayHomeScored,
		//AwayHomeConceided: tt.fields.AwayHomeConceided,
		//HomeCorners:       tt.fields.HomeCorners,
		//AwayCorners:       tt.fields.AwayCorners,
		//HomeYellowCards:   tt.fields.HomeYellowCards,
		//AwayYellowCards:   tt.fields.AwayYellowCards,
		//HomeRedCards:      tt.fields.HomeRedCards,
		//AwayRedCards:      tt.fields.AwayRedCards,
		//HomeIndex:         tt.fields.HomeIndex,
		//AwayIndex:         tt.fields.AwayIndex,
		//TotalIndex:        tt.fields.TotalIndex,
		HomeOdd:  1.72,
		DrawOdd:  4.0,
		AwayOdd:  4.5,
		OverOdd:  1.57,
		UnderOdd: 2.37,
	}
	homeOp := 1
	awayOp := 0
	got, got1, got2 := f.CalculateMetrics(homeOp, awayOp)
	want := 1.0
	want1 := 2.0
	want2 := 3.0
	if got != want {
		t.Errorf("CalculateMetrics() got = %v, want %v", got, want)
	}
	if got1 != want1 {
		t.Errorf("CalculateMetrics() got1 = %v, want %v", got1, want1)
	}
	if got2 != want2 {
		t.Errorf("CalculateMetrics() got2 = %v, want %v", got2, want2)
	}
}

func TestFixture_getGoalIndex(t *testing.T) {
	f := Fixture{}
	var want float64
	want = 100.0
	if got := f.getGoalIndex(60, 20, 3); got != want {
		t.Errorf("getGoalIndex() = %v, want %v", got, want)
	}
	want = 67.0
	if got := f.getGoalIndex(60, 30, 3); got != want {
		t.Errorf("getGoalIndex() = %v, want %v", got, want)
	}
	want = 0.0
	if got := f.getGoalIndex(0, 30, 3); got != want {
		t.Errorf("getGoalIndex() = %v, want %v", got, want)
	}
	want = 50.0
	if got := f.getGoalIndex(45, 30, 3); got != want {
		t.Errorf("getGoalIndex() = %v, want %v", got, want)
	}
}

func TestAllFixtures(t *testing.T) {
	files := getFiles()
	var fixture Fixture
	for _, i := range files {
		//open json file
		jsonText, _ := ioutil.ReadFile(i)
		err := json.Unmarshal(jsonText, &fixture)
		if err != nil {
			continue
		}
		fixture.CalculateMetrics(1, 2)
	}
}

func getFiles() []string {
	var files []string
	root := "../testdata/fixtures/"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	return files
}
