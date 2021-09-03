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

func NewMatchResult() MatchResult {
	return MatchResult{MarketType{Name: string("MRES")}}
}

func NewOverUnder() OverUnder {
	return OverUnder{MarketType{Name: string("UO")}}
}

func NewBtts() Btts {
	return Btts{MarketType{Name: string("BTTS")}}
}
