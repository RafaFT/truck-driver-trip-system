package rest_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/rafaft/truck-driver-trip-system/adapter/controller/rest"
	repository "github.com/rafaft/truck-driver-trip-system/adapter/gateway/database"
	"github.com/rafaft/truck-driver-trip-system/adapter/presenter"
	"github.com/rafaft/truck-driver-trip-system/infrastructure/log"
	"github.com/rafaft/truck-driver-trip-system/tests/samples"
	"github.com/rafaft/truck-driver-trip-system/usecase"
)

func TestUpdateDriver(t *testing.T) {
	l := log.NewFakeLogger()
	repo := repository.NewDriverInMemory(samples.GetDrivers(t))
	uc := usecase.NewUpdateDriver(l, repo)
	p := presenter.NewUpdateDriver()
	c := rest.NewUpdateDriver(p, uc)

	// valid CPF from samples.GetDrivers()
	validCPF := "63503201238"

	// type for ignoring fields from response body
	type output struct {
		CNH        string `json:"cnh"`
		Gender     string `json:"gender"`
		HasVehicle bool   `json:"has_vehicle"`
		Name       string `json:"name"`
	}

	tests := []struct {
		rw       *httptest.ResponseRecorder
		r        *http.Request
		cpf      string
		wantCode int
		wantBody string
	}{
		// update CNH
		{
			rw:       httptest.NewRecorder(),
			r:        httptest.NewRequest(http.MethodPatch, "/", strings.NewReader(`{"cnh":"a"}`)),
			cpf:      validCPF,
			wantCode: http.StatusOK,
			wantBody: `{"cnh":"A","gender":"F","has_vehicle":true,"name":"otávio benício ricardo ramos"}`,
		},
		// update gender
		{
			rw:       httptest.NewRecorder(),
			r:        httptest.NewRequest(http.MethodPatch, "/", strings.NewReader(`{"gender":"O"}`)),
			cpf:      validCPF,
			wantCode: http.StatusOK,
			wantBody: `{"cnh":"A","gender":"O","has_vehicle":true,"name":"otávio benício ricardo ramos"}`,
		},
		// update has_vehicle
		{
			rw:       httptest.NewRecorder(),
			r:        httptest.NewRequest(http.MethodPatch, "/", strings.NewReader(`{"has_vehicle":false}`)),
			cpf:      validCPF,
			wantCode: http.StatusOK,
			wantBody: `{"cnh":"A","gender":"O","has_vehicle":false,"name":"otávio benício ricardo ramos"}`,
		},
		// update name
		{
			rw:       httptest.NewRecorder(),
			r:        httptest.NewRequest(http.MethodPatch, "/", strings.NewReader(`{"name":"Mariana Luciana Alice Silveira"}`)),
			cpf:      validCPF,
			wantCode: http.StatusOK,
			wantBody: `{"cnh":"A","gender":"O","has_vehicle":false,"name":"mariana luciana alice silveira"}`,
		},
		// re-set all document changes
		{
			rw:       httptest.NewRecorder(),
			r:        httptest.NewRequest(http.MethodPatch, "/", strings.NewReader(`{"cnh":"B","gender":"f","has_vehicle":true,"name":"Otávio Benício Ricardo Ramos"}`)),
			cpf:      validCPF,
			wantCode: http.StatusOK,
			wantBody: `{"cnh":"B","gender":"F","has_vehicle":true,"name":"otávio benício ricardo ramos"}`,
		},
	}

	for i, test := range tests {
		ctx := context.WithValue(test.r.Context(), rest.CPFKey("cpf"), test.cpf)
		c.ServeHTTP(test.rw, test.r.WithContext(ctx))

		if test.wantCode != test.rw.Code {
			t.Errorf("%d: wantCode=[%v] gotCode=[%v]", i, test.wantCode, test.rw.Code)
			continue
		}

		var result output
		err := json.Unmarshal(test.rw.Body.Bytes(), &result)
		if err != nil {
			t.Errorf("%d: Could not unmarshal response body", i)
			continue
		}

		b, _ := json.Marshal(&result)
		if got := string(b); got != test.wantBody {
			t.Errorf("%d: wantBody=[%v] gotBody=[%v]", i, test.wantBody, got)
		}
	}
}

func TestUpdateDriverErrors(t *testing.T) {
	l := log.NewFakeLogger()
	repo := repository.NewDriverInMemory(samples.GetDrivers(t))
	uc := usecase.NewUpdateDriver(l, repo)
	p := presenter.NewUpdateDriver()
	c := rest.NewUpdateDriver(p, uc)

	// valid CPF from samples.GetDrivers()
	validCPF := "63503201238"

	tests := []struct {
		rw       *httptest.ResponseRecorder
		r        *http.Request
		cpf      string
		wantCode int
		wantBody string
	}{
		// invalid JSON payload
		{
			rw:       httptest.NewRecorder(),
			r:        httptest.NewRequest(http.MethodGet, "/", nil),
			cpf:      validCPF,
			wantCode: http.StatusBadRequest,
			wantBody: `{"error":"Invalid JSON."}`,
		},
		{
			rw:       httptest.NewRecorder(),
			r:        httptest.NewRequest(http.MethodGet, "/", strings.NewReader(``)),
			cpf:      validCPF,
			wantCode: http.StatusBadRequest,
			wantBody: `{"error":"Invalid JSON."}`,
		},
		{
			rw:       httptest.NewRecorder(),
			r:        httptest.NewRequest(http.MethodGet, "/", strings.NewReader(`"name":"someone else}"`)),
			cpf:      validCPF,
			wantCode: http.StatusBadRequest,
			wantBody: `{"error":"Invalid JSON."}`,
		},
		// valid JSON payload, but of different type as expected
		{
			rw:       httptest.NewRecorder(),
			r:        httptest.NewRequest(http.MethodGet, "/", strings.NewReader(`"string"`)),
			cpf:      validCPF,
			wantCode: http.StatusBadRequest,
			wantBody: `{"error":"Expected JSON Object."}`,
		},
		{
			rw:       httptest.NewRecorder(),
			r:        httptest.NewRequest(http.MethodGet, "/", strings.NewReader(`2`)),
			cpf:      validCPF,
			wantCode: http.StatusBadRequest,
			wantBody: `{"error":"Expected JSON Object."}`,
		},
		{
			rw:       httptest.NewRecorder(),
			r:        httptest.NewRequest(http.MethodGet, "/", strings.NewReader(`false`)),
			cpf:      validCPF,
			wantCode: http.StatusBadRequest,
			wantBody: `{"error":"Expected JSON Object."}`,
		},
		{
			rw:       httptest.NewRecorder(),
			r:        httptest.NewRequest(http.MethodGet, "/", strings.NewReader(`[]`)),
			cpf:      validCPF,
			wantCode: http.StatusBadRequest,
			wantBody: `{"error":"Expected JSON Object."}`,
		},
		// valid JSON Object, but without fields
		{
			rw:       httptest.NewRecorder(),
			r:        httptest.NewRequest(http.MethodGet, "/", strings.NewReader(`{}`)),
			cpf:      validCPF,
			wantCode: http.StatusBadRequest,
			wantBody: `{"error":"Empty update."}`,
		},
		// valid JSON with correct field(s) but different types as expected
		{
			rw:       httptest.NewRecorder(),
			r:        httptest.NewRequest(http.MethodGet, "/", strings.NewReader(`{"cnh":false}`)),
			cpf:      validCPF,
			wantCode: http.StatusBadRequest,
			wantBody: `{"error":"Invalid value at 'cnh'. Expected string, got bool."}`,
		},
		{
			rw:       httptest.NewRecorder(),
			r:        httptest.NewRequest(http.MethodGet, "/", strings.NewReader(`{"has_vehicle":"false"}`)),
			cpf:      validCPF,
			wantCode: http.StatusBadRequest,
			wantBody: `{"error":"Invalid value at 'has_vehicle'. Expected bool, got string."}`,
		},
		// valid JSON with unexpected fields
		{
			rw:       httptest.NewRecorder(),
			r:        httptest.NewRequest(http.MethodGet, "/", strings.NewReader(`{"invalid":"f"}`)),
			cpf:      validCPF,
			wantCode: http.StatusBadRequest,
			wantBody: `{"error":"Unexpected JSON field: 'invalid'."}`,
		},
		{
			rw:       httptest.NewRecorder(),
			r:        httptest.NewRequest(http.MethodGet, "/", strings.NewReader(`{"cnh":"e","gender_":true}`)),
			cpf:      validCPF,
			wantCode: http.StatusBadRequest,
			wantBody: `{"error":"Unexpected JSON field: 'gender_'."}`,
		},
		// invalid CNH
		{
			rw:       httptest.NewRecorder(),
			r:        httptest.NewRequest(http.MethodGet, "/", strings.NewReader(`{"cnh":"1"}`)),
			cpf:      validCPF,
			wantCode: http.StatusBadRequest,
			wantBody: `{"error":"CNH invalid. cnh=[1]"}`,
		},
		// invalid Gender
		{
			rw:       httptest.NewRecorder(),
			r:        httptest.NewRequest(http.MethodGet, "/", strings.NewReader(`{"gender":"not a real gender"}`)),
			cpf:      validCPF,
			wantCode: http.StatusBadRequest,
			wantBody: `{"error":"Gender invalid. gender=[not a real gender]"}`,
		},
		// invalid Name
		{
			rw:       httptest.NewRecorder(),
			r:        httptest.NewRequest(http.MethodGet, "/", strings.NewReader(`{"name":""}`)),
			cpf:      validCPF,
			wantCode: http.StatusBadRequest,
			wantBody: `{"error":"Name invalid. name=[]"}`,
		},
		// driver not found
		{
			rw:       httptest.NewRecorder(),
			r:        httptest.NewRequest(http.MethodGet, "/", strings.NewReader(`{"gender":"f"}`)),
			cpf:      "82246003008",
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
