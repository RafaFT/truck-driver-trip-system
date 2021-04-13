package entity_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/rafaft/truck-driver-trip-system/entity"
)

func TestNewBirthDate(t *testing.T) {
	tests := []struct {
		input time.Time
		want  entity.BirthDate
		err   error
	}{
		{
			entity.MinBirthDate().AddDate(0, 0, -1),
			entity.BirthDate{time.Time{}},
			entity.ErrInvalidBirthDate{},
		},
		{
			time.Time{},
			entity.BirthDate{time.Time{}},
			entity.ErrInvalidBirthDate{},
		},
		{
			entity.MinBirthDate(),
			entity.BirthDate{entity.MinBirthDate()},
			nil,
		},
		{
			time.Date(2020, 3, 1, 0, 0, 0, 0, time.UTC),
			entity.BirthDate{time.Date(2020, 3, 1, 0, 0, 0, 0, time.UTC)},
			nil,
		},
	}

	for _, test := range tests {
		got, gotErr := entity.NewBirthDate(test.input)

		if !got.Time.Equal(test.want.Time) || reflect.TypeOf(test.err) != reflect.TypeOf(gotErr) {
			t.Errorf("[input: %v] [want: %v] [err: %v] [got: %v] [gotErr: %v]",
				test.input, test.want, test.err, got, gotErr,
			)
		}
	}
}

func TestCalculateAge(t *testing.T) {
	now := time.Now()

	tests := []struct {
		input time.Time
		want  int
	}{
		{now.AddDate(-2020, 0, 0), 2020},
		{now, 0},
		{now.AddDate(0, -1, 0), 0},
		{now.AddDate(-1, 0, 1), 0},
		{now.AddDate(-1, 0, 0), 1},
		{now.AddDate(-1, -1, 0), 1},
		{now.AddDate(-31, 0, 0), 31},
		{now.AddDate(-20, -12, 1), 20},
	}

	for _, test := range tests {
		bd, _ := entity.NewBirthDate(test.input)

		if got := bd.CalculateAge(); got != test.want {
			t.Errorf("[input: %v] [want: %v] [got: %v]",
				test.input, test.want, got,
			)
		}
	}
}
