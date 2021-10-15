package entity_test

import (
	"reflect"
	"testing"

	"github.com/rafaft/truck-driver-trip-system/entity"
)

func TestNewCPF(t *testing.T) {
	tests := []struct {
		input string
		want  entity.CPF
		err   error
	}{
		// invalid CPF values
		{"", "", entity.ErrInvalidCPF{}},
		{"not even trying", "", entity.ErrInvalidCPF{}},
		{"12345678901", "", entity.ErrInvalidCPF{}},
		{"00000000000", "", entity.ErrInvalidCPF{}},
		{"10804773069", "", entity.ErrInvalidCPF{}},
		{"643.512.830-84", "", entity.ErrInvalidCPF{}},
		{"64351283084a", "", entity.ErrInvalidCPF{}},
		{"64351f283084", "", entity.ErrInvalidCPF{}},
		// valid cases
		{"64351283084", entity.CPF("64351283084"), nil},
		{"10804773068", entity.CPF("10804773068"), nil},
		{"14316382004", entity.CPF("14316382004"), nil},
		{"54692539020", entity.CPF("54692539020"), nil},
	}

	for _, test := range tests {
		got, gotErr := entity.NewCPF(test.input)

		if test.want != got || reflect.TypeOf(gotErr) != reflect.TypeOf(test.err) {
			t.Errorf("[input: %v] [want: %v] [err: %v] [got: %v] [gotErr: %v]",
				test.input, test.want, test.err, got, gotErr,
			)
		}
	}
}
