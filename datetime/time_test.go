package datetime

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestParseTime(t *testing.T) {
	const milliSec = 1_000_000
	const microSec = 1_000

	type args struct {
		date string
	}
	tests := []struct {
		args    args
		wantT   time.Time
		wantErr bool
	}{
		{args{"2022-01-02"}, time.Date(2022, 1, 2, 0, 0, 0, 0, time.UTC), false},
		{args{"1981-03-04"}, time.Date(1981, 3, 4, 0, 0, 0, 0, time.UTC), false},
		{args{"1981-03-04T23:45"}, time.Date(1981, 3, 4, 23, 45, 0, 0, time.UTC), false},

		{args{"1981-03-04T20:45:11"}, time.Date(1981, 3, 4, 20, 45, 11, 0, time.UTC), false},
		{args{"1981-03-04T20:45:11Z"}, time.Date(1981, 3, 4, 20, 45, 11, 0, time.UTC), false},
		{args{"1981-03-04T20:45:11+03:00"}, time.Date(1981, 3, 4, 17, 45, 11, 0, time.UTC), false},
		{args{"1981-03-04T20:45:11-03:00"}, time.Date(1981, 3, 4, 23, 45, 11, 0, time.UTC), false},

		{args{"1981-03-04T11:01:59.123"}, time.Date(1981, 3, 4, 11, 1, 59, 123*milliSec, time.UTC), false},
		{args{"1981-03-04T11:01:59.123Z"}, time.Date(1981, 3, 4, 11, 1, 59, 123*milliSec, time.UTC), false},
		{args{"1981-03-04T11:01:59.123+03:00"}, time.Date(1981, 3, 4, 8, 1, 59, 123*milliSec, time.UTC), false},
		{args{"1981-03-04T11:01:59.123-03:00"}, time.Date(1981, 3, 4, 14, 1, 59, 123*milliSec, time.UTC), false},

		{args{"1981-03-04T23:45:11.123456"}, time.Date(1981, 3, 4, 23, 45, 11, 123456*microSec, time.UTC), false},
		{args{"1981-03-04T16:17:18.123456Z"}, time.Date(1981, 3, 4, 16, 17, 18, 123456*microSec, time.UTC), false},
		{args{"1981-03-04T16:17:18.123456+03:00"}, time.Date(1981, 3, 4, 13, 17, 18, 123456*microSec, time.UTC), false},
		{args{"1981-03-04T16:17:18.123456-03:00"}, time.Date(1981, 3, 4, 19, 17, 18, 123456*microSec, time.UTC), false},

		{args{""}, time.Time{}, true},
		{args{"1/1/2023"}, time.Time{}, true},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			gotT, err := ParseTime(tt.args.date)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseDateTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !gotT.Equal(tt.wantT) {
				t.Errorf("ParseDateTime() gotT = %v, want %v", gotT, tt.wantT)
			}
		})
	}
}

type TestMarshalStruct struct {
	D SerializedTime `json:"d"`
}

func TestUnmarshalling(t *testing.T) {
	const milliSec = 1_000_000
	const microSec = 1_000

	tests := []struct {
		jsons   string
		wantT   time.Time
		wantErr bool
	}{
		{"2022-01-02", time.Date(2022, 1, 2, 0, 0, 0, 0, time.UTC), false},
		{"1981-03-04", time.Date(1981, 3, 4, 0, 0, 0, 0, time.UTC), false},
		{"1981-03-04T23:45", time.Date(1981, 3, 4, 23, 45, 0, 0, time.UTC), false},

		{"1981-03-04T20:45:11", time.Date(1981, 3, 4, 20, 45, 11, 0, time.UTC), false},
		{"1981-03-04T20:45:11Z", time.Date(1981, 3, 4, 20, 45, 11, 0, time.UTC), false},
		{"1981-03-04T20:45:11+03:00", time.Date(1981, 3, 4, 17, 45, 11, 0, time.UTC), false},
		{"1981-03-04T20:45:11-03:00", time.Date(1981, 3, 4, 23, 45, 11, 0, time.UTC), false},

		{"1981-03-04T11:01:59.123", time.Date(1981, 3, 4, 11, 1, 59, 123*milliSec, time.UTC), false},
		{"1981-03-04T11:01:59.123Z", time.Date(1981, 3, 4, 11, 1, 59, 123*milliSec, time.UTC), false},
		{"1981-03-04T11:01:59.123+03:00", time.Date(1981, 3, 4, 8, 1, 59, 123*milliSec, time.UTC), false},
		{"1981-03-04T11:01:59.123-03:00", time.Date(1981, 3, 4, 14, 1, 59, 123*milliSec, time.UTC), false},

		{"1981-03-04T23:45:11.123456", time.Date(1981, 3, 4, 23, 45, 11, 123456*microSec, time.UTC), false},
		{"1981-03-04T16:17:18.123456Z", time.Date(1981, 3, 4, 16, 17, 18, 123456*microSec, time.UTC), false},
		{"1981-03-04T16:17:18.123456+03:00", time.Date(1981, 3, 4, 13, 17, 18, 123456*microSec, time.UTC), false},
		{"1981-03-04T16:17:18.123456-03:00", time.Date(1981, 3, 4, 19, 17, 18, 123456*microSec, time.UTC), false},

		{"", time.Time{}, false},
		{"1/1/2023", time.Time{}, true},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			jsonString := fmt.Sprintf(`{"d": "%s"}`, tt.jsons)
			var testStruct TestMarshalStruct
			err := json.Unmarshal([]byte(jsonString), &testStruct)
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !testStruct.D.Equal(tt.wantT) {
				t.Errorf("gotT = %v, want %v", testStruct.D, tt.wantT)
			}
		})
	}
}
