package usecase

import (
	"context"

	"github.com/rafaft/truck-driver-trip-system/entity"
)

// interactor implements the input port
type DriverInteractor struct {
	logger    Logger
	repo      entity.DriverRepository
	presenter DriverPresenter
}

// input port
type DriverUseCase interface {
	CreateDriver(context.Context, CreateDriverInput) (CreateDriverOutput, error)
	DeleteDriver(context.Context, string) error
	UpdateDriver(context.Context, string, UpdateDriverInput) (UpdateDriverOutput, error)
}

// output port
type DriverPresenter interface {
	CreateDriverOutput(*entity.Driver) CreateDriverOutput
	DeleteDriverOutput(error) error
	UpdateDriverOutput(*entity.Driver) UpdateDriverOutput
}

func NewDriverInteractor(logger Logger, repo entity.DriverRepository, presenter DriverPresenter) DriverUseCase {
	return DriverInteractor{
		logger:    logger,
		repo:      repo,
		presenter: presenter,
	}
}
