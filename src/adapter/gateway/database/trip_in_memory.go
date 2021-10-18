package repository

import (
	"context"

	"github.com/rafaft/truck-driver-trip-system/entity"
	"github.com/rafaft/truck-driver-trip-system/usecase"
)

// trip repository in memory implementation
type InMemoryTrips struct {
	Trips map[string]entity.Trip
}

func NewTripInMemory(trips []*entity.Trip) usecase.TripRepository {
	internalTrips := make(map[string]entity.Trip, len(trips))

	for _, trip := range trips {
		internalTrips[trip.ID()] = *trip
	}

	return &InMemoryTrips{
		Trips: internalTrips,
	}
}

func (t *InMemoryTrips) Delete(ctx context.Context, id string) error {
	_, ok := t.Trips[id]
	if !ok {
		return entity.NewErrTripNotFound(id)
	}

	delete(t.Trips, id)

	return nil
}

func (t *InMemoryTrips) Find(ctx context.Context, q usecase.FindTripsQuery) ([]*entity.Trip, error) {
	limit := len(t.Trips)
	if q.Limit != nil {
		// TODO: since limit is uint, it's possible limit value may be overflowed
		if inputLimit := int(*q.Limit); inputLimit < limit {
			limit = inputLimit
		}
	}

	trips := make([]*entity.Trip, 0, limit)
	for _, trip := range t.Trips {
		if len(trips) == limit {
			break
		}

		if q.CPF != nil && trip.CPF() != *q.CPF {
			continue
		}
		if q.HasLoad != nil && trip.HasLoad() != *q.HasLoad {
			continue
		}
		if q.Vehicle != nil && trip.Vehicle() != *q.Vehicle {
			continue
		}

		trips = append(trips, &trip)
	}

	return trips, nil
}

func (t *InMemoryTrips) FindByID(ctx context.Context, id string) (*entity.Trip, error) {
	trip, ok := t.Trips[id]
	if !ok {
		return nil, entity.NewErrTripNotFound(id)
	}

	return &trip, nil
}

func (t *InMemoryTrips) Save(ctx context.Context, trip *entity.Trip) error {
	_, ok := t.Trips[trip.ID()]
	if ok {
		return entity.NewErrTripAlreadyExists(trip.ID())
	}

	t.Trips[trip.ID()] = *trip

	return nil
}

func (t *InMemoryTrips) Update(ctx context.Context, trip *entity.Trip) error {
	_, ok := t.Trips[trip.ID()]
	if !ok {
		return entity.NewErrTripNotFound(trip.ID())
	}

	t.Trips[trip.ID()] = *trip

	return nil
}
