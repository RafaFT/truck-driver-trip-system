package usecase

import (
	"context"
	"time"

	"github.com/rafaft/truck-driver-trip-system/entity"
)

type (
	// input port
	GetTrips interface {
		Execute(context.Context, GetTripsInput) ([]*GetTripsOutput, error)
	}

	GetTripsRepo interface {
		Find(context.Context, FindTripsQuery) ([]*entity.Trip, error)
	}

	FindTripsQuery struct {
		CPF     *entity.CPF
		HasLoad *bool
		Limit   *uint
		Vehicle *entity.Vehicle
	}

	// input port implementation - interactor
	getTrips struct {
		logger Logger
		repo   GetTripsRepo
	}

	GetTripsInput struct {
		CPF         *string
		HasLoad     *bool
		Limit       *uint
		VehicleCode *int
	}

	GetTripsOutput struct {
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

func NewFindTripsQuery(cpf *string, hasLoad *bool, limit *uint, vehicleCode *int) (FindTripsQuery, error) {
	var q FindTripsQuery

	if cpf != nil {
		CPF, err := entity.NewCPF(*cpf)
		if err != nil {
			return q, err
		}

		q.CPF = &CPF
	}

	if vehicleCode != nil {
		Vehicle, err := entity.NewVehicle(*vehicleCode)
		if err != nil {
			return q, err
		}

		q.Vehicle = &Vehicle
	}

	q.HasLoad = hasLoad
	q.Limit = limit

	return q, nil
}

// NewGetTrips returns input port implementation for getting Trips
func NewGetTrips(logger Logger, repo GetTripsRepo) GetTrips {
	return getTrips{
		logger: logger,
		repo:   repo,
	}
}

func (ti getTrips) Execute(ctx context.Context, rawQ GetTripsInput) ([]*GetTripsOutput, error) {
	q, err := NewFindTripsQuery(rawQ.CPF, rawQ.HasLoad, rawQ.Limit, rawQ.VehicleCode)
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
