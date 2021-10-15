package entity_test

import (
	"reflect"
	"testing"

	"github.com/rafaft/truck-driver-trip-system/entity"
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
		err      error
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
			entity.ErrInvalidLatitude{},
		},
		{
			locationInput{
				lat:  90.0000001,
				long: 0,
			},
			0,
			0,
			entity.ErrInvalidLatitude{},
		},
		// invalid longitude values
		{
			locationInput{
				lat:  0,
				long: -180.0000001,
			},
			0,
			0,
			entity.ErrInvalidLongitude{},
		},
		{
			locationInput{
				lat:  0,
				long: 180.0000001,
			},
			0,
			0,
			entity.ErrInvalidLongitude{},
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
		got, gotError := entity.NewLocation(test.input.lat, test.input.long)

		if reflect.TypeOf(test.err) != reflect.TypeOf(gotError) {
			t.Errorf("%d: [input: %v] [wantError: %T] [gotError: %v]",
				i, test.input, test.err, gotError,
			)
			continue
		}

		if got.Latitude() != test.wantLat || got.Longitude() != test.wantLong {
			t.Errorf("%d: [input: %v] [wantLat: %v] [wantLong: %v] [gotLat: %v] [gotLong: %v]",
				i, test.input, test.wantLat, test.wantLong, got.Latitude(), got.Longitude(),
			)
		}
	}
}
