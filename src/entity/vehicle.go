package entity

type Vehicle string

const (
	// cars
	car Vehicle = "CAR"

	// motorcycles
	motorcycle Vehicle = "MOTORCYCLE"

	// trucks
	truck      Vehicle = "TRUCK"
	_34Truck   Vehicle = "3/4Truck"
	stumpTruck Vehicle = "STUMP_TRUCK"
)

var vehicleCodes = map[int]Vehicle{
	// cars: 0-99
	0: car,

	// motorcycles: 100-199
	100: motorcycle,

	// trucks: 200-299
	200: truck,
	201: _34Truck,
	202: stumpTruck,
}

func NewVehicle(vehicleCode int) (Vehicle, error) {
	vehicle, ok := vehicleCodes[vehicleCode]
	if !ok {
		return "", newErrInvalidVehicleCode(vehicleCode)
	}

	return vehicle, nil
}
