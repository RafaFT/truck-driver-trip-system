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

func TestGetTripsErrors(t *testing.T) {
	logger := log.NewFakeLogger()
	tripRepo := repository.NewTripInMemory(nil)
	uc := usecase.NewGetTrips(logger, tripRepo)

	tests := []struct {
		input usecase.GetTripsInput
		err   error
	}{
		{
			input: usecase.GetTripsInput{
				CPF:         getStrPointer("123"), // invalid CPF
				HasLoad:     getBoolPointer(true),
				Limit:       getUintPointer(5),
				VehicleCode: getIntPointer(0),
			},
			err: entity.ErrInvalidCPF{},
		},
		{
			input: usecase.GetTripsInput{
				CPF:         getStrPointer("48027285062"),
				HasLoad:     getBoolPointer(true),
				Limit:       getUintPointer(5),
				VehicleCode: getIntPointer(-1), // invalid code
			},
			err: entity.ErrInvalidVehicleCode{},
		},
	}

	for i, test := range tests {
		_, gotErr := uc.Execute(context.Background(), test.input)

		if reflect.TypeOf(test.err) != reflect.TypeOf(gotErr) {
			t.Errorf("%d: [wantErr: %v] [gotErr: %v]", i, test.err, gotErr)
		}
	}
}

func TestGetTrips(t *testing.T) {
	logger := log.NewFakeLogger()
	tripRepo := repository.NewTripInMemory(samples.GetTrips(t))
	uc := usecase.NewGetTrips(logger, tripRepo)

	tests := []struct {
		input   usecase.GetTripsInput
		wantLen int
		err     error
	}{
		{
			input: usecase.GetTripsInput{
				CPF: getStrPointer("51598390031"),
			},
			wantLen: 3,
			err:     nil,
		},
		{
			input: usecase.GetTripsInput{
				HasLoad: getBoolPointer(false),
			},
			wantLen: 1,
			err:     nil,
		},
		{
			input: usecase.GetTripsInput{
				CPF:     getStrPointer("51598390031"),
				HasLoad: getBoolPointer(false),
			},
			wantLen: 0,
			err:     nil,
		},
		{
			input: usecase.GetTripsInput{
				CPF:         getStrPointer("51598390031"),
				VehicleCode: getIntPointer(0),
			},
			wantLen: 1,
			err:     nil,
		},
		{
			input: usecase.GetTripsInput{
				Limit: getUintPointer(2),
			},
			wantLen: 2,
			err:     nil,
		},
	}

	for i, test := range tests {
		got, gotErr := uc.Execute(context.Background(), test.input)

		if gotErr != nil {
			t.Errorf("%d: unexpected error -> [gotErr: %v]", i, gotErr)
			continue
		}

		if len(got) != test.wantLen {
			t.Errorf("%d: [wantLen: %v] != [gotLen: %v]", i, test.wantLen, len(got))
		}
	}
}
