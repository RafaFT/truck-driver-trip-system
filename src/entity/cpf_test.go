package entity

import (
	"errors"
	"testing"
)

func TestNewCPF(t *testing.T) {
	tests := []struct {
		input   string
		want    CPF
		wantErr error
	}{
		// invalid input
		{"", "", newErrInvalidCPF("")},
		{"not even trying", "", newErrInvalidCPF("not even trying")},
		{"12345678901", "", newErrInvalidCPF("12345678901")},
		{"00000000000", "", newErrInvalidCPF("00000000000")},
		{"10804773069", "", newErrInvalidCPF("10804773069")},
		{"643.512.830-84", "", newErrInvalidCPF("643.512.830-84")},
		{"64351283084a", "", newErrInvalidCPF("64351283084a")},
		{"64351f283084", "", newErrInvalidCPF("64351f283084")},
		// valid input
		{"64351283084", CPF("64351283084"), nil},
		{"10804773068", CPF("10804773068"), nil},
		{"14316382004", CPF("14316382004"), nil},
		{"54692539020", CPF("54692539020"), nil},
	}

	for i, test := range tests {
		got, gotErr := NewCPF(test.input)

		if !errors.Is(test.wantErr, gotErr) {
			t.Errorf("%d: [input: %v] [wantErr: %v] [gotErr: %v]", i, test.input, test.wantErr, gotErr)
			continue
		}

		if test.want != got {
			t.Errorf("%d: [input: %v] [want: %v] [got: %v]", i, test.input, test.want, got)
		}
	}
}
