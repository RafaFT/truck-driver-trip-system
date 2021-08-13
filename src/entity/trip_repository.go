package entity

import "context"

type TripRepository interface {
	Find(context.Context, FindTripsQuery) ([]*Trip, error)
	FindByID(context.Context, string) (*Trip, error)
	Save(context.Context, *Trip) error
}
