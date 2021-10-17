package entity

import (
	"errors"
	"testing"
	"time"
)

func TestNewBirthDate(t *testing.T) {
	moscowLocation, _ := time.LoadLocation("Europe/Moscow")
	torontoLocation, _ := time.LoadLocation("America/Toronto")

	tests := []struct {
		input   time.Time
		want    BirthDate
		wantErr error
	}{
		// invalid input
		{
			time.Time{},
			BirthDate{},
			newErrInvalidBirthDate(time.Time{}),
		},
		{
			minBirthDate.AddDate(0, 0, -1),
			BirthDate{},
			newErrInvalidBirthDate(minBirthDate.AddDate(0, 0, -1)),
		},
		{ // +3 UTC offset
			time.Date(1950, 1, 1, 2, 59, 59, 0, moscowLocation),
			BirthDate{},
			newErrInvalidBirthDate(time.Date(1950, 1, 1, 2, 59, 59, 0, moscowLocation)),
		},
		// valid input
		{
			minBirthDate,
			BirthDate{minBirthDate},
			nil,
		},
		{
			time.Date(2020, 3, 1, 0, 0, 0, 0, time.UTC),
			BirthDate{time.Date(2020, 3, 1, 0, 0, 0, 0, time.UTC)},
			nil,
		},
		{ // -5 UTC offset
			time.Date(1949, 12, 31, 19, 0, 0, 0, torontoLocation),
			BirthDate{time.Date(1950, 1, 1, 0, 0, 0, 0, time.UTC)},
			nil,
		},
	}

	for i, test := range tests {
		got, gotErr := NewBirthDate(test.input)

		if !errors.Is(test.wantErr, gotErr) {
			t.Errorf("%d: [input: %v] [wantErr: %v] [gotErr: %v]", i, test.input, test.wantErr, gotErr)
			continue
		}

		if !got.Time.Equal(test.want.Time) {
			t.Errorf("%d: [input: %v] [want: %v] [got: %v]", i, test.input, test.want, got)
		}
	}
}

func TestCalculateAge(t *testing.T) {
	now := time.Now().UTC()

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

	for i, test := range tests {
		bd, _ := NewBirthDate(test.input)

		if got := bd.age(); got != test.want {
			t.Errorf("%d: [input: %v] [want: %v] [got: %v]", i, test.input, test.want, got)
		}
	}
}
