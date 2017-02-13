package hammertime

import (
	"fmt"
	"testing"
	"time"
)

func TestMain(t *testing.T) {
	withoutDST := time.Date(2017, 3, 25, 15, 0, 0, 0, time.UTC)
	withDST := time.Date(2017, 3, 26, 15, 0, 0, 0, time.UTC)
	addedDay := withoutDST.AddDate(0, 0, 1)
	if !addedDay.Equal(withDST) {
		t.Error("DST error", addedDay, withDST)
	}

	withoutDST = time.Date(2017, 3, 25, 15, 0, 0, 0, time.Local)
	withDST = time.Date(2017, 3, 26, 15, 0, 0, 0, time.Local)
	addedDay = withoutDST.AddDate(0, 0, 1)
	if !addedDay.Equal(withDST) {
		t.Error("DST error", addedDay, withDST)
	}

	fmt.Println(withoutDST, addedDay, withDST)
}
