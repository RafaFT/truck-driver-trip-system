package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rafaft/truck-driver-trip-system/entity"
	"github.com/rafaft/truck-driver-trip-system/usecase"
)

// output port (out of place, according to clean architecture, this interface should be declared on usecase layer)
type UpdateDriverPresenter interface {
	Output(*usecase.UpdateDriverOutput) []byte
	OutputError(error) []byte
}

type updateDriverInput struct {
	CNH        *string `json:"cnh"`
	Gender     *string `json:"gender"`
	HasVehicle *bool   `json:"has_vehicle"`
	Name       *string `json:"name"`
}

type UpdateDriverController struct {
	p  UpdateDriverPresenter
	uc usecase.UpdateDriverUseCase
}

func NewUpdateDriverController(p UpdateDriverPresenter, uc usecase.UpdateDriverUseCase) UpdateDriverController {
	return UpdateDriverController{
		p:  p,
		uc: uc,
	}
}

func (c UpdateDriverController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var input updateDriverInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(c.p.OutputError(err))
		return
	}

	if input.CNH == nil && input.Gender == nil && input.HasVehicle == nil && input.Name == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(c.p.OutputError(fmt.Errorf("update needs at least one valid field")))
		return
	}

	cpf := r.Context().Value(CPFKey("cpf")).(string)
	ucInput := usecase.UpdateDriverInput{
		CNH:        input.CNH,
		Gender:     input.Gender,
		HasVehicle: input.HasVehicle,
		Name:       input.Name,
	}

	output, err := c.uc.Execute(r.Context(), cpf, ucInput)
	if err != nil {
		switch err.(type) {
		case entity.ErrDriverNotFound,
			entity.ErrInvalidCPF:
			w.WriteHeader(http.StatusNotFound)
			w.Header().Del("content-type")
		case entity.ErrInvalidBirthDate,
			entity.ErrInvalidCNH,
			entity.ErrInvalidGender,
			entity.ErrInvalidName:
			w.WriteHeader(http.StatusBadRequest)
			w.Write(c.p.OutputError(err))
		default:
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(c.p.OutputError(ErrInternalServerError))
		}

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(c.p.Output(output))
}
