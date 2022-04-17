package usecase

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/rafaft/truck-driver-trip-system/entity"
)

type mockGetTripRepo struct {
	trip *entity.Trip
	err  error
}

func (r mockGetTripRepo) FindByID(ctx context.Context, id string) (*entity.Trip, error) {
	return r.trip, r.err
}

func TestGetTrip(t *testing.T) {
	trip, _ := entity.NewTrip(
		"47bc0c57-adb7-47da-8886-6ff92d484d06",
		entity.TripInput{
			CPF:             "36428032023",
			StartDate:       time.Date(1998, 2, 28, 17, 42, 0, 0, time.UTC),
			EndDate:         time.Date(1998, 2, 28, 23, 15, 0, 0, time.UTC),
			HasLoad:         true,
			OriginLat:       78.5362,
			OriginLong:      -29.93141,
			DestinationLat:  -73.52,
			DestinationLong: 0,
			VehicleCode:     1,
		},
	)

	tests := []struct {
		id   string
		repo mockGetTripRepo
		want *GetTripOutput
	}{
		{
			"47bc0c57-adb7-47da-8886-6ff92d484d06",
			mockGetTripRepo{
				trip: trip,
			},
			&GetTripOutput{
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
		},
	}

	for i, test := range tests {
		uc := NewGetTrip(fakeLogger{}, test.repo)
		got, gotErr := uc.Execute(context.Background(), test.id)

		if got == nil || gotErr != nil {
			t.Errorf("%d: [input: %v] [wantErr: <nil>] [gotErr: %v]", i, test.id, gotErr)
			continue
		}

		if !reflect.DeepEqual(*got, *test.want) {
			t.Errorf("%d: [want: %v] [got: %v]", i, test.want, got)
		}
	}
}

func TestGetTripErr(t *testing.T) {
	networkErr := errors.New("some repository network error")

	tests := []struct {
		id      string
		repo    mockGetTripRepo
		wantErr error
	}{
		{
			"",
			mockGetTripRepo{},
			entity.ErrInvalidID,
		},
		{
			"someID",
			mockGetTripRepo{
				err: entity.NewErrTripNotFound("someID"),
			},
			entity.NewErrTripNotFound("someID"),
		},
		{
			"validID",
			mockGetTripRepo{
				err: networkErr,
			},
			networkErr,
		},
	}

	for i, test := range tests {
		uc := NewGetTrip(fakeLogger{}, test.repo)
		_, gotErr := uc.Execute(context.Background(), test.id)

		if !errors.Is(gotErr, test.wantErr) {
			t.Errorf("%d: [wantErr: %v] [gotErr: %v]", i, test.wantErr, gotErr)
		}
	}
}
