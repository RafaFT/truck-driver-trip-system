package entity

import (
	"errors"
	"testing"
)

func TestVehicle(t *testing.T) {
	tests := []struct {
		input   int
		want    Vehicle
		wantErr error
	}{
		// invalid input
		{-1, "", NewErrInvalidVehicleCode(-1)},
		{300, "", NewErrInvalidVehicleCode(300)},
		// valid input
		{0, Truck, nil},
		{1, Truck_34, nil},
		{2, StumpTruck, nil},
	}

	for i, test := range tests {
		got, gotErr := NewVehicle(test.input)

		if !errors.Is(test.wantErr, gotErr) {
			t.Errorf("%d: [input: %v] [wantErr: %v] [gotErr: %v]", i, test.input, test.wantErr, gotErr)
			continue
		}

		if test.want != got {
			t.Errorf("%d: [input: %v] [want: %v] [got: %v]", i, test.input, test.want, got)
		}
	}
}
