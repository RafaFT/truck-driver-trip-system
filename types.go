package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type VehicleType uint8
type CNHType string
type Gender string

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
	Gender     *Gender    `firestore:"gender" json:"gender,omitempty"`
	HasVehicle *bool      `firestore:"has_vehicle" json:"has_vehicle,omitempty"`
	CNHType    *CNHType   `firestore:"cnh_type" json:"cnh_type,omitempty"`
}

func (d *Driver) ValidateDriver() error {
	if d.CPF == nil ||
		d.Name == nil ||
		d.BirthDate == nil ||
		d.Gender == nil ||
		d.HasVehicle == nil ||
		d.CNHType == nil {
		fields := "['name', 'birth_date', 'gender', 'has_vehicle', 'cnh_type']"
		return fmt.Errorf("Driver must have all fields: %s", fields)
	}

	return nil
}

func (cnh *CNHType) UnmarshalJSON(b []byte) error {
	var sCNH string
	json.Unmarshal(b, &sCNH)

	sCNH = strings.ToUpper(sCNH)
	validCNHTypes := "ABCDE"
	if len(sCNH) != 1 || !strings.Contains(validCNHTypes, sCNH) {
		return fmt.Errorf("'cnh_type' must be 'A', 'B', 'C', 'D' or 'E'")
	}

	*cnh = CNHType(sCNH)
	return nil
}

func (gender *Gender) UnmarshalJSON(b []byte) error {
	var sGender string
	json.Unmarshal(b, &sGender)

	sGender = strings.ToUpper(sGender)
	validGenders := "FMO" // "O" stands for other
	if len(sGender) != 1 || !strings.Contains(validGenders, sGender) {
		return fmt.Errorf("'gender' must be 'F', 'M' or 'O'")
	}

	*gender = Gender(sGender)
	return nil
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

type errorJSON struct {
	Error string `json:"error"`
}
