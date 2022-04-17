package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/rafaft/truck-driver-trip-system/entity"
)

type mockGetDriverByCPFRepo struct {
	driver *entity.Driver
	err    error
}

func (d mockGetDriverByCPFRepo) FindByCPF(ctx context.Context, cpf entity.CPF) (*entity.Driver, error) {
	return d.driver, d.err
}

func TestGetDriverByCPF(t *testing.T) {
	now := time.Now().UTC()
	driver, _ := entity.NewDriver(
		"77163670303",
		"allana louise bianca nogueira",
		"o",
		"E",
		now.AddDate(-23, 0, 0),
		true,
	)

	tests := []struct {
		cpf  string
		repo mockGetDriverByCPFRepo
		want GetDriverByCPFOutput
	}{
		{
			"77163670303",
			mockGetDriverByCPFRepo{
				driver: driver,
			},
			GetDriverByCPFOutput{
				Age:        23,
				BirthDate:  now.AddDate(-23, 0, 0),
				CNH:        "E",
				CPF:        "77163670303",
				Gender:     "O",
				HasVehicle: true,
				Name:       "allana louise bianca nogueira",
			},
		},
	}

	for i, test := range tests {
		uc := NewGetDriverByCPF(fakeLogger{}, test.repo)
		got, gotErr := uc.Execute(context.Background(), test.cpf)

		if got == nil || gotErr != nil {
			t.Errorf("%d: [input: %v] [wantErr: <nil>] [gotErr: %v]", i, test.cpf, gotErr)
			continue
		}

		if test.want != *got {
			t.Errorf("%d: [input: %v] [want: %v] [got: %v]", i, test.cpf, test.want, got)
		}
	}
}

func TestGetDriverByCPFErr(t *testing.T) {
	tests := []struct {
		cpf     string
		repo    mockGetDriverByCPFRepo
		wantErr error
	}{
		{
			"375invalid79791250",
			mockGetDriverByCPFRepo{},
			entity.NewErrInvalidCPF("375invalid79791250"),
		},
		{
			"14084152242",
			mockGetDriverByCPFRepo{
				err: entity.NewErrDriverNotFound("14084152242"),
			},
			entity.NewErrDriverNotFound("14084152242"),
		},
	}

	for i, test := range tests {
		uc := NewGetDriverByCPF(fakeLogger{}, test.repo)
		_, gotErr := uc.Execute(context.Background(), test.cpf)

		if !errors.Is(gotErr, test.wantErr) {
			t.Errorf("%d: [wantErr: %v] [gotErr: %v]", i, test.wantErr, gotErr)
		}
	}
}
