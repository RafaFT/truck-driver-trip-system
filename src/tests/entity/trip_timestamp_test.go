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
		input       tripTSInput
		wantStart   time.Time
		wantEnd     time.Time
		wantSeconds int
		err         error
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
			1,
			nil,
		},
		// one hour trip
		{
			tripTSInput{
				startDate: time.Date(1968, 1, 1, 18, 0, 0, 0, time.UTC),
				endDate:   time.Date(1968, 1, 1, 19, 0, 0, 0, time.UTC),
			},
			time.Date(1968, 1, 1, 18, 0, 0, 0, time.UTC),
			time.Date(1968, 1, 1, 19, 0, 0, 0, time.UTC),
			3600,
			nil,
		},
		// one day trip
		{
			tripTSInput{
				startDate: time.Date(1968, 1, 1, 21, 0, 0, 0, time.UTC),
				endDate:   time.Date(1968, 1, 2, 21, 0, 0, 0, time.UTC),
			},
			time.Date(1968, 1, 1, 21, 0, 0, 0, time.UTC),
			time.Date(1968, 1, 2, 21, 0, 0, 0, time.UTC),
			86400,
			nil,
		},
		{
			tripTSInput{
				startDate: time.Date(2019, 12, 31, 0, 0, 0, 0, time.UTC),
				endDate:   time.Date(2021, 8, 6, 11, 36, 49, 0, time.UTC),
			},
			time.Date(2019, 12, 31, 0, 0, 0, 0, time.UTC),
			time.Date(2021, 8, 6, 11, 36, 49, 0, time.UTC),
			50499409,
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
			test.wantSeconds != got.Seconds() {
			t.Errorf("%d: [input: %v] [wantStart: %v] [wantEnd: %v] [wantSeconds: %v] [got: %v]",
				i, test.input, test.wantStart, test.wantEnd, test.wantSeconds, got,
			)
		}
	}
}
