package entity

import "time"

type Trip struct {
	id          string
	driverCPF   CPF
	hasLoad     bool
	origin      Location
	destination Location
	timeStamp   tripTS
	vehicle     Vehicle
}

type TripInput struct {
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

func NewTrip(id string, input TripInput) (*Trip, error) {
	if len(id) == 0 {
		return nil, ErrInvalidID
	}

	cpf, err := NewCPF(input.CPF)
	if err != nil {
		return &Trip{}, err
	}

	origin, err := NewLocation(input.OriginLat, input.OriginLong)
	if err != nil {
		return &Trip{}, err
	}

	destination, err := NewLocation(input.DestinationLat, input.DestinationLong)
	if err != nil {
		return &Trip{}, err
	}

	vehicle, err := NewVehicle(input.VehicleCode)
	if err != nil {
		return &Trip{}, err
	}

	tripTS, err := newTripTS(input.StartDate, input.EndDate)
	if err != nil {
		return &Trip{}, err
	}

	return &Trip{
		id:          id,
		driverCPF:   cpf,
		hasLoad:     input.HasLoad,
		origin:      origin,
		destination: destination,
		timeStamp:   tripTS,
		vehicle:     vehicle,
	}, nil
}

// getters
func (t *Trip) ID() string {
	return t.id
}

func (t *Trip) CPF() CPF {
	return t.driverCPF
}

func (t *Trip) EndDate() time.Time {
	return t.timeStamp.end
}

func (t *Trip) HasLoad() bool {
	return t.hasLoad
}

func (t *Trip) Origin() Location {
	return t.origin
}

func (t *Trip) Destination() Location {
	return t.destination
}

func (t *Trip) StartDate() time.Time {
	return t.timeStamp.start
}

func (t *Trip) Vehicle() Vehicle {
	return t.vehicle
}

func (t *Trip) Duration() time.Duration {
	return t.timeStamp.duration
}

// setters
func (t *Trip) SetHasLoad(hasLoad bool) {
	t.hasLoad = hasLoad
}

func (t *Trip) SetVehicle(vehicleCode int) error {
	vehicle, err := NewVehicle(vehicleCode)
	if err != nil {
		return err
	}

	t.vehicle = vehicle
	return nil
}

func (t *Trip) SetOrigin(lat, long float64) error {
	loc, err := NewLocation(lat, long)
	if err != nil {
		return err
	}

	t.origin = loc

	return nil
}

func (t *Trip) SetDestination(lat, long float64) error {
	loc, err := NewLocation(lat, long)
	if err != nil {
		return err
	}

	t.destination = loc

	return nil
}

func (t *Trip) SetTS(start, end time.Time) error {
	tripTS, err := newTripTS(start, end)
	if err != nil {
		return err
	}

	t.timeStamp = tripTS

	return nil
}
