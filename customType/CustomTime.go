package customType

import (
	"strings"
	"time"
)

type CustomTime time.Time


var dateFormats = []string{
	"2006-01-02",
	"2006-01-02T15:04:05Z",
	"2006-01-02 15:04:05",
	"02-01-2006", // format lokal
}

func (c *CustomTime) UnmarshalJSON(b []byte) error {
	str := string(b)
	str = strings.Trim(str, `"`) 
	var parseErr error

	for _, format := range dateFormats {
		parsedTime, err := time.Parse(format, str)
		if err == nil {
			*c = CustomTime(parsedTime)
			return nil
		}
		parseErr = err
	}

	return parseErr
}

func (c CustomTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(c).Format("2006-01-02") + `"`), nil
}

func (c CustomTime) ToTime() time.Time {
	return time.Time(c)
}

func (c CustomTime) IsZero() bool {
	return time.Time(c).IsZero()
}
