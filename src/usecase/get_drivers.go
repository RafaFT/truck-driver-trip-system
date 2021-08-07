package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/rafaft/truck-driver-trip-system/entity"
)

// input port
type GetDrivers interface {
	Execute(context.Context, GetDriversQuery) ([]*GetDriversOutput, error)
}

// input port implementation - Interactor
type getDrivers struct {
	logger Logger
	repo   entity.DriverRepository
}

type GetDriversQuery struct {
	CNH        *string
	Gender     *string
	HasVehicle *bool
	Limit      *uint
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

func NewGetDrivers(logger Logger, repo entity.DriverRepository) GetDrivers {
	return getDrivers{
		logger: logger,
		repo:   repo,
	}
}

func (di getDrivers) Execute(ctx context.Context, rawQ GetDriversQuery) ([]*GetDriversOutput, error) {
	q, err := entity.NewFindDriversQuery(rawQ.CNH, rawQ.Gender, rawQ.HasVehicle, rawQ.Limit)
	if err != nil {
		di.logger.Debug(err.Error())
		return nil, err
	}

	logQ, _ := json.MarshalIndent(q, "", "\t")
	di.logger.Debug(fmt.Sprintf("FindDriversQuery: %s", string(logQ)))

	drivers, err := di.repo.Find(ctx, q)
	if err != nil {
		di.logger.Error(err.Error())
		return nil, err
	}

	output := make([]*GetDriversOutput, len(drivers))
	for i, driver := range drivers {
		output[i] = &GetDriversOutput{
			Age:        driver.Age(),
			BirthDate:  driver.BirthDate().Time,
			CNH:        string(driver.CNH()),
			CPF:        string(driver.CPF()),
			Gender:     string(driver.Gender()),
			HasVehicle: driver.HasVehicle(),
			Name:       string(driver.Name()),
		}
	}

	return output, nil
}
