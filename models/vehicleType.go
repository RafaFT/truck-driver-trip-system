package models

import (
	"encoding/json"
	"fmt"
)

type VehicleType int

func (vt *VehicleType) UnmarshalJSON(b []byte) error {
	vehicleTypes := map[int]string{
		1: "TRUCK_34",
		2: "TRUCK_TOCO",
		3: "TRUCK",
		4: "SIMPLE_TRUCK",
		5: "EXTENDED_TRAILER",
	}

	var vehicleType int
	err := json.Unmarshal(b, &vehicleType)
	if err != nil {
		return err
	}

	_, exist := vehicleTypes[vehicleType]
	if !exist {
		return fmt.Errorf("invalid vehicle_type")
	}

	*vt = VehicleType(vehicleType)
	return nil
}
