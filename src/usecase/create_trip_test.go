package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/rafaft/truck-driver-trip-system/entity"
)

type mockCreateTripRepo struct {
	err error
}

func (m mockCreateTripRepo) Save(context.Context, *entity.Trip) error {
	return m.err
}

type mockCreateTripDriverRepo struct {
	err error
}

func (m mockCreateTripDriverRepo) FindByCPF(context.Context, entity.CPF) (*entity.Driver, error) {
	return nil, m.err
}

func TestCreateTrip(t *testing.T) {
	tests := []struct {
		input CreateTripInput
		want  CreateTripOutput
	}{
		{
			CreateTripInput{
				CPF:             "33510345398",
				StartDate:       time.Date(1997, 6, 23, 5, 5, 0, 0, time.UTC),
				EndDate:         time.Date(1997, 6, 23, 14, 37, 0, 0, time.UTC),
				HasLoad:         false,
				OriginLat:       -32.3736,
				OriginLong:      31.0709193,
				DestinationLat:  -44.51865,
				DestinationLong: 26.61581,
				VehicleCode:     0,
			},
			CreateTripOutput{
				ID:              "", // random UUID
				StartDate:       time.Date(1997, 6, 23, 5, 5, 0, 0, time.UTC),
				EndDate:         time.Date(1997, 6, 23, 14, 37, 0, 0, time.UTC),
				Duration:        time.Date(1997, 6, 23, 14, 37, 0, 0, time.UTC).Sub(time.Date(1997, 6, 23, 5, 5, 0, 0, time.UTC)),
				HasLoad:         false,
				OriginLat:       -32.3736,
				OriginLong:      31.0709193,
				DestinationLat:  -44.51865,
				DestinationLong: 26.61581,
				Vehicle:         string(entity.Truck),
			},
		},
	}

	for i, test := range tests {
		uc := NewCreateTrip(FakeLogger{}, mockCreateTripRepo{}, mockCreateTripDriverRepo{})
		got, gotErr := uc.Execute(context.Background(), test.input)

		if got == nil || gotErr != nil {
			t.Errorf("%d: [input: %v] [wantErr: <nil>] [gotErr: %v]", i, test.input, gotErr)
			continue
		}

		got.ID = ""
		if test.want != *got {
			t.Errorf("%d: [input: %v] [want: %v] [got: %v]", i, test.input, test.want, got)
		}
	}
}

func TestCreateTripErr(t *testing.T) {
	tests := []struct {
		input      CreateTripInput
		driverRepo mockCreateTripDriverRepo
		wantErr    error
	}{
		{
			CreateTripInput{
				CPF:             "12345678901", // invalid CPF
				StartDate:       time.Date(2000, 1, 1, 18, 25, 0, 0, time.UTC),
				EndDate:         time.Date(2000, 1, 1, 22, 55, 0, 0, time.UTC),
				HasLoad:         true,
				OriginLat:       0,
				OriginLong:      0,
				DestinationLat:  0,
				DestinationLong: 0,
				VehicleCode:     1,
			},
			mockCreateTripDriverRepo{},
			entity.NewErrInvalidCPF("12345678901"),
		},
		{
			CreateTripInput{
				CPF:             "74658482029",
				StartDate:       time.Date(1950, 1, 1, 18, 25, 0, 0, time.UTC), // invalid
				EndDate:         time.Date(2000, 1, 1, 22, 55, 0, 0, time.UTC),
				HasLoad:         true,
				OriginLat:       90,
				OriginLong:      0,
				DestinationLat:  0,
				DestinationLong: 0,
				VehicleCode:     1,
			},
			mockCreateTripDriverRepo{},
			entity.NewErrInvalidTripStartDate(time.Date(1950, 1, 1, 18, 25, 0, 0, time.UTC)),
		},
		{
			CreateTripInput{
				CPF:             "74658482029",
				StartDate:       time.Date(2000, 1, 1, 18, 25, 0, 0, time.UTC),
				EndDate:         time.Date(1999, 1, 1, 22, 55, 0, 0, time.UTC), // invalid
				HasLoad:         true,
				OriginLat:       90,
				OriginLong:      0,
				DestinationLat:  0,
				DestinationLong: 0,
				VehicleCode:     1,
			},
			mockCreateTripDriverRepo{},
			entity.NewErrInvalidTripEndDate(time.Date(1999, 1, 1, 22, 55, 0, 0, time.UTC)),
		},
		{
			CreateTripInput{
				CPF:             "74658482029",
				StartDate:       time.Date(2000, 1, 1, 18, 25, 0, 0, time.UTC),
				EndDate:         time.Date(2000, 1, 1, 22, 55, 0, 0, time.UTC),
				HasLoad:         true,
				OriginLat:       91, // invalid
				OriginLong:      0,
				DestinationLat:  0,
				DestinationLong: 0,
				VehicleCode:     1,
			},
			mockCreateTripDriverRepo{},
			entity.NewErrInvalidLatitude(91),
		},
		{
			CreateTripInput{
				CPF:             "74658482029",
				StartDate:       time.Date(2000, 1, 1, 18, 25, 0, 0, time.UTC),
				EndDate:         time.Date(2000, 1, 1, 22, 55, 0, 0, time.UTC),
				HasLoad:         true,
				OriginLat:       90,
				OriginLong:      0,
				DestinationLat:  0,
				DestinationLong: -181, // invalid
				VehicleCode:     1,
			},
			mockCreateTripDriverRepo{},
			entity.NewErrInvalidLongitude(-181),
		},
		{
			CreateTripInput{
				CPF:             "74658482029",
				StartDate:       time.Date(2000, 1, 1, 18, 25, 0, 0, time.UTC),
				EndDate:         time.Date(2000, 1, 1, 22, 55, 0, 0, time.UTC),
				HasLoad:         true,
				OriginLat:       90,
				OriginLong:      0,
				DestinationLat:  0,
				DestinationLong: 0,
				VehicleCode:     -5, // invalid
			},
			mockCreateTripDriverRepo{},
			entity.NewErrInvalidVehicleCode(-5),
		},
		{
			CreateTripInput{
				CPF:             "74658482029",
				StartDate:       time.Date(2000, 1, 1, 18, 25, 0, 0, time.UTC),
				EndDate:         time.Date(2000, 1, 1, 22, 55, 0, 0, time.UTC),
				HasLoad:         true,
				OriginLat:       0,
				OriginLong:      0,
				DestinationLat:  0,
				DestinationLong: 0,
				VehicleCode:     1,
			},
			mockCreateTripDriverRepo{
				err: entity.NewErrDriverNotFound("74658482029"),
			},
			entity.NewErrDriverNotFound("74658482029"),
		},
	}

	for i, test := range tests {
		uc := NewCreateTrip(FakeLogger{}, mockCreateTripRepo{}, test.driverRepo)
		_, gotErr := uc.Execute(context.Background(), test.input)

		if !errors.Is(gotErr, test.wantErr) {
			t.Errorf("%d: [input: %v] [wantErr: %v] [gotErr: %v]", i, test.input, test.wantErr, gotErr)
		}
	}
}
