package datetime

import (
	"encoding/json"
	"testing"
	"time"
)

type testSerializedTime struct {
	T SerializedTime `json:"t,omitempty"`
}

const milliSec = 1_000_000

func TestSerializedTime_UnmarshalJSON(t *testing.T) {

	tests := []struct {
		name         string
		marshalled   string
		wantErr      bool
		unmarshalled time.Time
	}{
		{"1", `{"t": "2023-01-01"}`, false, time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)},
		{"2", `{"t": "1981-03-04T11:01:59.123Z"}`, false, time.Date(1981, 3, 4, 11, 1, 59, 123*milliSec, time.UTC)},
		{"3", `{"t": ""}`, false, time.Time{}},
		{"4", `{}`, false, time.Time{}},
		{"5", `{"t": "not a time"}`, true, time.Time{}},
		{"6", `{"t": null}`, false, time.Time{}},
		{"7", `{"t": 123}`, true, time.Time{}},
		{"8", `{"t": 0.0}`, true, time.Time{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := testSerializedTime{}
			if err := json.Unmarshal([]byte(tt.marshalled), &d); (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
			}
			if d.T.Time != tt.unmarshalled {
				t.Errorf("expected time=%v, got %v", tt.unmarshalled, d.T.Time)
			}
		})
	}
}

func TestSerializedTime_MarshalJSON(t *testing.T) {
	tests := []struct {
		name       string
		t          time.Time
		marshalled string
	}{
		{"1", time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC), `{"t":"2023-01-01"}`},
		{"2", time.Date(1981, 3, 4, 11, 1, 59, 123*milliSec, time.UTC), `{"t":"1981-03-04T11:01:59.123Z"}`},
		{"3", time.Time{}, `{"t":""}`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := testSerializedTime{
				T: SerializedTime{
					Time: tt.t,
				},
			}
			data, err := json.Marshal(&d)
			if err != nil {
				t.Errorf("error = %v", err)
			}
			if string(data) != tt.marshalled {
				t.Errorf("expected %s, got %s", tt.marshalled, data)
			}
		})
	}
}
