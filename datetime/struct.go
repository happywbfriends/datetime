package datetime

import (
	"strings"
	"time"
)

/*
	Обертка для хранения времени в полях структуры. Нужна, чтобы при использовании `json.Unmarshal` десериализатор
	понимал все доступные формы задания времени.

	При сериализации `json.Marshal` время будет сериализовано в полную форму. Это сделано потому, что многие смежные
	системы могут не понимать сокращенных форм.

	Пример

		type MyStruct struct {
			ProperlyParsedTime ParsedTime `json:"time"`
		}
*/
type ParsedTime struct {
	time.Time
}

func (d *ParsedTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	if s == "null" || s == "" {
		*d = ParsedTime{time.Time{}}
		return nil
	}

	t, err := ParseTime(s)
	if err != nil {
		return err
	}

	*d = ParsedTime{t}
	return nil
}

func (d *ParsedTime) MarshalJSON() ([]byte, error) {
	if d.IsZero() {
		return []byte(`""`), nil
	}

	return []byte(`"` + SerializeTime(d.Time, false) + `"`), nil
}
