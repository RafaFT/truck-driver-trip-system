package entity

import (
	"time"
)

var minTripStartDate = minBirthDate.AddDate(minimumDriverAge, 0, 0)

type TripTS struct {
	start    time.Time
	end      time.Time
	duration time.Duration
}

func NewTripTS(start, end time.Time) (TripTS, error) {
	if start.Before(minTripStartDate) {
		return TripTS{}, newErrInvalidTripStartDate(start)
	}

	if !end.After(start) {
		return TripTS{}, newErrInvalidTripEndDate(end)
	}

	return TripTS{
		start:    start,
		end:      end,
		duration: end.Sub(start),
	}, nil
}

func (t TripTS) Start() time.Time {
	return t.start
}

func (t TripTS) End() time.Time {
	return t.end
}

func (t TripTS) Duration() time.Duration {
	return t.duration
}
