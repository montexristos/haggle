package fixtureModels

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
)

// Event the event model
type Event struct {
	Tournament            string `json:"tournament"`
	Date                  string `json:"date"`
	Time                  string `json:"time"`
	HomeTeam              string
	AwayTeam              string
	FullTimeHomeTeamGoals int
	FullTimeAwayTeamGoals int
	HomeTeamCorners       int
	AwayTeamCorners       int
	HomeTeamYellowCards   int
	AwayTeamYellowCards   int
	HomeTeamRedCards      int
	AwayTeamRedCards      int
	HomeOdd               string
	DrawOdd               string
	AwayOdd               string
	Over2                 string
	Under2                string
}

func mapLabel(label string) string {

	m := map[string]string{
		"Div":        "Tournament",
		"Date":       "Date",
		"Time":       "Time",
		"HomeTeam":   "HomeTeam",
		"AwayTeam":   "AwayTeam",
		"Home":       "HomeTeam",
		"Away":       "AwayTeam",
		"FTHG":       "FullTimeHomeTeamGoals",
		"FTAG":       "FullTimeAwayTeamGoals",
		"HG":         "FullTimeHomeTeamGoals",
		"AG":         "FullTimeAwayTeamGoals",
		"FTR":        "fullTimeResult",
		"HTHG":       "halfTimeHomeTeamGoals",
		"HTAG":       "HalfTimeAwayTeamGoals",
		"HTR":        "HalfTimeResult",
		"Attendance": "CrowdAttendance",
		"Referee":    "MatchReferee",
		"HS":         "HomeTeamShots",
		"AS":         "AwayTeamShots",
		"HST":        "HomeTeamShotsonTarget",
		"AST":        "AwayTeamShotsonTarget",
		"HHW":        "HomeTeamHitWoodwork",
		"AHW":        "AwayTeamHitWoodwork",
		"HC":         "HomeTeamCorners",
		"AC":         "AwayTeamCorners",
		"HF":         "HomeTeamFoulsCommitted",
		"AF":         "AwayTeamFoulsCommitted",
		"HFKC":       "HomeTeamFreeKicksConceded",
		"AFKC":       "AwayTeamFreeKicksConceded",
		"HO":         "HomeTeamOffsides",
		"AO":         "AwayTeamOffsides",
		"HY":         "HomeTeamYellowCards",
		"AY":         "AwayTeamYellowCards",
		"HR":         "HomeTeamRedCards",
		"AR":         "AwayTeamRedCards",
		"HBP":        "HomeTeamBookingsPoints",
		"ABP":        "AwayTeamBookingsPoints",
		"B365H":      "HomeOdd",
		"B365D":      "DrawOdd",
		"B365A":      "AwayOdd",
		"BSH":        "BlueSquarehomewinodds",
		"BSD":        "BlueSquaredrawodds",
		"BSA":        "BlueSquareawaywinodds",
		"BWH":        "Bet&Winhomewinodds",
		"BWD":        "Bet&Windrawodds",
		"BWA":        "Bet&Winawaywinodds",
		"GBH":        "Gamebookershomewinodds",
		"GBD":        "Gamebookersdrawodds",
		"GBA":        "Gamebookersawaywinodds",
		"IWH":        "Interwettenhomewinodds",
		"IWD":        "Interwettendrawodds",
		"IWA":        "Interwettenawaywinodds",
		"LBH":        "Ladbrokeshomewinodds",
		"LBD":        "Ladbrokesdrawodds",
		"LBA":        "Ladbrokesawaywinodds",
		"PSHandPH":   "Pinnaclehomewinodds",
		"PSDandPD":   "Pinnacledrawodds",
		"PSAandPA":   "Pinnacleawaywinodds",
		"SOH":        "SportingOddshomewinodds",
		"SOD":        "SportingOddsdrawodds",
		"SOA":        "SportingOddsawaywinodds",
		"SBH":        "Sportingbethomewinodds",
		"SBD":        "Sportingbetdrawodds",
		"SBA":        "Sportingbetawaywinodds",
		"SJH":        "StanJameshomewinodds",
		"SJD":        "StanJamesdrawodds",
		"SJA":        "StanJamesawaywinodds",
		"SYH":        "Stanleybethomewinodds",
		"SYD":        "Stanleybetdrawodds",
		"SYA":        "Stanleybetawaywinodds",
		"VCH":        "VCBethomewinodds",
		"VCD":        "VCBetdrawodds",
		"VCA":        "VCBetawaywinodds",
		"WHH":        "WilliamHillhomewinodds",
		"WHD":        "WilliamHilldrawodds",
		"WHA":        "WilliamHillawaywinodds",
		"Bb1X2":      "NumberofBetBrainbookmakersusedtocalculatematchoddsaveragesandmaximums",
		"BbMxH":      "Betbrainmaximumhomewinodds",
		"BbAvH":      "Betbrainaveragehomewinodds",
		"BbMxD":      "Betbrainmaximumdrawodds",
		"BbAvD":      "Betbrainaveragedrawwinodds",
		"BbMxA":      "Betbrainmaximumawaywinodds",
		"BbAvA":      "Betbrainaverageawaywinodds",
		"MaxH":       "Oddsportalmaximumhomewinodds",
		"MaxD":       "Oddsportalmaximumdrawwinodds",
		"MaxA":       "Oddsportalmaximumawaywinodds",
		"AvgH":       "Oddsportalaveragehomewinodds",
		"AvgD":       "Oddsportalaveragedrawwinodds",
		"AvgA":       "Oddsportalaverageawaywinodds",
		"BbOU":       "NumberofBetBrainbookmakersusedtocalculateover/under2.5",
		"BbMx>2.5":   "Betbrainmaximumover 2.5 goals",
		"BbAv>2.5":   "Betbrainaverageover 2.5 goals",
		"BbMx<2.5":   "Betbrainmaximumunder 2.5 goals",
		"BbAv<2.5":   "Betbrainaverageunder 2.5 goals",
		"GB>2.5":     "Gamebookersover 2.5 goals",
		"GB<2.5":     "Gamebookersunder 2.5 goals",
		"B365>2.5":   "Over2",
		"B365<2.5":   "Under2",
		"P>2.5":      "Pinnacleover 2.5 goals",
		"P<2.5":      "Pinnacleunder 2.5 goals",
		"Max>2.5":    "Oddsportalmaximumover 2.5 goals",
		"Max<2.5":    "Oddsportalmaximumunder 2.5 goals",
		"Avg>2.5":    "Oddsportalaverageover 2.5 goals",
		"Avg<2.5":    "Oddsportalaverageunder 2.5 goals",
		"BbAH":       "NumberofBetBrainbookmakersusedtoAsianhandicapaveragesandmaximums",
		"BbAHh":      "Betbrainsizeofhandicap",
		"AHh":        "Oddsportalsizeofhandicap",
		"BbMxAHH":    "BetbrainmaximumAsianhandicaphometeamodds",
		"BbAvAHH":    "BetbrainaverageAsianhandicaphometeamodds",
		"BbMxAHA":    "BetbrainmaximumAsianhandicapawayteamodds",
		"BbAvAHA":    "BetbrainaverageAsianhandicapawayteamodds",
		"GBAHH":      "GamebookersAsianhandicaphometeamodds",
		"GBAHA":      "GamebookersAsianhandicapawayteamodds",
		"GBAH":       "Gamebookerssizeofhandicap (hometeam)",
		"LBAHH":      "LadbrokesAsianhandicaphometeamodds",
		"LBAHA":      "LadbrokesAsianhandicapawayteamodds",
		"LBAH":       "Ladbrokessizeofhandicap (hometeam)",
		"B365AHH":    "Bet365 Asianhandicaphometeamodds",
		"B365AHA":    "Bet365 Asianhandicapawayteamodds",
		"B365AH":     "Bet365 sizeofhandicap (hometeam)",
		"PAHH":       "PinnacleAsianhandicaphometeamodds",
		"PAHA":       "PinnacleAsianhandicapawayteamodds",
		"MaxAHH":     "OddsportalmaximumAsianhandicaphometeamodds",
		"MaxAHA":     "OddsportalmaximumAsianhandicapawayteamodds",
		"AvgAHH":     "OddsportalaverageAsianhandicaphometeamodds",
		"AvgAHA":     "OddsportalaverageAsianhandicapawayteamodds",
	}
	if value, found := m[label]; found {
		return value
	}
	return ""
}

func (e *Event) Parse(eventDTO []string, headers []string, tournament *Tournament, week Week) (Fixture, error) {
	val := reflect.ValueOf(e)
	for index, key := range headers {
		f := (val.Elem()).FieldByName(mapLabel(key))
		if f.CanSet() {
			// change value of N
			if f.Kind() == reflect.String {

				f.SetString(eventDTO[index])
			}

			if f.Kind() == reflect.Int {
				v, _ := strconv.ParseInt(eventDTO[index], 10, 64)
				f.SetInt(v)
			}
		}
	}
	eventDate, _ := time.Parse(`02/01/2006 15:04`, fmt.Sprintf("%s %s", e.Date, e.Time))
	homeOdd, _ := strconv.ParseFloat(e.HomeOdd, 32)
	drawOdd, _ := strconv.ParseFloat(e.DrawOdd, 32)
	awayOdd, _ := strconv.ParseFloat(e.AwayOdd, 32)
	overOdd, _ := strconv.ParseFloat(e.Over2, 32)
	underOdd, _ := strconv.ParseFloat(e.Under2, 32)
	return Fixture{
		HomeTeamName:    e.HomeTeam,
		AwayTeamName:    e.AwayTeam,
		Date:            eventDate,
		Score:           fmt.Sprintf(`%d - %d`, e.FullTimeHomeTeamGoals, e.FullTimeAwayTeamGoals),
		HomeScored:      e.FullTimeHomeTeamGoals,
		HomeConceided:   e.FullTimeAwayTeamGoals,
		AwayScored:      e.FullTimeAwayTeamGoals,
		AwayConceided:   e.FullTimeHomeTeamGoals,
		HomeCorners:     e.HomeTeamCorners,
		AwayCorners:     e.AwayTeamCorners,
		HomeYellowCards: e.HomeTeamYellowCards,
		AwayYellowCards: e.AwayTeamYellowCards,
		HomeRedCards:    e.HomeTeamRedCards,
		AwayRedCards:    e.AwayTeamRedCards,
		HomeOdd:         float32(homeOdd),
		DrawOdd:         float32(drawOdd),
		AwayOdd:         float32(awayOdd),
		OverOdd:         float32(overOdd),
		UnderOdd:        float32(underOdd),
	}, nil
}

func (e Event) ParseHeader(eventDTO []string) ([]string, error) {

	headers := make([]string, 0)
	for _, header := range eventDTO {
		headers = append(headers, header)
	}
	return headers, nil
}
