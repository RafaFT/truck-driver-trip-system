package presenter

import (
	"encoding/json"

	"github.com/rafaft/truck-driver-trip-system/adapter/controller/rest"
	"github.com/rafaft/truck-driver-trip-system/usecase"
)

type getDriversOutput struct {
	Age        *int    `json:"age,omitempty"`
	BirthDate  *string `json:"birth_date,omitempty"`
	CNH        *string `json:"cnh,omitempty"`
	CPF        *string `json:"cpf,omitempty"`
	Gender     *string `json:"gender,omitempty"`
	HasVehicle *bool   `json:"has_vehicle,omitempty"`
	Name       *string `json:"name,omitempty"`
}

type getDriversOutputError struct {
	Error string `json:"error"`
}

// output port (presenter) implementation
type GetDriversPresenter struct {
}

func NewGetDriversPresenter() rest.GetDriversPresenter {
	return GetDriversPresenter{}
}

func (p GetDriversPresenter) Output(drivers []*usecase.GetDriversOutput, fields ...string) []byte {
	fieldsMap := make(map[string]bool)
	for _, field := range fields {
		if len(field) != 0 {
			fieldsMap[field] = true
		}
	}
	output := make([]*getDriversOutput, len(drivers))

	for i, driver := range drivers {
		var d getDriversOutput

		if _, ok := fieldsMap["age"]; len(fieldsMap) == 0 || ok {
			d.Age = &driver.Age
		}
		if _, ok := fieldsMap["birth_date"]; len(fieldsMap) == 0 || ok {
			birthDate := driver.BirthDate.Format("2006-01-02")
			d.BirthDate = &birthDate
		}
		if _, ok := fieldsMap["cnh"]; len(fieldsMap) == 0 || ok {
			d.CNH = &driver.CNH
		}
		if _, ok := fieldsMap["cpf"]; len(fieldsMap) == 0 || ok {
			d.CPF = &driver.CPF
		}
		if _, ok := fieldsMap["gender"]; len(fieldsMap) == 0 || ok {
			d.Gender = &driver.Gender
		}
		if _, ok := fieldsMap["has_vehicle"]; len(fieldsMap) == 0 || ok {
			d.HasVehicle = &driver.HasVehicle
		}
		if _, ok := fieldsMap["name"]; len(fieldsMap) == 0 || ok {
			d.Name = &driver.Name
		}

		output[i] = &d
	}

	b, _ := json.Marshal(&output)

	return b
}

func (p GetDriversPresenter) OutputError(err error) []byte {
	output := getDriversOutputError{
		Error: err.Error(),
	}

	b, _ := json.Marshal(&output)

	return b
}
