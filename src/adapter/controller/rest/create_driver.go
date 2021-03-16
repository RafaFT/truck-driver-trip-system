package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/rafaft/truck-driver-trip-system/entity"
	"github.com/rafaft/truck-driver-trip-system/usecase"
)

type createDriverInput struct {
	BirthDate  *time.Time `json:"birth_date"`
	CNH        *string    `json:"cnh"`
	CPF        *string    `json:"cpf"`
	Gender     *string    `json:"gender"`
	HasVehicle *bool      `json:"has_vehicle"`
	Name       *string    `json:"name"`
}

func (cd createDriverInput) writeUCInput(ucInput *usecase.CreateDriverInput) error {
	missingFields := make([]string, 0, 6)

	if cd.BirthDate == nil {
		missingFields = append(missingFields, "birth_date")
	}
	if cd.CNH == nil {
		missingFields = append(missingFields, "cnh")
	}
	if cd.CPF == nil {
		missingFields = append(missingFields, "cpf")
	}
	if cd.Gender == nil {
		missingFields = append(missingFields, "gender")
	}
	if cd.HasVehicle == nil {
		missingFields = append(missingFields, "has_vehicle")
	}
	if cd.Name == nil {
		missingFields = append(missingFields, "name")
	}

	if len(missingFields) > 0 {
		return fmt.Errorf("missing fields: [%s]", strings.Join(missingFields, ", "))
	}

	ucInput.BirthDate = *cd.BirthDate
	ucInput.CNH = *cd.CNH
	ucInput.CPF = *cd.CPF
	ucInput.Gender = *cd.Gender
	ucInput.HasVehicle = *cd.HasVehicle
	ucInput.Name = *cd.Name

	return nil
}

type CreateDriverController struct {
	url string
	uc  usecase.CreateDriverUseCase
}

func NewCreateDriverController(url string, uc usecase.CreateDriverUseCase) CreateDriverController {
	return CreateDriverController{
		url: url,
		uc:  uc,
	}
}

func (c CreateDriverController) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var input createDriverInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "invalid request body"}`))
		return
	}

	var ucInput usecase.CreateDriverInput
	if err := input.writeUCInput(&ucInput); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}

	output, err := c.uc.Execute(r.Context(), ucInput)
	if err != nil {
		var code int
		var msg string

		switch err.(type) {
		case entity.ErrInvalidAge,
			entity.ErrInvalidBirthDate,
			entity.ErrInvalidCNH,
			entity.ErrInvalidCPF,
			entity.ErrInvalidGender,
			entity.ErrInvalidName:
			code = http.StatusBadRequest
			msg = err.Error()
		case entity.ErrDriverAlreadyExists:
			code = http.StatusConflict
			msg = err.Error()
		default:
			code = http.StatusInternalServerError
			msg = "internal server error"
		}

		w.WriteHeader(code)
		w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, msg)))
		return
	}

	response, _ := json.Marshal(&output)

	w.WriteHeader(http.StatusCreated)
	w.Write(response)
	w.Header().Set("location", fmt.Sprintf("%s/%s", c.url, input.CPF))
}
