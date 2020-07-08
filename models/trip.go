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

func NewTrip(b []byte) (*Trip, error) {
	var trip Trip
	err := json.Unmarshal(b, &trip)
	if err != nil {
		return nil, err
	}

	if trip.DriverID == nil ||
		trip.HasLoad == nil ||
		trip.VehicleType == nil ||
		trip.Time == nil ||
		trip.Origin == nil ||
		trip.Destination == nil {
		fields := "['driver_id', 'has_load', 'vehicle_type', 'time', 'origin', 'destination']"
		return nil, fmt.Errorf("Trip must have all fields: %s", fields)
	}

	// a trip's ID is it's timestamp on string format
	trip.ID = trip.Time.Format("20060102150405")

	return &trip, nil
}
