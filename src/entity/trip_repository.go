package entity

import "context"

type TripRepository interface {
	Save(context.Context, *Trip) error
}
