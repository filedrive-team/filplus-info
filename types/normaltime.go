package types

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type NormalTime struct {
	time.Time
}

func NewNormalTime(t time.Time) NormalTime {
	nt := NormalTime{}
	nt.Time = t
	return nt
}

func NewNormalTimeFromUnix(timestamp int64) NormalTime {
	nt := NormalTime{}
	nt.Time = time.Unix(timestamp, 0)
	return nt
}

func (t NormalTime) MarshalJSON() ([]byte, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return []byte(`"-"`), nil
	}
	tune := t.Format(`"2006-01-02 15:04:05"`)
	return []byte(tune), nil
}

func (t *NormalTime) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" || string(data) == `"-"` {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	var err error
	t.Time, err = time.ParseInLocation(`"2006-01-02 15:04:05"`, string(data), time.Local)
	return err
}

func (t NormalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

func (t *NormalTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = NormalTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

func (t *NormalTime) String() string {
	if t.IsEmpty() {
		return ""
	}
	return t.Time.Format("2006-01-02 15:04:05")
}

func (t *NormalTime) YearMonthDate() string {
	if t.IsEmpty() {
		return ""
	}
	return t.Time.Format("2006-01-02")
}

func (t *NormalTime) IsEmpty() bool {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return true
	}
	return false
}
