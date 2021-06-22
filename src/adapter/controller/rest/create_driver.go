package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"regexp"
	"strings"

	"github.com/rafaft/truck-driver-trip-system/entity"
	"github.com/rafaft/truck-driver-trip-system/usecase"
)

type URLKey string

var unknownJSONField = regexp.MustCompile(`\"(.*)\"$`)

// output port (out of place, according to clean architecture, this interface should be declared on usecase layer)
type CreateDriverPresenter interface {
	Output(*usecase.CreateDriverOutput) []byte
	OutputError(error) []byte
}

type createDriverInput struct {
	BirthDate  *ISO8601Date `json:"birth_date"`
	CNH        *string      `json:"cnh"`
	CPF        *string      `json:"cpf"`
	Gender     *string      `json:"gender"`
	HasVehicle *bool        `json:"has_vehicle"`
	Name       *string      `json:"name"`
}

func (cd *createDriverInput) UnmarshalJSON(b []byte) error {
	// create and use alias type to prevent infinite recursion
	type createDriverInput_ createDriverInput
	var cd_ createDriverInput_

	// Use json.Decoder instead of Unmarshal, to prevent unexpected fields
	decoder := json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&cd_)
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

	*cd = createDriverInput(cd_)

	if err := cd.hasFieldsMissing(); err != nil {
		return err
	}

	return nil
}

func (cd *createDriverInput) hasFieldsMissing() error {
	v := reflect.ValueOf(*cd)
	t := reflect.TypeOf(*cd)

	s := make([][2]string, 0, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		// This logic only works if all mandatory fields are pointer types and
		// if all fields have json tags.
		if v.Field(i).IsNil() {
			if jsonTags := strings.Split(t.Field(i).Tag.Get("json"), ","); len(jsonTags) > 0 {
				s = append(s, [2]string{
					jsonTags[0],
					t.Field(i).Type.Elem().Name(),
				})
			}
		}
	}

	if len(s) > 0 {
		return newErrMissingJSONFields(s)
	}

	return nil
}

type CreateDriverController struct {
	p  CreateDriverPresenter
	uc usecase.CreateDriver
}

func NewCreateDriver(p CreateDriverPresenter, uc usecase.CreateDriver) CreateDriverController {
	return CreateDriverController{
		p:  p,
		uc: uc,
	}
}

func (c CreateDriverController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

	var input createDriverInput
	if err := json.Unmarshal(b, &input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(c.p.OutputError(err))
		return
	}

	ucInput := usecase.CreateDriverInput{
		BirthDate:  input.BirthDate.Time,
		CNH:        *input.CNH,
		CPF:        *input.CPF,
		Gender:     *input.Gender,
		HasVehicle: *input.HasVehicle,
		Name:       *input.Name,
	}

	output, err := c.uc.Execute(r.Context(), ucInput)
	if err != nil {
		var code int

		switch err.(type) {
		case entity.ErrInvalidAge,
			entity.ErrInvalidBirthDate,
			entity.ErrInvalidCNH,
			entity.ErrInvalidCPF,
			entity.ErrInvalidGender,
			entity.ErrInvalidName:
			code = http.StatusBadRequest
		case entity.ErrDriverAlreadyExists:
			code = http.StatusConflict
		default:
			code = http.StatusInternalServerError
			err = ErrInternalServerError
		}

		w.WriteHeader(code)
		w.Write(c.p.OutputError(err))
		return
	}

	w.Header().Set("location", fmt.Sprintf("%s/%s", r.Context().Value(URLKey("url")).(string), *input.CPF))
	w.WriteHeader(http.StatusCreated)
	w.Write(c.p.Output(output))
}
