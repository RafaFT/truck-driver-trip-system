package entity

import (
	"time"
)

var minTripStartDate = minBirthDate.AddDate(minimumDriverAge, 0, 0)

type tripTS struct {
	start    time.Time
	end      time.Time
	duration time.Duration
}

func newTripTS(start, end time.Time) (tripTS, error) {
	if start.Before(minTripStartDate) {
		return tripTS{}, NewErrInvalidTripStartDate(start)
	}

	if !end.After(start) {
		return tripTS{}, NewErrInvalidTripEndDate(end)
	}

	return tripTS{
		start:    start,
		end:      end,
		duration: end.Sub(start),
	}, nil
}
