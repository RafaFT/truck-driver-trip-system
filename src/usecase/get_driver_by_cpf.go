package usecase

import (
	"context"
	"time"

	"github.com/rafaft/truck-driver-trip-system/entity"
)

type (
	// input port
	GetDriverByCPF interface {
		Execute(context.Context, string) (*GetDriverByCPFOutput, error)
	}

	GetDriverByCPFRepo interface {
		FindByCPF(context.Context, entity.CPF) (*entity.Driver, error)
	}

	// input port implementation - Interactor
	getDriverByCPF struct {
		logger Logger
		repo   GetDriverByCPFRepo
	}

	// output data - type
	GetDriverByCPFOutput struct {
		Age        int
		BirthDate  time.Time
		CNH        string
		CPF        string
		Gender     string
		HasVehicle bool
		Name       string
	}
)

func NewGetDriverByCPF(logger Logger, repo GetDriverByCPFRepo) GetDriverByCPF {
	return getDriverByCPF{
		logger: logger,
		repo:   repo,
	}
}

func (di getDriverByCPF) Execute(ctx context.Context, cpf string) (*GetDriverByCPFOutput, error) {
	driverCPF, err := entity.NewCPF(cpf)
	if err != nil {
		di.logger.Debug(err.Error())
		return nil, err
	}

	driver, err := di.repo.FindByCPF(ctx, driverCPF)
	if err != nil {
		switch err.(type) {
		case entity.ErrDriverNotFound:
			di.logger.Debug(err.Error())
		default:
			di.logger.Error(err.Error())
		}

		return nil, err
	}

	return &GetDriverByCPFOutput{
		Age:        driver.Age(),
		BirthDate:  driver.BirthDate().Time,
		CNH:        string(driver.CNH()),
		CPF:        string(driver.CPF()),
		Gender:     string(driver.Gender()),
		HasVehicle: driver.HasVehicle(),
		Name:       string(driver.Name()),
	}, nil
}
