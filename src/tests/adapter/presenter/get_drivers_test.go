package presenter_test

import (
	"bytes"
	"testing"
	"time"

	"github.com/rafaft/truck-driver-trip-system/adapter/presenter"
	"github.com/rafaft/truck-driver-trip-system/usecase"
)

func TestGetDrivers(t *testing.T) {
	now := time.Now()
	p := presenter.NewGetDrivers()

	tests := []struct {
		input  []*usecase.GetDriversOutput
		fields []string
		want   []byte
	}{
		{
			input:  nil,
			fields: nil,
			want:   []byte(`[]`),
		},
		{
			input:  nil,
			fields: []string{"age", "name"},
			want:   []byte(`[]`),
		},
		{
			input: []*usecase.GetDriversOutput{
				{
					Age:        18,
					CPF:        "63503201238",
					BirthDate:  now.AddDate(-18, 0, 0),
					CNH:        "B",
					Gender:     "F",
					HasVehicle: true,
					Name:       "Otávio Benício Ricardo Ramos",
				},
			},
			fields: nil,
			want:   []byte(`[{"age":18,"birth_date":"2003-05-29","cnh":"B","cpf":"63503201238","gender":"F","has_vehicle":true,"name":"Otávio Benício Ricardo Ramos"}]`),
		},
		{
			input: []*usecase.GetDriversOutput{
				{
					Age:        71,
					CPF:        "33510345398",
					BirthDate:  now.AddDate(-71, 0, 0),
					CNH:        "a",
					Gender:     "m",
					HasVehicle: true,
					Name:       "Kaique João Teixeira",
				},
				{
					Age:        65,
					CPF:        "52742089403",
					BirthDate:  now.AddDate(-65, 0, 0),
					CNH:        "b",
					Gender:     "f",
					HasVehicle: false,
					Name:       "Lorenzo Ian Carlos Eduardo Drumond",
				},
			},
			fields: []string{},
			want:   []byte(`[{},{}]`),
		},
		{
			input: []*usecase.GetDriversOutput{
				{
					Age:        60,
					CPF:        "70286951150",
					BirthDate:  now.AddDate(-60, 0, 0),
					CNH:        "c",
					Gender:     "o",
					HasVehicle: true,
					Name:       "Gustavo Francisco Cardoso",
				},
				{
					Age:        57,
					CPF:        "94982599769",
					BirthDate:  now.AddDate(-57, 0, 0),
					CNH:        "d",
					Gender:     "m",
					HasVehicle: false,
					Name:       "Fátima Vanessa Monteiro",
				},
			},
			fields: []string{"name", "has_vehicle"},
			want:   []byte(`[{"has_vehicle":true,"name":"Gustavo Francisco Cardoso"},{"has_vehicle":false,"name":"Fátima Vanessa Monteiro"}]`),
		},
		{
			input: []*usecase.GetDriversOutput{
				{
					Age:        52,
					CPF:        "08931283849",
					BirthDate:  now.AddDate(-52, 0, 0),
					CNH:        "E",
					Gender:     "F",
					HasVehicle: true,
					Name:       "Ana Giovanna Porto",
				},
				{
					Age:        46,
					CPF:        "31803413603",
					BirthDate:  now.AddDate(-46, 0, 0),
					CNH:        "A",
					Gender:     "O",
					HasVehicle: false,
					Name:       "Rayssa Emanuelly Andrea Viana",
				},
			},
			fields: []string{"name", "has_vehicle", "gender", "age"},
			want:   []byte(`[{"age":52,"gender":"F","has_vehicle":true,"name":"Ana Giovanna Porto"},{"age":46,"gender":"O","has_vehicle":false,"name":"Rayssa Emanuelly Andrea Viana"}]`),
		},
		{
			input: []*usecase.GetDriversOutput{
				{
					Age:        41,
					CPF:        "72595478710",
					BirthDate:  now.AddDate(-41, 0, 0),
					CNH:        "B",
					Gender:     "M",
					HasVehicle: true,
					Name:       "Raimundo Erick Nicolas Souza",
				},
				{
					Age:        37,
					CPF:        "94665431728",
					BirthDate:  now.AddDate(-37, 0, 0),
					CNH:        "C",
					Gender:     "F",
					HasVehicle: false,
					Name:       "Bernardo Rafael Julio Figueiredo",
				},
			},
			fields: []string{"name", "has_vehicle", "gender", "age", "cnh", "cpf"},
			want:   []byte(`[{"age":41,"cnh":"B","cpf":"72595478710","gender":"M","has_vehicle":true,"name":"Raimundo Erick Nicolas Souza"},{"age":37,"cnh":"C","cpf":"94665431728","gender":"F","has_vehicle":false,"name":"Bernardo Rafael Julio Figueiredo"}]`),
		},
	}

	for _, test := range tests {
		got := p.Output(test.input, test.fields...)

		if diff := bytes.Compare(test.want, got); diff != 0 {
			t.Errorf("fields=[%v] want=[%v] got=[%v] diffCount=[%d]", test.fields, string(test.want), string(got), diff)
		}
	}
}
