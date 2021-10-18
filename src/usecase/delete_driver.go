package usecase

import (
	"context"
	"fmt"

	"github.com/rafaft/truck-driver-trip-system/entity"
)

type (
	// input port
	DeleteDriver interface {
		Execute(context.Context, string) error
	}

	DeleteDriverRepo interface {
		DeleteByCPF(context.Context, entity.CPF) error
	}

	// input port implementation - Interactor
	deleteDriver struct {
		logger Logger
		repo   DeleteDriverRepo
	}
)

func NewDeleteDriver(logger Logger, repo DeleteDriverRepo) DeleteDriver {
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
