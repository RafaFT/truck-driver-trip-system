package presenter_test

import (
	"bytes"
	"testing"
	"time"

	"github.com/rafaft/truck-driver-trip-system/adapter/presenter"
	"github.com/rafaft/truck-driver-trip-system/usecase"
)

func TestCreateDriver(t *testing.T) {
	p := presenter.NewCreateDriver()

	tests := []struct {
		input *usecase.CreateDriverOutput
		want  []byte
	}{
		{
			input: &usecase.CreateDriverOutput{
				Age:        71,
				BirthDate:  time.Date(1997, time.Month(2), 17, 21, 1, 2, 3, time.UTC),
				CNH:        "C",
				CPF:        "33510345398",
				CreatedAt:  time.Date(1997, time.Month(2), 17, 0, 0, 0, 0, time.UTC),
				Gender:     "F",
				HasVehicle: false,
				Name:       "Kaique João Teixeira",
			},
			want: []byte(`{"age":71,"birth_date":"1997-02-17","cnh":"C","cpf":"33510345398","created_at":"1997-02-17T00:00:00Z","gender":"F","has_vehicle":false,"name":"Kaique João Teixeira"}`),
		},
	}

	for _, test := range tests {
		got := p.Output(test.input)

		if diff := bytes.Compare(test.want, got); diff != 0 {
			t.Errorf("input=[%v] want=[%v] got=[%v] diffCount=[%d]", test.input, string(test.want), string(got), diff)
		}
	}
}
