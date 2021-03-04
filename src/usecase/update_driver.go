package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/rafaft/truck-driver-trip-system/entity"
)

// input port
type UpdateDriverUseCase interface {
	Execute(context.Context, string, UpdateDriverInput) (UpdateDriverOutput, error)
}

// input port implementation
type UpdateDriverInteractor struct {
	logger    Logger
	presenter UpdateDriverPresenter
	repo      entity.DriverRepository
}

// output port
type UpdateDriverPresenter interface {
	Output(*entity.Driver) UpdateDriverOutput
}

type UpdateDriverInput struct {
	CNH        *string
	Gender     *string
	HasVehicle *bool
	Name       *string
}

type UpdateDriverOutput struct {
	BirthDate  time.Time `json:"birth_date"`
	CNH        string    `json:"cnh"`
	CPF        string    `json:"cpf"`
	Gender     string    `json:"gender"`
	HasVehicle bool      `json:"has_vehicle"`
	Name       string    `json:"name"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func NewUpdateDriverInteractor(logger Logger, presenter UpdateDriverPresenter, repo entity.DriverRepository) UpdateDriverUseCase {
	return UpdateDriverInteractor{
		logger:    logger,
		presenter: presenter,
		repo:      repo,
	}
}

func (di UpdateDriverInteractor) Execute(ctx context.Context, cpf string, input UpdateDriverInput) (UpdateDriverOutput, error) {
	driverCPF, err := entity.NewCPF(cpf)
	if err != nil {
		di.logger.Debug(err.Error())
		return di.presenter.Output(nil), err
	}

	driver, err := di.repo.FindDriverByCPF(ctx, driverCPF)
	if err != nil {
		di.logger.Debug(err.Error())
		return di.presenter.Output(nil), err
	}

	if input.CNH != nil {
		if err := driver.SetCNHType(*input.CNH); err != nil {
			di.logger.Debug(err.Error())
			return di.presenter.Output(nil), err
		}
	}

	if input.Gender != nil {
		if err := driver.SetGender(*input.Gender); err != nil {
			di.logger.Debug(err.Error())
			return di.presenter.Output(nil), err
		}
	}

	if input.HasVehicle != nil {
		driver.SetHasVehicle(*input.HasVehicle)
	}

	if input.Name != nil {
		if err := driver.SetName(*input.Name); err != nil {
			di.logger.Debug(err.Error())
			return di.presenter.Output(nil), err
		}
	}

	err = di.repo.UpdateDriver(ctx, driver)
	if err != nil {
		di.logger.Error(err.Error())
		return di.presenter.Output(nil), err
	}

	di.logger.Info(fmt.Sprintf("driver updated. cpf=[%s], update=[%v]", cpf, input))

	return di.presenter.Output(driver), nil
}
