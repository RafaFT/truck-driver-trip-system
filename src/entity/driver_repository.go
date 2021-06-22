package entity

import (
	"context"
)

type DriverRepository interface {
	DeleteDriverByCPF(context.Context, CPF) error
	FindDriverByCPF(context.Context, CPF) (*Driver, error)
	FindDrivers(context.Context, FindDriversQuery) ([]*Driver, error)
	SaveDriver(context.Context, *Driver) error
	UpdateDriver(context.Context, *Driver) error
}
