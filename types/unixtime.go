package types

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"time"
)

type UnixTime struct {
	time.Time
}

func NewUnixTime(t time.Time) NormalTime {
	nt := NormalTime{}
	nt.Time = t
	return nt
}

func NewUnixTimeFromUnix(timestamp int64) NormalTime {
	nt := NormalTime{}
	nt.Time = time.Unix(timestamp, 0)
	return nt
}

func (t UnixTime) MarshalJSON() ([]byte, error) {
	var zeroTime time.Time
	if t.Time.Unix() == zeroTime.Unix() {
		return []byte(`"-"`), nil
	}
	tune := fmt.Sprint(t.Unix())
	return []byte(tune), nil
}

func (t *UnixTime) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" || string(data) == `"-"` {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	v, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		t.Time = time.Unix(v, 0)
	}
	return err
}

func (t UnixTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.Unix() == zeroTime.Unix() {
		return nil, nil
	}
	return t.Time, nil
}

func (t *UnixTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = UnixTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

func (t *UnixTime) String() string {
	if t.IsEmpty() {
		return ""
	}
	return t.Time.Format("2006-01-02 15:04:05")
}

func (t *UnixTime) YearMonthDate() string {
	if t.IsEmpty() {
		return ""
	}
	return t.Time.Format("2006-01-02")
}

func (t *UnixTime) IsEmpty() bool {
	var zeroTime time.Time
	if t.Time.Unix() == zeroTime.Unix() {
		return true
	}
	return false
}
