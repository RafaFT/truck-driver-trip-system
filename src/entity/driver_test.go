package entity

import (
	"testing"
	"time"
)

func TestNewCPF(t *testing.T) {
	tests := []struct {
		input string
		want  CPF
		err   error
	}{
		// invalid CPF values
		{"", "", newErrInvalidCPF("")},
		{"not even trying", "", newErrInvalidCPF("not even trying")},
		{"12345678901", "", newErrInvalidCPF("12345678901")},
		{"00000000000", "", newErrInvalidCPF("00000000000")},
		{"10804773069", "", newErrInvalidCPF("10804773069")},
		{"643.512.830-84", "", newErrInvalidCPF("643.512.830-84")},
		{"64351283084a", "", newErrInvalidCPF("64351283084a")},
		{"64351f283084", "", newErrInvalidCPF("64351f283084")},
		// valid cases
		{"64351283084", CPF("64351283084"), nil},
		{"10804773068", CPF("10804773068"), nil},
	}

	for _, test := range tests {
		got, gotErr := NewCPF(test.input)

		if test.want != got || gotErr != test.err {
			t.Errorf("[input: %v] [want: %v] [err: %v] [got: %v] [gotErr: %v]",
				test.input, test.want, test.err, got, gotErr,
			)
		}
	}
}

func TestCalculateAge(t *testing.T) {
	baseDate := time.Date(2021, time.Month(2), 21, 20, 12, 0, 0, time.UTC)

	tests := []struct {
		baseDate  time.Time
		birthDate time.Time
		want      int
	}{
		{baseDate, time.Time{}, 2020},
		{baseDate, baseDate, 0},
		{baseDate, time.Date(2021, time.Month(1), 1, 9, 59, 59, 59, time.UTC), 0},
		{baseDate, time.Date(2020, time.Month(2), 22, 10, 0, 0, 1, time.UTC), 0},
		{baseDate, time.Date(2020, time.Month(2), 21, 10, 0, 0, 1, time.UTC), 1},
		{baseDate, time.Date(2020, time.Month(2), 20, 10, 0, 0, 1, time.UTC), 1},
		{baseDate, time.Date(1989, time.Month(12), 12, 0, 0, 0, 0, time.UTC), 31},
		{baseDate, time.Date(2000, time.Month(6), 25, 0, 0, 0, 0, time.UTC), 20},
	}

	for _, test := range tests {
		got := calculateAge(test.baseDate, test.birthDate)

		if got != test.want {
			t.Errorf("[baseDate: %v] [birthDate: %v] [want: %v] [got: %v]",
				test.baseDate, test.birthDate, test.want, got,
			)
		}
	}
}

func TestGender(t *testing.T) {
	tests := []struct {
		input string
		want  string
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
		{"M", "M", nil},
		{"F", "F", nil},
		{"O", "O", nil},
		{"m", "M", nil},
		{"f", "F", nil},
		{"o", "O", nil},
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
		want  string
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
		{"A", "A", nil},
		{"B", "B", nil},
		{"C", "C", nil},
		{"d", "D", nil},
		{"e", "E", nil},
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
		want  string
		error error
	}{
		// invalid input
		{"", "", ErrInvalidName},
		// valid input
		{"Rafael Trad", "rafael trad", nil},
		{"rafael trad", "rafael trad", nil},
		{"12345", "12345", nil},
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

func TestNewTruckDriver(t *testing.T) {
	now := time.Now()

	type Input struct {
		cpf        string
		name       string
		gender     string
		cnhType    string
		birthDate  time.Time
		hasVehicle bool
	}
	tests := []struct {
		input Input
		want  *TruckDriver
		error error
	}{
		{
			Input{
				cpf:        "369063555110", // invalid CPF
				name:       "Noah Oliver Cauã da Rocha",
				gender:     "M",
				cnhType:    "A",
				birthDate:  time.Date(1950, time.Month(8), 25, 0, 0, 0, 0, time.UTC),
				hasVehicle: false,
			},
			nil,
			newErrInvalidCPF("369063555110"),
		},
		{
			Input{
				cpf:        "33617661688",
				name:       "", // invalid Name
				gender:     "M",
				cnhType:    "B",
				birthDate:  time.Date(1994, time.Month(4), 1, 0, 0, 0, 0, time.UTC),
				hasVehicle: true,
			},
			nil,
			ErrInvalidName,
		},
		{
			Input{
				cpf:        "07405451756",
				name:       "Larissa Juliana Moura",
				gender:     "H", // invalid Gender
				cnhType:    "C",
				birthDate:  time.Date(1973, time.Month(8), 20, 0, 0, 0, 0, time.UTC),
				hasVehicle: false,
			},
			nil,
			ErrInvalidGender,
		},
		{
			Input{
				cpf:        "48858994000",
				name:       "Isabelly Luiza das Neves",
				gender:     "F",
				cnhType:    "G", // invalid CNHType
				birthDate:  time.Date(1953, time.Month(6), 20, 0, 0, 0, 0, time.UTC),
				hasVehicle: true,
			},
			nil,
			ErrInvalidCNHType,
		},
		{
			Input{
				cpf:        "44854153253",
				name:       "Lavínia Milena Valentina de Paula",
				gender:     "O",
				cnhType:    "D",
				birthDate:  time.Now().AddDate(-17, 0, 0), // invalid birthDate
				hasVehicle: false,
			},
			nil,
			ErrInvalidBirthDate,
		},
		{
			Input{
				cpf:        "22349860442",
				name:       "Alexandre Thiago Caleb Ferreira",
				gender:     "m",
				cnhType:    "a",
				birthDate:  time.Date(1979, time.Month(5), 6, 0, 0, 0, 0, time.UTC),
				hasVehicle: true,
			},
			&TruckDriver{
				cpf:        "22349860442",
				name:       "alexandre thiago caleb ferreira",
				gender:     "M",
				cnhType:    "A",
				birthDate:  time.Date(1979, time.Month(5), 6, 0, 0, 0, 0, time.UTC),
				hasVehicle: true,
			},
			nil,
		},
		{
			Input{
				cpf:        "59706144757",
				name:       "Ricardo Igor Luiz Barbosa",
				gender:     "o",
				cnhType:    "d",
				birthDate:  now.AddDate(-18, 0, 0),
				hasVehicle: false,
			},
			&TruckDriver{
				cpf:        "59706144757",
				name:       "ricardo igor luiz barbosa",
				gender:     "O",
				cnhType:    "D",
				birthDate:  now.AddDate(-18, 0, 0),
				hasVehicle: false,
			},
			nil,
		},
	}

	for i, test := range tests {
		got, err := NewTruckDriver(
			test.input.cpf,
			test.input.name,
			test.input.gender,
			test.input.cnhType,
			test.input.birthDate,
			test.input.hasVehicle,
		)

		if err != test.error {
			t.Errorf("%d: [wantErr: %v] [gotError: %v]", i, test.error, err)
		} else if test.want != nil {
			if got.CPF() != test.want.CPF() ||
				got.Name() != test.want.Name() ||
				got.Gender() != test.want.Gender() ||
				got.CNHType() != test.want.CNHType() ||
				!(time.Time(got.BirthDate()).Equal(test.want.BirthDate())) ||
				got.HasVehicle() != test.want.HasVehicle() {
				t.Errorf("%d: [want: %v] [got: %v]", i, test.want, got)
			}
		}
	}
}
