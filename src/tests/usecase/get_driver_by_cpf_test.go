package usecase_test

import (
	"context"
	"reflect"
	"testing"
	"time"

	repository "github.com/rafaft/truck-driver-trip-system/adapter/gateway/database"
	"github.com/rafaft/truck-driver-trip-system/entity"
	"github.com/rafaft/truck-driver-trip-system/infrastructure/log"
	"github.com/rafaft/truck-driver-trip-system/tests/samples"
	"github.com/rafaft/truck-driver-trip-system/usecase"
)

func TestGetDriverByCPF(t *testing.T) {
	l := log.NewFakeLogger()
	r := repository.NewDriverInMemory(samples.GetDrivers(t))
	uc := usecase.NewGetDriverByCPF(l, r)

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
			input: "77163670303",
			want: &usecase.GetDriverByCPFOutput{
				Age:        23,
				BirthDate:  time.Now().AddDate(-23, 0, 0),
				CNH:        "E",
				CPF:        "77163670303",
				Gender:     "O",
				HasVehicle: true,
				Name:       "allana louise bianca nogueira",
			},
			err: nil,
		},
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
