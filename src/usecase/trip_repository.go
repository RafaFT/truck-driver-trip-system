package usecase

import (
	"context"

	"github.com/rafaft/truck-driver-trip-system/entity"
)

type TripRepository interface {
	Delete(context.Context, string) error
	Find(context.Context, FindTripsQuery) ([]*entity.Trip, error)
	FindByID(context.Context, string) (*entity.Trip, error)
	Save(context.Context, *entity.Trip) error
	Update(context.Context, *entity.Trip) error
}
