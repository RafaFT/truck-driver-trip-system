package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/rafaft/truck-driver-trip-system/entity"
)

// input port
type CreateTrip interface {
	Execute(context.Context, CreateTripInput) (*CreateTripOutput, error)
}

// input port implementation - interactor
type createTrip struct {
	logger     Logger
	tripRepo   entity.TripRepository
	driverRepo entity.DriverRepository
}

type CreateTripInput struct {
	CPF             string
	StartDate       time.Time
	EndDate         time.Time
	HasLoad         bool
	OriginLat       float64
	OriginLong      float64
	DestinationLat  float64
	DestinationLong float64
	VehicleCode     int
}

type CreateTripOutput struct {
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

// NewCreateTrip returns input port implementation for creating new Trips
func NewCreateTrip(logger Logger, tripRepo entity.TripRepository, driverRepo entity.DriverRepository) CreateTrip {
	return createTrip{
		logger:     logger,
		tripRepo:   tripRepo,
		driverRepo: driverRepo,
	}
}

func (ti createTrip) Execute(ctx context.Context, input CreateTripInput) (*CreateTripOutput, error) {
	var output CreateTripOutput

	tripUUID, err := uuid.NewRandom()
	if err != nil {
		ti.logger.Error("Could not generate random UUID for new trip")
		return &output, err
	}

	tripInput := entity.TripInput{
		CPF:             input.CPF,
		StartDate:       input.StartDate,
		EndDate:         input.EndDate,
		HasLoad:         input.HasLoad,
		OriginLat:       input.OriginLat,
		OriginLong:      input.OriginLong,
		DestinationLat:  input.DestinationLat,
		DestinationLong: input.DestinationLong,
		VehicleCode:     input.VehicleCode,
	}

	trip, err := entity.NewTrip(tripUUID.String(), tripInput)
	if err != nil {
		ti.logger.Debug(err.Error())
		return &output, err
	}

	_, err = ti.driverRepo.FindByCPF(ctx, trip.CPF())
	if err != nil {
		switch err.(type) {
		case entity.ErrDriverNotFound:
			ti.logger.Debug(err.Error())
		default:
			ti.logger.Error(err.Error())
		}

		return &output, err
	}

	err = ti.tripRepo.Save(ctx, trip)
	if err != nil {
		ti.logger.Error(err.Error())
		return &output, err
	}

	logInput, _ := json.MarshalIndent(input, "", "\t")
	ti.logger.Info(fmt.Sprintf("new trip created. input=[%s]", string(logInput)))

	output = CreateTripOutput{
		ID:              trip.ID(),
		StartDate:       trip.StartDate(),
		EndDate:         trip.EndDate(),
		HasLoad:         trip.HasLoad(),
		OriginLat:       trip.Origin().Latitude(),
		OriginLong:      trip.Origin().Longitude(),
		DestinationLat:  trip.Destination().Latitude(),
		DestinationLong: trip.Destination().Longitude(),
		Vehicle:         string(trip.Vehicle()),
	}

	return &output, nil
}
