package entity

import (
	"errors"
	"testing"
)

func TestLocation(t *testing.T) {
	type locationInput struct {
		lat  float64
		long float64
	}

	tests := []struct {
		input    locationInput
		wantLat  float64
		wantLong float64
		wantErr  error
	}{
		// invalid input
		// invalid latitude values
		{
			locationInput{
				lat:  -90.0000001,
				long: 0,
			},
			0,
			0,
			newErrInvalidLatitude(-90.0000001),
		},
		{
			locationInput{
				lat:  90.0000001,
				long: 0,
			},
			0,
			0,
			newErrInvalidLatitude(90.0000001),
		},
		// invalid longitude values
		{
			locationInput{
				lat:  0,
				long: -180.0000001,
			},
			0,
			0,
			newErrInvalidLongitude(-180.0000001),
		},
		{
			locationInput{
				lat:  0,
				long: 180.0000001,
			},
			0,
			0,
			newErrInvalidLongitude(180.0000001),
		},
		// valid values
		{
			locationInput{
				lat:  0,
				long: 180.00000001, // values are truncated at 7 digits
			},
			0,
			180,
			nil,
		},
		{
			locationInput{
				lat:  0,
				long: 0,
			},
			0,
			0,
			nil,
		},
		{
			locationInput{
				lat:  -90,
				long: 180,
			},
			-90,
			180,
			nil,
		},
		{
			locationInput{
				lat:  90,
				long: -180,
			},
			90,
			-180,
			nil,
		},
		{
			locationInput{
				lat:  -45.12345678901234567890,
				long: 90.01234567890123456789,
			},
			-45.1234567,
			90.0123456,
			nil,
		},
	}

	for i, test := range tests {
		got, gotErr := NewLocation(test.input.lat, test.input.long)

		if !errors.Is(test.wantErr, gotErr) {
			t.Errorf("%d: [input: %v] [wantErr: %v] [gotErr: %v]", i, test.input, test.wantErr, gotErr)
			continue
		}

		if got.Latitude() != test.wantLat || got.Longitude() != test.wantLong {
			t.Errorf("%d: [input: %v] [wantLat: %v] [gotLat: %v] [wantLong: %v] [gotLong: %v]",
				i, test.input, test.wantLat, got.Latitude(), test.wantLong, got.Longitude(),
			)
		}
	}
}
