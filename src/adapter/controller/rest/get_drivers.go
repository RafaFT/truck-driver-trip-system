package rest

import (
	"errors"
	"net/http"
	"net/url"
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

func NewGetDrivers(p GetDriversPresenter, uc usecase.GetDrivers) GetDriversController {
	return GetDriversController{
		p:  p,
		uc: uc,
	}
}

func (c GetDriversController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	params, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(c.p.OutputError(ErrInvalidQueryString))
		return
	}

	for p := range params {
		if _, ok := getDriversParameters[p]; !ok {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(c.p.OutputError(newErrUnexpectedParameter(p)))
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
			ucErr = newErrInvalidParameterType("cnh", "CNH")
		case entity.ErrInvalidGender:
			ucErr = newErrInvalidParameterType("gender", "Gender")
		default:
			code = http.StatusInternalServerError
			ucErr = ErrInternalServerError
		}

		w.WriteHeader(code)
		w.Write(c.p.OutputError(ucErr))
		return
	}

	// TODO: how to decouple controllers from presenters?
	var driverFields []string
	if fields := params.Get("fields"); fields != "" {
		driverFields = strings.Split(fields, ",")
	}

	w.WriteHeader(http.StatusOK)
	w.Write(c.p.Output(output, driverFields...))
}

func getGetDriversQuery(v url.Values) (usecase.GetDriversQuery, error) {
	var q usecase.GetDriversQuery

	if rawHasVehicle := v.Get("has_vehicle"); rawHasVehicle != "" {
		if hasVehicle, err := strconv.ParseBool(rawHasVehicle); err == nil {
			q.HasVehicle = &hasVehicle
		} else {
			return q, newErrInvalidParameterType("has_vehicle", "bool")
		}
	}

	if rawLimit := v.Get("limit"); rawLimit != "" {
		if limit, err := strconv.ParseUint(rawLimit, 10, 64); err == nil {
			ulimit := uint(limit)
			q.Limit = &ulimit
		} else {
			return q, newErrInvalidParameterType("limit", "uint")
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
