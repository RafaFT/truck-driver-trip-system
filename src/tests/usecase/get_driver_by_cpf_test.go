package usecase_test

import (
	"context"
	"reflect"
	"testing"
	"time"

	repo "github.com/rafaft/truck-driver-trip-system/adapter/gateway/database"
	"github.com/rafaft/truck-driver-trip-system/entity"
	"github.com/rafaft/truck-driver-trip-system/infrastructure/logger"
	"github.com/rafaft/truck-driver-trip-system/usecase"
)

func TestGetDriverByCPF(t *testing.T) {
	l := logger.NewFakeLogger()
	r := repo.NewDriverInMemory()
	uc := usecase.NewGetDriverByCPFInteractor(l, r)

	tests := []struct {
		input string
		want  *usecase.GetDriverByCPFOutput
		err   error
	}{
		{
			input: "375invalid79791250", // invalid CPF
			want:  nil,
			err:   entity.ErrInvalidCPF{},
		},
		{
			input: "14084152242", // driver doesn't exist
			want:  nil,
			err:   entity.ErrDriverNotFound{},
		},
		{
			input: "85296235762",
			want: &usecase.GetDriverByCPFOutput{
				Age:        43,
				BirthDate:  time.Now().AddDate(-43, 0, 0),
				CNH:        "B",
				CPF:        "85296235762",
				Gender:     "O",
				HasVehicle: true,
				Name:       "analu kamilly sophie oliveira",
			},
			err: nil,
		},
	}

	driver, err := entity.NewTruckDriver(
		"85296235762",
		"Analu Kamilly Sophie Oliveira",
		"o",
		"b",
		time.Now().AddDate(-43, 0, 0),
		true,
	)
	if err != nil {
		t.Fatalf("could not initialize driver for get_driver_by_cpf_test")
	}

	if err := r.SaveDriver(context.Background(), driver); err != nil {
		t.Fatalf("could not save driver for delete_driver_test")
	}

	for i, test := range tests {
		got, gotErr := uc.Execute(context.Background(), test.input)

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
