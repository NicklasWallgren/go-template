package types

import (
	"strings"
	"time"
)

type Date time.Time

const formatYYYYMMDD = "2006-01-02"

func (d *Date) UnmarshalJSON(b []byte) error {
	date := strings.Trim(string(b), "\"")
	time, err := time.Parse(formatYYYYMMDD, date)
	if err != nil {
		return err
	}
	*d = Date(time)
	return nil
}

func (d Date) MarshalJSON() ([]byte, error) {
	return []byte("\"" + time.Time(d).Format(formatYYYYMMDD) + "\""), nil
}

func (d Date) ToTime() time.Time {
	return time.Time(d)
}

func (d Date) Format(s string) string {
	t := time.Time(d)
	return t.Format(s)
}
