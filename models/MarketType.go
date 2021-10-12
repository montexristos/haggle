package models

type MarketType struct {
	Name string
}

type MatchResult struct {
	MarketType
}

type OverUnder struct {
	MarketType
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

type UnderOverHome struct {
	MarketType
}

type UnderOverAway struct {
	MarketType
}
type UnderOverHalf struct {
	MarketType
}
type FirstGoalEarly struct {
	MarketType
}
type UnderOverCorners struct {
	MarketType
}

func NewMatchResult() MatchResult {
	return MatchResult{MarketType{Name: string("MRES")}}
}

func NewOverUnder() OverUnder {
	return OverUnder{MarketType{Name: string("UO")}}
}

func NewBtts() Btts {
	return Btts{MarketType{Name: string("BTTS")}}
}

func NewNextTeamToScore() NextTeamToScore {
	return NextTeamToScore{MarketType{Name: string("INTS")}}
}

func NewDoubleChance() DoubleChance {
	return DoubleChance{MarketType{Name: string("DBLC")}}
}

func NewDrawNoBet() DrawNoBet {
	return DrawNoBet{MarketType{Name: string("DNOB")}}
}

func NewUnderOverHome() UnderOverHome {
	return UnderOverHome{MarketType{Name: string("OUHG")}}
}

func NewUnderOverAway() UnderOverAway {
	return UnderOverAway{MarketType{Name: string("OUAG")}}
}

func NewUnderOverHalf() UnderOverHalf {
	return UnderOverHalf{MarketType{Name: string("OUH1")}}
}

func NewFirstGoalEarly() FirstGoalEarly {
	return FirstGoalEarly{MarketType{Name: string("FG28")}}
}

func NewUnderOverCorners() UnderOverCorners {
	return UnderOverCorners{MarketType{Name: string("UOCR")}}
}
