package hammertime

import (
	"time"
)

func AddOneSecond(t time.Time) time.Time {
	return AddSeconds(t, 1)
}

func SubOneSecond(t time.Time) time.Time {
	return SubSeconds(t, 1)
}

func SubSeconds(t time.Time, n int) time.Time {
	return AddSeconds(t, -n)
}

func AddSeconds(t time.Time, n int) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second()+n, t.Nanosecond(), t.Location())
}
