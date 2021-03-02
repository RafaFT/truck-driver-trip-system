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
	now := time.Now()

	tests := []struct {
		input time.Time
		want  int
	}{
		{now.AddDate(-2020, 0, 0), 2020},
		{now, 0},
		{now.AddDate(0, -1, 0), 0},
		{now.AddDate(-1, 0, 1), 0},
		{now.AddDate(-1, 0, 0), 1},
		{now.AddDate(-1, -1, 0), 1},
		{now.AddDate(-31, 0, 0), 31},
		{now.AddDate(-20, -12, 1), 20},
	}

	for _, test := range tests {
		bd, _ := NewBirthDate(test.input)
		got := bd.CalculateAge()

		if got != test.want {
			t.Errorf("[input: %v] [want: %v] [got: %v]",
				test.input, test.want, got,
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
		{"", "", newErrInvalidGender("")},
		{"not even trying", "", newErrInvalidGender("not even trying")},
		{"a", "", newErrInvalidGender("a")},
		{"5", "", newErrInvalidGender("5")},
		{"Ó", "", newErrInvalidGender("Ó")},
		{"ô", "", newErrInvalidGender("ô")},
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
		want  CNH
		error error
	}{
		// invalid input
		{"", "", newErrInvalidCNH("")},
		{"not even trying", "", newErrInvalidCNH("not even trying")},
		{"f", "", newErrInvalidCNH("f")},
		{"0", "", newErrInvalidCNH("0")},
		{"é", "", newErrInvalidCNH("é")},
		{"ẽ", "", newErrInvalidCNH("ẽ")},
		{"ç", "", newErrInvalidCNH("ç")},
		// valid input
		{"A", "A", nil},
		{"B", "B", nil},
		{"C", "C", nil},
		{"d", "D", nil},
		{"e", "E", nil},
	}

	for _, test := range tests {
		got, gotError := NewCNH(test.input)

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
		{"", "", newErrInvalidName("")},
		{"12345", "", newErrInvalidName("12345")},
		{"Húgo Diego Barros ", "", newErrInvalidName("Húgo Diego Barros ")},   // trailing space
		{" Húgo Diego Barros", "", newErrInvalidName(" Húgo Diego Barros")},   // leading space
		{"Húgo  Diego  Barros", "", newErrInvalidName("Húgo  Diego  Barros")}, // double spacing
		{"Húgo Diego Barros 1", "", newErrInvalidName("Húgo Diego Barros 1")}, // digit
		{"Húgo\tDiego\vBarros", "", newErrInvalidName("Húgo\tDiego\vBarros")}, // use of tab/v space
		{"张伟 ", "", newErrInvalidName("张伟 ")},                                 // chine name with trailing space
		// valid input
		{"Húgo Diego Barros", "húgo diego barros", nil},
		{"张伟", "张伟", nil},
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
			newErrInvalidName(""),
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
			newErrInvalidGender("H"),
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
			newErrInvalidCNH("G"),
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
			newErrInvalidAge(17),
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
				cpf:        CPF("22349860442"),
				name:       Name("alexandre thiago caleb ferreira"),
				gender:     Gender("M"),
				cnh:        CNH("A"),
				birthDate:  BirthDate{time.Date(1979, time.Month(5), 6, 0, 0, 0, 0, time.UTC)},
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
				cpf:        CPF("59706144757"),
				name:       Name("ricardo igor luiz barbosa"),
				gender:     Gender("O"),
				cnh:        CNH("D"),
				birthDate:  BirthDate{now.AddDate(-18, 0, 0)},
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
				!got.BirthDate().Equal(test.want.BirthDate().Time) ||
				got.HasVehicle() != test.want.HasVehicle() {
				t.Errorf("%d: [want: %v] [got: %v]", i, test.want, got)
			}
		}
	}
}
