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

func TestCreateTripErrors(t *testing.T) {
	now := time.Now()
	logger := log.NewFakeLogger()
	tripRepo := repository.NewTripInMemory(nil)
	driverRepo := repository.NewDriverInMemory(nil)
	uc := usecase.NewCreateTrip(logger, tripRepo, driverRepo)

	tests := []struct {
		input usecase.CreateTripInput
		want  *usecase.CreateTripOutput
		err   error
	}{
		{
			input: usecase.CreateTripInput{
				CPF:             "42265277040", // invalid
				StartDate:       now.AddDate(0, 0, -1),
				EndDate:         now,
				HasLoad:         true,
				OriginLat:       0,
				OriginLong:      0,
				DestinationLat:  0,
				DestinationLong: 0,
				VehicleCode:     1,
			},
			want: nil,
			err:  entity.ErrInvalidCPF{},
		},
		{
			input: usecase.CreateTripInput{
				CPF:             "78066022085",
				StartDate:       time.Date(1968, 1, 1, 0, 0, 0, 0, time.UTC).AddDate(0, 0, -1), // invalid
				EndDate:         now,
				HasLoad:         true,
				OriginLat:       0,
				OriginLong:      0,
				DestinationLat:  0,
				DestinationLong: 0,
				VehicleCode:     1,
			},
			want: nil,
			err:  entity.ErrInvalidTripStartDate{},
		},
		{
			input: usecase.CreateTripInput{
				CPF:             "10259412090",
				StartDate:       now,
				EndDate:         now, // invalid
				HasLoad:         true,
				OriginLat:       0,
				OriginLong:      0,
				DestinationLat:  0,
				DestinationLong: 0,
				VehicleCode:     1,
			},
			want: nil,
			err:  entity.ErrInvalidTripEndDate{},
		},
		{
			input: usecase.CreateTripInput{
				CPF:             "90106226061",
				StartDate:       now.AddDate(0, 0, -1),
				EndDate:         now,
				HasLoad:         true,
				OriginLat:       -91, // invalid
				OriginLong:      0,
				DestinationLat:  0,
				DestinationLong: 0,
				VehicleCode:     1,
			},
			want: nil,
			err:  entity.ErrInvalidLatitude{},
		},
		{
			input: usecase.CreateTripInput{
				CPF:             "82905682078",
				StartDate:       now.AddDate(0, 0, -1),
				EndDate:         now,
				HasLoad:         true,
				OriginLat:       0,
				OriginLong:      -181, // invalid
				DestinationLat:  0,
				DestinationLong: 0,
				VehicleCode:     1,
			},
			want: nil,
			err:  entity.ErrInvalidLongitude{},
		},
		{
			input: usecase.CreateTripInput{
				CPF:             "51362416088",
				StartDate:       now.AddDate(0, 0, -1),
				EndDate:         now,
				HasLoad:         true,
				OriginLat:       0,
				OriginLong:      0,
				DestinationLat:  91, // invalid
				DestinationLong: 0,
				VehicleCode:     1,
			},
			want: nil,
			err:  entity.ErrInvalidLatitude{},
		},
		{
			input: usecase.CreateTripInput{
				CPF:             "77297976075",
				StartDate:       now.AddDate(0, 0, -1),
				EndDate:         now,
				HasLoad:         true,
				OriginLat:       0,
				OriginLong:      0,
				DestinationLat:  0,
				DestinationLong: -181, // invalid
				VehicleCode:     1,
			},
			want: nil,
			err:  entity.ErrInvalidLongitude{},
		},
		{
			input: usecase.CreateTripInput{
				CPF:             "64947151099",
				StartDate:       now.AddDate(0, 0, -1),
				EndDate:         now,
				HasLoad:         true,
				OriginLat:       0,
				OriginLong:      0,
				DestinationLat:  0,
				DestinationLong: 0,
				VehicleCode:     35,
			},
			want: nil,
			err:  entity.ErrInvalidVehicleCode{},
		},
		{
			input: usecase.CreateTripInput{
				CPF:             "32041794003", // there is no driver with this CPF
				StartDate:       now.AddDate(0, 0, -1),
				EndDate:         now,
				HasLoad:         true,
				OriginLat:       0,
				OriginLong:      0,
				DestinationLat:  0,
				DestinationLong: 0,
				VehicleCode:     1,
			},
			want: nil,
			err:  entity.ErrDriverNotFound{},
		},
	}

	for i, test := range tests {
		_, gotErr := uc.Execute(context.Background(), test.input)

		if reflect.TypeOf(test.err) != reflect.TypeOf(gotErr) {
			t.Errorf("%d: [err: %v] [gotErr: %v]", i, test.err, gotErr)
			continue
		}
	}
}

func TestCreateTrip(t *testing.T) {
	logger := log.NewFakeLogger()
	tripRepo := repository.NewTripInMemory(nil)
	driverRepo := repository.NewDriverInMemory(samples.GetDrivers(t))
	uc := usecase.NewCreateTrip(logger, tripRepo, driverRepo)

	tests := []struct {
		input usecase.CreateTripInput
		want  *usecase.CreateTripOutput
		err   error
	}{
		{
			input: usecase.CreateTripInput{
				CPF:             "33510345398",
				StartDate:       time.Date(2000, 1, 1, 18, 25, 0, 0, time.UTC),
				EndDate:         time.Date(2000, 1, 1, 22, 55, 0, 0, time.UTC),
				HasLoad:         true,
				OriginLat:       0,
				OriginLong:      0,
				DestinationLat:  0,
				DestinationLong: 0,
				VehicleCode:     1,
			},
			want: &usecase.CreateTripOutput{
				ID:              "", // random UUID
				StartDate:       time.Date(2000, 1, 1, 18, 25, 0, 0, time.UTC),
				EndDate:         time.Date(2000, 1, 1, 22, 55, 0, 0, time.UTC),
				Duration:        time.Date(2000, 1, 1, 22, 55, 0, 0, time.UTC).Sub(time.Date(2000, 1, 1, 18, 25, 0, 0, time.UTC)),
				HasLoad:         true,
				OriginLat:       0,
				OriginLong:      0,
				DestinationLat:  0,
				DestinationLong: 0,
				Vehicle:         string(entity.Truck_34),
			},
			err: nil,
		},
		{
			input: usecase.CreateTripInput{
				CPF:             "52742089403",
				StartDate:       time.Date(2000, 12, 25, 04, 57, 0, 0, time.UTC),
				EndDate:         time.Date(2021, 10, 1, 22, 55, 0, 0, time.UTC),
				HasLoad:         false,
				OriginLat:       87.1234567,
				OriginLong:      -4.9876543,
				DestinationLat:  14.564738,
				DestinationLong: -179.98765,
				VehicleCode:     2,
			},
			want: &usecase.CreateTripOutput{
				ID:              "", // random UUID
				StartDate:       time.Date(2000, 12, 25, 04, 57, 0, 0, time.UTC),
				EndDate:         time.Date(2021, 10, 1, 22, 55, 0, 0, time.UTC),
				Duration:        time.Date(2021, 10, 1, 22, 55, 0, 0, time.UTC).Sub(time.Date(2000, 12, 25, 04, 57, 0, 0, time.UTC)),
				HasLoad:         false,
				OriginLat:       87.1234567,
				OriginLong:      -4.9876543,
				DestinationLat:  14.564738,
				DestinationLong: -179.98765,
				Vehicle:         string(entity.StumpTruck),
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

		if test.want.HasLoad != got.HasLoad ||
			test.want.OriginLat != got.OriginLat ||
			test.want.OriginLong != got.OriginLong ||
			test.want.DestinationLat != got.DestinationLat ||
			test.want.DestinationLong != got.DestinationLong ||
			test.want.Vehicle != got.Vehicle ||
			!test.want.StartDate.Equal(got.StartDate) ||
			!test.want.EndDate.Equal(got.EndDate) ||
			test.want.Duration != got.Duration {
			t.Errorf("%d: [input: %v] [want: %v] [got: %v]", i, test.input, test.want, got)
		}
	}
}
