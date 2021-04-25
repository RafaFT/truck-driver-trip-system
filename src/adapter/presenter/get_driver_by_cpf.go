package presenter

import (
	"encoding/json"

	"github.com/rafaft/truck-driver-trip-system/adapter/controller/rest"
	"github.com/rafaft/truck-driver-trip-system/usecase"
)

type getDriverByCPFOutput struct {
	Age        *int    `json:"age,omitempty"`
	BirthDate  *string `json:"birth_date,omitempty"`
	CNH        *string `json:"cnh,omitempty"`
	CPF        *string `json:"cpf,omitempty"`
	Gender     *string `json:"gender,omitempty"`
	HasVehicle *bool   `json:"has_vehicle,omitempty"`
	Name       *string `json:"name,omitempty"`
}

type getDriverByCPFOutputError struct {
	Error string `json:"error"`
}

// output port implementation - Presenter
type getDriverByCPF struct {
}

func NewGetDriverByCPF() rest.GetDriverByCPFPresenter {
	return getDriverByCPF{}
}

func (p getDriverByCPF) Output(driver *usecase.GetDriverByCPFOutput, fields ...string) []byte {
	var output getDriverByCPFOutput

	if containsField("age", fields) {
		output.Age = &driver.Age
	}
	if containsField("age", fields) {
		birthDate := driver.BirthDate.Format("2006-01-02")
		output.BirthDate = &birthDate
	}
	if containsField("cnh", fields) {
		output.CNH = &driver.CNH
	}
	if containsField("cpf", fields) {
		output.CPF = &driver.CPF
	}
	if containsField("gender", fields) {
		output.Gender = &driver.Gender
	}
	if containsField("has_vehicle", fields) {
		output.HasVehicle = &driver.HasVehicle
	}
	if containsField("cpf", fields) {
		output.Name = &driver.Name
	}

	b, _ := json.Marshal(&output)

	return b
}

func (p getDriverByCPF) OutputError(err error) []byte {
	output := getDriverByCPFOutputError{
		Error: err.Error(),
	}

	b, _ := json.Marshal(&output)

	return b
}

func containsField(field string, fields []string) bool {
	for _, v := range fields {
		if v == field {
			return true
		}
	}
	return false
}
