package rest

import (
	"errors"
	"fmt"
	"net/http"
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
	uc usecase.GetDriversUseCase
}

func NewGetDriversController(p GetDriversPresenter, uc usecase.GetDriversUseCase) GetDriversController {
	return GetDriversController{
		p:  p,
		uc: uc,
	}
}

func (c GetDriversController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var q usecase.GetDriversQuery

	formError := r.ParseForm()
	if formError == nil {
		if rawHasVehicle := r.Form.Get("has_vehicle"); rawHasVehicle != "" {
			if hasVehicle, err := strconv.ParseBool(rawHasVehicle); err == nil {
				q.HasVehicle = &hasVehicle
			} else {
				err := newErrInvalidParameterValue("has_vehicle", rawHasVehicle, reflect.TypeOf((*bool)(nil)).Elem())
				w.WriteHeader(http.StatusBadRequest)
				w.Write(c.p.OutputError(err))
				return
			}
		}
		if rawLimit := r.Form.Get("limit"); rawLimit != "" {
			if limit, err := strconv.ParseUint(rawLimit, 10, 64); err == nil {
				ulimit := uint(limit)
				q.Limit = &ulimit
			} else {
				err := newErrInvalidParameterValue("limit", rawLimit, reflect.TypeOf((*uint)(nil)).Elem())
				w.WriteHeader(http.StatusBadRequest)
				w.Write(c.p.OutputError(err))
				return
			}
		}
		if cnh := r.Form.Get("cnh"); cnh != "" {
			q.CNH = &cnh
		}
		if gender := r.Form.Get("gender"); gender != "" {
			q.Gender = &gender
		}
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
			ucErr = fmt.Errorf("internal server error")
		}

		w.WriteHeader(code)
		w.Write(c.p.OutputError(ucErr))
		return
	}

	// TODO: how to decouple controllers from presenters?
	var driverFields []string
	if formError == nil {
		driverFields = strings.Split(r.Form.Get("fields"), ",")
	}

	w.WriteHeader(http.StatusOK)
	w.Write(c.p.Output(output, driverFields...))
}
