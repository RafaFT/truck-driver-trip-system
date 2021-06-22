package presenter_test

import (
	"bytes"
	"testing"
	"time"

	"github.com/rafaft/truck-driver-trip-system/adapter/presenter"
	"github.com/rafaft/truck-driver-trip-system/usecase"
)

func TestGetDriverByCPF(t *testing.T) {
	p := presenter.NewGetDriverByCPF()
	input := &usecase.GetDriverByCPFOutput{
		Age:        71,
		BirthDate:  time.Date(1997, time.Month(2), 17, 21, 1, 2, 3, time.UTC),
		CNH:        "C",
		CPF:        "33510345398",
		Gender:     "F",
		HasVehicle: false,
		Name:       "Kaique João Teixeira",
	}

	tests := []struct {
		fields []string
		want   []byte
	}{
		{
			fields: nil,
			want:   []byte(`{"age":71,"birth_date":"1997-02-17","cnh":"C","cpf":"33510345398","gender":"F","has_vehicle":false,"name":"Kaique João Teixeira"}`),
		},
		{
			fields: []string{},
			want:   []byte(`{}`),
		},
		{
			fields: []string{"age"},
			want:   []byte(`{"age":71}`),
		},
		{
			fields: []string{"age", "birth_date"},
			want:   []byte(`{"age":71,"birth_date":"1997-02-17"}`),
		},
		{
			fields: []string{"age","birth_date","cnh"},
			want:   []byte(`{"age":71,"birth_date":"1997-02-17","cnh":"C"}`),
		},
		{
			fields: []string{"age","birth_date","cnh","cpf"},
			want:   []byte(`{"age":71,"birth_date":"1997-02-17","cnh":"C","cpf":"33510345398"}`),
		},
		{
			fields: []string{"age","birth_date","cnh","cpf","gender"},
			want:   []byte(`{"age":71,"birth_date":"1997-02-17","cnh":"C","cpf":"33510345398","gender":"F"}`),
		},
		{
			fields: []string{"age","birth_date","cnh","cpf","gender","has_vehicle"},
			want:   []byte(`{"age":71,"birth_date":"1997-02-17","cnh":"C","cpf":"33510345398","gender":"F","has_vehicle":false}`),
		},
		{
			fields: []string{"age","birth_date","cnh","cpf","gender","has_vehicle","name"},
			want:   []byte(`{"age":71,"birth_date":"1997-02-17","cnh":"C","cpf":"33510345398","gender":"F","has_vehicle":false,"name":"Kaique João Teixeira"}`),
		},
	}

	for _, test := range tests {
		got := p.Output(input, test.fields...)

		if diff := bytes.Compare(test.want, got); diff != 0 {
			t.Errorf("fields=[%v] want=[%v] got=[%v] diffCount=[%d]", test.fields, string(test.want), string(got), diff)
		}
	}
}
