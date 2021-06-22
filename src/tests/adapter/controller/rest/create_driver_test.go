package rest_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/rafaft/truck-driver-trip-system/adapter/controller/rest"
	repository "github.com/rafaft/truck-driver-trip-system/adapter/gateway/database"
	"github.com/rafaft/truck-driver-trip-system/adapter/presenter"
	"github.com/rafaft/truck-driver-trip-system/infrastructure/log"
	"github.com/rafaft/truck-driver-trip-system/tests/samples"
	"github.com/rafaft/truck-driver-trip-system/usecase"
)

func TestCreateDriver(t *testing.T) {
	const target = "https://127.0.0.1"

	l := log.NewFakeLogger()
	repo := repository.NewDriverInMemory(nil)
	uc := usecase.NewCreateDriver(l, repo)
	p := presenter.NewCreateDriver()
	c := rest.NewCreateDriver(p, uc)

	// type for ignoring created_at field from response body
	type output struct {
		Age        int    `json:"age"`
		BirthDate  string `json:"birth_date"`
		CNH        string `json:"cnh"`
		CPF        string `json:"cpf"`
		Gender     string `json:"gender"`
		HasVehicle bool   `json:"has_vehicle"`
		Name       string `json:"name"`
	}

	tests := []struct {
		rw           *httptest.ResponseRecorder
		r            *http.Request
		wantCode     int
		wantBody     string
		wantLocation string
	}{
		{
			rw:           httptest.NewRecorder(),
			r:            httptest.NewRequest(http.MethodPost, target, strings.NewReader(`{"cpf":"52742089403","birth_date":"1990-01-01","cnh":"b","gender":"f","has_vehicle":false,"name":"Lorenzo Ian Carlos Eduardo Drumond"}`)),
			wantCode:     http.StatusCreated,
			wantBody:     `{"age":31,"birth_date":"1990-01-01","cnh":"B","cpf":"52742089403","gender":"F","has_vehicle":false,"name":"lorenzo ian carlos eduardo drumond"}`,
			wantLocation: fmt.Sprintf("%s/52742089403", target),
		},
	}

	for i, test := range tests {
		ctx := context.WithValue(test.r.Context(), rest.URLKey("url"), test.r.Host)
		c.ServeHTTP(test.rw, test.r.WithContext(ctx))

		if test.wantCode != test.rw.Code {
			t.Errorf("%d: wantCode=[%v] gotCode=[%v]", i, test.wantCode, test.rw.Code)
			continue
		}

		if gotLocation := test.rw.Header()["Location"]; len(gotLocation) != 1 || gotLocation[0] != test.wantLocation {
			t.Errorf("%d: wantLocation=[%v] gotFullLocation=[%v]", i, test.wantLocation, gotLocation)
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

func TestCreateDriverErrors(t *testing.T) {
	l := log.NewFakeLogger()
	repo := repository.NewDriverInMemory(samples.GetDrivers(t))
	uc := usecase.NewCreateDriver(l, repo)
	p := presenter.NewCreateDriver()
	c := rest.NewCreateDriver(p, uc)

	invalidBirthDateAge := time.Now().AddDate(-17, 0, 0).Format("2006-01-02")
	invalidBirthDate := time.Date(1900, time.Month(1), 1, 1, 1, 1, 1, time.UTC).Format("2006-01-02")

	tests := []struct {
		rw       *httptest.ResponseRecorder
		r        *http.Request
		wantCode int
		wantBody string
	}{
		// invalid JSON payload
		{
			rw:       httptest.NewRecorder(),
			r:        httptest.NewRequest(http.MethodGet, "/", nil),
			wantCode: http.StatusBadRequest,
			wantBody: `{"error":"Invalid JSON."}`,
		},
		{
			rw:       httptest.NewRecorder(),
			r:        httptest.NewRequest(http.MethodGet, "/", strings.NewReader(``)),
			wantCode: http.StatusBadRequest,
			wantBody: `{"error":"Invalid JSON."}`,
		},
		{
			rw:       httptest.NewRecorder(),
			r:        httptest.NewRequest(http.MethodGet, "/", strings.NewReader(`{"name":"someone"`)),
			wantCode: http.StatusBadRequest,
			wantBody: `{"error":"Invalid JSON."}`,
		},
		// valid JSON payload, but of different type as expected
		{
			rw:       httptest.NewRecorder(),
			r:        httptest.NewRequest(http.MethodGet, "/", strings.NewReader(`"valid string"`)),
			wantCode: http.StatusBadRequest,
			wantBody: `{"error":"Expected JSON Object."}`,
		},
		{
			rw:       httptest.NewRecorder(),
			r:        httptest.NewRequest(http.MethodGet, "/", strings.NewReader(`1`)),
			wantCode: http.StatusBadRequest,
			wantBody: `{"error":"Expected JSON Object."}`,
		},
		{
			rw:       httptest.NewRecorder(),
			r:        httptest.NewRequest(http.MethodGet, "/", strings.NewReader(`true`)),
			wantCode: http.StatusBadRequest,
			wantBody: `{"error":"Expected JSON Object."}`,
		},
		{
			rw:       httptest.NewRecorder(),
			r:        httptest.NewRequest(http.MethodGet, "/", strings.NewReader(`[]`)),
			wantCode: http.StatusBadRequest,
			wantBody: `{"error":"Expected JSON Object."}`,
		},
		// valid JSON Object, but missing fields
		{
			rw:       httptest.NewRecorder(),
			r:        httptest.NewRequest(http.MethodGet, "/", strings.NewReader(`{}`)),
			wantCode: http.StatusBadRequest,
			wantBody: `{"error":"Missing JSON fields. [birth_date: string (ISO-8601 date as 'YYYY-MM-DD')], [cnh: string], [cpf: string], [gender: string], [has_vehicle: bool], [name: string]."}`,
		},
		{
			rw:       httptest.NewRecorder(),
			r:        httptest.NewRequest(http.MethodGet, "/", strings.NewReader(`{"birth_date":"1990-01-01","name":"not real","gender":"f","cnh":"a"}`)),
			wantCode: http.StatusBadRequest,
			wantBody: `{"error":"Missing JSON fields. [cpf: string], [has_vehicle: bool]."}`,
		},
		// valid JSON with correct field(s) but different types as expected
		{
			rw:       httptest.NewRecorder(),
			r:        httptest.NewRequest(http.MethodGet, "/", strings.NewReader(`{"birth_date":true}`)),
			wantCode: http.StatusBadRequest,
			wantBody: `{"error":"Invalid value at 'birth_date'. Expected string (ISO-8601 date as 'YYYY-MM-DD'), got bool."}`,
		},
		{
			rw:       httptest.NewRecorder(),
			r:        httptest.NewRequest(http.MethodGet, "/", strings.NewReader(`{"has_vehicle":"true"}`)),
			wantCode: http.StatusBadRequest,
			wantBody: `{"error":"Invalid value at 'has_vehicle'. Expected bool, got string."}`,
		},
		// valid JSON with unexpected fields
		{
			rw:       httptest.NewRecorder(),
			r:        httptest.NewRequest(http.MethodGet, "/", strings.NewReader(`{"birth_date2":"1990-01-01","birth_date3":true}`)),
			wantCode: http.StatusBadRequest,
			wantBody: `{"error":"Unexpected JSON field: 'birth_date2'."}`,
		},
		{
			rw:       httptest.NewRecorder(),
			r:        httptest.NewRequest(http.MethodGet, "/", strings.NewReader(`{"cpf":"70286951150","birth_date":"1990-01-01","cnh":"c","gender":"o","has_vehicle":true,"name":"gustavo","name_":"gustavo"}`)),
			wantCode: http.StatusBadRequest,
			wantBody: `{"error":"Unexpected JSON field: 'name_'."}`,
		},
		// invalid age
		{
			rw: httptest.NewRecorder(),
			r: httptest.NewRequest(http.MethodGet, "/", strings.NewReader(
				fmt.Sprintf(
					`{"cpf":"69048144620","birth_date":"%s","cnh":"b","gender":"o","has_vehicle":false,"name":"Fernanda Beatriz"}`,
					invalidBirthDateAge,
				),
			)),
			wantCode: http.StatusBadRequest,
			wantBody: `{"error":"Age invalid. age=[17]"}`,
		},
		// invalid birth_date
		{
			rw: httptest.NewRecorder(),
			r: httptest.NewRequest(http.MethodGet, "/", strings.NewReader(
				fmt.Sprintf(
					`{"cpf":"69048144620","birth_date":"%s","cnh":"b","gender":"o","has_vehicle":false,"name":"Fernanda Beatriz"}`,
					invalidBirthDate,
				),
			)),
			wantCode: http.StatusBadRequest,
			wantBody: `{"error":"Birth Date invalid. birthDate=[1900-01-01 00:00:00 +0000 UTC]"}`,
		},
		// invalid CNH
		{
			rw:       httptest.NewRecorder(),
			r:        httptest.NewRequest(http.MethodGet, "/", strings.NewReader(`{"cpf":"69048144620","birth_date":"1990-01-01","cnh":"f","gender":"o","has_vehicle":false,"name":"Fernanda Beatriz"}`)),
			wantCode: http.StatusBadRequest,
			wantBody: `{"error":"CNH invalid. cnh=[f]"}`,
		},
		// invalid CPF
		{
			rw:       httptest.NewRecorder(),
			r:        httptest.NewRequest(http.MethodGet, "/", strings.NewReader(`{"cpf":"69048144621","birth_date":"1990-01-01","cnh":"b","gender":"o","has_vehicle":false,"name":"Fernanda Beatriz"}`)),
			wantCode: http.StatusBadRequest,
			wantBody: `{"error":"CPF invalid. cpf=[69048144621]"}`,
		},
		// invalid Gender
		{
			rw:       httptest.NewRecorder(),
			r:        httptest.NewRequest(http.MethodGet, "/", strings.NewReader(`{"cpf":"69048144620","birth_date":"1990-01-01","cnh":"b","gender":"9","has_vehicle":false,"name":"Fernanda Beatriz"}`)),
			wantCode: http.StatusBadRequest,
			wantBody: `{"error":"Gender invalid. gender=[9]"}`,
		},
		// invalid Name
		{
			rw:       httptest.NewRecorder(),
			r:        httptest.NewRequest(http.MethodGet, "/", strings.NewReader(`{"cpf":"69048144620","birth_date":"1990-01-01","cnh":"b","gender":"o","has_vehicle":false,"name":" Fernanda Beatriz"}`)),
			wantCode: http.StatusBadRequest,
			wantBody: `{"error":"Name invalid. name=[ Fernanda Beatriz]"}`,
		},
		// driver already exists
		{
			rw:       httptest.NewRecorder(),
			r:        httptest.NewRequest(http.MethodGet, "/", strings.NewReader(`{"cpf":"64053595061","birth_date":"1990-01-01","cnh":"b","gender":"o","has_vehicle":false,"name":"Fernanda Beatriz"}`)),
			wantCode: http.StatusConflict,
			wantBody: `{"error":"Driver already exists. cpf=[64053595061]"}`,
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
