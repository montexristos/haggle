package fixtureModels

type Stats struct {
	Over0         int     `json:"over0"`
	Over1         int     `json:"over1"`
	Over2         int     `json:"over2"`
	Over3         int     `json:"over3"`
	Over4         int     `json:"over4"`
	Under1        int     `json:"under1"`
	Under2        int     `json:"under2"`
	Under3        int     `json:"under3"`
	Under4        int     `json:"under4"`
	Threshold0    float64 `json:"threshold0"`
	Threshold1    float64 `json:"threshold1"`
	Threshold2    float64 `json:"threshold2"`
	Threshold3    float64 `json:"threshold3"`
	Threshold4    float64 `json:"threshold4"`
	OverCorrect0  int     `json:"overCorrect0"`
	OverCorrect1  int     `json:"overCorrect1"`
	OverCorrect2  int     `json:"overCorrect2"`
	OverCorrect3  int     `json:"overCorrect3"`
	OverCorrect4  int     `json:"overCorrect4"`
	UnderCorrect0 int     `json:"underCorrect0"`
	UnderCorrect1 int     `json:"underCorrect1"`
	UnderCorrect2 int     `json:"underCorrect2"`
	UnderCorrect3 int     `json:"underCorrect3"`
	UnderCorrect4 int     `json:"underCorrect4"`
	OverWrong0    int     `json:"overWrong0"`
	OverWrong1    int     `json:"overWrong1"`
	OverWrong2    int     `json:"overWrong2"`
	OverWrong3    int     `json:"overWrong3"`
	OverWrong4    int     `json:"overWrong4"`
	UnderWrong0   int     `json:"underWrong0"`
	UnderWrong1   int     `json:"underWrong1"`
	UnderWrong2   int     `json:"underWrong2"`
	UnderWrong3   int     `json:"underWrong3"`
	UnderWrong4   int     `json:"underWrong4"`
	CorrectCards  int     `json:"correctCards"`
	WrongCards    int     `json:"wrongCards"`
	CorrectCorner int     `json:"correctCorner"`
	WrongCorner   int     `json:"wrongCorner"`
}

//NewStats returns an initialized stats object
func NewStats() *Stats {
	stats := Stats{}
	stats.Over0 = 0
	stats.Over1 = 0
	stats.Over2 = 0
	stats.Over3 = 0
	stats.Over4 = 0
	stats.OverCorrect0 = 0
	stats.OverCorrect1 = 0
	stats.OverCorrect2 = 0
	stats.OverCorrect3 = 0
	stats.OverCorrect4 = 0
	stats.OverWrong0 = 0
	stats.OverWrong1 = 0
	stats.OverWrong2 = 0
	stats.OverWrong3 = 0
	stats.OverWrong4 = 0
	stats.UnderCorrect0 = 0
	stats.UnderCorrect1 = 0
	stats.UnderCorrect2 = 0
	stats.UnderCorrect3 = 0
	stats.UnderCorrect4 = 0
	stats.UnderWrong0 = 0
	stats.UnderWrong1 = 0
	stats.UnderWrong2 = 0
	stats.UnderWrong3 = 0
	stats.UnderWrong4 = 0
	stats.Under1 = 0
	stats.Under2 = 0
	stats.Under3 = 0
	stats.Under4 = 0
	stats.CorrectCards = 0
	stats.WrongCards = 0
	stats.CorrectCorner = 0
	stats.WrongCorner = 0
	stats.Threshold0 = 50.0
	stats.Threshold1 = 70.0
	stats.Threshold2 = 75.0
	stats.Threshold3 = 80.0
	stats.Threshold4 = 85.0
	return &stats
}
