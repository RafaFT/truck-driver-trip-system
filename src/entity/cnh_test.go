package entity_test

import (
	"reflect"
	"testing"

	"github.com/rafaft/truck-driver-trip-system/entity"
)

func TestCNHType(t *testing.T) {
	tests := []struct {
		input string
		want  entity.CNH
		err   error
	}{
		// invalid input
		{"", "", entity.ErrInvalidCNH{}},
		{"not even trying", "", entity.ErrInvalidCNH{}},
		{"f", "", entity.ErrInvalidCNH{}},
		{"0", "", entity.ErrInvalidCNH{}},
		{"é", "", entity.ErrInvalidCNH{}},
		{"ẽ", "", entity.ErrInvalidCNH{}},
		{"ç", "", entity.ErrInvalidCNH{}},
		// valid input
		{"A", "A", nil},
		{"B", "B", nil},
		{"C", "C", nil},
		{"d", "D", nil},
		{"e", "E", nil},
	}

	for _, test := range tests {
		got, gotError := entity.NewCNH(test.input)

		if test.want != got || reflect.TypeOf(test.err) != reflect.TypeOf(gotError) {
			t.Errorf("[input: %v] [want: %v] [error: %v] [got: %v] [gotError: %v]",
				test.input, test.want, test.err, got, gotError,
			)
		}
	}
}
