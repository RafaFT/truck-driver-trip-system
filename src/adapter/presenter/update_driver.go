package presenter

import (
	"encoding/json"
	"time"

	"github.com/rafaft/truck-driver-trip-system/adapter/controller/rest"
	"github.com/rafaft/truck-driver-trip-system/usecase"
)

type updateDriverOutput struct {
	BirthDate  string    `json:"birthDate"`
	CNH        string    `json:"cnh"`
	CPF        string    `json:"cpf"`
	Gender     string    `json:"gender"`
	HasVehicle bool      `json:"has_vehicle"`
	Name       string    `json:"name"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type updateDriverOutputError struct {
	Error string `json:"error"`
}

// output port (presenter) implementation
type UpdateDriverPresenter struct {
}

func NewUpdateDriverPresenter() rest.UpdateDriverPresenter {
	return UpdateDriverPresenter{}
}

func (p UpdateDriverPresenter) Output(driver *usecase.UpdateDriverOutput) []byte {
	var output updateDriverOutput

	output.BirthDate = driver.BirthDate.Format("2006-01-02")
	output.CNH = driver.CNH
	output.CPF = driver.CPF
	output.Gender = driver.Gender
	output.HasVehicle = driver.HasVehicle
	output.Name = driver.Name
	output.UpdatedAt = driver.UpdatedAt

	b, _ := json.Marshal(&output)

	return b
}

func (p UpdateDriverPresenter) OutputError(err error) []byte {
	output := updateDriverOutputError{
		Error: err.Error(),
	}

	b, _ := json.Marshal(&output)

	return b
}
