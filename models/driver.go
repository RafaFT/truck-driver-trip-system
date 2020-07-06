package models

import (
	"fmt"
	"time"
)

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
