package entity

import (
	"time"
)

var minTripStartDate = minBirthDate.AddDate(minimumDriverAge, 0, 0)

type TripTimeStamp struct {
	start   time.Time
	end     time.Time
	seconds int64
}

func NewTripTimeStamp(start, end time.Time) (TripTimeStamp, error) {
	if !start.After(minTripStartDate) {
		return TripTimeStamp{}, newErrInvalidTripStartDate(start)
	}

	if end.Before(start) {
		return TripTimeStamp{}, newErrInvalidTripEndDate(end)
	}

	seconds := int64(end.Sub(start).Seconds())

	return TripTimeStamp{
		start:   start,
		end:     end,
		seconds: seconds,
	}, nil
}
