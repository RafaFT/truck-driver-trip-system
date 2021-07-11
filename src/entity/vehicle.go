package entity

type Vehicle string

const (
	truck      Vehicle = "TRUCK"
	_34Truck   Vehicle = "3/4Truck"
	stumpTruck Vehicle = "STUMP_TRUCK"
)

var vehicleCodes = map[int]Vehicle{
	0: truck,
	1: _34Truck,
	2: stumpTruck,
}

func NewVehicle(vehicleCode int) (Vehicle, error) {
	vehicle, ok := vehicleCodes[vehicleCode]
	if !ok {
		return "", newErrInvalidVehicleCode(vehicleCode)
	}

	return vehicle, nil
}
