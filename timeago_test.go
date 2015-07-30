// Copyright 2013 Simon HEGE. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package timeago

import (
	"testing"
	"time"
)

//Base time for testing
var tBase = time.Date(2013, 8, 30, 12, 0, 0, 0, time.UTC)

//Test data for TestFormatReference
var formatReferenceTests = []struct {
	t        time.Time // input time
	ref      time.Time // input reference
	cfg      Config    //input cfguage
	expected string    // expected result
}{
	//Lang
	{tBase, tBase, NoMax(English), "about a second ago"},
	{tBase, tBase, NoMax(French), "il y a environ une seconde"},
	{tBase, tBase, NoMax(Chinese), "1 秒前"},

	//Thresholds
	{tBase, tBase.Add(1*time.Second + 500000000).Add(-1), NoMax(English), "about a second ago"},
	{tBase, tBase.Add(1*time.Second + 500000000), NoMax(English), "2 seconds ago"},
	{tBase, tBase.Add(1 * time.Minute).Add(-500000001), NoMax(English), "59 seconds ago"},
	{tBase, tBase.Add(1 * time.Minute), NoMax(English), "about a minute ago"},
	{tBase, tBase.Add(1*time.Minute + 30*time.Second).Add(-1), NoMax(English), "about a minute ago"},
	{tBase, tBase.Add(1*time.Minute + 30*time.Second), NoMax(English), "2 minutes ago"},
	{tBase, tBase.Add(59*time.Minute + 30*time.Second).Add(-1), NoMax(English), "59 minutes ago"},
	{tBase, tBase.Add(59*time.Minute + 30*time.Second), NoMax(English), "about an hour ago"},
	{tBase, tBase.Add(90 * time.Minute).Add(-1), NoMax(English), "about an hour ago"},
	{tBase, tBase.Add(90 * time.Minute), NoMax(English), "2 hours ago"},
	{tBase, tBase.Add(23*time.Hour + 30*time.Minute).Add(-1), NoMax(English), "23 hours ago"},
	{tBase, tBase.Add(23*time.Hour + 30*time.Minute), NoMax(English), "one day ago"},
	{tBase, tBase.Add(36 * time.Hour).Add(-1), NoMax(English), "one day ago"},
	{tBase, tBase.Add(36 * time.Hour), NoMax(English), "2 days ago"},
	{tBase, tBase.Add(30 * 24 * time.Hour).Add(-12*time.Hour - 1), NoMax(English), "29 days ago"},
	{tBase, tBase.Add(30 * 24 * time.Hour), NoMax(English), "one month ago"},
	{tBase, tBase.Add(45 * 24 * time.Hour).Add(-1), NoMax(English), "one month ago"},
	{tBase, tBase.Add(45 * 24 * time.Hour), NoMax(English), "2 months ago"},
	{tBase, tBase.Add(Year).Add(-30 * Day), NoMax(English), "11 months ago"},
	{tBase, tBase.Add(Year), NoMax(English), "one year ago"},
	{tBase, tBase.Add(547 * Day), NoMax(English), "one year ago"},
	{tBase, tBase.Add(548 * Day), NoMax(English), "2 years ago"},
	{tBase, tBase.Add(10 * Year), NoMax(English), "10 years ago"},

	//Max
	{tBase, tBase.Add(90 * time.Minute).Add(-1), NoMax(English), "about an hour ago"},
	{tBase, tBase.Add(90 * time.Minute).Add(-1), WithMax(English, 90*time.Minute, ""), "about an hour ago"},
	{tBase, tBase.Add(90 * time.Minute), WithMax(English, 90*time.Minute, "2006-01-02"), "2013-08-30"},

	//Future
	{tBase.Add(24 * time.Hour), tBase, NoMax(English), "in one day"},
}

//Test the FormatReference method
func TestFormatReference(t *testing.T) {
	for i, tt := range formatReferenceTests {
		actual := tt.cfg.FormatReference(tt.t, tt.ref)
		if actual != tt.expected {
			t.Errorf("%d) FormatReference(%s,%s): expected '%s', actual '%s'", i+1, tt.t, tt.ref, tt.expected, actual)
		}
	}
}

type testData struct {
	inpStr string
	inpCnt int
}

func TestNbParamInFormat(t *testing.T) {

	inpTestData := []testData{
		{"testing param0", 0},
		{"testing param1 %d", 1},
		{"testing param2 %d %v %%", 2},
		{"testing param3 %v  param %v param %v", 3},
		{"testing param4 %v%v  param %b param %c param %q %v %d param %v test %t %v", 10},
	}

	for _, itm := range inpTestData {
		t.Log("Testing with params count:", itm.inpCnt)
		outCnt := nbParamInFormat(itm.inpStr)
		if outCnt != itm.inpCnt {
			t.Error("Expected: ", itm.inpCnt, " actual: ", outCnt)
		}
	}

}
