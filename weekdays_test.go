package weekdays

import (
	"testing"
	"time"
)

func TestLastOccurence(t *testing.T) {
	before := time.Date(2017, 2, 13, 0, 0, 0, 0, time.UTC)

	var (
		monday    = time.Date(2017, 2, 6, 0, 0, 0, 0, time.UTC)
		tuesday   = time.Date(2017, 2, 7, 0, 0, 0, 0, time.UTC)
		wednesday = time.Date(2017, 2, 8, 0, 0, 0, 0, time.UTC)
		thursday  = time.Date(2017, 2, 9, 0, 0, 0, 0, time.UTC)
		friday    = time.Date(2017, 2, 10, 0, 0, 0, 0, time.UTC)
		saturday  = time.Date(2017, 2, 11, 0, 0, 0, 0, time.UTC)
		sunday    = time.Date(2017, 2, 12, 0, 0, 0, 0, time.UTC)
	)

	if out := LastOccurence(time.Monday, before); out != monday {
		t.Error("nope for monday", out)
	}

	if out := LastOccurence(time.Tuesday, before); out != tuesday {
		t.Error("nope for tuesday", out)
	}

	if out := LastOccurence(time.Wednesday, before); out != wednesday {
		t.Error("nope for wednesday", out)
	}

	if out := LastOccurence(time.Thursday, before); out != thursday {
		t.Error("nope for thursday", out)
	}

	if out := LastOccurence(time.Friday, before); out != friday {
		t.Error("nope for friday", out)
	}

	if out := LastOccurence(time.Saturday, before); out != saturday {
		t.Error("nope for saturday", out)
	}

	if out := LastOccurence(time.Sunday, before); out != sunday {
		t.Error("nope for sunday", out)
	}
}
