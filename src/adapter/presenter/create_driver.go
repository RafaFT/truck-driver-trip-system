package presenter

import (
	"time"

	"github.com/rafaft/truck-driver-trip-system/entity"
	"github.com/rafaft/truck-driver-trip-system/usecase"
)

// output port (presenter) implementation
type CreateDriverPresenter struct {
}

func NewCreateDriverPresenter() usecase.CreateDriverPresenter {
	return CreateDriverPresenter{}
}

func (p CreateDriverPresenter) Output(driver *entity.Driver) usecase.CreateDriverOutput {
	var output usecase.CreateDriverOutput

	if driver == nil {
		return output
	}

	output.BirthDate = driver.BirthDate().Format("2006-01-02")
	output.CNH = string(driver.CNHType())
	output.CPF = string(driver.CPF())
	// This field (createdAt) should probably come from the repository response.
	// But for that, the repository methods should return DTO's, which
	// is an extra layer I'm not sure it's worth doing:
	// https://softwareengineering.stackexchange.com/questions/376447/what-data-should-a-repository-return
	output.CreatedAt = time.Now()
	output.Gender = string(driver.Gender())
	output.HasVehicle = driver.HasVehicle()
	output.Name = string(driver.Name())

	return output
}
