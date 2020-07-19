package time

import (
	"time"
)

// TimeFormat 2006-01-02 15:04:05
const TimeFormat = "2006-01-02 15:04:05"

// TimeFormatTZ 2006-01-02 15:04:05
const TimeFormatTZ = "2006-01-02 15:04:05 -0700 MST"

// ChinaTimeZone UTC +8 timezone
var ChinaTimeZone = func() *time.Location {
	tz, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		tz = time.UTC
	}
	return tz
}()

// FromString from time string like 2006-01-02 15:04:05
func FromString(str string) int64 {
	t, err := time.Parse(TimeFormat, str)
	if err != nil {
		t = time.Unix(0, 0)
	}
	return t.Unix()
}

// ToString transfer time.Time to 2006-01-02 15:04:05 string
func ToString(t int64) string {
	return time.Unix(t, 0).Format(TimeFormat)
}
