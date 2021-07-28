package entity

import (
	"time"
)

var minTripStartDate = minBirthDate.AddDate(minimumDriverAge, 0, 0)

type TripTS struct {
	start   time.Time
	end     time.Time
	seconds uint32
}

func NewTripTS(start, end time.Time) (TripTS, error) {
	if !start.After(minTripStartDate) {
		return TripTS{}, newErrInvalidTripStartDate(start)
	}

	if end.Before(start) {
		return TripTS{}, newErrInvalidTripEndDate(end)
	}

	seconds := uint32(end.Sub(start).Seconds())

	return TripTS{
		start:   start,
		end:     end,
		seconds: seconds,
	}, nil
}
