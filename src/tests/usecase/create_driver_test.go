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

func TestCreateDriver(t *testing.T) {
	now := time.Now()
	l := log.NewFakeLogger()
	r := repository.NewDriverInMemory(nil)
	uc := usecase.NewCreateDriver(l, r)

	tests := []struct {
		input usecase.CreateDriverInput
		want  *usecase.CreateDriverOutput
		err   error
	}{
		// invalid input
		{
			input: usecase.CreateDriverInput{
				BirthDate:  now, // invalid BirthDate
				CNH:        "a",
				CPF:        "15279541028",
				Gender:     "f",
				HasVehicle: true,
				Name:       "Jennifer Marlene Freitas",
			},
			want: nil,
			err:  entity.ErrInvalidAge{},
		},
		{
			input: usecase.CreateDriverInput{
				BirthDate:  now.AddDate(-20, 0, 0),
				CNH:        "7", // invalid CNH
				CPF:        "75703960223",
				Gender:     "m",
				HasVehicle: false,
				Name:       "Benjamin Thomas Monteiro",
			},
			want: nil,
			err:  entity.ErrInvalidCNH{},
		},
		{
			input: usecase.CreateDriverInput{
				BirthDate:  now.AddDate(-47, 0, 0),
				CNH:        "c",
				CPF:        "75703960223",
				Gender:     "", // invalid gender
				HasVehicle: true,
				Name:       "Márcio Iago Igor Nascimento",
			},
			want: nil,
			err:  entity.ErrInvalidGender{},
		},
		{
			input: usecase.CreateDriverInput{
				BirthDate:  now.AddDate(0, -300, 0),
				CNH:        "d",
				CPF:        "55800346100",
				Gender:     "o",
				HasVehicle: false,
				Name:       " Vera Laís Ferreira", // invalid name
			},
			want: nil,
			err:  entity.ErrInvalidName{},
		},
		// valid input
		{
			input: usecase.CreateDriverInput{
				BirthDate:  now.AddDate(-18, 0, 0),
				CNH:        "e",
				CPF:        "73552587020",
				Gender:     "m",
				HasVehicle: true,
				Name:       "Henrique Bernardo Rafael Silveira",
			},
			want: &usecase.CreateDriverOutput{
				Age:        18,
				BirthDate:  time.Date(now.AddDate(-18, 0, 0).Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC),
				CNH:        "E",
				CPF:        "73552587020",
				Gender:     "M",
				HasVehicle: true,
				Name:       "henrique bernardo rafael silveira",
			},
			err: nil,
		},
		{
			input: usecase.CreateDriverInput{
				BirthDate:  now.AddDate(-30, 0, 0),
				CNH:        "a",
				CPF:        "45298982530",
				Gender:     "f",
				HasVehicle: false,
				Name:       "Eduarda Sônia Rebeca Oliveira",
			},
			want: &usecase.CreateDriverOutput{
				Age:        30,
				BirthDate:  time.Date(now.AddDate(-30, 0, 0).Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC),
				CNH:        "A",
				CPF:        "45298982530",
				Gender:     "F",
				HasVehicle: false,
				Name:       "eduarda sônia rebeca oliveira",
			},
			err: nil,
		},
		{
			input: usecase.CreateDriverInput{
				BirthDate:  now.AddDate(-41, 0, 0),
				CNH:        "b",
				CPF:        "94587146218",
				Gender:     "O",
				HasVehicle: true,
				Name:       "Bárbara Andreia Vieira",
			},
			want: &usecase.CreateDriverOutput{
				Age:        41,
				BirthDate:  time.Date(now.AddDate(-41, 0, 0).Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC),
				CNH:        "B",
				CPF:        "94587146218",
				Gender:     "O",
				HasVehicle: true,
				Name:       "bárbara andreia vieira",
			},
			err: nil,
		},
		{
			input: usecase.CreateDriverInput{
				BirthDate:  now.AddDate(-65, 0, 0),
				CNH:        "C",
				CPF:        "26226702575",
				Gender:     "m",
				HasVehicle: false,
				Name:       "Cláudio Lorenzo Luan Santos",
			},
			want: &usecase.CreateDriverOutput{
				Age:        65,
				BirthDate:  time.Date(now.AddDate(-65, 0, 0).Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC),
				CNH:        "C",
				CPF:        "26226702575",
				Gender:     "M",
				HasVehicle: false,
				Name:       "cláudio lorenzo luan santos",
			},
			err: nil,
		},
		{
			input: usecase.CreateDriverInput{
				BirthDate:  now.AddDate(-71, 0, 0),
				CNH:        "d",
				CPF:        "70874163404",
				Gender:     "f",
				HasVehicle: true,
				Name:       "Márcia Analu Antônia Assis",
			},
			want: &usecase.CreateDriverOutput{
				Age:        71,
				BirthDate:  time.Date(now.AddDate(-71, 0, 0).Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC),
				CNH:        "D",
				CPF:        "70874163404",
				Gender:     "F",
				HasVehicle: true,
				Name:       "márcia analu antônia assis",
			},
			err: nil,
		},
	}

	for i, test := range tests {
		got, gotErr := uc.Execute(context.Background(), test.input)

		if reflect.TypeOf(test.err) != reflect.TypeOf(gotErr) {
			t.Errorf("%d: [err: %v] [gotErr: %v]", i, test.err, gotErr)
			continue
		}

		if test.want != nil {
			if test.want.Age != got.Age ||
				test.want.BirthDate.Year() != got.BirthDate.Year() ||
				test.want.BirthDate.Month() != got.BirthDate.Month() ||
				test.want.BirthDate.Day() != got.BirthDate.Day() ||
				test.want.CNH != got.CNH ||
				test.want.CPF != got.CPF ||
				!got.CreatedAt.After(now) ||
				test.want.Gender != got.Gender ||
				test.want.HasVehicle != got.HasVehicle ||
				test.want.Name != got.Name {
				t.Errorf("%d: [input: %v] [want: %v] [got: %v]", i, test.input, test.want, got)
			}
		}
	}
}
