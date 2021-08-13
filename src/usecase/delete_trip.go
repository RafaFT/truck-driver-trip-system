package usecase

import (
	"context"
	"fmt"

	"github.com/rafaft/truck-driver-trip-system/entity"
)

// input port
type DeleteTrip interface {
	Execute(context.Context, string) error
}

// input port implementation - Interactor
type deleteTrip struct {
	logger Logger
	repo   entity.TripRepository
}

func NewDeleteTrip(logger Logger, repo entity.TripRepository) DeleteTrip {
	return deleteTrip{
		logger: logger,
		repo:   repo,
	}
}

func (ti deleteTrip) Execute(ctx context.Context, id string) error {
	if len(id) == 0 {
		return entity.ErrInvalidID
	}

	err := ti.repo.Delete(ctx, id)
	if err != nil {
		switch err.(type) {
		case entity.ErrTripNotFound:
			ti.logger.Debug(err.Error())
		default:
			ti.logger.Error(err.Error())
		}

		return err
	}

	ti.logger.Info(fmt.Sprintf("trip deleted. id=[%s]", id))

	return nil
}
