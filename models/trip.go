package models

import (
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/genproto/googleapis/type/latlng"
)

// custom types for field validation
type Latitute int
type Longitude int

// Trip type for Firestore Trips collection
type Trip struct {
	CPF         *CPF           `firestore:"cpf" json:"cpf,omitempty"`
	HasLoad     *bool          `firestore:"has_load" json:"has_load,omitempty"`
	VehicleType *VehicleType   `firestore:"vehicle_type" json:"vehicle_type,omitempty"`
	Time        *time.Time     `firestore:"time" json:"time,omitempty"`
	Origin      *latlng.LatLng `firestore:"origin" json:"origin,omitempty"`
	Destination *latlng.LatLng `firestore:"destination" json:"destination,omitempty"`
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
