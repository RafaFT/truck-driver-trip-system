package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/rafaft/truck-driver-trip-system/entity"
)

type (
	// input port
	GetDrivers interface {
		Execute(context.Context, GetDriversQuery) ([]*GetDriversOutput, error)
	}

	GetDriverRepo interface {
		Find(context.Context, FindDriversQuery) ([]*entity.Driver, error)
	}

	FindDriversQuery struct {
		CNH        *entity.CNH
		Gender     *entity.Gender
		HasVehicle *bool
		Limit      *uint
	}

	// input port implementation - Interactor
	getDrivers struct {
		logger Logger
		repo   GetDriverRepo
	}

	GetDriversQuery struct {
		CNH        *string
		Gender     *string
		HasVehicle *bool
		Limit      *uint
	}

	// output data - type
	GetDriversOutput struct {
		Age        int
		BirthDate  time.Time
		CNH        string
		CPF        string
		Gender     string
		HasVehicle bool
		Name       string
	}
)

func NewFindDriversQuery(cnh, gender *string, hasVehicle *bool, limit *uint) (FindDriversQuery, error) {
	errorMsg := "Invalid FindDriversQuery: %w"
	var q FindDriversQuery

	if cnh != nil {
		cnhT, err := entity.NewCNH(*cnh)
		if err != nil {
			return q, fmt.Errorf(errorMsg, err)
		}
		q.CNH = &cnhT
	}

	if gender != nil {
		genderT, err := entity.NewGender(*gender)
		if err != nil {
			return q, fmt.Errorf(errorMsg, err)
		}
		q.Gender = &genderT
	}

	if hasVehicle != nil {
		q.HasVehicle = hasVehicle
	}

	if limit != nil {
		q.Limit = limit
	}

	return q, nil
}

func NewGetDrivers(logger Logger, repo GetDriverRepo) GetDrivers {
	return getDrivers{
		logger: logger,
		repo:   repo,
	}
}

func (di getDrivers) Execute(ctx context.Context, rawQ GetDriversQuery) ([]*GetDriversOutput, error) {
	q, err := NewFindDriversQuery(rawQ.CNH, rawQ.Gender, rawQ.HasVehicle, rawQ.Limit)
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
