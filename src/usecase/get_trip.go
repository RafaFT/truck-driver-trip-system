package usecase

import (
	"context"
	"time"

	"github.com/rafaft/truck-driver-trip-system/entity"
)

type (
	// input port
	GetTrip interface {
		Execute(context.Context, string) (*GetTripOutput, error)
	}

	GetTripRepo interface {
		FindByID(context.Context, string) (*entity.Trip, error)
	}

	// input port implementation - interactor
	getTrip struct {
		logger Logger
		repo   GetTripRepo
	}

	GetTripOutput struct {
		ID              string
		StartDate       time.Time
		EndDate         time.Time
		Duration        time.Duration
		HasLoad         bool
		OriginLat       float64
		OriginLong      float64
		DestinationLat  float64
		DestinationLong float64
		Vehicle         string
	}
)

// NewGetTrip returns input port implementation for getting Trip by ID
func NewGetTrip(logger Logger, repo GetTripRepo) GetTrip {
	return getTrip{
		logger: logger,
		repo:   repo,
	}
}

func (ti getTrip) Execute(ctx context.Context, id string) (*GetTripOutput, error) {
	if len(id) == 0 {
		return nil, entity.ErrInvalidID
	}

	trip, err := ti.repo.FindByID(ctx, id)
	if err != nil {
		switch err.(type) {
		case entity.ErrTripNotFound:
			ti.logger.Debug(err.Error())
		default:
			ti.logger.Error(err.Error())
		}

		return nil, err
	}

	return &GetTripOutput{
		ID:              trip.ID(),
		StartDate:       trip.StartDate(),
		EndDate:         trip.EndDate(),
		Duration:        trip.Duration(),
		HasLoad:         trip.HasLoad(),
		OriginLat:       trip.Origin().Latitude(),
		OriginLong:      trip.Origin().Longitude(),
		DestinationLat:  trip.Destination().Latitude(),
		DestinationLong: trip.Destination().Longitude(),
		Vehicle:         string(trip.Vehicle()),
	}, nil
}
