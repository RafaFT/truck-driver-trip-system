package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/rafaft/truck-driver-trip-system/entity"
)

type mockDeleteDriverRepo struct {
	err error
}

func (d mockDeleteDriverRepo) DeleteByCPF(ctx context.Context, cpf entity.CPF) error {
	return d.err
}

func TestDeleteDriver(t *testing.T) {
	tests := []struct {
		cpf     string
		repo    mockDeleteDriverRepo
		wantErr error
	}{
		// invalid input
		{
			"49628536021",
			mockDeleteDriverRepo{},
			entity.NewErrInvalidCPF("49628536021"),
		},
		{
			"47128632000",
			mockDeleteDriverRepo{
				err: entity.NewErrDriverNotFound(entity.CPF("47128632000")),
			},
			entity.NewErrDriverNotFound(entity.CPF("47128632000")),
		},
		// valid input
		{
			"63503201238",
			mockDeleteDriverRepo{},
			nil,
		},
	}

	for i, test := range tests {
		uc := NewDeleteDriver(FakeLogger{}, test.repo)
		gotErr := uc.Execute(context.Background(), test.cpf)

		if !errors.Is(gotErr, test.wantErr) {
			t.Errorf("%d: [wantErr: %v] [gotErr: %v]", i, test.wantErr, gotErr)
		}
	}
}
