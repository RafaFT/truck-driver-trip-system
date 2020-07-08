package models

import (
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/genproto/googleapis/type/latlng"
)

// Trip type for Firestore Trips collection
type Trip struct {
	ID          string         `firestore:"id" json:"id,omitempty"`
	DriverID    *DriverID      `firestore:"driver_id" json:"driver_id,omitempty"`
	HasLoad     *bool          `firestore:"has_load" json:"has_load,omitempty"`
	VehicleType *VehicleType   `firestore:"vehicle_type" json:"vehicle_type,omitempty"`
	Time        *time.Time     `firestore:"time" json:"time,omitempty"`
	Origin      *latlng.LatLng `firestore:"origin" json:"origin,omitempty"`
	Destination *latlng.LatLng `firestore:"destination" json:"destination,omitempty"`
}

func (t *Trip) ValidateTrip() error {
	if t.DriverID == nil ||
		t.HasLoad == nil ||
		t.VehicleType == nil ||
		t.Time == nil ||
		t.Origin == nil ||
		t.Destination == nil {
		fields := "['driver_id', 'has_load', 'vehicle_type', 'time', 'origin', 'destination']"
		return fmt.Errorf("Trip must have all fields: %s", fields)
	}

	return nil
}

func (t *Trip) SetID() error {
	if t.Time == nil {
		return fmt.Errorf("cannot set trip ID with field Time==nil")
	}
	// a trip's ID is it's timestamp on string format
	t.ID = t.Time.Format("20060102150405")
	return nil
}

func NewTrip(b []byte) (*Trip, error) {
	var trip Trip
	err := json.Unmarshal(b, &trip)
	if err != nil {
		return nil, err
	}

	err = trip.ValidateTrip()
	if err != nil {
		return nil, err
	}

	trip.SetID()

	return &trip, nil
}
