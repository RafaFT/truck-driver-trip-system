package entity

import (
	"errors"
	"testing"
)

func TestName(t *testing.T) {
	tests := []struct {
		input   string
		want    Name
		wantErr error
	}{
		// invalid input
		{"", "", newErrInvalidName("")},
		{"12345", "", newErrInvalidName("12345")},
		{"Húgo Diego Barros ", "", newErrInvalidName("Húgo Diego Barros ")},   // trailing space
		{" Húgo Diego Barros", "", newErrInvalidName(" Húgo Diego Barros")},   // leading space
		{"Húgo  Diego  Barros", "", newErrInvalidName("Húgo  Diego  Barros")}, // double spacing
		{"Húgo Diego Barros 1", "", newErrInvalidName("Húgo Diego Barros 1")}, // digit
		{"Húgo Diego Barros!", "", newErrInvalidName("Húgo Diego Barros!")},   // use of symbol
		{"Húgo\tDiego\vBarros", "", newErrInvalidName("Húgo\tDiego\vBarros")}, // use of tab/v space
		{"张伟 ", "", newErrInvalidName("张伟 ")},                                 // chinese name with trailing space
		// valid input
		{"Húgo Diego Barros", Name("húgo diego barros"), nil},
		{"张伟", Name("张伟"), nil},
	}

	for i, test := range tests {
		got, gotErr := NewName(test.input)

		if !errors.Is(test.wantErr, gotErr) {
			t.Errorf("%d: [input: %v] [wantErr: %v] [gotErr: %v]", i, test.input, test.wantErr, gotErr)
			continue
		}

		if test.want != got {
			t.Errorf("%d: [input: %v] [want: %v] [got: %v]", i, test.input, test.want, got)
		}
	}
}
