package hammertime

import (
	"time"
)

func AddOneHour(t time.Time) time.Time {
	return AddHours(t, 1)
}

func SubOneHour(t time.Time) time.Time {
	return SubHours(t, 1)
}

func SubHours(t time.Time, n int) time.Time {
	return AddHours(t, -n)
}

func AddHours(t time.Time, n int) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour()+n, t.Minute(), t.Second(), t.Nanosecond(), t.Location())
}
