package usecase

import (
	"context"
	"time"
)

// input port
type DriverUseCase interface {
	CreateDriver(context.Context, CreateDriverInput) (CreateDriverOutput, error)
	DeleteDriver(context.Context, string) (time.Time, error)
	UpdateDriver(context.Context, string, UpdateDriverInput) (UpdateDriverOutput, error)
}
