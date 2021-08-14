package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/rafaft/truck-driver-trip-system/entity"
)

// input port
type UpdateTrip interface {
	Execute(context.Context, string, UpdateTripInput) (*UpdateTripOutput, error)
}

// input port implementation - Interactor
type updateTrip struct {
	logger Logger
	repo   entity.TripRepository
}

type UpdateTripInput struct {
	StartDate       *time.Time
	EndDate         *time.Time
	HasLoad         *bool
	OriginLat       *float64
	OriginLong      *float64
	DestinationLat  *float64
	DestinationLong *float64
	VehicleCode     *int
}

type UpdateTripOutput struct {
	ID              string
	StartDate       time.Time
	EndDate         time.Time
	HasLoad         bool
	OriginLat       float64
	OriginLong      float64
	DestinationLat  float64
	DestinationLong float64
	Vehicle         string
}

func NewUpdateTrip(logger Logger, repo entity.TripRepository) UpdateTrip {
	return updateTrip{
		logger: logger,
		repo:   repo,
	}
}

func (ti updateTrip) Execute(ctx context.Context, tripID string, input UpdateTripInput) (*UpdateTripOutput, error) {
	if len(tripID) == 0 {
		return nil, entity.ErrInvalidID
	}

	trip, err := ti.repo.FindByID(ctx, tripID)
	if err != nil {
		switch err.(type) {
		case entity.ErrTripNotFound:
			ti.logger.Debug(err.Error())
		default:
			ti.logger.Error(err.Error())
		}

		return nil, err
	}

	if input.StartDate != nil || input.EndDate != nil {
		start := trip.StartDate()
		end := trip.EndDate()

		if input.StartDate != nil {
			start = *input.StartDate
		}
		if input.EndDate != nil {
			end = *input.EndDate
		}

		if err := trip.SetTS(start, end); err != nil {
			ti.logger.Debug(err.Error())
			return nil, err
		}
	}

	if input.HasLoad != nil {
		trip.SetHasLoad(*input.HasLoad)
	}

	if input.OriginLat != nil || input.OriginLong != nil {
		lat := trip.Origin().Latitude()
		long := trip.Origin().Longitude()

		if input.OriginLat != nil {
			lat = *input.OriginLat
		}
		if input.OriginLong != nil {
			long = *input.OriginLong
		}

		if err := trip.SetOrigin(lat, long); err != nil {
			ti.logger.Debug(err.Error())
			return nil, err
		}
	}

	if input.DestinationLat != nil || input.DestinationLong != nil {
		lat := trip.Destination().Latitude()
		long := trip.Destination().Longitude()

		if input.DestinationLat != nil {
			lat = *input.DestinationLat
		}
		if input.DestinationLong != nil {
			long = *input.DestinationLong
		}

		if err := trip.SetDestination(lat, long); err != nil {
			ti.logger.Debug(err.Error())
			return nil, err
		}
	}

	if input.VehicleCode != nil {
		if err := trip.SetVehicle(*input.VehicleCode); err != nil {
			ti.logger.Debug(err.Error())
			return nil, err
		}
	}

	err = ti.repo.Update(ctx, trip)
	if err != nil {
		ti.logger.Error(err.Error())
		return nil, err
	}

	logInput, _ := json.MarshalIndent(input, "", "\t")
	ti.logger.Info(fmt.Sprintf("trip updated. id=[%s], update=[%s]", tripID, logInput))

	return &UpdateTripOutput{
		ID:              trip.ID(),
		StartDate:       trip.StartDate(),
		EndDate:         trip.EndDate(),
		HasLoad:         trip.HasLoad(),
		OriginLat:       trip.Origin().Latitude(),
		OriginLong:      trip.Origin().Longitude(),
		DestinationLat:  trip.Destination().Latitude(),
		DestinationLong: trip.Destination().Longitude(),
		Vehicle:         string(trip.Vehicle()),
	}, nil
}
