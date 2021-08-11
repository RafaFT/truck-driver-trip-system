package entity

import "context"

type TripRepository interface {
	Find(context.Context, FindTripsQuery) ([]*Trip, error)
	Save(context.Context, *Trip) error
}
