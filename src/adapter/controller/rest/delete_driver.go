package rest

import (
	"net/http"

	"github.com/rafaft/truck-driver-trip-system/entity"
	"github.com/rafaft/truck-driver-trip-system/usecase"
)

// output port (out of place, according to clean architecture, this interface should be declared on usecase layer)
type DeleteDriverPresenter interface {
	OutputError(error) []byte
}

type DeleteDriverByCPFController struct {
	p  DeleteDriverPresenter
	uc usecase.DeleteDriverUseCase
}

func NewDeleteDriverByCPFController(p DeleteDriverPresenter, uc usecase.DeleteDriverUseCase) DeleteDriverByCPFController {
	return DeleteDriverByCPFController{
		p:  p,
		uc: uc,
	}
}

func (c DeleteDriverByCPFController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cpf := r.Context().Value(CPFKey("cpf")).(string)

	err := c.uc.Execute(r.Context(), cpf)
	if err != nil {
		switch err.(type) {
		case entity.ErrDriverNotFound,
			entity.ErrInvalidCPF:
			w.WriteHeader(http.StatusNotFound)
			w.Header().Del("content-type")
		default:
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(c.p.OutputError(ErrInternalServerError))
		}

		return
	}

	w.WriteHeader(http.StatusNoContent)
}
