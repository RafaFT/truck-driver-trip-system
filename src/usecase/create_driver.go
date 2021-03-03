package usecase

import (
	"context"
	"time"

	"github.com/rafaft/truck-driver-trip-system/entity"
)

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

func (di DriverInteractor) CreateDriver(ctx context.Context, input CreateDriverInput) (CreateDriverOutput, error) {
	driver, err := entity.NewTruckDriver(
		input.CPF,
		input.Name,
		input.Gender,
		input.CNH,
		input.BirthDate,
		input.HasVehicle,
	)

	if err != nil {
		return di.presenter.CreateDriverOutput(nil), err
	}

	err = di.repo.SaveDriver(ctx, driver)
	if err != nil {
		return di.presenter.CreateDriverOutput(nil), err
	}

	return di.presenter.CreateDriverOutput(driver), nil
}
