package entity_test

import (
	"reflect"
	"testing"

	"github.com/rafaft/truck-driver-trip-system/entity"
)

func TestVehicle(t *testing.T) {
	tests := []struct {
		input int
		want  entity.Vehicle
		err   error
	}{
		// invalid input
		{-1, "", entity.ErrInvalidVehicleCode{}},
		{300, "", entity.ErrInvalidVehicleCode{}},
		// valid input
		{0, entity.Truck, nil},
		{1, entity.Truck_34, nil},
		{2, entity.StumpTruck, nil},
	}

	for _, test := range tests {
		got, gotError := entity.NewVehicle(test.input)

		if test.want != got || reflect.TypeOf(test.err) != reflect.TypeOf(gotError) {
			t.Errorf("[input: %v] [want: %v] [error: %v] [got: %v] [gotError: %v]",
				test.input, test.want, test.err, got, gotError,
			)
		}
	}
}
