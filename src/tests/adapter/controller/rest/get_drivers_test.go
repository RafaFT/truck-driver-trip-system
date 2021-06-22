package rest_test

import (
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

func TestGetDrivers(t *testing.T) {
	l := log.NewFakeLogger()
	repo := repository.NewDriverInMemory(samples.GetDrivers(t))
	uc := usecase.NewGetDrivers(l, repo)
	p := presenter.NewGetDrivers()
	c := rest.NewGetDrivers(p, uc)

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
		wantCode int
		wantBody string
		count    int
	}{
		{
			rw:       httptest.NewRecorder(),
			r:        httptest.NewRequest(http.MethodGet, "/", nil),
			wantCode: http.StatusOK,
			wantBody: "",
			count:    20,
		},
		{
			rw:       httptest.NewRecorder(),
			r:        httptest.NewRequest(http.MethodGet, "/?gender=m&fields=doesNotexist", nil),
			wantCode: http.StatusOK,
			wantBody: `[{},{},{},{},{},{},{}]`,
			count:    7,
		},
		{
			rw:       httptest.NewRecorder(),
			r:        httptest.NewRequest(http.MethodGet, "/?has_vehicle=false", nil),
			wantCode: http.StatusOK,
			wantBody: "",
			count:    10,
		},
		{
			rw:       httptest.NewRecorder(),
			r:        httptest.NewRequest(http.MethodGet, "/?cnh=C", nil),
			wantCode: http.StatusOK,
			wantBody: "",
			count:    4,
		},
		{
			rw:       httptest.NewRecorder(),
			r:        httptest.NewRequest(http.MethodGet, "/?limit=0&cnh=c", nil),
			wantCode: http.StatusOK,
			wantBody: "",
			count:    0,
		},
		{
			rw:       httptest.NewRecorder(),
			r:        httptest.NewRequest(http.MethodGet, "/?gender=o&cnh=A", nil),
			wantCode: http.StatusOK,
			wantBody: `[{"age":46,"cnh":"A","cpf":"31803413603","gender":"O","has_vehicle":false,"name":"rayssa emanuelly andrea viana"}]`,
			count:    1,
		},
		{
			rw:       httptest.NewRecorder(),
			r:        httptest.NewRequest(http.MethodGet, "/?cnh=d&gender=m&has_vehicle=true&fields=name,cpf,age,,,", nil),
			wantCode: http.StatusOK,
			wantBody: `[{"age":53,"cpf":"31300454652","name":"danilo bryan mateus melo"}]`,
			count:    1,
		},
	}

	for i, test := range tests {
		c.ServeHTTP(test.rw, test.r)

		if test.wantCode != test.rw.Code {
			t.Errorf("%d: wantCode=[%v] gotCode=[%v]", i, test.wantCode, test.rw.Code)
			continue
		}

		response := make([]output, 0)
		err := json.Unmarshal(test.rw.Body.Bytes(), &response)
		if err != nil {
			t.Errorf("%d: Could not unmarshal response body: %s", i, err)
			continue
		}

		if len(response) != test.count {
			t.Errorf("%d: expectedCount=[%d] gotCount=[%d]", i, test.count, len(response))
			continue
		}

		if test.wantBody != "" {
			if gotBody, _ := json.Marshal(response); test.wantBody != string(gotBody) {
				t.Errorf("%d: wantBody=[%s] gotBody=[%s]", i, test.wantBody, gotBody)
			}
		}
	}
}

func TestGetDriversErrors(t *testing.T) {
	l := log.NewFakeLogger()
	repo := repository.NewDriverInMemory(nil)
	uc := usecase.NewGetDrivers(l, repo)
	p := presenter.NewGetDrivers()
	c := rest.NewGetDrivers(p, uc)

	tests := []struct {
		rw       *httptest.ResponseRecorder
		r        *http.Request
		wantCode int
		wantBody string
	}{
		// invalid query parameter values
		{
			rw:       httptest.NewRecorder(),
			r:        httptest.NewRequest(http.MethodGet, "/?cnh=1", nil),
			wantCode: http.StatusBadRequest,
			wantBody: `{"error":"Invalid value at 'cnh'. Expected CNH."}`,
		},
		{
			rw:       httptest.NewRecorder(),
			r:        httptest.NewRequest(http.MethodGet, "/?gender=a", nil),
			wantCode: http.StatusBadRequest,
			wantBody: `{"error":"Invalid value at 'gender'. Expected Gender."}`,
		},
		{
			rw:       httptest.NewRecorder(),
			r:        httptest.NewRequest(http.MethodGet, "/?limit=-5", nil),
			wantCode: http.StatusBadRequest,
			wantBody: `{"error":"Invalid value at 'limit'. Expected positive integer."}`,
		},
		{
			rw:       httptest.NewRecorder(),
			r:        httptest.NewRequest(http.MethodGet, "/?has_vehicle=notABoolean", nil),
			wantCode: http.StatusBadRequest,
			wantBody: `{"error":"Invalid value at 'has_vehicle'. Expected bool."}`,
		},
		{ // has_vehicle check happens before limit. Check ends at first wrong value found
			rw:       httptest.NewRecorder(),
			r:        httptest.NewRequest(http.MethodGet, "/?limit=a&has_vehicle=nottrue", nil),
			wantCode: http.StatusBadRequest,
			wantBody: `{"error":"Invalid value at 'has_vehicle'. Expected bool."}`,
		},
		// unexpected/unrecognized parameter key
		{
			rw:       httptest.NewRecorder(),
			r:        httptest.NewRequest(http.MethodGet, "/?date=2020-01-01", nil),
			wantCode: http.StatusBadRequest,
			wantBody: `{"error":"Unexpected query parameter: date."}`,
		},
	}

	for i, test := range tests {
		c.ServeHTTP(test.rw, test.r)

		if test.wantCode != test.rw.Code {
			t.Errorf("%d: wantCode=[%v] gotCode=[%v]", i, test.wantCode, test.rw.Code)
			continue
		}

		if got := test.rw.Body.String(); got != test.wantBody {
			t.Errorf("%d: wantBody=[%v] gotBody=[%v]", i, test.wantBody, got)
		}
	}
}
