package main

import "time"

type VehicleType uint8

const (
	TRUCK_34         VehicleType = 1
	TRUCK_TOCO       VehicleType = 2
	TRUCK            VehicleType = 3
	SIMPLE_TRUCK     VehicleType = 4
	EXTENDED_TRAILER VehicleType = 5
)

// Driver type for Firestore Drivers collection
type Driver struct {
	CPF        *string    `firestore:"cpf" json:"cpf,omitempty"`
	Name       *string    `firestore:"name" json:"name,omitempty"`
	BirthDate  *time.Time `firestore:"birth_date" json:"birth_date,omitempty"`
	Age        *int       `firestore:"-" json:"age,omitempty"`
	Gender     *string    `firestore:"gender" json:"gender,omitempty"`
	HasVehicle *bool      `firestore:"has_vehicle" json:"has_vehicle,omitempty"`
	CNHType    *string    `firestore:"cnh_type" json:"cnh_type,omitempty"`
}

// Trip type for Firestore Trips collection
type Trip struct {
	HasLoad     *bool        `firestore:"has_load" json:"has_load"`
	VehicleType *VehicleType `firestore:"vehicle_type" json:"vehicle_type"`
	Origin      *Location    `firestore:"origin" json:"origin"`
	Destination *Location    `firestore:"destination" json:"destination"`
}

type Location struct {
	Latitute  *int
	Longitude *int
}
