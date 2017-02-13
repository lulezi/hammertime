package hammertime

import (
	"testing"
)

func TestIsLeapYear(t *testing.T) {
	noLeapYears := []int{1999, 2001, 2100}
	leapYears := []int{2000, 2004}

	for _, e := range noLeapYears {
		if IsLeapYear(e) {
			t.Errorf("%d is not a leap year", e)
		}
	}

	for _, e := range leapYears {
		if !IsLeapYear(e) {
			t.Errorf("%d is a leap year", e)
		}
	}
}
