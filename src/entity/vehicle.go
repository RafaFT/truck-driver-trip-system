package entity

type Vehicle string

const (
	Truck      Vehicle = "TRUCK"
	Truck_34   Vehicle = "3/4Truck"
	StumpTruck Vehicle = "STUMP_TRUCK"
)

var vehicleCodes = map[int]Vehicle{
	0: Truck,
	1: Truck_34,
	2: StumpTruck,
}

func NewVehicle(vehicleCode int) (Vehicle, error) {
	vehicle, ok := vehicleCodes[vehicleCode]
	if !ok {
		return "", newErrInvalidVehicleCode(vehicleCode)
	}

	return vehicle, nil
}
