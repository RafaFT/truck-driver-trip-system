package entity_test

import (
	"reflect"
	"testing"

	"github.com/rafaft/truck-driver-trip-system/entity"
)

func TestGender(t *testing.T) {
	tests := []struct {
		input string
		want  entity.Gender
		err   error
	}{
		// invalid input
		{"", "", entity.ErrInvalidGender{}},
		{"not even trying", "", entity.ErrInvalidGender{}},
		{"a", "", entity.ErrInvalidGender{}},
		{"5", "", entity.ErrInvalidGender{}},
		{"Ó", "", entity.ErrInvalidGender{}},
		{"ô", "", entity.ErrInvalidGender{}},
		// valid input
		{"M", entity.Gender("M"), nil},
		{"F", entity.Gender("F"), nil},
		{"O", entity.Gender("O"), nil},
		{"m", entity.Gender("M"), nil},
		{"f", entity.Gender("F"), nil},
		{"o", entity.Gender("O"), nil},
	}

	for _, test := range tests {
		got, gotError := entity.NewGender(test.input)

		if test.want != got || reflect.TypeOf(test.err) != reflect.TypeOf(gotError) {
			t.Errorf("[input: %v] [want: %v] [error: %v] [got: %v] [gotError: %v]",
				test.input, test.want, test.err, got, gotError,
			)
		}
	}
}
