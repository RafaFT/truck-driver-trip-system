package usecase_test

import (
	"context"
	"reflect"
	"testing"

	repository "github.com/rafaft/truck-driver-trip-system/adapter/gateway/database"
	"github.com/rafaft/truck-driver-trip-system/entity"
	"github.com/rafaft/truck-driver-trip-system/infrastructure/log"
	"github.com/rafaft/truck-driver-trip-system/tests/samples"
	"github.com/rafaft/truck-driver-trip-system/usecase"
)

func TestDeleteTripErrors(t *testing.T) {
	logger := log.NewFakeLogger()
	tripRepo := repository.NewTripInMemory(nil)
	uc := usecase.NewDeleteTrip(logger, tripRepo)

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
		gotErr := uc.Execute(context.Background(), test.input)

		if reflect.TypeOf(test.err) != reflect.TypeOf(gotErr) {
			t.Errorf("%d: [err: %v] [gotErr: %v]", i, test.err, gotErr)
		}
	}
}

func TestDeleteTrip(t *testing.T) {
	logger := log.NewFakeLogger()
	tripRepo := repository.NewTripInMemory(samples.GetTrips(t))
	uc := usecase.NewDeleteTrip(logger, tripRepo)

	tests := []struct {
		input string
		err   error
	}{
		{
			input: "47bc0c57-adb7-47da-8886-6ff92d484d06",
			err:   nil,
		},
		{
			input: "794b9937-7afa-449e-9662-92271d44cb81",
			err:   nil,
		},
	}

	for i, test := range tests {
		gotErr := uc.Execute(context.Background(), test.input)

		if gotErr != nil {
			t.Errorf("%d: unexpected error -> [gotErr: %v]", i, gotErr)
		}
	}
}
