package usecase_test

import (
	"context"
	"reflect"
	"testing"
	"time"

	repository "github.com/rafaft/truck-driver-trip-system/adapter/gateway/database"
	"github.com/rafaft/truck-driver-trip-system/entity"
	"github.com/rafaft/truck-driver-trip-system/infrastructure/log"
	"github.com/rafaft/truck-driver-trip-system/tests/samples"
	"github.com/rafaft/truck-driver-trip-system/usecase"
)

func TestGetTripErrors(t *testing.T) {
	logger := log.NewFakeLogger()
	tripRepo := repository.NewTripInMemory(nil)
	uc := usecase.NewGetTrip(logger, tripRepo)

	tests := []struct {
		input string
		err   error
	}{
		{
			input: "",
			err:   entity.ErrInvalidID,
		},
		{
			input: "a676b5ad-5ffa-4917-a62e-d0933e53c1bb",
			err:   entity.ErrTripNotFound{},
		},
	}

	for i, test := range tests {
		_, gotErr := uc.Execute(context.Background(), test.input)

		if reflect.TypeOf(test.err) != reflect.TypeOf(gotErr) {
			t.Errorf("%d: [wantErr: %v] [gotErr: %v]", i, test.err, gotErr)
		}
	}
}

func TestGetTrip(t *testing.T) {
	logger := log.NewFakeLogger()
	tripRepo := repository.NewTripInMemory(samples.GetTrips(t))
	uc := usecase.NewGetTrip(logger, tripRepo)

	tests := []struct {
		input string
		want  usecase.GetTripOutput
		err   error
	}{
		{
			input: "47bc0c57-adb7-47da-8886-6ff92d484d06",
			want: usecase.GetTripOutput{
				ID:              "47bc0c57-adb7-47da-8886-6ff92d484d06",
				StartDate:       time.Date(1998, 2, 28, 17, 42, 0, 0, time.UTC),
				EndDate:         time.Date(1998, 2, 28, 23, 15, 0, 0, time.UTC),
				Duration:        time.Date(1998, 2, 28, 23, 15, 0, 0, time.UTC).Sub(time.Date(1998, 2, 28, 17, 42, 0, 0, time.UTC)),
				HasLoad:         true,
				OriginLat:       78.5362,
				OriginLong:      -29.93141,
				DestinationLat:  -73.52,
				DestinationLong: 0,
				Vehicle:         "3/4Truck",
			},
			err: nil,
		},
		{
			input: "794b9937-7afa-449e-9662-92271d44cb81",
			want: usecase.GetTripOutput{
				ID:              "794b9937-7afa-449e-9662-92271d44cb81",
				StartDate:       time.Date(2004, 7, 31, 21, 12, 0, 0, time.UTC),
				EndDate:         time.Date(2004, 8, 1, 6, 23, 0, 0, time.UTC),
				Duration:        time.Date(2004, 8, 1, 6, 23, 0, 0, time.UTC).Sub(time.Date(2004, 7, 31, 21, 12, 0, 0, time.UTC)),
				HasLoad:         true,
				OriginLat:       -11.7471003,
				OriginLong:      -57.6569773,
				DestinationLat:  58.6049797,
				DestinationLong: -137.9583005,
				Vehicle:         "TRUCK",
			},
			err: nil,
		},
	}

	for i, test := range tests {
		got, gotErr := uc.Execute(context.Background(), test.input)

		if gotErr != nil {
			t.Errorf("%d: unexpected error -> [gotErr: %v]", i, gotErr)
			continue
		}

		if *got != test.want {
			t.Errorf("%d: [want: %v] != [got: %v]", i, test.want, *got)
		}
	}
}
