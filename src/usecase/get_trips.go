package usecase

import (
	"context"
	"time"

	"github.com/rafaft/truck-driver-trip-system/entity"
)

// input port
type GetTrips interface {
	Execute(context.Context, GetTripsInput) ([]*GetTripsOutput, error)
}

// input port implementation - interactor
type getTrips struct {
	logger Logger
	repo   entity.TripRepository
}

type GetTripsInput struct {
	CPF         *string
	HasLoad     *bool
	Limit       *uint
	VehicleCode *int
}

type GetTripsOutput struct {
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

// NewGetTrips returns input port implementation for getting Trips
func NewGetTrips(logger Logger, repo entity.TripRepository) GetTrips {
	return getTrips{
		logger: logger,
		repo:   repo,
	}
}

func (ti getTrips) Execute(ctx context.Context, rawQ GetTripsInput) ([]*GetTripsOutput, error) {
	q, err := entity.NewFindTripsQuery(rawQ.CPF, rawQ.HasLoad, rawQ.Limit, rawQ.VehicleCode)
	if err != nil {
		ti.logger.Debug(err.Error())
		return nil, err
	}

	trips, err := ti.repo.Find(ctx, q)
	if err != nil {
		ti.logger.Error(err.Error())
		return nil, err
	}

	output := make([]*GetTripsOutput, len(trips))
	for i, trip := range trips {
		output[i] = &GetTripsOutput{
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
		}
	}

	return output, nil
}
