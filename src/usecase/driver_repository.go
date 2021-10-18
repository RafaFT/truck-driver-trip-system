package usecase

import (
	"context"

	"github.com/rafaft/truck-driver-trip-system/entity"
)

type DriverRepository interface {
	DeleteByCPF(context.Context, entity.CPF) error
	FindByCPF(context.Context, entity.CPF) (*entity.Driver, error)
	Find(context.Context, FindDriversQuery) ([]*entity.Driver, error)
	Save(context.Context, *entity.Driver) error
	Update(context.Context, *entity.Driver) error
}
