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
	CreateDriver interface {
		Execute(context.Context, CreateDriverInput) (*CreateDriverOutput, error)
	}

	CreateDriverRepo interface {
		Save(context.Context, *entity.Driver) error
	}

	// input port implementation - Interactor
	createDriver struct {
		logger Logger
		repo   CreateDriverRepo
	}

	CreateDriverInput struct {
		BirthDate  time.Time
		CNH        string
		CPF        string
		Gender     string
		HasVehicle bool
		Name       string
	}

	CreateDriverOutput struct {
		Age        int
		BirthDate  time.Time
		CNH        string
		CPF        string
		CreatedAt  time.Time
		Gender     string
		HasVehicle bool
		Name       string
	}
)

func NewCreateDriver(logger Logger, repo CreateDriverRepo) CreateDriver {
	return createDriver{
		logger: logger,
		repo:   repo,
	}
}

func (di createDriver) Execute(ctx context.Context, input CreateDriverInput) (*CreateDriverOutput, error) {
	driver, err := entity.NewDriver(
		input.CPF,
		input.Name,
		input.Gender,
		input.CNH,
		input.BirthDate,
		input.HasVehicle,
	)

	if err != nil {
		di.logger.Debug(err.Error())
		return nil, err
	}

	err = di.repo.Save(ctx, driver)
	if err != nil {
		switch err.(type) {
		case entity.ErrDriverAlreadyExists:
			di.logger.Debug(err.Error())
		default:
			di.logger.Error(err.Error())
		}
		return nil, err
	}

	logInput, _ := json.MarshalIndent(input, "", "\t")
	di.logger.Info(fmt.Sprintf("new driver created. input=[%s]", string(logInput)))

	result := CreateDriverOutput{
		Age:        driver.Age(),
		BirthDate:  driver.BirthDate().Time,
		CNH:        string(driver.CNH()),
		CPF:        string(driver.CPF()),
		CreatedAt:  time.Now(),
		Gender:     string(driver.Gender()),
		HasVehicle: driver.HasVehicle(),
		Name:       string(driver.Name()),
	}

	return &result, nil
}
