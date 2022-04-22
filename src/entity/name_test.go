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
		{"", "", NewErrInvalidName("")},
		{"12345", "", NewErrInvalidName("12345")},
		{"Húgo Diego Barros ", "", NewErrInvalidName("Húgo Diego Barros ")},         // trailing space
		{" Húgo Diego Barros", "", NewErrInvalidName(" Húgo Diego Barros")},         // leading space
		{"Húgo  Diego  Barros", "", NewErrInvalidName("Húgo  Diego  Barros")},       // double spacing
		{"Húgo Diego Barros 1", "", NewErrInvalidName("Húgo Diego Barros 1")},       // digit
		{"Húgo Diego Barros!", "", NewErrInvalidName("Húgo Diego Barros!")},         // use of symbol
		{"Húgo\tDiego\vBarros", "", NewErrInvalidName("Húgo\tDiego\vBarros")},       // use of tab/v space
		{"张伟 ", "", NewErrInvalidName("张伟 ")},                                       // chinese name with trailing space
		{strings.Repeat("a", 128), "", NewErrInvalidName(strings.Repeat("a", 128))}, // max length
		// valid input
		{strings.Repeat("a", 127), Name(strings.Repeat("a", 127)), nil},
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
