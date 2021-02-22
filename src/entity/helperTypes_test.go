package entity

import (
	"testing"
)

func TestCPF(t *testing.T) {
	tests := []struct {
		input string
		want  CPF
		error error
	}{
		// invalid input
		{"", "", ErrInvalidCPF},
		{"not even trying", "", ErrInvalidCPF},
		{"abcdefghijk", "", ErrInvalidCPF},
		{"a12345678901", "", ErrInvalidCPF},
		{"123.456.789-01", "", ErrInvalidCPF},
		{"1234567890a", "", ErrInvalidCPF},
		// valid input
		{"12345678901", CPF("12345678901"), nil},
		{"01234567891", CPF("01234567891"), nil},
	}

	for _, test := range tests {
		got, gotError := NewCPF(test.input)

		if test.want != got || test.error != gotError {
			t.Errorf("[input: %v] [want: %v] [error: %v] [got: %v] [gotError: %v]",
				test.input, test.want, test.error, got, gotError,
			)
		}
	}
}

func TestGender(t *testing.T) {
	tests := []struct {
		input string
		want  Gender
		error error
	}{
		// invalid input
		{"", "", ErrInvalidGender},
		{"not even trying", "", ErrInvalidGender},
		{"a", "", ErrInvalidGender},
		{"5", "", ErrInvalidGender},
		{"Ó", "", ErrInvalidGender},
		{"ô", "", ErrInvalidGender},
		// valid input
		{"M", Gender("M"), nil},
		{"F", Gender("F"), nil},
		{"O", Gender("O"), nil},
		{"m", Gender("M"), nil},
		{"f", Gender("F"), nil},
		{"o", Gender("O"), nil},
	}

	for _, test := range tests {
		got, gotError := NewGender(test.input)

		if test.want != got || test.error != gotError {
			t.Errorf("[input: %v] [want: %v] [error: %v] [got: %v] [gotError: %v]",
				test.input, test.want, test.error, got, gotError,
			)
		}
	}
}

func TestCNHType(t *testing.T) {
	tests := []struct {
		input string
		want  CNHType
		error error
	}{
		// invalid input
		{"", "", ErrInvalidCNHType},
		{"not even trying", "", ErrInvalidCNHType},
		{"f", "", ErrInvalidCNHType},
		{"0", "", ErrInvalidCNHType},
		{"é", "", ErrInvalidCNHType},
		{"ẽ", "", ErrInvalidCNHType},
		{"ç", "", ErrInvalidCNHType},
		// valid input
		{"A", CNHType("A"), nil},
		{"B", CNHType("B"), nil},
		{"C", CNHType("C"), nil},
		{"d", CNHType("D"), nil},
		{"e", CNHType("E"), nil},
	}

	for _, test := range tests {
		got, gotError := NewCNHType(test.input)

		if test.want != got || test.error != gotError {
			t.Errorf("[input: %v] [want: %v] [error: %v] [got: %v] [gotError: %v]",
				test.input, test.want, test.error, got, gotError,
			)
		}
	}
}

func TestName(t *testing.T) {
	tests := []struct {
		input string
		want  Name
		error error
	}{
		// invalid input
		{"", "", ErrInvalidName},
		// valid input
		{"Rafael Trad", Name("rafael trad"), nil},
		{"rafael trad", Name("rafael trad"), nil},
		{"12345", Name("12345"), nil},
	}

	for _, test := range tests {
		got, gotError := NewName(test.input)

		if test.want != got || test.error != gotError {
			t.Errorf("[input: %v] [want: %v] [error: %v] [got: %v] [gotError: %v]",
				test.input, test.want, test.error, got, gotError,
			)
		}
	}
}
