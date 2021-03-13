package rest

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
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "invalid request body"}`))
		return
	}

	var input usecase.CreateDriverInput
	err = json.Unmarshal(body, &input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "invalid request body"}`))
		return
	}

	output, err := c.uc.Execute(r.Context(), input)
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
