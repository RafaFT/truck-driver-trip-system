package usecase

import (
	"context"

	"github.com/rafaft/truck-driver-trip-system/entity"
)

// input port - interface
type GetDriverByCPFUseCase interface {
	Execute(context.Context, string) (*GetDriverByCPFOutput, error)
}

// input port implementation - interactor
type GetDriverByCPFInteractor struct {
	logger    Logger
	presenter GetDriverByCPFPresenter
	repo      entity.DriverRepository
}

// output port - (presenter) interface
type GetDriverByCPFPresenter interface {
	Output(*entity.Driver) *GetDriverByCPFOutput
}

// output data - type
type GetDriverByCPFOutput struct {
	BirthDate  *string `json:"birth_date,omitempty"`
	CNH        *string `json:"cnh,omitempty"`
	Gender     *string `json:"gender,omitempty"`
	HasVehicle *bool   `json:"has_vehicle,omitempty"`
	Name       *string `json:"name,omitempty"`
}

func NewGetDriverByCPFInteractor(logger Logger, presenter GetDriverByCPFPresenter, repo entity.DriverRepository) GetDriverByCPFUseCase {
	return GetDriverByCPFInteractor{
		logger:    logger,
		presenter: presenter,
		repo:      repo,
	}
}

func (di GetDriverByCPFInteractor) Execute(ctx context.Context, cpf string) (*GetDriverByCPFOutput, error) {
	driverCPF, err := entity.NewCPF(cpf)
	if err != nil {
		di.logger.Debug(err.Error())
		return di.presenter.Output(nil), err
	}

	driver, err := di.repo.FindDriverByCPF(ctx, driverCPF)
	if err != nil {
		switch err.(type) {
		case entity.ErrDriverNotFound:
			di.logger.Debug(err.Error())
		default:
			di.logger.Error(err.Error())
		}

		return di.presenter.Output(nil), err
	}

	return di.presenter.Output(driver), nil
}
