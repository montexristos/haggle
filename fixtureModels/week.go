package fixtureModels

import (
	"encoding/json"
	"time"
)

// Week holds two timestamps start and end of a week
type Week struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

// MarshalJSON custom json encode
func (w *Week) MarshalJSON() ([]byte, error) {
	layout := "02-01-2006"
	return json.Marshal(&struct {
		Start string `json:"start"`
		End   string `json:"end"`
	}{
		Start: w.Start.Format(layout),
		End:   w.End.Format(layout),
	})
}
