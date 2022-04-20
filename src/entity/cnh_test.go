package entity

import (
	"errors"
	"testing"
)

func TestCNHType(t *testing.T) {
	tests := []struct {
		input   string
		want    CNH
		wantErr error
	}{
		// invalid input
		{"", "", NewErrInvalidCNH("")},
		{"not even trying", "", NewErrInvalidCNH("not even trying")},
		{"f", "", NewErrInvalidCNH("f")},
		{"0", "", NewErrInvalidCNH("0")},
		{"é", "", NewErrInvalidCNH("é")},
		{"ẽ", "", NewErrInvalidCNH("ẽ")},
		{"ç", "", NewErrInvalidCNH("ç")},
		// valid input
		{"A", CNH("A"), nil},
		{"B", CNH("B"), nil},
		{"C", CNH("C"), nil},
		{"d", CNH("D"), nil},
		{"e", CNH("E"), nil},
	}

	for i, test := range tests {
		got, gotErr := NewCNH(test.input)

		if !errors.Is(test.wantErr, gotErr) {
			t.Errorf("%d: [input: %v] [wantErr: %v] [gotErr: %v]", i, test.input, test.wantErr, gotErr)
			continue
		}

		if test.want != got {
			t.Errorf("%d: [input: %v] [want: %v] [got: %v]", i, test.input, test.want, got)
		}
	}
}

func BenchmarkCNH(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewCNH("c")
	}
}
