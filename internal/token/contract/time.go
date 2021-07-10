package contract

import (
	"encoding/json"
	"time"
)

// Time define costume time type
type Time struct {
	time.Time
}

// TimeNow is a function to return current time
func TimeNow() *Time {
	return &Time{time.Now()}
}

// AddDates is a function to return current time plus dates
func (t *Time) AddDates(year, month, day int) *Time {
	return &Time{t.AddDate(year, month, day)}
}

// UnixTime is a function to convert timestamp to unix format
func (t *Time) UnixTime() int64 {
	return t.Unix()
}

// MarshalJSON is a function to marshaling field time
func (t *Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Unix())
}

// UnmarshalJSON is a function to unmarshaling function for time-related claims.
func (t *Time) UnmarshalJSON(b []byte) error {
	var unix *int64
	if err := json.Unmarshal(b, &unix); err != nil {
		return err
	}
	if unix == nil {
		return nil
	}
	t.Time = time.Unix(*unix, 0)
	return nil
}
