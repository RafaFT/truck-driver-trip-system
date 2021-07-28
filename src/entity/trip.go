package entity

import "time"

type Trip struct {
	driverCPF   CPF
	hasLoad     bool
	origin      Location
	destination Location
	timeStamp   TripTS
	vehicle     Vehicle
}

type TripInput struct {
	CPF             string
	startDate       time.Time
	endDate         time.Time
	hasLoad         bool
	originLat       float64
	originLong      float64
	destinationLat  float64
	destinationLong float64
	vehicleCode     int
}

func NewTrip(input TripInput) (*Trip, error) {
	cpf, err := NewCPF(input.CPF)
	if err != nil {
		return &Trip{}, err
	}

	origin, err := NewLocation(input.originLat, input.originLong)
	if err != nil {
		return &Trip{}, err
	}

	destination, err := NewLocation(input.destinationLat, input.destinationLong)
	if err != nil {
		return &Trip{}, err
	}

	vehicle, err := NewVehicle(input.vehicleCode)
	if err != nil {
		return &Trip{}, err
	}

	tripTS, err := NewTripTS(input.startDate, input.endDate)
	if err != nil {
		return &Trip{}, err
	}

	return &Trip{
		driverCPF:   cpf,
		hasLoad:     input.hasLoad,
		origin:      origin,
		destination: destination,
		timeStamp:   tripTS,
		vehicle:     vehicle,
	}, nil
}

// getters
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
	tripTS, err := NewTripTS(start, end)
	if err != nil {
		return err
	}

	t.timeStamp = tripTS

	return nil
}
