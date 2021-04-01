package presenter

import (
	"encoding/json"
	"time"

	"github.com/rafaft/truck-driver-trip-system/adapter/controller/rest"
	"github.com/rafaft/truck-driver-trip-system/usecase"
)

type createDriverOutput struct {
	Age        int       `json:"age"`
	BirthDate  string    `json:"birth_date"`
	CNH        string    `json:"cnh"`
	CPF        string    `json:"cpf"`
	CreatedAt  time.Time `json:"created_at"`
	Gender     string    `json:"gender"`
	HasVehicle bool      `json:"has_vehicle"`
	Name       string    `json:"name"`
}

type createDriverOutputError struct {
	Error string `json:"error"`
}

// output port (presenter) implementation
type CreateDriverPresenter struct {
}

func NewCreateDriverPresenter() rest.CreateDriverPresenter {
	return CreateDriverPresenter{}
}

func (p CreateDriverPresenter) Output(driver *usecase.CreateDriverOutput) []byte {
	var output createDriverOutput

	output.Age = driver.Age
	output.BirthDate = driver.BirthDate.Format("2006-01-02")
	output.CNH = driver.CNH
	output.CPF = driver.CPF
	// This field (createdAt) should probably come from the repository response.
	// But for that, the repository methods should return DTO's, which
	// is an extra layer I'm not sure it's worth doing:
	// https://softwareengineering.stackexchange.com/questions/376447/what-data-should-a-repository-return
	output.CreatedAt = driver.CreatedAt
	output.Gender = driver.Gender
	output.HasVehicle = driver.HasVehicle
	output.Name = driver.Name

	b, _ := json.Marshal(&output)

	return b
}

func (p CreateDriverPresenter) OutputError(err error) []byte {
	output := createDriverOutputError{
		Error: err.Error(),
	}

	b, _ := json.Marshal(&output)

	return b
}
