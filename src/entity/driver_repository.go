package entity

import (
	"context"
)

type DriverRepository interface {
	DeleteByCPF(context.Context, CPF) error
	FindByCPF(context.Context, CPF) (*Driver, error)
	Find(context.Context, FindDriversQuery) ([]*Driver, error)
	Save(context.Context, *Driver) error
	Update(context.Context, *Driver) error
}
