package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/rafaft/truck-driver-trip-system/entity"
)

type mockCreateDriverRepo struct {
	err error
}

func (m mockCreateDriverRepo) Save(ctx context.Context, driver *entity.Driver) error {
	return m.err
}

func TestCreateDriver(t *testing.T) {
	now := time.Now().UTC()

	tests := []struct {
		input CreateDriverInput
		repo  mockCreateDriverRepo
		want  CreateDriverOutput
	}{
		{
			CreateDriverInput{
				BirthDate:  now.AddDate(-18, 0, 0),
				CNH:        "e",
				CPF:        "73552587020",
				Gender:     "m",
				HasVehicle: true,
				Name:       "Henrique Bernardo Rafael Silveira",
			},
			mockCreateDriverRepo{},
			CreateDriverOutput{
				Age:        18,
				BirthDate:  time.Date(now.AddDate(-18, 0, 0).Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), now.Nanosecond(), time.UTC),
				CNH:        "E",
				CPF:        "73552587020",
				Gender:     "M",
				HasVehicle: true,
				Name:       "henrique bernardo rafael silveira",
			},
		},
		{
			CreateDriverInput{
				BirthDate:  now.AddDate(-30, 0, 0),
				CNH:        "a",
				CPF:        "45298982530",
				Gender:     "f",
				HasVehicle: false,
				Name:       "Eduarda Sônia Rebeca Oliveira",
			},
			mockCreateDriverRepo{},
			CreateDriverOutput{
				Age:        30,
				BirthDate:  time.Date(now.AddDate(-30, 0, 0).Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), now.Nanosecond(), time.UTC),
				CNH:        "A",
				CPF:        "45298982530",
				Gender:     "F",
				HasVehicle: false,
				Name:       "eduarda sônia rebeca oliveira",
			},
		},
		{
			CreateDriverInput{
				BirthDate:  now.AddDate(-41, 0, 0),
				CNH:        "b",
				CPF:        "94587146218",
				Gender:     "O",
				HasVehicle: true,
				Name:       "Bárbara Andreia Vieira",
			},
			mockCreateDriverRepo{},
			CreateDriverOutput{
				Age:        41,
				BirthDate:  time.Date(now.AddDate(-41, 0, 0).Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), now.Nanosecond(), time.UTC),
				CNH:        "B",
				CPF:        "94587146218",
				Gender:     "O",
				HasVehicle: true,
				Name:       "bárbara andreia vieira",
			},
		},
	}

	for i, test := range tests {
		uc := NewCreateDriver(FakeLogger{}, test.repo)
		got, gotErr := uc.Execute(context.Background(), test.input)

		if got == nil || gotErr != nil {
			t.Errorf("%d: [input: %v] [wantErr: <nil>] [gotErr: %v]", i, test.input, gotErr)
			continue
		}

		// reset both CreatedAt fields as their value are out of the test's control
		got.CreatedAt = time.Time{}
		test.want.CreatedAt = time.Time{}
		if test.want != *got {
			t.Errorf("%d: [input: %v] [want: %v] [got: %v]", i, test.input, test.want, got)
		}
	}
}

func TestCreateDriverErr(t *testing.T) {
	now := time.Now().UTC()

	tests := []struct {
		input   CreateDriverInput
		repo    mockCreateDriverRepo
		wantErr error
	}{
		{
			CreateDriverInput{
				BirthDate:  now, // invalid BirthDate
				CNH:        "a",
				CPF:        "15279541028",
				Gender:     "f",
				HasVehicle: true,
				Name:       "Jennifer Marlene Freitas",
			},
			mockCreateDriverRepo{},
			entity.NewErrInvalidAge(0),
		},
		{
			CreateDriverInput{
				BirthDate:  now.AddDate(-20, 0, 0),
				CNH:        "7", // invalid CNH
				CPF:        "75703960223",
				Gender:     "m",
				HasVehicle: false,
				Name:       "Benjamin Thomas Monteiro",
			},
			mockCreateDriverRepo{},
			entity.NewErrInvalidCNH("7"),
		},
		{
			CreateDriverInput{
				BirthDate:  now.AddDate(-47, 0, 0),
				CNH:        "c",
				CPF:        "75703960223",
				Gender:     "", // invalid gender
				HasVehicle: true,
				Name:       "Márcio Iago Igor Nascimento",
			},
			mockCreateDriverRepo{},
			entity.NewErrInvalidGender(""),
		},
		{
			CreateDriverInput{
				BirthDate:  now.AddDate(0, -300, 0),
				CNH:        "d",
				CPF:        "55800346100",
				Gender:     "o",
				HasVehicle: false,
				Name:       " Vera Laís Ferreira", // invalid name
			},
			mockCreateDriverRepo{},
			entity.NewErrInvalidName(" Vera Laís Ferreira"),
		},
		{
			CreateDriverInput{
				BirthDate:  now.AddDate(-18, 0, 0),
				CNH:        "e",
				CPF:        "73552587020",
				Gender:     "m",
				HasVehicle: true,
				Name:       "Henrique Bernardo Rafael Silveira",
			},
			mockCreateDriverRepo{
				err: entity.NewErrDriverAlreadyExists("73552587020"),
			},
			entity.NewErrDriverAlreadyExists("73552587020"),
		},
	}

	for i, test := range tests {
		uc := NewCreateDriver(FakeLogger{}, test.repo)
		_, gotErr := uc.Execute(context.Background(), test.input)

		if !errors.Is(gotErr, test.wantErr) {
			t.Errorf("%d: [input: %v] [wantErr: %v] [gotErr: %v]", i, test.input, test.wantErr, gotErr)
		}
	}
}
