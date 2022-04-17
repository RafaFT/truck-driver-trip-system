package entity

import (
	"errors"
	"testing"
)

func TestGender(t *testing.T) {
	tests := []struct {
		input   string
		want    Gender
		wantErr error
	}{
		// invalid input
		{"", "", NewErrInvalidGender("")},
		{"not even trying", "", NewErrInvalidGender("not even trying")},
		{"a", "", NewErrInvalidGender("a")},
		{"5", "", NewErrInvalidGender("5")},
		{"Ó", "", NewErrInvalidGender("Ó")},
		{"ô", "", NewErrInvalidGender("ô")},
		// valid input
		{"M", Gender("M"), nil},
		{"F", Gender("F"), nil},
		{"O", Gender("O"), nil},
		{"m", Gender("M"), nil},
		{"f", Gender("F"), nil},
		{"o", Gender("O"), nil},
	}

	for i, test := range tests {
		got, gotErr := NewGender(test.input)

		if !errors.Is(test.wantErr, gotErr) {
			t.Errorf("%d: [input: %v] [wantErr: %v] [gotErr: %v]", i, test.input, test.wantErr, gotErr)
			continue
		}

		if test.want != got {
			t.Errorf("%d: [input: %v] [want: %v] [got: %v]", i, test.input, test.want, got)
		}
	}
}
