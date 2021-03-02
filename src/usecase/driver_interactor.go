package usecase

import (
	"context"

	"github.com/rafaft/truck-driver-trip-system/entity"
)

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
