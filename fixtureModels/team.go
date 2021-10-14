package fixtureModels

// Team the event model
type Team struct {
	Name                 string        `json:"name"`
	GoalsScored          float64       `json:"goalsScored"`
	GoalsConceided       float64       `json:"goalsConceided"`
	HomeGoalsScored      float64       `json:"homeGoalsScored"`
	HomeGoalsConceided   float64       `json:"homeGoalsConceided"`
	AwayGoalsScored      float64       `json:"awayGoalsScored"`
	AwayGoalsConceided   float64       `json:"awayGoalsConceided"`
	NoEvents             int           `json:"noEvents"`
	HomeEvents           int           `json:"homeEvents"`
	AwayEvents           int           `json:"awayEvents"`
	Points               int           `json:"points"`
	Corners              int           `json:"corners"`
	HomeCorners          int           `json:"homeCorners"`
	HomeAwayCorners      int           `json:"homeAwayCorners"`
	AwayCorners          int           `json:"awayCorners"`
	AwayHomeCorners      int           `json:"awayHomeCorners"`
	YellowCards          int           `json:"yellowCards"`
	RedCards             int           `json:"redCards"`
	TotalCards           int           `json:"totalCards"`
	Fixtures             []FixtureJson `json:"fixtures"`
	Over0                int           `json:"over0"`
	Over1                int           `json:"over1"`
	Over2                int           `json:"over2"`
	Over3                int           `json:"over3"`
	Over4                int           `json:"over4"`
	Under1               int           `json:"under1"`
	Under2               int           `json:"under2"`
	Under3               int           `json:"under3"`
	Under4               int           `json:"under4"`
	Gg                   int           `json:"gg"`
	HomeGg               int           `json:"homeGg"`
	AwayGg               int           `json:"awayGg"`
	HomeMatchesScored    int           `json:"homeMatchesScored"`
	AwayMatchesScored    int           `json:"awayMatchesScored"`
	HomeMatchesConceided int           `json:"homeMatchesConceided"`
	AwayMatchesConceided int           `json:"awayMatchesConceided"`
	HomeIndex            float32       `json:"homeIndex"`
	AwayIndex            float32       `json:"awayIndex"`
	CardIndex            float32       `json:"cardIndex"`
}

// TeamEvent the event model
type TeamEvent struct {
	Tournament string `json:"tournament"`
	Date       string `json:"date"`
	HomeTeam   string
	AwayTeam   string
}
