package presenter

import (
	"encoding/json"

	"github.com/rafaft/truck-driver-trip-system/adapter/controller/rest"
	"github.com/rafaft/truck-driver-trip-system/usecase"
)

type getDriverByCPFOutput struct {
	Age        int    `json:"age"`
	BirthDate  string `json:"birth_date"`
	CNH        string `json:"cnh"`
	CPF        string `json:"cpf"`
	Gender     string `json:"gender"`
	HasVehicle bool   `json:"has_vehicle"`
	Name       string `json:"name"`
}

type getDriverByCPFOutputError struct {
	Error string `json:"error"`
}

// output port (presenter) implementation
type getDriverByCPFPresenter struct {
}

func NewGetDriverByCPFPresenter() rest.GetDriverByCPFPresenter {
	return getDriverByCPFPresenter{}
}

func (p getDriverByCPFPresenter) Output(driver *usecase.GetDriverByCPFOutput) []byte {
	var output getDriverByCPFOutput

	output.Age = driver.Age
	output.BirthDate = driver.BirthDate.Format("2006-01-02")
	output.CNH = driver.CNH
	output.CPF = driver.CPF
	output.Gender = driver.Gender
	output.HasVehicle = driver.HasVehicle
	output.Name = driver.Name

	b, _ := json.Marshal(&output)

	return b
}

func (p getDriverByCPFPresenter) OutputError(err error) []byte {
	output := getDriverByCPFOutputError{
		Error: err.Error(),
	}

	b, _ := json.Marshal(&output)

	return b
}
