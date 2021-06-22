package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

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

func (ud *updateDriverInput) UnmarshalJSON(b []byte) error {
	// create and use alias type to prevent infinite recursion
	type updateDriverInput_ updateDriverInput
	var ud_ updateDriverInput_

	// Use json.Decoder instead of Unmarshal, to prevent unexpected fields
	decoder := json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&ud_)
	if err != nil {
		if jsonTypeErr, ok := err.(*json.UnmarshalTypeError); ok {
			if jsonTypeErr.Field == "" && jsonTypeErr.Struct == "" {
				return ErrExpectedJSONObject
			}

			return newErrInvalidJSONFieldType(jsonTypeErr.Field, jsonTypeErr.Type.Name(), jsonTypeErr.Value)
		}

		if strings.HasPrefix(err.Error(), "json: unknown field") {
			if match := unknownJSONField.FindStringSubmatch(err.Error()); match != nil {
				return newErrUnexpectedJSONField(match[1])
			}
		}

		return err
	}

	*ud = updateDriverInput(ud_)

	return nil
}

type UpdateDriverController struct {
	p  UpdateDriverPresenter
	uc usecase.UpdateDriver
}

func NewUpdateDriver(p UpdateDriverPresenter, uc usecase.UpdateDriver) UpdateDriverController {
	return UpdateDriverController{
		p:  p,
		uc: uc,
	}
}

func (c UpdateDriverController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(c.p.OutputError(ErrInvalidBody))
		return
	}

	if !json.Valid(b) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(c.p.OutputError(ErrInvalidJSON))
		return
	}

	var input updateDriverInput
	if err := json.Unmarshal(b, &input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(c.p.OutputError(err))
		return
	}

	if input.CNH == nil && input.Gender == nil && input.HasVehicle == nil && input.Name == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(c.p.OutputError(fmt.Errorf("Empty update.")))
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
		case entity.ErrInvalidCNH,
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
