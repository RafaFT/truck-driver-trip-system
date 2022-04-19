package usecase

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/rafaft/truck-driver-trip-system/entity"
)

type mockUpdateDriverRepo struct {
	driver    *entity.Driver
	findErr   error
	updateErr error
}

func (r mockUpdateDriverRepo) FindByCPF(ctx context.Context, cpf entity.CPF) (*entity.Driver, error) {
	return r.driver, r.findErr
}

func (r mockUpdateDriverRepo) Update(ctx context.Context, driver *entity.Driver) error {
	return r.updateErr
}

func TestUpdateDriver(t *testing.T) {
	now := time.Now()
	driver, _ := entity.NewDriver("31803413603", "name", "M", "b", now.AddDate(-46, 0, 0), true)

	tests := []struct {
		cpf   string
		input UpdateDriverInput
		repo  UpdateDriverRepo
		want  *UpdateDriverOutput
	}{
		{
			"31803413603",
			UpdateDriverInput{
				CNH:        getStrPointer("A"),
				Gender:     getStrPointer("o"),
				HasVehicle: getBoolPointer(false),
				Name:       getStrPointer("new name"),
			},
			mockUpdateDriverRepo{
				driver: driver,
			},
			&UpdateDriverOutput{
				Age:        46,
				BirthDate:  now.AddDate(-46, 0, 0).UTC(),
				CNH:        "A",
				CPF:        "31803413603",
				Gender:     "O",
				HasVehicle: false,
				Name:       "new name",
			},
		},
	}

	for i, test := range tests {
		uc := NewUpdateDriver(FakeLogger{}, test.repo)
		got, gotErr := uc.Execute(context.Background(), test.cpf, test.input)

		if got == nil || gotErr != nil {
			t.Errorf("%d: [input: %v] [wantErr: <nil>] [gotErr: %v]", i, test.input, gotErr)
			continue
		}

		got.UpdatedAt = time.Time{}
		test.want.UpdatedAt = time.Time{}
		if !reflect.DeepEqual(test.want, got) {
			t.Errorf("%d: [input: %v] [want: %v] [got: %v]", i, test.input, test.want, got)
		}
	}
}

func TestUpdateDriverErr(t *testing.T) {
	now := time.Now()
	networkErr := errors.New("some network error")
	driver, _ := entity.NewDriver("81183605048", "name", "M", "b", now.AddDate(-46, 0, 0), true)

	tests := []struct {
		cpf     string
		input   UpdateDriverInput
		repo    UpdateDriverRepo
		wantErr error
	}{
		{
			"52742089401",
			UpdateDriverInput{},
			mockUpdateDriverRepo{},
			entity.NewErrInvalidCPF("52742089401"),
		},
		{
			"00645063045",
			UpdateDriverInput{},
			mockUpdateDriverRepo{
				findErr: entity.NewErrDriverNotFound("00645063045"),
			},
			entity.NewErrDriverNotFound("00645063045"),
		},
		{
			"52742089403",
			UpdateDriverInput{
				CNH: getStrPointer("8"),
			},
			mockUpdateDriverRepo{},
			entity.NewErrInvalidCNH("8"),
		},
		{
			"94982599769",
			UpdateDriverInput{
				Gender: getStrPointer("l"),
			},
			mockUpdateDriverRepo{},
			entity.NewErrInvalidGender("l"),
		},
		{
			"08931283849",
			UpdateDriverInput{
				Name: getStrPointer("invalid  name because  of  extra  space"),
			},
			mockUpdateDriverRepo{},
			entity.NewErrInvalidName("invalid  name because  of  extra  space"),
		},
		{
			"81183605048",
			UpdateDriverInput{
				Name: getStrPointer("new name"),
			},
			mockUpdateDriverRepo{
				driver:    driver,
				updateErr: networkErr,
			},
			networkErr,
		},
	}

	for i, test := range tests {
		uc := NewUpdateDriver(FakeLogger{}, test.repo)
		_, gotErr := uc.Execute(context.Background(), test.cpf, test.input)

		if !errors.Is(gotErr, test.wantErr) {
			t.Errorf("%d: [input: %v] [wantErr: %v] [gotErr: %v]", i, test.input, test.wantErr, gotErr)
		}
	}
}
