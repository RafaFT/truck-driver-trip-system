package usecase

import (
	"context"

	"github.com/rafaft/truck-driver-trip-system/entity"
)

func (di DriverInteractor) DeleteDriver(ctx context.Context, cpf string) error {
	driverCPF, err := entity.NewCPF(cpf)
	if err != nil {
		return err
	}

	err = di.repo.DeleteDriverByCPF(ctx, driverCPF)
	if err != nil {
		return err
	}

	return nil
}
