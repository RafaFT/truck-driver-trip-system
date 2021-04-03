package rest

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

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
		q.CNH = r.Form.Get("cnh")
		q.Gender = r.Form.Get("gender")
		if hasVehicle, err := strconv.ParseBool(r.Form.Get("has_vehicle")); err == nil {
			q.HasVehicle = &hasVehicle
		}
		if limit, err := strconv.ParseUint(r.Form.Get("limit"), 10, 64); err == nil {
			ulimit := uint(limit)
			q.Limit = &ulimit
		}
	}

	output, err := c.uc.Execute(r.Context(), q)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(c.p.OutputError(fmt.Errorf("internal server error")))
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
