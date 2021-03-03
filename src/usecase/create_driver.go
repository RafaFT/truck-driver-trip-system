package usecase

import (
	"context"
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
	BirthDate  time.Time `json:"birth_date"`
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
		return di.presenter.Output(nil), err
	}

	err = di.repo.SaveDriver(ctx, driver)
	if err != nil {
		return di.presenter.Output(nil), err
	}

	return di.presenter.Output(driver), nil
}
