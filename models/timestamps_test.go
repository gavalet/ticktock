package models_test

import (
	"cmd/ticktock/models"
	"net/http"
	"testing"
)

func TestCreateTimestamps(t *testing.T) {
	tests := []struct {
		name      string
		period    string
		tz        string
		t1Str     string
		t2Str     string
		expected  []string
		status    int
		expectErr bool
	}{
		//set test's parameters
		{
			name:   "Valid - Tz:Europe/Athens - Period 1h",
			period: "1h",
			tz:     "Europe/Athens",
			t1Str:  "20210714T204603Z",
			t2Str:  "20210715T123456Z",
			expected: []string{"20210714T210000Z",
				"20210714T220000Z",
				"20210714T230000Z",
				"20210715T000000Z",
				"20210715T010000Z",
				"20210715T020000Z",
				"20210715T030000Z",
				"20210715T040000Z",
				"20210715T050000Z",
				"20210715T060000Z",
				"20210715T070000Z",
				"20210715T080000Z",
				"20210715T090000Z",
				"20210715T100000Z",
				"20210715T110000Z",
				"20210715T120000Z"},
			status: http.StatusOK,
		},
		{
			name:   "Valid - Tz:Europe/Athens - Period 1d",
			period: "1d",
			tz:     "Europe/Athens",
			t1Str:  "20211010T204603Z",
			t2Str:  "20211115T123456Z",
			expected: []string{"20211010T210000Z",
				"20211011T210000Z", "20211012T210000Z",
				"20211013T210000Z",
				"20211014T210000Z",
				"20211015T210000Z",
				"20211016T210000Z",
				"20211017T210000Z",
				"20211018T210000Z",
				"20211019T210000Z",
				"20211020T210000Z",
				"20211021T210000Z",
				"20211022T210000Z",
				"20211023T210000Z",
				"20211024T210000Z",
				"20211025T210000Z",
				"20211026T210000Z",
				"20211027T210000Z",
				"20211028T210000Z",
				"20211029T210000Z",
				"20211030T210000Z",
				"20211031T220000Z",
				"20211101T220000Z",
				"20211102T220000Z",
				"20211103T220000Z",
				"20211104T220000Z",
				"20211105T220000Z",
				"20211106T220000Z",
				"20211107T220000Z",
				"20211108T220000Z",
				"20211109T220000Z",
				"20211110T220000Z",
				"20211111T220000Z",
				"20211112T220000Z",
				"20211113T220000Z",
				"20211114T220000Z"},
			status: http.StatusOK,
		},
		{
			name:   "Valid - Tz:Europe/Athens - Period 1month",
			period: "1mo",
			tz:     "Europe/Athens",
			t1Str:  "20210214T204603Z",
			t2Str:  "20211115T123456Z",
			expected: []string{"20210228T220000Z",
				"20210331T210000Z",
				"20210430T210000Z",
				"20210531T210000Z",
				"20210630T210000Z",
				"20210731T210000Z",
				"20210831T210000Z",
				"20210930T210000Z",
				"20211031T220000Z"},
			status: http.StatusOK,
		},
		{
			name:   "Valid - Tz:Europe/Athens - Period 1 year",
			period: "1y",
			tz:     "Europe/Athens",
			t1Str:  "20180214T204603Z",
			t2Str:  "20211115T123456Z",
			expected: []string{"20181231T220000Z",
				"20191231T220000Z",
				"20201231T220000Z"},
			status: http.StatusOK,
		},
		{
			name:     "Invalid - period",
			period:   "2y",
			tz:       "Europe/Athens",
			t1Str:    "20180214T204603Z",
			t2Str:    "20211115T123456Z",
			expected: []string{},
			status:   http.StatusBadRequest,
		},
		{
			name:     "Invalid - t2 < t1",
			period:   "1y",
			tz:       "Europe/Athens",
			t1Str:    "20211115T123456Z",
			t2Str:    "20180214T204603Z",
			expected: []string{},
			status:   http.StatusBadRequest,
		},
		{
			name:     "Invalid -time zone does not exists",
			period:   "1y",
			tz:       "Europe/Ksidas",
			t1Str:    "20180214T204603Z",
			t2Str:    "20211115T123456Z",
			expected: []string{},
			status:   http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//run
			result, status, err := models.GetTimestamps("TestReqID", test.period, test.tz, test.t1Str, test.t2Str)

			//assert
			if test.expectErr && err == nil {
				t.Errorf("Expected an error, but got nil")
				return
			}

			if test.expectErr && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if status != test.status {
				t.Errorf("Expected status %d, but got %d", test.status, status)
			}

			if !compareSlices(result, test.expected) {
				t.Errorf("Expected timestamps %v, but got %v", test.expected, result)
			}
		})
	}
}

func compareSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
