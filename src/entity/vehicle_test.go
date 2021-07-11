package entity

import (
	"reflect"
	"testing"
)

func TestVehicle(t *testing.T) {
	tests := []struct {
		input int
		want  Vehicle
		err   error
	}{
		// invalid input
		{-1, "", ErrInvalidVehicleCode{}},
		{300, "", ErrInvalidVehicleCode{}},
		// valid input
		{0, truck, nil},
		{1, _34Truck, nil},
		{2, stumpTruck, nil},
	}

	for _, test := range tests {
		got, gotError := NewVehicle(test.input)

		if test.want != got || reflect.TypeOf(test.err) != reflect.TypeOf(gotError) {
			t.Errorf("[input: %v] [want: %v] [error: %v] [got: %v] [gotError: %v]",
				test.input, test.want, test.err, got, gotError,
			)
		}
	}
}
