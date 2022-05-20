package models

import (
	"database/sql/driver"
	"errors"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
)

type NullTime struct {
	time.Time
	Valid bool
}

func (nt *NullTime) Scan(value interface{}) error {
	nt.Time, nt.Valid = value.(time.Time)

	return nil
}

func (nt NullTime) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, errors.New("timestamp provided is not valid")
	}

	return nt.Time, nil
}

// TimestampToTime returns a Time object from a protobuf Timestamp
func TimestampToTime(ts *timestamp.Timestamp) time.Time {
	if ts == nil {
		return time.Time{}
	}
	t, err := ptypes.Timestamp(ts)
	if err != nil {
		log.Printf("Error converting timestamp %v", err)
	}

	return t
}

// TimestampToNullTime returns a models.NullTime from a protobuf Timestamp
func TimestampToNullTime(ts *timestamp.Timestamp) NullTime {
	return TimeToNullTime(TimestampToTime(ts))
}

// TimeToNullTime coverts a Time into a NullTime
func TimeToNullTime(t time.Time) NullTime {
	if t.IsZero() {
		return NullTime{}
	}

	return NullTime{
		Time:  t,
		Valid: true,
	}
}
