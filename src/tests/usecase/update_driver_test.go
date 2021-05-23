package usecase_test

import (
	"context"
	"reflect"
	"testing"
	"time"

	repository "github.com/rafaft/truck-driver-trip-system/adapter/gateway/database"
	"github.com/rafaft/truck-driver-trip-system/entity"
	"github.com/rafaft/truck-driver-trip-system/infrastructure/log"
	"github.com/rafaft/truck-driver-trip-system/usecase"
)

func TestUpdateDriver(t *testing.T) {
	now := time.Now()

	l := log.NewFakeLogger()
	r := repository.NewDriverInMemory(getDriversSample(t))
	uc := usecase.NewUpdateDriver(l, r)

	tests := []struct {
		cpf   string
		input usecase.UpdateDriverInput
		want  *usecase.UpdateDriverOutput
		err   error
	}{
		// invalid
		{
			cpf:   "52742089401", // invalid CPF
			input: usecase.UpdateDriverInput{},
			want:  nil,
			err:   entity.ErrInvalidCPF{},
		},
		{
			cpf: "52742089403",
			input: usecase.UpdateDriverInput{
				CNH: getStrPointer("8"), // invalid CNH
			},
			want: nil,
			err:  entity.ErrInvalidCNH{},
		},
		{
			cpf: "94982599769",
			input: usecase.UpdateDriverInput{
				Gender: getStrPointer("l"), // invalid Gender
			},
			want: nil,
			err:  entity.ErrInvalidGender{},
		},
		{
			cpf: "08931283849",
			input: usecase.UpdateDriverInput{
				Name: getStrPointer("invalid  name because  of  extra  space"), // invalid name
			},
			want: nil,
			err:  entity.ErrInvalidName{},
		},
		{
			cpf:   "08298213092", // CPF not registered
			input: usecase.UpdateDriverInput{},
			want:  nil,
			err:   entity.ErrDriverNotFound{},
		},
		// valid
		{
			cpf: "31803413603",
			input: usecase.UpdateDriverInput{
				CNH: getStrPointer("b"),
			},
			want: &usecase.UpdateDriverOutput{
				Age:        46,
				BirthDate:  now.AddDate(-46, 0, 0),
				CNH:        "B",
				CPF:        "31803413603",
				Gender:     "O",
				HasVehicle: false,
				Name:       "rayssa emanuelly andrea viana",
			},
			err: nil,
		},
		{
			cpf: "72595478710",
			input: usecase.UpdateDriverInput{
				Gender: getStrPointer("f"),
			},
			want: &usecase.UpdateDriverOutput{
				Age:        41,
				BirthDate:  now.AddDate(-41, 0, 0),
				CNH:        "B",
				CPF:        "72595478710",
				Gender:     "F",
				HasVehicle: true,
				Name:       "raimundo erick nicolas souza",
			},
			err: nil,
		},
		{
			cpf: "27188079463",
			input: usecase.UpdateDriverInput{
				HasVehicle: getBoolPointer(false),
			},
			want: &usecase.UpdateDriverOutput{
				Age:        33,
				BirthDate:  now.AddDate(-33, 0, 0),
				CNH:        "D",
				CPF:        "27188079463",
				Gender:     "O",
				HasVehicle: false,
				Name:       "thales marcos fogaça",
			},
			err: nil,
		},
		{
			cpf: "56820381930",
			input: usecase.UpdateDriverInput{
				Name: getStrPointer("Jessica Torres"),
			},
			want: &usecase.UpdateDriverOutput{
				Age:        28,
				BirthDate:  now.AddDate(-28, 0, 0),
				CNH:        "A",
				CPF:        "56820381930",
				Gender:     "F",
				HasVehicle: true,
				Name:       "jessica torres",
			},
			err: nil,
		},
		{
			cpf: "56820381930",
			input: usecase.UpdateDriverInput{
				Name: getStrPointer("Jessica Torres"),
			},
			want: &usecase.UpdateDriverOutput{
				Age:        28,
				BirthDate:  now.AddDate(-28, 0, 0),
				CNH:        "A",
				CPF:        "56820381930",
				Gender:     "F",
				HasVehicle: true,
				Name:       "jessica torres",
			},
			err: nil,
		},
		{
			cpf: "17844926805",
			input: usecase.UpdateDriverInput{
				CNH:        getStrPointer("D"),
				Gender:     getStrPointer("O"),
				HasVehicle: getBoolPointer(true),
				Name:       getStrPointer("allana louise bianca"),
			},
			want: &usecase.UpdateDriverOutput{
				Age:        20,
				BirthDate:  now.AddDate(-20, 0, 0),
				CNH:        "D",
				CPF:        "17844926805",
				Gender:     "O",
				HasVehicle: true,
				Name:       "allana louise bianca",
			},
			err: nil,
		},
	}

	for i, test := range tests {
		got, gotErr := uc.Execute(context.Background(), test.cpf, test.input)

		if reflect.TypeOf(test.err) != reflect.TypeOf(gotErr) {
			t.Errorf("%d: [wantErr: %v] [gotErr: %v]", i, test.err, gotErr)
			continue
		}

		if test.want != nil {
			if test.want.Age != got.Age ||
				test.want.BirthDate.Year() != got.BirthDate.Year() ||
				test.want.BirthDate.Month() != got.BirthDate.Month() ||
				test.want.BirthDate.Day() != got.BirthDate.Day() ||
				test.want.CNH != got.CNH ||
				test.want.CPF != got.CPF ||
				test.want.Gender != got.Gender ||
				test.want.HasVehicle != got.HasVehicle ||
				test.want.Name != got.Name {
				t.Errorf("%d: [input: %v] [want: %v] [got: %v]", i, test.input, test.want, got)
			}
		}
	}
}
