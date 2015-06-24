// dbfreader_test.go
package godbf

import (
	"fmt"
	"path"
	"strconv"
	"testing"
)

func TestPercentile(t *testing.T) {
	localDirectory := "C:\\tmp\\BR_20140901_2"
	prefix := "segments"
	i := 3
	zero := "0"
	localDbfOptimal := path.Join(localDirectory, fmt.Sprintf("%s_%d_optimal.dbf", prefix, i))
	localDbf := path.Join(localDirectory, fmt.Sprintf("%s_%d.dbf", prefix, i))
	dbfTable, err := NewFromFile(localDbfOptimal, "UTF8")
	if err != nil {
		t.Fail()
	}
	for i := 0; i < dbfTable.NumberOfRecords(); i++ {
		currentSegmentId, _ := strconv.ParseUint(dbfTable.FieldValueUtf8(i, 0), 0, 32)
		segmentMaxSpeedKph, _ := strconv.ParseUint(dbfTable.FieldValueUtf8(i, 6), 0, 16)
		jamthresholdspeed, _ := strconv.ParseUint(dbfTable.FieldValueUtf8(i, 7), 0, 16)
		dbfTable.SetFieldValueUtf8(i, 1, fmt.Sprint(currentSegmentId))   // speed
		dbfTable.SetFieldValueUtf8(i, 2, fmt.Sprint(segmentMaxSpeedKph)) // coverage
		dbfTable.SetFieldValueUtf8(i, 3, fmt.Sprint(jamthresholdspeed))  // los
		dbfTable.SetFieldValueUtf8(i, 4, zero)                           // is optimal ?

	}
	writeerr := dbfTable.SaveFile(localDbf)
	if writeerr != nil {
		t.Fail()
	}
}
