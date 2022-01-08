package models

type MarketType struct {
	Name string
}

type MatchResult struct {
	MarketType
}

type OverUnder struct {
	MarketType
	Handicap float64
}

type Btts struct {
	MarketType
}

type NextTeamToScore struct {
	MarketType
}
type DoubleChance struct {
	MarketType
}
type DrawNoBet struct {
	MarketType
}

type HomeNoBet struct {
	MarketType
}

type UnderOverHome struct {
	MarketType
	Handicap float64
}

type UnderOverAway struct {
	MarketType
	Handicap float64
}
type UnderOverHalf struct {
	MarketType
	Handicap float64
}
type FirstGoalEarly struct {
	MarketType
}
type UnderOverCorners struct {
	MarketType
	Handicap float64
}
type BttsOrOver struct {
	MarketType
}
type RacePoints struct {
	MarketType
	Handicap float64
}

func NewMatchResult() MatchResult {
	return MatchResult{MarketType{Name: "MRES"}}
}

func NewOverUnder() OverUnder {
	return OverUnder{
		MarketType: MarketType{Name: "OU"},
		Handicap:   2.5,
	}
}
func NewOverUnderHandicap(hc float64) OverUnder {
	return OverUnder{
		MarketType: MarketType{Name: "OU"},
		Handicap:   hc,
	}
}

func NewBtts() Btts {
	return Btts{MarketType{Name: "BTTS"}}
}

func NewNextTeamToScore() NextTeamToScore {
	return NextTeamToScore{MarketType{Name: "INTS"}}
}

func NewDoubleChance() DoubleChance {
	return DoubleChance{MarketType{Name: "DBLC"}}
}

func NewDrawNoBet() DrawNoBet {
	return DrawNoBet{MarketType{Name: "DNOB"}}
}
func NewHomeNoBet() HomeNoBet {
	return HomeNoBet{MarketType{Name: "DNOB"}}
}

func NewUnderOverHome(hc float64) UnderOverHome {
	return UnderOverHome{MarketType{Name: string("OUHG")}, hc}
}

func NewUnderOverAway(hc float64) UnderOverAway {
	return UnderOverAway{MarketType{Name: string("OUAG")}, hc}
}

func NewUnderOverHalf(hc float64) UnderOverHalf {
	return UnderOverHalf{MarketType{Name: string("OUH1")}, hc}
}

func NewFirstGoalEarly() FirstGoalEarly {
	return FirstGoalEarly{MarketType{Name: string("FG28")}}
}

func NewUnderOverCorners(hc float64) UnderOverCorners {
	return UnderOverCorners{MarketType{Name: string("OUCR")}, hc}
}
func NewBttsOrOver(hc float64) BttsOrOver {
	return BttsOrOver{MarketType{Name: string("BTTSOROV")}}
}

func NewRacePoints(hc float64) RacePoints {
	return RacePoints{MarketType{Name: string("")}, hc}
}
