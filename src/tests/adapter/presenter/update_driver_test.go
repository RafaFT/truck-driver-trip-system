package presenter_test

import (
	"bytes"
	"testing"
	"time"

	"github.com/rafaft/truck-driver-trip-system/adapter/presenter"
	"github.com/rafaft/truck-driver-trip-system/usecase"
)

func TestUpdateDriver(t *testing.T) {
	p := presenter.NewUpdateDriver()

	tests := []struct {
		input *usecase.UpdateDriverOutput
		want  []byte
	}{
		{
			input: &usecase.UpdateDriverOutput{
				Age:        37,
				BirthDate:  time.Date(1972, time.Month(11), 29, 21, 1, 2, 3, time.UTC),
				CNH:        "A",
				CPF:        "31803413603",
				Gender:     "O",
				HasVehicle: true,
				Name:       "Rayssa Emanuelly Andrea Viana",
				UpdatedAt:  time.Date(2021, time.Month(5), 29, 20, 43, 2, 3, time.UTC),
			},
			want: []byte(`{"age":37,"birth_date":"1972-11-29","cnh":"A","cpf":"31803413603","gender":"O","has_vehicle":true,"name":"Rayssa Emanuelly Andrea Viana","updated_at":"2021-05-29T20:43:02.000000003Z"}`),
		},
	}

	for _, test := range tests {
		got := p.Output(test.input)

		if diff := bytes.Compare(test.want, got); diff != 0 {
			t.Errorf("fields=[%v] want=[%v] got=[%v] diffCount=[%d]", test.input, string(test.want), string(got), diff)
		}
	}
}
