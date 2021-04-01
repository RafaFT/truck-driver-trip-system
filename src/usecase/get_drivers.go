package usecase

import (
	"context"
	"time"

	"github.com/rafaft/truck-driver-trip-system/entity"
)

// input port - interface
type GetDriversUseCase interface {
	Execute(context.Context) ([]*GetDriversOutput, error)
}

// input port implementation - interactor
type GetDriversInteractor struct {
	logger Logger
	repo   entity.DriverRepository
}

// output data - type
type GetDriversOutput struct {
	Age        int
	BirthDate  time.Time
	CNH        string
	CPF        string
	Gender     string
	HasVehicle bool
	Name       string
}

func NewGetDriversInteractor(logger Logger, repo entity.DriverRepository) GetDriversUseCase {
	return GetDriversInteractor{
		logger: logger,
		repo:   repo,
	}
}

func (di GetDriversInteractor) Execute(ctx context.Context) ([]*GetDriversOutput, error) {
	drivers, err := di.repo.FindDrivers(ctx)
	if err != nil {
		di.logger.Error(err.Error())
		return nil, err
	}

	output := make([]*GetDriversOutput, len(drivers))
	for i, driver := range drivers {
		output[i] = &GetDriversOutput{
			Age:        driver.Age(),
			BirthDate:  driver.BirthDate().Time,
			CNH:        string(driver.CNHType()),
			CPF:        string(driver.CPF()),
			Gender:     string(driver.Gender()),
			HasVehicle: driver.HasVehicle(),
			Name:       string(driver.Name()),
		}
	}

	return output, nil
}
