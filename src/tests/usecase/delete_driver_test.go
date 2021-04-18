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

func TestDeleteDriver(t *testing.T) {
	l := logger.NewFakeLogger()
	r := repo.NewDriverInMemory()
	uc := usecase.NewDeleteDriverInteractor(l, r)

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
			input: "55177294013", // driver exists and should be deleted
			want:  nil,
		},
		{
			input: "55177294013", // once deleted, it should be not found!
			want:  entity.ErrDriverNotFound{},
		},
	}

	driver, err := entity.NewTruckDriver(
		"55177294013",
		"Amanda Rayssa Oliveira",
		"f",
		"b",
		time.Now().AddDate(-53, 0, 0),
		false,
	)
	if err != nil {
		t.Fatalf("could not initialize driver for delete_driver_test")
	}

	if err := r.SaveDriver(context.Background(), driver); err != nil {
		t.Fatalf("could not save driver for delete_driver_test")
	}

	for i, test := range tests {
		got := uc.Execute(context.Background(), test.input)

		if reflect.TypeOf(test.want) != reflect.TypeOf(got) {
			t.Errorf("%d: [want: %v] [got: %v]", i, test.want, got)
		}
	}
}
