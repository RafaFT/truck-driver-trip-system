package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/rafaft/truck-driver-trip-system/entity"
	"github.com/rafaft/truck-driver-trip-system/usecase"
)

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

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(`{"error": "invalid request body"}`))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var input usecase.CreateDriverInput
	err = json.Unmarshal(body, &input)
	if err != nil {
		w.Write([]byte(`{"error": "invalid request body"}`))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	output, err := c.uc.Execute(r.Context(), input)
	if err != nil {
		var code int
		var msg string

		switch err.(type) {
		case entity.ErrInvalidAge:
		case entity.ErrInvalidBirthDate:
		case entity.ErrInvalidCNH:
		case entity.ErrInvalidCPF:
		case entity.ErrInvalidGender:
		case entity.ErrInvalidName:
			code = http.StatusBadRequest
			msg = err.Error()
		case entity.ErrDriverAlreadyExists:
			code = http.StatusConflict
			msg = err.Error()
		default:
			code = http.StatusInternalServerError
			msg = "internal server error"
		}

		w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, msg)))
		w.WriteHeader(code)
		return
	}

	response, _ := json.Marshal(&output)

	w.Write(response)
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("location", fmt.Sprintf("%s/%s", c.url, input.CPF))
}
