package usecase_test

import (
	"context"
	"errors"
	"reflect"
	"testing"

	repository "github.com/rafaft/truck-driver-trip-system/adapter/gateway/database"
	"github.com/rafaft/truck-driver-trip-system/entity"
	"github.com/rafaft/truck-driver-trip-system/infrastructure/log"
	"github.com/rafaft/truck-driver-trip-system/tests/samples"
	"github.com/rafaft/truck-driver-trip-system/usecase"
)

func TestGetDrivers(t *testing.T) {
	l := log.NewFakeLogger()
	r := repository.NewDriverInMemory(samples.GetDrivers(t))
	uc := usecase.NewGetDrivers(l, r)

	tests := []struct {
		input usecase.GetDriversQuery
		// want    []*usecase.GetDriversOutput
		wantLen int
		err     error
	}{
		// invalid
		{
			input: usecase.GetDriversQuery{
				CNH: getStrPointer("f"), // invalid CNH
			},
			// want:    nil,
			wantLen: -1,
			err:     entity.ErrInvalidCNH{},
		},
		{
			input: usecase.GetDriversQuery{
				Gender: getStrPointer("a"), // invalid Gender
			},
			// want:    nil,
			wantLen: -1,
			err:     entity.ErrInvalidGender{},
		},
		// valid
		{
			input: usecase.GetDriversQuery{},
			// want:    nil,
			wantLen: 20,
			err:     nil,
		},
		{
			input: usecase.GetDriversQuery{
				Limit: getUintPointer(0),
			},
			// want:    nil,
			wantLen: 0,
			err:     nil,
		},
		{
			input: usecase.GetDriversQuery{
				CNH: getStrPointer("a"),
			},
			// want:    nil,
			wantLen: 4,
			err:     nil,
		},
		{
			input: usecase.GetDriversQuery{
				CNH:    getStrPointer("a"),
				Gender: getStrPointer("M"),
			},
			// want:    nil,
			wantLen: 2,
			err:     nil,
		},
		{
			input: usecase.GetDriversQuery{
				CNH:    getStrPointer("A"),
				Gender: getStrPointer("m"),
				Limit:  getUintPointer(1),
			},
			// want:    nil,
			wantLen: 1,
			err:     nil,
		},
		{
			input: usecase.GetDriversQuery{
				HasVehicle: getBoolPointer(true),
			},
			// want:    nil,
			wantLen: 10,
			err:     nil,
		},
		{
			input: usecase.GetDriversQuery{
				CNH:        getStrPointer("e"),
				Gender:     getStrPointer("F"),
				HasVehicle: getBoolPointer(false),
			},
			// want: []*usecase.GetDriversOutput{
			// 	&usecase.GetDriversOutput{
			// 		Age:        30,
			// 		BirthDate:  time.Now().AddDate(-30, 0, 0),
			// 		CNH:        "E",
			// 		CPF:        "45464490388",
			// 		Gender:     "F",
			// 		HasVehicle: false,
			// 		Name:       "bárbara valentina ana barros",
			// 	},
			// },
			wantLen: 1,
			err:     nil,
		},
	}

	for i, test := range tests {
		got, gotErr := uc.Execute(context.Background(), test.input)

		if err := errors.Unwrap(gotErr); reflect.TypeOf(test.err) != reflect.TypeOf(err) {
			t.Errorf("%d: [wantErr: %v] [gotErr: %v]", i, test.err, err)
			continue
		}

		// test.want == -1 signals that the result length check should be ignored
		if test.wantLen != -1 && test.wantLen != len(got) {
			t.Errorf("%d: [wantLen: %v] [gotLen: %v]", i, test.wantLen, len(got))
			continue
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
