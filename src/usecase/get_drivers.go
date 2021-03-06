package usecase

import (
	"context"

	"github.com/rafaft/truck-driver-trip-system/entity"
)

// input port - interface
type GetDriversUseCase interface {
	Execute(context.Context) ([]*GetDriversOutput, error)
}

// input port implementation - interactor
type GetDriversInteractor struct {
	logger    Logger
	presenter GetDriversPresenter
	repo      entity.DriverRepository
}

// output port - (presenter) interface
type GetDriversPresenter interface {
	Output([]*entity.Driver) []*GetDriversOutput
}

// output data - type
type GetDriversOutput struct {
	BirthDate  *string `json:"birth_date,omitempty"`
	CNH        *string `json:"cnh,omitempty"`
	CPF        *string `json:"cpf,omitempty"`
	Gender     *string `json:"gender,omitempty"`
	HasVehicle *bool   `json:"has_vehicle,omitempty"`
	Name       *string `json:"name,omitempty"`
}

func NewGetDriversInteractor(logger Logger, presenter GetDriversPresenter, repo entity.DriverRepository) GetDriversUseCase {
	return GetDriversInteractor{
		logger:    logger,
		presenter: presenter,
		repo:      repo,
	}
}

func (di GetDriversInteractor) Execute(ctx context.Context) ([]*GetDriversOutput, error) {
	drivers, err := di.repo.FindDrivers(ctx)
	if err != nil {
		di.logger.Error(err.Error())
		return di.presenter.Output(nil), err
	}

	return di.presenter.Output(drivers), nil
}
