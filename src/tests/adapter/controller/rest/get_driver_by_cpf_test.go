package rest_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rafaft/truck-driver-trip-system/adapter/controller/rest"
	repository "github.com/rafaft/truck-driver-trip-system/adapter/gateway/database"
	"github.com/rafaft/truck-driver-trip-system/adapter/presenter"
	"github.com/rafaft/truck-driver-trip-system/infrastructure/log"
	"github.com/rafaft/truck-driver-trip-system/tests/samples"
	"github.com/rafaft/truck-driver-trip-system/usecase"
)

func TestGetDriverByCPF(t *testing.T) {
	l := log.NewFakeLogger()
	repo := repository.NewDriverInMemory(samples.GetDrivers(t))
	uc := usecase.NewGetDriverByCPF(l, repo)
	p := presenter.NewGetDriverByCPF()
	c := rest.NewGetDriverByCPF(p, uc)

	// type for easier marshalling
	type output struct {
		// compare only age, instead of birth_date
		Age        *int    `json:"age,omitempty"`
		CNH        *string `json:"cnh,omitempty"`
		CPF        *string `json:"cpf,omitempty"`
		Gender     *string `json:"gender,omitempty"`
		HasVehicle *bool   `json:"has_vehicle,omitempty"`
		Name       *string `json:"name,omitempty"`
	}

	tests := []struct {
		rw       *httptest.ResponseRecorder
		r        *http.Request
		cpf      string
		wantCode int
		wantBody string
	}{
		{
			rw:       httptest.NewRecorder(),
			r:        httptest.NewRequest(http.MethodGet, "/", nil),
			cpf:      "33510345398",
			wantCode: http.StatusOK,
			wantBody: `{"age":71,"cnh":"A","cpf":"33510345398","gender":"M","has_vehicle":true,"name":"kaique joão teixeira"}`,
		},
		{
			rw:       httptest.NewRecorder(),
			r:        httptest.NewRequest(http.MethodGet, "/?fields=", nil),
			cpf:      "27188079463",
			wantCode: http.StatusOK,
			wantBody: `{"age":33,"cnh":"D","cpf":"27188079463","gender":"O","has_vehicle":true,"name":"thales marcos fogaça"}`,
		},
		{
			rw:       httptest.NewRecorder(),
			r:        httptest.NewRequest(http.MethodGet, "/?fields=age,cpf", nil),
			cpf:      "08931283849",
			wantCode: http.StatusOK,
			wantBody: `{"age":52,"cpf":"08931283849"}`,
		},
		{
			rw:       httptest.NewRecorder(),
			r:        httptest.NewRequest(http.MethodGet, "/?fields=,", nil),
			cpf:      "69048144620",
			wantCode: http.StatusOK,
			wantBody: `{}`,
		},
	}

	for i, test := range tests {
		ctx := context.WithValue(test.r.Context(), rest.CPFKey("cpf"), test.cpf)
		c.ServeHTTP(test.rw, test.r.WithContext(ctx))

		if test.wantCode != test.rw.Code {
			t.Errorf("%d: wantCode=[%v] gotCode=[%v] gotBody=[%v]", i, test.wantCode, test.rw.Code, test.rw.Body.String())
			continue
		}

		var response output
		err := json.Unmarshal(test.rw.Body.Bytes(), &response)
		if err != nil {
			t.Errorf("%d: Could not unmarshal response body: %s", i, err)
			continue
		}

		if gotBody, _ := json.Marshal(response); test.wantBody != string(gotBody) {
			t.Errorf("%d: wantBody=[%s] gotBody=[%s]", i, test.wantBody, gotBody)
		}
	}
}

func TestGetDriverByCPFErrors(t *testing.T) {
	l := log.NewFakeLogger()
	repo := repository.NewDriverInMemory(nil)
	uc := usecase.NewGetDriverByCPF(l, repo)
	p := presenter.NewGetDriverByCPF()
	c := rest.NewGetDriverByCPF(p, uc)

	tests := []struct {
		rw       *httptest.ResponseRecorder
		r        *http.Request
		cpf      string
		wantCode int
		wantBody string
	}{
		// unexpected/unrecognized parameter key
		{
			rw:       httptest.NewRecorder(),
			r:        httptest.NewRequest(http.MethodGet, "/?gender=o&cnh=b", nil),
			cpf:      "70286951150",
			wantCode: http.StatusBadRequest,
			wantBody: `{"error":"Unexpected query parameter: gender."}`,
		},
		// cpf not found/invalid
		{
			rw:       httptest.NewRecorder(),
			r:        httptest.NewRequest(http.MethodGet, "/", nil),
			cpf:      "00000000000",
			wantCode: http.StatusNotFound,
			wantBody: ``,
		},
	}

	for i, test := range tests {
		ctx := context.WithValue(test.r.Context(), rest.CPFKey("cpf"), test.cpf)
		c.ServeHTTP(test.rw, test.r.WithContext(ctx))

		if test.wantCode != test.rw.Code {
			t.Errorf("%d: wantCode=[%v] gotCode=[%v]", i, test.wantCode, test.rw.Code)
			continue
		}

		if got := test.rw.Body.String(); got != test.wantBody {
			t.Errorf("%d: wantBody=[%v] gotBody=[%v]", i, test.wantBody, got)
		}
	}
}
