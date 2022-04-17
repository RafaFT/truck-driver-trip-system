package usecase

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/rafaft/truck-driver-trip-system/entity"
)

type mockUpdateTripRepo struct {
	trip *entity.Trip
	err  error
}

func (r mockUpdateTripRepo) FindByID(ctx context.Context, id string) (*entity.Trip, error) {
	return r.trip, r.err
}

func (r mockUpdateTripRepo) Update(ctx context.Context, trip *entity.Trip) error {
	return r.err
}

func TestUpdateTrip(t *testing.T) {
	networkErr := errors.New("some network error")
	trip, _ := entity.NewTrip(
		"794b9937-7afa-449e-9662-92271d44cb81",
		entity.TripInput{
			CPF:             "67730115077",
			StartDate:       time.Date(2022, 4, 16, 15, 35, 0, 0, time.UTC),
			EndDate:         time.Date(2022, 4, 16, 21, 2, 0, 0, time.UTC),
			HasLoad:         true,
			OriginLat:       -11,
			OriginLong:      -57.6569773,
			DestinationLat:  58.61,
			DestinationLong: -137.9583005,
			VehicleCode:     2,
		},
	)

	tests := []struct {
		tripID  string
		input   UpdateTripInput
		repo    UpdateTripRepo
		want    UpdateTripOutput
		wantErr error
	}{
		// invalid
		{
			tripID:  "", // id invalid
			input:   UpdateTripInput{},
			repo:    mockUpdateTripRepo{},
			want:    UpdateTripOutput{},
			wantErr: entity.ErrInvalidID,
		},
		{
			tripID: "6e900ee0-49af-4e34-83f8-373af7a6bf18",
			input:  UpdateTripInput{},
			repo: mockUpdateTripRepo{
				err: entity.NewErrTripNotFound("6e900ee0-49af-4e34-83f8-373af7a6bf18"),
			},
			want:    UpdateTripOutput{},
			wantErr: entity.NewErrTripNotFound("6e900ee0-49af-4e34-83f8-373af7a6bf18"),
		},
		{
			tripID: "2d601808-ac67-4214-8b65-30635627bc1b",
			input:  UpdateTripInput{},
			repo: mockUpdateTripRepo{
				err: networkErr,
			},
			want:    UpdateTripOutput{},
			wantErr: networkErr,
		},
		{
			tripID: "47bc0c57-adb7-47da-8886-6ff92d484d06",
			input: UpdateTripInput{
				StartDate: getDatePointer(time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)), // invalid
			},
			repo: mockUpdateTripRepo{
				trip: trip,
			},
			want:    UpdateTripOutput{},
			wantErr: entity.NewErrInvalidTripStartDate(time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)),
		},
		{
			tripID: "47bc0c57-adb7-47da-8886-6ff92d484d06",
			input: UpdateTripInput{
				StartDate: getDatePointer(time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)),
				EndDate:   getDatePointer(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)), // invalid
			},
			repo: mockUpdateTripRepo{
				trip: trip,
			},
			want:    UpdateTripOutput{},
			wantErr: entity.NewErrInvalidTripEndDate(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
		},
		{
			tripID: "ed22d58b-46ee-420e-9b1a-d8932c65f2d8",
			input: UpdateTripInput{
				DestinationLat: getFloatPointer(90.1), // invalid
			},
			repo: mockUpdateTripRepo{
				trip: trip,
			},
			want:    UpdateTripOutput{},
			wantErr: entity.NewErrInvalidLatitude(90.1),
		},
		{
			tripID: "ed22d58b-46ee-420e-9b1a-d8932c65f2d8",
			input: UpdateTripInput{
				OriginLong: getFloatPointer(-181), // invalid
			},
			repo: mockUpdateTripRepo{
				trip: trip,
			},
			want:    UpdateTripOutput{},
			wantErr: entity.NewErrInvalidLongitude(-181),
		},
		{
			tripID: "dc86341e-fd07-4c37-8f7e-da3dbf4319d4",
			input: UpdateTripInput{
				VehicleCode: getIntPointer(-5), // invalid
			},
			repo: mockUpdateTripRepo{
				trip: trip,
			},
			want:    UpdateTripOutput{},
			wantErr: entity.NewErrInvalidVehicleCode(-5),
		},
		// valid
		{
			tripID: "794b9937-7afa-449e-9662-92271d44cb81",
			input:  UpdateTripInput{},
			repo: mockUpdateTripRepo{
				trip: trip,
			},
			want: UpdateTripOutput{
				ID:              "794b9937-7afa-449e-9662-92271d44cb81",
				StartDate:       time.Date(2022, 4, 16, 15, 35, 0, 0, time.UTC),
				EndDate:         time.Date(2022, 4, 16, 21, 2, 0, 0, time.UTC),
				Duration:        time.Date(2022, 4, 16, 21, 2, 0, 0, time.UTC).Sub(time.Date(2022, 4, 16, 15, 35, 0, 0, time.UTC)),
				HasLoad:         true,
				OriginLat:       -11,
				OriginLong:      -57.6569773,
				DestinationLat:  58.61,
				DestinationLong: -137.9583005,
				Vehicle:         "STUMP_TRUCK",
			},
			wantErr: nil,
		},
		{
			tripID: "794b9937-7afa-449e-9662-92271d44cb81",
			input: UpdateTripInput{
				StartDate:       getDatePointer(time.Date(2005, 1, 1, 0, 0, 0, 0, time.UTC)),
				EndDate:         getDatePointer(time.Date(2005, 1, 1, 1, 0, 0, 0, time.UTC)),
				HasLoad:         getBoolPointer(false),
				OriginLat:       getFloatPointer(-11),
				OriginLong:      getFloatPointer(-57.65697731), // should not change
				DestinationLat:  getFloatPointer(58.61),
				DestinationLong: getFloatPointer(-137.95830051), // should not change
			},
			repo: mockUpdateTripRepo{
				trip: trip,
			},
			want: UpdateTripOutput{
				ID:              "794b9937-7afa-449e-9662-92271d44cb81",
				StartDate:       time.Date(2005, 1, 1, 0, 0, 0, 0, time.UTC),
				EndDate:         time.Date(2005, 1, 1, 1, 0, 0, 0, time.UTC),
				Duration:        time.Date(2005, 1, 1, 1, 0, 0, 0, time.UTC).Sub(time.Date(2005, 1, 1, 0, 0, 0, 0, time.UTC)),
				HasLoad:         false,
				OriginLat:       -11,
				OriginLong:      -57.6569773,
				DestinationLat:  58.61,
				DestinationLong: -137.9583005,
				Vehicle:         "STUMP_TRUCK",
			},
			wantErr: nil,
		},
	}

	for i, test := range tests {
		uc := NewUpdateTrip(fakeLogger{}, test.repo)
		got, gotErr := uc.Execute(context.Background(), test.tripID, test.input)

		if !errors.Is(gotErr, test.wantErr) {
			t.Errorf("%d: [wantErr: %v] [gotErr: %v]", i, test.wantErr, gotErr)
			continue
		}

		if test.wantErr == nil {
			if got == nil || !reflect.DeepEqual(test.want, *got) {
				t.Errorf("%d: [want: %v] [got: %v]", i, test.want, got)
			}
		}
	}
}
