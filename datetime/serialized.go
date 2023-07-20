package datetime

import (
	"strings"
	"time"
)

type SerializedTime struct {
	time.Time
}

func (d *SerializedTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	if s == "null" || s == "" {
		*d = SerializedTime{time.Time{}}
		return nil
	}

	t, err := ParseTime(s)
	if err != nil {
		return err
	}

	*d = SerializedTime{t}
	return nil
}

func (d *SerializedTime) MarshalJSON() ([]byte, error) {
	if d.IsZero() {
		return []byte(`""`), nil
	}

	return []byte(`"` + SerializeTime(d.Time) + `"`), nil
}
