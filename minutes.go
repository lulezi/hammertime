package hammertime

import (
	"time"
)

func AddOneMinute(t time.Time) time.Time {
	return AddMinutes(t, 1)
}

func SubOneMinute(t time.Time) time.Time {
	return SubMinutes(t, 1)
}

func SubMinutes(t time.Time, n int) time.Time {
	return AddMinutes(t, -n)
}

func AddMinutes(t time.Time, n int) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute()+n, t.Second(), t.Nanosecond(), t.Location())
}
