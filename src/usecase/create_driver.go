package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/rafaft/truck-driver-trip-system/entity"
)

// input port
type CreateDriver interface {
	Execute(context.Context, CreateDriverInput) (*CreateDriverOutput, error)
}

// input port implementation - Interactor
type createDriver struct {
	logger Logger
	repo   entity.DriverRepository
}

type CreateDriverInput struct {
	BirthDate  time.Time
	CNH        string
	CPF        string
	Gender     string
	HasVehicle bool
	Name       string
}

type CreateDriverOutput struct {
	Age        int
	BirthDate  time.Time
	CNH        string
	CPF        string
	CreatedAt  time.Time
	Gender     string
	HasVehicle bool
	Name       string
}

func NewCreateDriver(logger Logger, repo entity.DriverRepository) CreateDriver {
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

	err = di.repo.SaveDriver(ctx, driver)
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
