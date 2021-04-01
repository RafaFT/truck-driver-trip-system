package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/rafaft/truck-driver-trip-system/entity"
)

// input port
type UpdateDriverUseCase interface {
	Execute(context.Context, string, UpdateDriverInput) (*UpdateDriverOutput, error)
}

// input port implementation
type UpdateDriverInteractor struct {
	logger Logger
	repo   entity.DriverRepository
}

type UpdateDriverInput struct {
	CNH        *string
	Gender     *string
	HasVehicle *bool
	Name       *string
}

type UpdateDriverOutput struct {
	Age        int
	BirthDate  time.Time
	CNH        string
	CPF        string
	Gender     string
	HasVehicle bool
	Name       string
	UpdatedAt  time.Time
}

func NewUpdateDriverInteractor(logger Logger, repo entity.DriverRepository) UpdateDriverUseCase {
	return UpdateDriverInteractor{
		logger: logger,
		repo:   repo,
	}
}

func (di UpdateDriverInteractor) Execute(ctx context.Context, cpf string, input UpdateDriverInput) (*UpdateDriverOutput, error) {
	driverCPF, err := entity.NewCPF(cpf)
	if err != nil {
		di.logger.Debug(err.Error())
		return nil, err
	}

	driver, err := di.repo.FindDriverByCPF(ctx, driverCPF)
	if err != nil {
		switch err.(type) {
		case entity.ErrDriverNotFound:
			di.logger.Debug(err.Error())
		default:
			di.logger.Error(err.Error())
		}

		return nil, err
	}

	if input.CNH != nil {
		if err := driver.SetCNHType(*input.CNH); err != nil {
			di.logger.Debug(err.Error())
			return nil, err
		}
	}

	if input.Gender != nil {
		if err := driver.SetGender(*input.Gender); err != nil {
			di.logger.Debug(err.Error())
			return nil, err
		}
	}

	if input.HasVehicle != nil {
		driver.SetHasVehicle(*input.HasVehicle)
	}

	if input.Name != nil {
		if err := driver.SetName(*input.Name); err != nil {
			di.logger.Debug(err.Error())
			return nil, err
		}
	}

	err = di.repo.UpdateDriver(ctx, driver)
	if err != nil {
		di.logger.Error(err.Error())
		return nil, err
	}

	logInput, _ := json.MarshalIndent(input, "", "\t")
	di.logger.Info(fmt.Sprintf("driver updated. cpf=[%s], update=[%s]", cpf, logInput))

	return &UpdateDriverOutput{
		Age:        driver.Age(),
		BirthDate:  driver.BirthDate().Time,
		CNH:        string(driver.CNHType()),
		CPF:        string(driver.CPF()),
		Gender:     string(driver.Gender()),
		HasVehicle: driver.HasVehicle(),
		Name:       string(driver.Name()),
		UpdatedAt:  time.Now(),
	}, nil
}
