package presenter

import (
	"encoding/json"

	"github.com/rafaft/truck-driver-trip-system/adapter/controller/rest"
	"github.com/rafaft/truck-driver-trip-system/usecase"
)

type getDriversOutput struct {
	BirthDate  string `json:"birth_date"`
	CNH        string `json:"cnh"`
	CPF        string `json:"cpf"`
	Gender     string `json:"gender"`
	HasVehicle bool   `json:"has_vehicle"`
	Name       string `json:"name"`
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

func (p GetDriversPresenter) Output(drivers []*usecase.GetDriversOutput) []byte {
	output := make([]*getDriversOutput, len(drivers))

	for i, driver := range drivers {
		output[i] = &getDriversOutput{
			BirthDate:  driver.BirthDate.Format("2006-01-02"),
			CNH:        driver.CNH,
			CPF:        driver.CPF,
			Gender:     driver.Gender,
			HasVehicle: driver.HasVehicle,
			Name:       driver.Name,
		}
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
