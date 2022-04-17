package entity

import (
	"errors"
	"testing"
	"time"
)

func TestNewTruckDriver(t *testing.T) {
	now := time.Now()
	moscowLocation, _ := time.LoadLocation("Europe/Moscow")

	type DriverInput struct {
		cpf        string
		name       string
		gender     string
		cnh        string
		birthDate  time.Time
		hasVehicle bool
	}

	tests := []struct {
		input   DriverInput
		want    Driver
		wantErr error
	}{
		// invalid input
		{
			DriverInput{
				cpf:        "369063555110", // invalid CPF
				name:       "Noah Oliver Cauã da Rocha",
				gender:     "M",
				cnh:        "A",
				birthDate:  time.Date(1950, time.Month(8), 25, 0, 0, 0, 0, time.UTC),
				hasVehicle: false,
			},
			Driver{},
			NewErrInvalidCPF("369063555110"),
		},
		{
			DriverInput{
				cpf:        "33617661688",
				name:       "", // invalid Name
				gender:     "M",
				cnh:        "B",
				birthDate:  time.Date(1994, time.Month(4), 1, 0, 0, 0, 0, time.UTC),
				hasVehicle: true,
			},
			Driver{},
			NewErrInvalidName(""),
		},
		{
			DriverInput{
				cpf:        "07405451756",
				name:       "Larissa Juliana Moura",
				gender:     "H", // invalid Gender
				cnh:        "C",
				birthDate:  time.Date(1973, time.Month(8), 20, 0, 0, 0, 0, time.UTC),
				hasVehicle: false,
			},
			Driver{},
			NewErrInvalidGender("H"),
		},
		{
			DriverInput{
				cpf:        "48858994000",
				name:       "Isabelly Luiza das Neves",
				gender:     "F",
				cnh:        "G", // invalid CNHType
				birthDate:  time.Date(1953, time.Month(6), 20, 0, 0, 0, 0, time.UTC),
				hasVehicle: true,
			},
			Driver{},
			NewErrInvalidCNH("G"),
		},
		{
			DriverInput{
				cpf:        "44854153253",
				name:       "Lavínia Milena Valentina de Paula",
				gender:     "O",
				cnh:        "D",
				birthDate:  now.AddDate(-17, 0, 0), // invalid birthDate
				hasVehicle: false,
			},
			Driver{},
			NewErrInvalidAge(17),
		},
		// valid input
		{ // +3 UTC offset
			DriverInput{
				cpf:        "22349860442",
				name:       "Alexandre Thiago Caleb Ferreira",
				gender:     "m",
				cnh:        "a",
				birthDate:  time.Date(1979, time.Month(5), 6, 0, 0, 0, 0, moscowLocation),
				hasVehicle: true,
			},
			Driver{
				cpf:        CPF("22349860442"),
				name:       Name("alexandre thiago caleb ferreira"),
				gender:     Gender("M"),
				cnh:        CNH("A"),
				birthDate:  BirthDate{time.Date(1979, time.Month(5), 5, 21, 0, 0, 0, time.UTC)},
				hasVehicle: true,
			},
			nil,
		},
		{
			DriverInput{
				cpf:        "59706144757",
				name:       "Ricardo Igor Luiz Barbosa",
				gender:     "o",
				cnh:        "d",
				birthDate:  now.AddDate(-18, 0, 0),
				hasVehicle: false,
			},
			Driver{
				cpf:        CPF("59706144757"),
				name:       Name("ricardo igor luiz barbosa"),
				gender:     Gender("O"),
				cnh:        CNH("D"),
				birthDate:  BirthDate{now.AddDate(-18, 0, 0).UTC()},
				hasVehicle: false,
			},
			nil,
		},
	}

	for i, test := range tests {
		got, gotErr := NewDriver(
			test.input.cpf,
			test.input.name,
			test.input.gender,
			test.input.cnh,
			test.input.birthDate,
			test.input.hasVehicle,
		)

		if !errors.Is(test.wantErr, gotErr) {
			t.Errorf("%d: [input: %v] [wantErr: %v] [gotErr: %v]", i, test.input, test.wantErr, gotErr)
			continue
		}

		if got != nil && test.want != *got {
			t.Errorf("%d: [input: %v] [want: %v] [got: %v]", i, test.input, test.want, got)
		}
	}
}
