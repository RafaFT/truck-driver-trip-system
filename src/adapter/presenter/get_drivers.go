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

// output port implementation - Presenter
type getDrivers struct {
}

func NewGetDrivers() rest.GetDriversPresenter {
	return getDrivers{}
}

func (p getDrivers) Output(drivers []*usecase.GetDriversOutput, fields ...string) []byte {
	age := fields == nil || containsField("age", fields)
	birthDate := fields == nil || containsField("birth_date", fields)
	cnh := fields == nil || containsField("cnh", fields)
	cpf := fields == nil || containsField("cpf", fields)
	gender := fields == nil || containsField("gender", fields)
	hasVehicle := fields == nil || containsField("has_vehicle", fields)
	name := fields == nil || containsField("name", fields)

	output := make([]*getDriversOutput, len(drivers))
	for i, driver := range drivers {
		var d getDriversOutput

		if age {
			d.Age = &driver.Age
		}
		if birthDate {
			birthDate := driver.BirthDate.Format("2006-01-02")
			d.BirthDate = &birthDate
		}
		if cnh {
			d.CNH = &driver.CNH
		}
		if cpf {
			d.CPF = &driver.CPF
		}
		if gender {
			d.Gender = &driver.Gender
		}
		if hasVehicle {
			d.HasVehicle = &driver.HasVehicle
		}
		if name {
			d.Name = &driver.Name
		}

		output[i] = &d
	}

	b, _ := json.Marshal(&output)

	return b
}

func (p getDrivers) OutputError(err error) []byte {
	output := getDriversOutputError{
		Error: err.Error(),
	}

	b, _ := json.Marshal(&output)

	return b
}
