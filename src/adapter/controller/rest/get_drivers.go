package rest

import (
	"errors"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"

	"github.com/rafaft/truck-driver-trip-system/entity"
	"github.com/rafaft/truck-driver-trip-system/usecase"
)

// output port (out of place, according to clean architecture, this interface should be declared on usecase layer)
type GetDriversPresenter interface {
	Output([]*usecase.GetDriversOutput, ...string) []byte
	OutputError(error) []byte
}

type GetDriversController struct {
	p  GetDriversPresenter
	uc usecase.GetDrivers
}

var getDriversParameters = map[string]struct{}{
	"cnh":         {},
	"fields":      {},
	"gender":      {},
	"has_vehicle": {},
	"limit":       {},
}

func NewGetDriversController(p GetDriversPresenter, uc usecase.GetDrivers) GetDriversController {
	return GetDriversController{
		p:  p,
		uc: uc,
	}
}

func (c GetDriversController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	params, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(c.p.OutputError(newErrParseQueryString(r.URL.RawQuery)))
		return
	}

	for p := range params {
		if _, ok := getDriversParameters[p]; !ok {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(c.p.OutputError(newErrUnknownParameter(p)))
			return
		}
	}

	q, err := getGetDriversQuery(params)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(c.p.OutputError(err))
		return
	}

	output, err := c.uc.Execute(r.Context(), q)
	if err != nil {
		var ucErr error
		code := http.StatusBadRequest

		switch errors.Unwrap(err).(type) {
		case entity.ErrInvalidCNH:
			ucErr = newErrInvalidParameterValue("cnh", r.Form.Get("cnh"), reflect.TypeOf((*entity.CNH)(nil)).Elem())
		case entity.ErrInvalidGender:
			ucErr = newErrInvalidParameterValue("gender", r.Form.Get("gender"), reflect.TypeOf((*entity.Gender)(nil)).Elem())
		default:
			code = http.StatusInternalServerError
			ucErr = ErrInternalServerError
		}

		w.WriteHeader(code)
		w.Write(c.p.OutputError(ucErr))
		return
	}

	// TODO: how to decouple controllers from presenters?
	driverFields := strings.Split(params.Get("fields"), ",")

	w.WriteHeader(http.StatusOK)
	w.Write(c.p.Output(output, driverFields...))
}

func getGetDriversQuery(v url.Values) (usecase.GetDriversQuery, error) {
	var q usecase.GetDriversQuery

	if rawHasVehicle := v.Get("has_vehicle"); rawHasVehicle != "" {
		if hasVehicle, err := strconv.ParseBool(rawHasVehicle); err == nil {
			q.HasVehicle = &hasVehicle
		} else {
			return q, newErrInvalidParameterValue("has_vehicle", rawHasVehicle, reflect.TypeOf((*bool)(nil)).Elem())
		}
	}

	if rawLimit := v.Get("limit"); rawLimit != "" {
		if limit, err := strconv.ParseUint(rawLimit, 10, 64); err == nil {
			ulimit := uint(limit)
			q.Limit = &ulimit
		} else {
			return q, newErrInvalidParameterValue("limit", rawLimit, reflect.TypeOf((*uint)(nil)).Elem())
		}
	}

	if cnh := v.Get("cnh"); cnh != "" {
		q.CNH = &cnh
	}

	if gender := v.Get("gender"); gender != "" {
		q.Gender = &gender
	}

	return q, nil
}
