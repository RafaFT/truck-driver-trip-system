package entity_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/rafaft/truck-driver-trip-system/entity"
)

func TestTripTS(t *testing.T) {
	type tripTSInput struct {
		startDate time.Time
		endDate   time.Time
	}

	tests := []struct {
		input        tripTSInput
		wantStart    time.Time
		wantEnd      time.Time
		wantDuration time.Duration
		err          error
	}{
		// invalid cases
		// start date before minTripStartDate
		{
			tripTSInput{
				startDate: time.Date(1968, 1, 1, 0, 0, 0, 0, time.UTC).Add(-time.Microsecond),
				endDate:   time.Now(),
			},
			time.Time{},
			time.Time{},
			0,
			entity.ErrInvalidTripStartDate{},
		},
		// endDate equal to StartDate - 0 second trip is invalid
		{
			tripTSInput{
				startDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				endDate:   time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			time.Time{},
			time.Time{},
			0,
			entity.ErrInvalidTripEndDate{},
		},
		// endDate before StartDate - negative second trip is invalid
		{
			tripTSInput{
				startDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				endDate:   time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC).Add(-time.Second),
			},
			time.Time{},
			time.Time{},
			0,
			entity.ErrInvalidTripEndDate{},
		},
		// valid cases
		// earliest possible one second trip
		{
			tripTSInput{
				startDate: time.Date(1968, 1, 1, 0, 0, 0, 0, time.UTC),
				endDate:   time.Date(1968, 1, 1, 0, 0, 1, 0, time.UTC),
			},
			time.Date(1968, 1, 1, 0, 0, 0, 0, time.UTC),
			time.Date(1968, 1, 1, 0, 0, 1, 0, time.UTC),
			time.Date(1968, 1, 1, 0, 0, 1, 0, time.UTC).Sub(time.Date(1968, 1, 1, 0, 0, 0, 0, time.UTC)),
			nil,
		},
	}

	for i, test := range tests {
		got, gotError := entity.NewTripTS(test.input.startDate, test.input.endDate)

		if reflect.TypeOf(test.err) != reflect.TypeOf(gotError) {
			t.Errorf("%d: [input: %v] [wantError: %v] [gotError: %v]",
				i, test.input, test.err, gotError,
			)
			continue
		}

		if !test.wantStart.Equal(got.Start()) ||
			!test.wantEnd.Equal(got.End()) ||
			test.wantDuration != got.Duration() {
			t.Errorf("%d: [input: %v] [wantStart: %v] [wantEnd: %v] [wantDuration: %v] [got: %v]",
				i, test.input, test.wantStart, test.wantEnd, test.wantDuration, got,
			)
		}
	}
}
