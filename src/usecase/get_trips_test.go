package usecase

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/rafaft/truck-driver-trip-system/entity"
)

type mockGetTripsRepo struct {
	trips []*entity.Trip
	err   error
}

func (r mockGetTripsRepo) Find(ctx context.Context, query FindTripsQuery) ([]*entity.Trip, error) {
	return r.trips, r.err
}

func TestGetTrips(t *testing.T) {
	now := time.Now().UTC()
	trip1, _ := entity.NewTrip(
		"ID1",
		entity.TripInput{
			CPF:             "31519028040",
			StartDate:       now.AddDate(-1, 0, -1),
			EndDate:         now.AddDate(-1, 0, 0),
			HasLoad:         true,
			OriginLat:       -23.5,
			OriginLong:      -46.6,
			DestinationLat:  10,
			DestinationLong: 11,
			VehicleCode:     0,
		},
	)
	trip2, _ := entity.NewTrip(
		"ID2",
		entity.TripInput{
			CPF:             "95580146051",
			StartDate:       now.AddDate(-2, -1, -1),
			EndDate:         now.AddDate(-2, -1, 0),
			HasLoad:         false,
			OriginLat:       1,
			OriginLong:      2,
			DestinationLat:  3,
			DestinationLong: 4,
			VehicleCode:     1,
		},
	)

	tests := []struct {
		query GetTripsInput
		repo  GetTripsRepo
		want  []*GetTripsOutput
	}{
		{
			GetTripsInput{},
			mockGetTripsRepo{},
			[]*GetTripsOutput{},
		},
		{
			GetTripsInput{
				Limit: getUintPointer(10),
			},
			mockGetTripsRepo{
				trips: []*entity.Trip{
					trip1,
					trip2,
				},
			},
			[]*GetTripsOutput{
				{
					ID:              "ID1",
					StartDate:       now.AddDate(-1, 0, -1),
					EndDate:         now.AddDate(-1, 0, 0),
					Duration:        now.AddDate(-1, 0, 0).Sub(now.AddDate(-1, 0, -1)),
					HasLoad:         true,
					OriginLat:       -23.5,
					OriginLong:      -46.6,
					DestinationLat:  10,
					DestinationLong: 11,
					Vehicle:         "TRUCK",
				},
				{
					ID:              "ID2",
					StartDate:       now.AddDate(-2, -1, -1),
					EndDate:         now.AddDate(-2, -1, 0),
					Duration:        now.AddDate(-2, -1, 0).Sub(now.AddDate(-2, -1, -1)),
					HasLoad:         false,
					OriginLat:       1,
					OriginLong:      2,
					DestinationLat:  3,
					DestinationLong: 4,
					Vehicle:         "3/4Truck",
				},
			},
		},
	}

	for i, test := range tests {
		uc := NewGetTrips(FakeLogger{}, test.repo)
		got, gotErr := uc.Execute(context.Background(), test.query)

		if got == nil || gotErr != nil {
			t.Errorf("%d: [input: %v] [wantErr: <nil>] [gotErr: %v]", i, test.query, gotErr)
			continue
		}

		if !reflect.DeepEqual(test.want, got) {
			t.Errorf("%d: [input: %v] [want: %v] [got: %v]", i, test.query, test.want, got)
		}
	}
}

func TestGetTripsErr(t *testing.T) {
	networkError := errors.New("some repo network error")

	tests := []struct {
		query   GetTripsInput
		repo    GetTripsRepo
		wantErr error
	}{
		{
			GetTripsInput{
				CPF: getStrPointer("12345678901"),
			},
			mockGetTripsRepo{},
			entity.NewErrInvalidCPF("12345678901"),
		},
		{
			GetTripsInput{
				VehicleCode: getIntPointer(-3),
			},
			mockGetTripsRepo{},
			entity.NewErrInvalidVehicleCode(-3),
		},
		{
			GetTripsInput{},
			mockGetTripsRepo{
				err: networkError,
			},
			networkError,
		},
	}

	for i, test := range tests {
		uc := NewGetTrips(FakeLogger{}, test.repo)
		_, gotErr := uc.Execute(context.Background(), test.query)

		if !errors.Is(gotErr, test.wantErr) {
			t.Errorf("%d: [wantErr: %v] [gotErr: %v]", i, test.wantErr, gotErr)
		}
	}
}
