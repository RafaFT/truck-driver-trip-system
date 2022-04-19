package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/rafaft/truck-driver-trip-system/entity"
)

type mockDeleteTripRepo struct {
	err error
}

func (d mockDeleteTripRepo) Delete(ctx context.Context, id string) error {
	return d.err
}

func TestDeleteTrip(t *testing.T) {
	tests := []struct {
		tripID  string
		repo    mockDeleteTripRepo
		wantErr error
	}{
		// invalid input
		{
			"",
			mockDeleteTripRepo{},
			entity.ErrInvalidID,
		},
		{
			"a676b5ad-5ffa-4917-a62e-d0933e53c1bb",
			mockDeleteTripRepo{
				err: entity.NewErrTripNotFound("a676b5ad-5ffa-4917-a62e-d0933e53c1bb"),
			},
			entity.NewErrTripNotFound("a676b5ad-5ffa-4917-a62e-d0933e53c1bb"),
		},
		// valid input
		{
			"794b9937-7afa-449e-9662-92271d44cb81",
			mockDeleteTripRepo{},
			nil,
		},
	}

	for i, test := range tests {
		uc := NewDeleteTrip(FakeLogger{}, test.repo)
		gotErr := uc.Execute(context.Background(), test.tripID)

		if !errors.Is(gotErr, test.wantErr) {
			t.Errorf("%d: [wantErr: %v] [gotErr: %v]", i, test.wantErr, gotErr)
		}
	}
}
