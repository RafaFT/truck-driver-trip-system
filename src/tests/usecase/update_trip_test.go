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

func TestUpdateTripErrors(t *testing.T) {
	logger := log.NewFakeLogger()
	tripRepo := repository.NewTripInMemory(samples.GetTrips(t))
	uc := usecase.NewUpdateTrip(logger, tripRepo)

	tests := []struct {
		tripID string
		input  usecase.UpdateTripInput
		err    error
	}{
		{
			tripID: "", // id invalid
			input:  usecase.UpdateTripInput{},
			err:    entity.ErrInvalidID,
		},
		{
			tripID: "6e900ee0-49af-4e34-83f8-373af7a6bf18",
			input:  usecase.UpdateTripInput{},
			err:    entity.ErrTripNotFound{},
		},
		{
			tripID: "47bc0c57-adb7-47da-8886-6ff92d484d06",
			input: usecase.UpdateTripInput{
				StartDate: getDatePointer(time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)), // invalid
			},
			err: entity.ErrInvalidTripStartDate{},
		},
		{
			tripID: "47bc0c57-adb7-47da-8886-6ff92d484d06",
			input: usecase.UpdateTripInput{
				StartDate: getDatePointer(time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)),
				EndDate:   getDatePointer(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)), // invalid
			},
			err: entity.ErrInvalidTripEndDate{},
		},
		{
			tripID: "ed22d58b-46ee-420e-9b1a-d8932c65f2d8",
			input: usecase.UpdateTripInput{
				OriginLat: getFloatPointer(-90.1), // invalid
			},
			err: entity.ErrInvalidLatitude{},
		},
		{
			tripID: "ed22d58b-46ee-420e-9b1a-d8932c65f2d8",
			input: usecase.UpdateTripInput{
				DestinationLat: getFloatPointer(90.1), // invalid
			},
			err: entity.ErrInvalidLatitude{},
		},
		{
			tripID: "ed22d58b-46ee-420e-9b1a-d8932c65f2d8",
			input: usecase.UpdateTripInput{
				OriginLong: getFloatPointer(-181), // invalid
			},
			err: entity.ErrInvalidLongitude{},
		},
		{
			tripID: "ed22d58b-46ee-420e-9b1a-d8932c65f2d8",
			input: usecase.UpdateTripInput{
				DestinationLong: getFloatPointer(180.0000001), // invalid
			},
			err: entity.ErrInvalidLongitude{},
		},
		{
			tripID: "dc86341e-fd07-4c37-8f7e-da3dbf4319d4",
			input: usecase.UpdateTripInput{
				VehicleCode: getIntPointer(-5), // invalid
			},
			err: entity.ErrInvalidVehicleCode{},
		},
	}

	for i, test := range tests {
		_, gotErr := uc.Execute(context.Background(), test.tripID, test.input)

		if reflect.TypeOf(test.err) != reflect.TypeOf(gotErr) {
			t.Errorf("%d: [wantErr: %v] [gotErr: %v]", i, test.err, gotErr)
		}
	}
}

func TestUpdateTrip(t *testing.T) {
	logger := log.NewFakeLogger()
	tripRepo := repository.NewTripInMemory(samples.GetTrips(t))
	uc := usecase.NewUpdateTrip(logger, tripRepo)

	tests := []struct {
		tripID string
		input  usecase.UpdateTripInput
		want   usecase.UpdateTripOutput
		err    error
	}{
		{ // empty update
			tripID: "47bc0c57-adb7-47da-8886-6ff92d484d06",
			input:  usecase.UpdateTripInput{},
			want: usecase.UpdateTripOutput{
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
		{ // start date
			tripID: "47bc0c57-adb7-47da-8886-6ff92d484d06",
			input: usecase.UpdateTripInput{
				StartDate: getDatePointer(time.Date(1998, 2, 28, 23, 14, 59, 0, time.UTC)),
			},
			want: usecase.UpdateTripOutput{
				ID:              "47bc0c57-adb7-47da-8886-6ff92d484d06",
				StartDate:       time.Date(1998, 2, 28, 23, 14, 59, 0, time.UTC),
				EndDate:         time.Date(1998, 2, 28, 23, 15, 0, 0, time.UTC),
				Duration:        time.Date(1998, 2, 28, 23, 15, 0, 0, time.UTC).Sub(time.Date(1998, 2, 28, 23, 14, 59, 0, time.UTC)),
				HasLoad:         true,
				OriginLat:       78.5362,
				OriginLong:      -29.93141,
				DestinationLat:  -73.52,
				DestinationLong: 0,
				Vehicle:         "3/4Truck",
			},
		},
		{ // end date
			tripID: "47bc0c57-adb7-47da-8886-6ff92d484d06",
			input: usecase.UpdateTripInput{
				EndDate: getDatePointer(time.Date(1998, 2, 28, 23, 16, 00, 0, time.UTC)),
			},
			want: usecase.UpdateTripOutput{
				ID:              "47bc0c57-adb7-47da-8886-6ff92d484d06",
				StartDate:       time.Date(1998, 2, 28, 23, 14, 59, 0, time.UTC),
				EndDate:         time.Date(1998, 2, 28, 23, 16, 0, 0, time.UTC),
				Duration:        time.Date(1998, 2, 28, 23, 16, 0, 0, time.UTC).Sub(time.Date(1998, 2, 28, 23, 14, 59, 0, time.UTC)),
				HasLoad:         true,
				OriginLat:       78.5362,
				OriginLong:      -29.93141,
				DestinationLat:  -73.52,
				DestinationLong: 0,
				Vehicle:         "3/4Truck",
			},
		},
		{ // start and end date
			tripID: "47bc0c57-adb7-47da-8886-6ff92d484d06",
			input: usecase.UpdateTripInput{
				StartDate: getDatePointer(time.Date(1998, 3, 28, 17, 42, 00, 0, time.UTC)),
				EndDate:   getDatePointer(time.Date(1998, 3, 28, 23, 15, 00, 0, time.UTC)),
			},
			want: usecase.UpdateTripOutput{
				ID:              "47bc0c57-adb7-47da-8886-6ff92d484d06",
				StartDate:       time.Date(1998, 3, 28, 17, 42, 00, 0, time.UTC),
				EndDate:         time.Date(1998, 3, 28, 23, 15, 00, 0, time.UTC),
				Duration:        time.Date(1998, 3, 28, 23, 15, 00, 0, time.UTC).Sub(time.Date(1998, 3, 28, 17, 42, 00, 0, time.UTC)),
				HasLoad:         true,
				OriginLat:       78.5362,
				OriginLong:      -29.93141,
				DestinationLat:  -73.52,
				DestinationLong: 0,
				Vehicle:         "3/4Truck",
			},
		},
		{ // has load
			tripID: "ed22d58b-46ee-420e-9b1a-d8932c65f2d8",
			input: usecase.UpdateTripInput{
				HasLoad: getBoolPointer(false),
			},
			want: usecase.UpdateTripOutput{
				ID:              "ed22d58b-46ee-420e-9b1a-d8932c65f2d8",
				StartDate:       time.Date(1998, 3, 1, 8, 0, 0, 0, time.UTC),
				EndDate:         time.Date(1998, 3, 1, 16, 51, 0, 0, time.UTC),
				Duration:        time.Date(1998, 3, 1, 16, 51, 0, 0, time.UTC).Sub(time.Date(1998, 3, 1, 8, 0, 0, 0, time.UTC)),
				HasLoad:         false,
				OriginLat:       -73.52,
				OriginLong:      0,
				DestinationLat:  -25.0,
				DestinationLong: 117.54321,
				Vehicle:         "3/4Truck",
			},
		},
		{ // origin lat and destination long
			tripID: "dc86341e-fd07-4c37-8f7e-da3dbf4319d4",
			input: usecase.UpdateTripInput{
				OriginLat:       getFloatPointer(38.123456789), // extra 8th and 9th digit should be ignored
				DestinationLong: getFloatPointer(-49),
			},
			want: usecase.UpdateTripOutput{
				ID:              "dc86341e-fd07-4c37-8f7e-da3dbf4319d4",
				StartDate:       time.Date(2001, 4, 7, 0, 0, 0, 0, time.UTC),
				EndDate:         time.Date(2001, 4, 7, 12, 30, 0, 0, time.UTC),
				Duration:        time.Date(2001, 4, 7, 12, 30, 0, 0, time.UTC).Sub(time.Date(2001, 4, 7, 0, 0, 0, 0, time.UTC)),
				HasLoad:         false,
				OriginLat:       38.1234567,
				OriginLong:      -25,
				DestinationLat:  -31.1769614,
				DestinationLong: -49,
				Vehicle:         "STUMP_TRUCK",
			},
		},
		{ // origin long and destination lat
			tripID: "dc86341e-fd07-4c37-8f7e-da3dbf4319d4",
			input: usecase.UpdateTripInput{
				OriginLong:     getFloatPointer(-25.1234567),
				DestinationLat: getFloatPointer(-31.00000001), // 8th digit should be ignored
			},
			want: usecase.UpdateTripOutput{
				ID:              "dc86341e-fd07-4c37-8f7e-da3dbf4319d4",
				StartDate:       time.Date(2001, 4, 7, 0, 0, 0, 0, time.UTC),
				EndDate:         time.Date(2001, 4, 7, 12, 30, 0, 0, time.UTC),
				Duration:        time.Date(2001, 4, 7, 12, 30, 0, 0, time.UTC).Sub(time.Date(2001, 4, 7, 0, 0, 0, 0, time.UTC)),
				HasLoad:         false,
				OriginLat:       38.1234567,
				OriginLong:      -25.1234567,
				DestinationLat:  -31,
				DestinationLong: -49,
				Vehicle:         "STUMP_TRUCK",
			},
		},
		{ // all
			tripID: "794b9937-7afa-449e-9662-92271d44cb81",
			input: usecase.UpdateTripInput{
				StartDate:       getDatePointer(time.Date(2005, 1, 1, 0, 0, 0, 0, time.UTC)),
				EndDate:         getDatePointer(time.Date(2005, 1, 1, 1, 0, 0, 0, time.UTC)),
				HasLoad:         getBoolPointer(false),
				OriginLat:       getFloatPointer(-11),
				OriginLong:      getFloatPointer(-57.65697731), // should not change
				DestinationLat:  getFloatPointer(58.61),
				DestinationLong: getFloatPointer(-137.95830051), // should not change
			},
			want: usecase.UpdateTripOutput{
				ID:              "794b9937-7afa-449e-9662-92271d44cb81",
				StartDate:       time.Date(2005, 1, 1, 0, 0, 0, 0, time.UTC),
				EndDate:         time.Date(2005, 1, 1, 1, 0, 0, 0, time.UTC),
				Duration:        time.Date(2005, 1, 1, 1, 0, 0, 0, time.UTC).Sub(time.Date(2005, 1, 1, 0, 0, 0, 0, time.UTC)),
				HasLoad:         false,
				OriginLat:       -11,
				OriginLong:      -57.6569773,
				DestinationLat:  58.61,
				DestinationLong: -137.9583005,
				Vehicle:         "TRUCK",
			},
		},
	}

	for i, test := range tests {
		got, gotErr := uc.Execute(context.Background(), test.tripID, test.input)

		if gotErr != nil {
			t.Errorf("%d: unexpected error -> [gotErr: %v]", i, gotErr)
			continue
		}

		if *got != test.want {
			t.Errorf("%d: [want: %+v] != [got: %+v]", i, test.want, *got)
		}
	}
}
