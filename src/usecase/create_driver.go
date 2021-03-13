package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/rafaft/truck-driver-trip-system/entity"
)

// input port
type CreateDriverUseCase interface {
	Execute(context.Context, CreateDriverInput) (CreateDriverOutput, error)
}

// input port implementation
type CreateDriverInteractor struct {
	logger    Logger
	presenter CreateDriverPresenter
	repo      entity.DriverRepository
}

// output port
type CreateDriverPresenter interface {
	Output(*entity.Driver) CreateDriverOutput
}

type CreateDriverInput struct {
	BirthDate  time.Time `json:"birth_date"`
	CNH        string    `json:"cnh"`
	CPF        string    `json:"cpf"`
	Gender     string    `json:"gender"`
	HasVehicle bool      `json:"has_vehicle"`
	Name       string    `json:"name"`
}

type CreateDriverOutput struct {
	BirthDate  string    `json:"birth_date"`
	CNH        string    `json:"cnh"`
	CPF        string    `json:"cpf"`
	CreatedAt  time.Time `json:"created_at"`
	Gender     string    `json:"gender"`
	HasVehicle bool      `json:"has_vehicle"`
	Name       string    `json:"name"`
}

func NewCreateDriverInteractor(logger Logger, presenter CreateDriverPresenter, repo entity.DriverRepository) CreateDriverUseCase {
	return CreateDriverInteractor{
		logger:    logger,
		presenter: presenter,
		repo:      repo,
	}
}

func (di CreateDriverInteractor) Execute(ctx context.Context, input CreateDriverInput) (CreateDriverOutput, error) {
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
		return di.presenter.Output(nil), err
	}

	err = di.repo.SaveDriver(ctx, driver)
	if err != nil {
		switch err.(type) {
		case entity.ErrDriverAlreadyExists:
			di.logger.Debug(err.Error())
		default:
			di.logger.Error(err.Error())
		}
		return di.presenter.Output(nil), err
	}

	di.logger.Info(fmt.Sprintf("new driver created. driver=[%v]", &driver))

	return di.presenter.Output(driver), nil
}
