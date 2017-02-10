package hammertime

import (
	"time"
)

func AddTime(t time.Time, hours, minutes, seconds, nanoseconds int) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour()+hours, t.Minute()+minutes, t.Second()+seconds, t.Nanosecond()+nanoseconds, t.Location())
}
