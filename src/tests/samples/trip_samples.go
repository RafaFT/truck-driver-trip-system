package samples

import (
	"testing"
	"time"

	"github.com/rafaft/truck-driver-trip-system/entity"
)

func GetTrips(t *testing.T) []*entity.Trip {
	tripsInput := []struct {
		ID              string
		DriverCPF       string
		StartDate       time.Time
		EndDate         time.Time
		HasLoad         bool
		OriginLat       float64
		OriginLong      float64
		DestinationLat  float64
		DestinationLong float64
		VehicleCode     int
	}{
		{
			ID:              "47bc0c57-adb7-47da-8886-6ff92d484d06",
			DriverCPF:       "51598390031",
			StartDate:       time.Date(1998, 2, 28, 17, 42, 0, 0, time.UTC),
			EndDate:         time.Date(1998, 2, 28, 23, 15, 0, 0, time.UTC),
			HasLoad:         true,
			OriginLat:       78.5362,
			OriginLong:      -29.93141,
			DestinationLat:  -73.52,
			DestinationLong: 0,
			VehicleCode:     1,
		},
		{
			ID:              "ed22d58b-46ee-420e-9b1a-d8932c65f2d8",
			DriverCPF:       "51598390031",
			StartDate:       time.Date(1998, 3, 1, 8, 0, 0, 0, time.UTC),
			EndDate:         time.Date(1998, 3, 1, 16, 51, 0, 0, time.UTC),
			HasLoad:         true,
			OriginLat:       -73.52,
			OriginLong:      0,
			DestinationLat:  -25.0,
			DestinationLong: 117.54321,
			VehicleCode:     1,
		},
		{
			ID:              "dc86341e-fd07-4c37-8f7e-da3dbf4319d4",
			DriverCPF:       "76812015059",
			StartDate:       time.Date(2001, 4, 7, 0, 0, 0, 0, time.UTC),
			EndDate:         time.Date(2001, 4, 7, 12, 30, 0, 0, time.UTC),
			HasLoad:         false,
			OriginLat:       39.2202146,
			OriginLong:      -25,
			DestinationLat:  -31.1769614,
			DestinationLong: -49.5604080,
			VehicleCode:     2,
		},
		{
			ID:              "794b9937-7afa-449e-9662-92271d44cb81",
			DriverCPF:       "51598390031",
			StartDate:       time.Date(2004, 7, 31, 21, 12, 0, 0, time.UTC),
			EndDate:         time.Date(2004, 8, 1, 6, 23, 0, 0, time.UTC),
			HasLoad:         true,
			OriginLat:       -11.7471003,
			OriginLong:      -57.6569773,
			DestinationLat:  58.6049797,
			DestinationLong: -137.9583005,
			VehicleCode:     0,
		},
	}

	trips := make([]*entity.Trip, len(tripsInput))
	for i, input := range tripsInput {
		trip, err := entity.NewTrip(
			input.ID,
			entity.TripInput{
				CPF:             input.DriverCPF,
				StartDate:       input.StartDate,
				EndDate:         input.EndDate,
				HasLoad:         input.HasLoad,
				OriginLat:       input.OriginLat,
				OriginLong:      input.OriginLong,
				DestinationLat:  input.DestinationLat,
				DestinationLong: input.DestinationLong,
				VehicleCode:     input.VehicleCode,
			},
		)

		if err != nil {
			t.Fatalf("%d: could not generate trip. [input: %v] [gotErr: %v]", i, input, err)
		}

		trips[i] = trip
	}

	return trips
}
