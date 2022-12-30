package values

import (
	"time"
)

type Date struct {
	t time.Time
}

func ParseDate(input string) (Date, error) {
	t, err := time.Parse(dateFormat, input)
	if err != nil {
		return Date{}, err
	}
	return Date{t: t}, nil
}

func MakeDate(year int, month time.Month, day int) Date {
	return Date{t: time.Date(year, month, day, 0, 0, 0, 0, time.UTC)}
}

const dateFormat = "2006-01-02"

func (d Date) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(dateFormat)+2)
	b = append(b, '"')
	b = d.t.AppendFormat(b, dateFormat)
	b = append(b, '"')
	return b, nil
}

func (d *Date) UnmarshalJSON(data []byte) error {
	value := string(data)
	if value == "null" {
		return nil
	}
	t, err := time.Parse(dateFormat, value[1:len(value)-1])
	if err != nil {
		return err
	}
	d.t = t
	return nil

}

func (d Date) String() string {
	return d.t.Format(dateFormat)
}
