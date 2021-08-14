package entity

import "context"

type TripRepository interface {
	Delete(context.Context, string) error
	Find(context.Context, FindTripsQuery) ([]*Trip, error)
	FindByID(context.Context, string) (*Trip, error)
	Save(context.Context, *Trip) error
	Update(context.Context, *Trip) error
}
