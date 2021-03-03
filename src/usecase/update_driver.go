package usecase

import (
	"context"
	"time"

	"github.com/rafaft/truck-driver-trip-system/entity"
)

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

func (di DriverInteractor) UpdateDriver(ctx context.Context, cpf string, input UpdateDriverInput) (UpdateDriverOutput, error) {
	driverCPF, err := entity.NewCPF(cpf)
	if err != nil {
		return di.presenter.UpdateDriverOutput(nil), err
	}

	driver, err := di.repo.FindDriverByCPF(ctx, driverCPF)
	if err != nil {
		return di.presenter.UpdateDriverOutput(nil), err
	}

	if input.CNH != nil {
		if err := driver.SetCNHType(*input.CNH); err != nil {
			return di.presenter.UpdateDriverOutput(nil), err
		}
	}

	if input.Gender != nil {
		if err := driver.SetGender(*input.Gender); err != nil {
			return di.presenter.UpdateDriverOutput(nil), err
		}
	}

	if input.HasVehicle != nil {
		driver.SetHasVehicle(*input.HasVehicle)
	}

	if input.Name != nil {
		if err := driver.SetName(*input.Name); err != nil {
			return di.presenter.UpdateDriverOutput(nil), err
		}
	}

	err = di.repo.UpdateDriver(ctx, driver)
	if err != nil {
		return di.presenter.UpdateDriverOutput(nil), err
	}

	return di.presenter.UpdateDriverOutput(driver), nil
}
