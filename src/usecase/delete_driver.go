package usecase

import (
	"context"
	"fmt"

	"github.com/rafaft/truck-driver-trip-system/entity"
)

// input port
type DeleteDriver interface {
	Execute(context.Context, string) error
}

// input port implementation - Interactor
type deleteDriver struct {
	logger Logger
	repo   entity.DriverRepository
}

func NewDeleteDriver(logger Logger, repo entity.DriverRepository) DeleteDriver {
	return deleteDriver{
		logger: logger,
		repo:   repo,
	}
}

func (di deleteDriver) Execute(ctx context.Context, cpf string) error {
	driverCPF, err := entity.NewCPF(cpf)
	if err != nil {
		di.logger.Debug(err.Error())
		return err
	}

	err = di.repo.DeleteByCPF(ctx, driverCPF)
	if err != nil {
		switch err.(type) {
		case entity.ErrDriverNotFound:
			di.logger.Debug(err.Error())
		default:
			di.logger.Error(err.Error())
		}

		return err
	}

	di.logger.Info(fmt.Sprintf("driver deleted. cpf=[%s]", cpf))

	return nil
}
