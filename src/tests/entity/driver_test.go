package entity_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/rafaft/truck-driver-trip-system/entity"
)

func TestNewTruckDriver(t *testing.T) {
	now := time.Now()

	type Input struct {
		cpf        string
		name       string
		gender     string
		cnh        string
		birthDate  time.Time
		hasVehicle bool
	}
	tests := []struct {
		input Input
		err   error
	}{
		{
			Input{
				cpf:        "369063555110", // invalid CPF
				name:       "Noah Oliver Cauã da Rocha",
				gender:     "M",
				cnh:        "A",
				birthDate:  time.Date(1950, time.Month(8), 25, 0, 0, 0, 0, time.UTC),
				hasVehicle: false,
			},
			entity.ErrInvalidCPF{},
		},
		{
			Input{
				cpf:        "33617661688",
				name:       "", // invalid Name
				gender:     "M",
				cnh:        "B",
				birthDate:  time.Date(1994, time.Month(4), 1, 0, 0, 0, 0, time.UTC),
				hasVehicle: true,
			},
			entity.ErrInvalidName{},
		},
		{
			Input{
				cpf:        "07405451756",
				name:       "Larissa Juliana Moura",
				gender:     "H", // invalid Gender
				cnh:        "C",
				birthDate:  time.Date(1973, time.Month(8), 20, 0, 0, 0, 0, time.UTC),
				hasVehicle: false,
			},
			entity.ErrInvalidGender{},
		},
		{
			Input{
				cpf:        "48858994000",
				name:       "Isabelly Luiza das Neves",
				gender:     "F",
				cnh:        "G", // invalid CNHType
				birthDate:  time.Date(1953, time.Month(6), 20, 0, 0, 0, 0, time.UTC),
				hasVehicle: true,
			},
			entity.ErrInvalidCNH{},
		},
		{
			Input{
				cpf:        "44854153253",
				name:       "Lavínia Milena Valentina de Paula",
				gender:     "O",
				cnh:        "D",
				birthDate:  time.Now().AddDate(-17, 0, 0), // invalid birthDate
				hasVehicle: false,
			},
			entity.ErrInvalidAge{},
		},
		{
			Input{
				cpf:        "22349860442",
				name:       "Alexandre Thiago Caleb Ferreira",
				gender:     "m",
				cnh:        "a",
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
				cnh:        "d",
				birthDate:  now.AddDate(-18, 0, 0),
				hasVehicle: false,
			},
			nil,
		},
	}

	for i, test := range tests {
		got, gotErr := entity.NewDriver(
			test.input.cpf,
			test.input.name,
			test.input.gender,
			test.input.cnh,
			test.input.birthDate,
			test.input.hasVehicle,
		)

		if reflect.TypeOf(test.err) != reflect.TypeOf(gotErr) {
			t.Errorf("%d: [wantErr: %v] [gotError: %v]", i, test.err, gotErr)
			continue
		}

		if test.err == nil {
			cpf, _ := entity.NewCPF(test.input.cpf)
			name, _ := entity.NewName(test.input.name)
			gender, _ := entity.NewGender(test.input.gender)
			cnh, _ := entity.NewCNH(test.input.cnh)
			birthDate, _ := entity.NewBirthDate(test.input.birthDate)

			if got.CPF() != cpf ||
				got.Name() != name ||
				got.Gender() != gender ||
				got.CNHType() != cnh ||
				!got.BirthDate().Equal(birthDate.Time) ||
				got.HasVehicle() != test.input.hasVehicle {
				t.Errorf("%d: [input: %v] [got: %v]", i, test.input, got)
			}
		}
	}
}
