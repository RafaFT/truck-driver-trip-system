package rest

import (
	"fmt"
	"net/http"

	"github.com/rafaft/truck-driver-trip-system/usecase"
)

// output port (out of place, according to clean architecture, this interface should be declared on usecase layer)
type GetDriversPresenter interface {
	Output([]*usecase.GetDriversOutput) []byte
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
	output, err := c.uc.Execute(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(c.p.OutputError(fmt.Errorf("internal server error")))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(c.p.Output(output))
}
