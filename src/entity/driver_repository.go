package entity

import (
	"context"
)

type DriverRepository interface {
	DeleteDriverByCPF(context.Context, CPF) error
	FindDriverByCPF(context.Context, CPF) (*Driver, error)
	FindDrivers(context.Context) ([]*Driver, error)
	UpdateDriver(context.Context, UpdateDriverInput) (*Driver, error)
}

type UpdateDriverInput struct {
	CNH        *CNH
	Gender     *Gender
	HasVehicle *bool
	Name       *Name
}
