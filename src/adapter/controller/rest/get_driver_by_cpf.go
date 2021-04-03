package rest

import (
	"fmt"
	"net/http"
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
	uc usecase.GetDriverByCPFUseCase
}

func NewGetDriverByCPFController(p GetDriverByCPFPresenter, uc usecase.GetDriverByCPFUseCase) GetDriverByCPFController {
	return GetDriverByCPFController{
		p:  p,
		uc: uc,
	}
}

func (c GetDriverByCPFController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cpf := r.Context().Value(CPFKey("cpf")).(string)

	output, err := c.uc.Execute(r.Context(), cpf)
	if err != nil {
		switch err.(type) {
		case entity.ErrDriverNotFound,
			entity.ErrInvalidCPF:
			w.WriteHeader(http.StatusNotFound)
		default:
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(c.p.OutputError(fmt.Errorf("internal server error")))
		}

		return
	}

	var driverFields []string
	err = r.ParseForm()
	if err == nil {
		driverFields = strings.Split(r.Form.Get("fields"), ",")
	}

	w.WriteHeader(http.StatusOK)
	w.Write(c.p.Output(output, driverFields...))
}
