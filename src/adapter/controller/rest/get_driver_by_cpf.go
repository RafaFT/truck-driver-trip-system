package rest

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/rafaft/truck-driver-trip-system/entity"
	"github.com/rafaft/truck-driver-trip-system/usecase"
)

type CPFKey string

// output port (out of place, according to clean architecture, this interface should be declared on usecase layer)
type GetDriverByCPFPresenter interface {
	Output(*usecase.GetDriverByCPFOutput, ...string) []byte
	OutputError(error) []byte
}

type GetDriverByCPFController struct {
	p  GetDriverByCPFPresenter
	uc usecase.GetDriverByCPF
}

var getDriverByCPFParameters = map[string]struct{}{
	"fields": {},
}

func NewGetDriverByCPF(p GetDriverByCPFPresenter, uc usecase.GetDriverByCPF) GetDriverByCPFController {
	return GetDriverByCPFController{
		p:  p,
		uc: uc,
	}
}

func (c GetDriverByCPFController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	params, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(c.p.OutputError(ErrInvalidQueryString))
		return
	}

	for p := range params {
		if _, ok := getDriverByCPFParameters[p]; !ok {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(c.p.OutputError(newErrUnexpectedParameter(p)))
			return
		}
	}

	cpf := r.Context().Value(CPFKey("cpf")).(string)

	output, err := c.uc.Execute(r.Context(), cpf)
	if err != nil {
		switch err.(type) {
		case entity.ErrDriverNotFound,
			entity.ErrInvalidCPF:
			w.WriteHeader(http.StatusNotFound)
		default:
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(c.p.OutputError(ErrInternalServerError))
		}

		return
	}

	// TODO: how to decouple controllers from presenters?
	var driverFields []string
	err = r.ParseForm()
	if err == nil {
		if fields := r.Form.Get("fields"); fields != "" {
			driverFields = strings.Split(fields, ",")
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Write(c.p.Output(output, driverFields...))
}
