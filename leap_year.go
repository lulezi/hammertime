package hammertime

import (
	"time"
)

func IsInLeapYear(t time.Time) bool {
	return IsLeapYear(t.Year())
}

func IsLeapYear(year int) bool {
	if year%400 == 0 {
		return true
	}

	if year%100 == 0 {
		return false
	}

	if year%4 == 0 {
		return true
	}

	return false
}
