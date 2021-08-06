package entity

import (
	"time"
)

var minTripStartDate = minBirthDate.AddDate(minimumDriverAge, 0, 0)

type TripTS struct {
	start   time.Time
	end     time.Time
	seconds int
}

func NewTripTS(start, end time.Time) (TripTS, error) {
	if start.Before(minTripStartDate) {
		return TripTS{}, newErrInvalidTripStartDate(start)
	}

	if !end.After(start) {
		return TripTS{}, newErrInvalidTripEndDate(end)
	}

	seconds := int(end.Sub(start).Seconds())

	return TripTS{
		start:   start,
		end:     end,
		seconds: seconds,
	}, nil
}

func (t TripTS) Start() time.Time {
	return t.start
}

func (t TripTS) End() time.Time {
	return t.end
}

func (t TripTS) Seconds() int {
	return t.seconds
}
