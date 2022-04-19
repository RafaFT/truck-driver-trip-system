package rest_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/rafaft/truck-driver-trip-system/adapter/controller/rest"
	repository "github.com/rafaft/truck-driver-trip-system/adapter/gateway/database"
	"github.com/rafaft/truck-driver-trip-system/adapter/presenter"
	"github.com/rafaft/truck-driver-trip-system/tests/samples"
	"github.com/rafaft/truck-driver-trip-system/usecase"
)

func TestDeleteDriver(t *testing.T) {
	l := usecase.FakeLogger{}
	repo := repository.NewDriverInMemory(samples.GetDrivers(t))
	uc := usecase.NewDeleteDriver(l, repo)
	p := presenter.NewDeleteDriver()
	c := rest.NewDeleteDriverByCPF(p, uc)

	tests := []struct {
		rw       *httptest.ResponseRecorder
		r        *http.Request
		cpf      string
		wantCode int
	}{
		{
			rw:       httptest.NewRecorder(),
			r:        httptest.NewRequest(http.MethodDelete, "/", nil),
			cpf:      "57765277677",
			wantCode: http.StatusNoContent,
		},
	}

	for i, test := range tests {
		ctx := context.WithValue(test.r.Context(), rest.CPFKey("cpf"), test.cpf)
		c.ServeHTTP(test.rw, test.r.WithContext(ctx))

		if test.wantCode != test.rw.Code {
			t.Errorf("%d: wantCode=[%v] gotCode=[%v]", i, test.wantCode, test.rw.Code)
			continue
		}
	}
}

func TestDeleteDriverErrors(t *testing.T) {
	l := usecase.FakeLogger{}
	repo := repository.NewDriverInMemory(samples.GetDrivers(t))
	uc := usecase.NewDeleteDriver(l, repo)
	p := presenter.NewDeleteDriver()
	c := rest.NewDeleteDriverByCPF(p, uc)

	tests := []struct {
		rw       *httptest.ResponseRecorder
		r        *http.Request
		cpf      string
		wantCode int
	}{
		// CPF does not exist
		{
			rw:       httptest.NewRecorder(),
			r:        httptest.NewRequest(http.MethodDelete, "/", nil),
			cpf:      "62670879055",
			wantCode: http.StatusNotFound,
		},
		// CPF invalid
		{
			rw:       httptest.NewRecorder(),
			r:        httptest.NewRequest(http.MethodDelete, "/", strings.NewReader(``)),
			cpf:      "not a real CPF",
			wantCode: http.StatusNotFound,
		},
	}

	for i, test := range tests {
		ctx := context.WithValue(test.r.Context(), rest.CPFKey("cpf"), test.cpf)
		c.ServeHTTP(test.rw, test.r.WithContext(ctx))

		if test.wantCode != test.rw.Code {
			t.Errorf("%d: wantCode=[%v] gotCode=[%v]", i, test.wantCode, test.rw.Code)
			continue
		}
	}
}
