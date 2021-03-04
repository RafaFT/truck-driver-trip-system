package usecase

import (
	"context"
	"fmt"

	"github.com/rafaft/truck-driver-trip-system/entity"
)

// input port
type DeleteDriverUseCase interface {
	Execute(context.Context, string) error
}

// input port implementation
type DeleteDriverInteractor struct {
	logger Logger
	// presenter DeleteDriverPresenter
	repo entity.DriverRepository
}

func NewDeleteDriverInteractor(logger Logger, repo entity.DriverRepository) DeleteDriverUseCase {
	return DeleteDriverInteractor{
		logger: logger,
		// presenter: presenter,
		repo: repo,
	}
}

func (di DeleteDriverInteractor) Execute(ctx context.Context, cpf string) error {
	driverCPF, err := entity.NewCPF(cpf)
	if err != nil {
		di.logger.Debug(err.Error())
		return err
	}

	err = di.repo.DeleteDriverByCPF(ctx, driverCPF)
	if err != nil {
		di.logger.Warning(err.Error())
		return err
	}

	di.logger.Info(fmt.Sprintf("driver deleted. cpf=[%s]", cpf))

	return nil
}
