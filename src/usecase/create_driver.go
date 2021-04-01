package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/rafaft/truck-driver-trip-system/entity"
)

// input port
type CreateDriverUseCase interface {
	Execute(context.Context, CreateDriverInput) (*CreateDriverOutput, error)
}

// input port implementation
type CreateDriverInteractor struct {
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
	BirthDate  time.Time
	CNH        string
	CPF        string
	CreatedAt  time.Time
	Gender     string
	HasVehicle bool
	Name       string
}

func NewCreateDriverInteractor(logger Logger, repo entity.DriverRepository) CreateDriverUseCase {
	return CreateDriverInteractor{
		logger: logger,
		repo:   repo,
	}
}

func (di CreateDriverInteractor) Execute(ctx context.Context, input CreateDriverInput) (*CreateDriverOutput, error) {
	driver, err := entity.NewTruckDriver(
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
		BirthDate:  driver.BirthDate().Time,
		CNH:        string(driver.CNHType()),
		CPF:        string(driver.CPF()),
		CreatedAt:  time.Now(),
		Gender:     string(driver.Gender()),
		HasVehicle: driver.HasVehicle(),
		Name:       string(driver.Name()),
	}

	return &result, nil
}
