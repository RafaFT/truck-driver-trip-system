package usecase

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/rafaft/truck-driver-trip-system/entity"
)

type mockGetDriverRepo struct {
	drivers []*entity.Driver
	err     error
}

func (repo mockGetDriverRepo) Find(ctx context.Context, query FindDriversQuery) ([]*entity.Driver, error) {
	return repo.drivers, repo.err
}

func TestGetDrivers(t *testing.T) {
	now := time.Now()
	driver1, _ := entity.NewDriver(
		"75839001074",
		"name",
		"F",
		"a",
		now.AddDate(-20, 0, 0),
		true,
	)
	driver2, _ := entity.NewDriver(
		"29321569006",
		"Joaquim",
		"O",
		"b",
		now.AddDate(-25, 0, 0),
		false,
	)

	tests := []struct {
		input GetDriversQuery
		repo  GetDriverRepo
		want  []*GetDriversOutput
	}{
		{
			GetDriversQuery{},
			mockGetDriverRepo{},
			[]*GetDriversOutput{},
		},
		{
			GetDriversQuery{},
			mockGetDriverRepo{
				drivers: []*entity.Driver{
					driver1,
					driver2,
				},
			},
			[]*GetDriversOutput{
				{
					Age:        20,
					BirthDate:  now.AddDate(-20, 0, 0).UTC(),
					CNH:        "A",
					CPF:        "75839001074",
					Gender:     "F",
					HasVehicle: true,
					Name:       "name",
				},
				{
					Age:        25,
					BirthDate:  now.AddDate(-25, 0, 0).UTC(),
					CNH:        "B",
					CPF:        "29321569006",
					Gender:     "O",
					HasVehicle: false,
					Name:       "joaquim",
				},
			},
		},
		{
			GetDriversQuery{
				CNH:        getStrPointer("B"),
				Gender:     getStrPointer("M"),
				HasVehicle: getBoolPointer(false),
				Limit:      getUintPointer(1),
			},
			mockGetDriverRepo{},
			[]*GetDriversOutput{},
		},
	}

	for i, test := range tests {
		uc := NewGetDrivers(FakeLogger{}, test.repo)
		got, gotErr := uc.Execute(context.Background(), test.input)

		if got == nil || gotErr != nil {
			t.Errorf("%d: [input: %v] [wantErr: <nil>] [gotErr: %v]", i, test.input, gotErr)
			continue
		}

		if !reflect.DeepEqual(test.want, got) {
			t.Errorf("%d:  [input: %v] [want: %v] [got: %v]", i, test.input, test.want, got)
		}
	}
}

func TestGetDriversErr(t *testing.T) {
	networkErr := errors.New("some network err")

	tests := []struct {
		input   GetDriversQuery
		repo    GetDriverRepo
		wantErr error
	}{
		{
			GetDriversQuery{
				CNH: getStrPointer("f"),
			},
			mockGetDriverRepo{},
			entity.NewErrInvalidCNH("f"),
		},
		{
			GetDriversQuery{
				Gender: getStrPointer(""),
			},
			mockGetDriverRepo{},
			entity.NewErrInvalidGender(""),
		},
		{
			GetDriversQuery{},
			mockGetDriverRepo{
				err: networkErr,
			},
			networkErr,
		},
	}

	for i, test := range tests {
		uc := NewGetDrivers(FakeLogger{}, test.repo)
		_, gotErr := uc.Execute(context.Background(), test.input)

		if !errors.Is(gotErr, test.wantErr) {
			t.Errorf("%d: [wantErr: %v] [gotErr: %v]", i, test.wantErr, gotErr)
		}
	}
}

func getBoolPointer(b bool) *bool {
	return &b
}

func getStrPointer(s string) *string {
	return &s
}

func getUintPointer(ui uint) *uint {
	return &ui
}

func getIntPointer(i int) *int {
	return &i
}

func getFloatPointer(f float64) *float64 {
	return &f
}

func getDatePointer(t time.Time) *time.Time {
	return &t
}
