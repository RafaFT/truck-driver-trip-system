package main

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"time"

	"google.golang.org/genproto/googleapis/type/latlng"
)

// custom types for field validation
type CNHType string
type CPF string
type Gender string
type VehicleType int
type Latitute int
type Longitude int

// Driver type for Firestore Drivers collection
type Driver struct {
	CPF        *CPF       `firestore:"cpf" json:"cpf,omitempty"`
	Name       *string    `firestore:"name" json:"name,omitempty"`
	BirthDate  *time.Time `firestore:"birth_date" json:"birth_date,omitempty"`
	Age        int        `firestore:"-" json:"age,omitempty"`
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
		fields := "['cpf', 'name', 'birth_date', 'gender', 'has_vehicle', 'cnh_type']"
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

func (cpf *CPF) UnmarshalJSON(b []byte) error {
	var sCPF string
	err := json.Unmarshal(b, &sCPF)
	if err != nil {
		return err
	}

	matched, err := regexp.MatchString(`^\d{11}$`, sCPF)
	if err != nil {
		return err
	}
	if !matched {
		return fmt.Errorf("invalid value for 'CPF'")
	}

	*cpf = CPF(sCPF)
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
	CPF         *CPF           `firestore:"cpf" json:"cpf,omitempty"`
	HasLoad     *bool          `firestore:"has_load" json:"has_load,omitempty"`
	VehicleType *VehicleType   `firestore:"vehicle_type" json:"vehicle_type,omitempty"`
	Time        *time.Time     `firestore:"time" json:"time,omitempty"`
	Origin      *latlng.LatLng `firestore:"origin" json:"origin,omitempty"`
	Destination *latlng.LatLng `firestore:"destination" json:"destination,omitempty"`
}

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

func (lat *Latitute) UnmarshalJSON(b []byte) error {
	var latitude int
	err := json.Unmarshal(b, &latitude)
	if err != nil {
		return err
	}

	if latitude < -90 || latitude > 90 {
		return fmt.Errorf("latitude must be between -90 to 90")
	}

	*lat = Latitute(latitude)
	return nil
}

func (long *Longitude) UnmarshalJSON(b []byte) error {
	var longitude int
	err := json.Unmarshal(b, &longitude)
	if err != nil {
		return err
	}

	if longitude < -180 || longitude > 180 {
		return fmt.Errorf("longitude must be between -180 to 180")
	}

	*long = Longitude(longitude)
	return nil
}

type errorJSON struct {
	Error string `json:"error"`
}
