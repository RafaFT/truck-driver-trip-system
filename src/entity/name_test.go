package entity_test

import (
	"reflect"
	"testing"

	"github.com/rafaft/truck-driver-trip-system/entity"
)

func TestName(t *testing.T) {
	tests := []struct {
		input string
		want  entity.Name
		err   error
	}{
		// invalid input
		{"", "", entity.ErrInvalidName{}},
		{"12345", "", entity.ErrInvalidName{}},
		{"Húgo Diego Barros ", "", entity.ErrInvalidName{}},  // trailing space
		{" Húgo Diego Barros", "", entity.ErrInvalidName{}},  // leading space
		{"Húgo  Diego  Barros", "", entity.ErrInvalidName{}}, // double spacing
		{"Húgo Diego Barros 1", "", entity.ErrInvalidName{}}, // digit
		{"Húgo\tDiego\vBarros", "", entity.ErrInvalidName{}}, // use of tab/v space
		{"张伟 ", "", entity.ErrInvalidName{}},                 // chinese name with trailing space
		// valid input
		{"Húgo Diego Barros", "húgo diego barros", nil},
		{"张伟", "张伟", nil},
	}

	for _, test := range tests {
		got, gotError := entity.NewName(test.input)

		if test.want != got || reflect.TypeOf(test.err) != reflect.TypeOf(gotError) {
			t.Errorf("[input: %v] [want: %v] [error: %v] [got: %v] [gotError: %v]",
				test.input, test.want, test.err, got, gotError,
			)
		}
	}
}
