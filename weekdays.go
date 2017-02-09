package hammertime

import (
	"time"
)

func LastOccurence(weekday time.Weekday, before time.Time) time.Time {
	before = before.AddDate(0, 0, -1)

	for before.Weekday() != weekday {
		before = before.AddDate(0, 0, -1)
	}

	return time.Date(before.Year(), before.Month(), before.Day(), 0, 0, 0, 0, time.UTC)
}
