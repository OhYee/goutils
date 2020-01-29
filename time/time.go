package time

import (
	"time"
)

// TimeFormat 2006-01-02 15:04:05
const TimeFormat = "2006-01-02 15:04:05"

// FromString from time string like 2006-01-02 15:04:05
func FromString(str string) time.Time {
	t, err := time.Parse(TimeFormat, str)
	if err != nil {
		t = time.Now()
	}
	return t
}

// ToString transfer time.Time to 2006-01-02 15:04:05 string
func ToString(t time.Time) string {
	return t.Format(TimeFormat)
}
