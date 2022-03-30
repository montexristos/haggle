package tools

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/montexristos/haggle/fixtureModels"
)

// GetHomeDateRange
func GetHomeDateRange() fixtureModels.Week {
	start, end := WeekRange(time.Now().ISOWeek())
	return fixtureModels.Week{
		Start: start,
		End:   end,
	}
}

func WeekStart(year, week int) time.Time {
	// Start from the middle of the year:
	t := time.Date(year, 7, 1, 0, 0, 0, 0, time.UTC)

	// Roll back to Monday:
	if wd := t.Weekday(); wd == time.Sunday {
		t = t.AddDate(0, 0, -6)
	} else {
		t = t.AddDate(0, 0, -int(wd)+1)
	}

	// Difference in weeks:
	_, w := t.ISOWeek()
	t = t.AddDate(0, 0, (week-w)*7)

	return t
}

func WeekRange(year, week int) (start, end time.Time) {
	start = WeekStart(year, week)
	end = start.AddDate(0, 0, 7)
	return
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()
	fmt.Print(string(content))
	out.Write(content)
	out.Sync()
	return err
}

func GetDate() string {
	return time.Now().Format("2006-01-02")
}
