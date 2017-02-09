package hammertime

import (
	"time"
)

func AddOneNanosecond(t time.Time) time.Time {
	return AddNanoseconds(t, 1)
}

func SubOneNanosecond(t time.Time) time.Time {
	return SubNanoseconds(t, 1)
}

func SubNanoseconds(t time.Time, n int) time.Time {
	return AddNanoseconds(t, -n)
}

func AddNanoseconds(t time.Time, n int) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond()+n, t.Location())
}
