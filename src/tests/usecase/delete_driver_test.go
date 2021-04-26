package usecase_test

import (
	"context"
	"reflect"
	"testing"

	repository "github.com/rafaft/truck-driver-trip-system/adapter/gateway/database"
	"github.com/rafaft/truck-driver-trip-system/entity"
	"github.com/rafaft/truck-driver-trip-system/infrastructure/log"
	"github.com/rafaft/truck-driver-trip-system/usecase"
)

func TestDeleteDriver(t *testing.T) {
	l := log.NewFakeLogger()
	r := repository.NewDriverInMemory(getDriversSample(t))
	uc := usecase.NewDeleteDriver(l, r)

	tests := []struct {
		input string
		want  error
	}{
		{
			input: "49628536021", // invalid CPF
			want:  entity.ErrInvalidCPF{},
		},
		{
			input: "47128632000", // CPF never existed
			want:  entity.ErrDriverNotFound{},
		},
		{
			input: "63503201238", // driver exists and should be deleted
			want:  nil,
		},
		{
			input: "63503201238", // once deleted, it should be not found!
			want:  entity.ErrDriverNotFound{},
		},
	}

	for i, test := range tests {
		got := uc.Execute(context.Background(), test.input)

		if reflect.TypeOf(test.want) != reflect.TypeOf(got) {
			t.Errorf("%d: [want: %v] [got: %v]", i, test.want, got)
		}
	}
}
