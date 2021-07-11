package entity

import (
	"github.com/shopspring/decimal"
)

var (
	maxLatitude = decimal.NewFromInt(90)
	minLatitude = decimal.NewFromInt(-90)

	maxLongitude = decimal.NewFromInt(180)
	minLongitude = decimal.NewFromInt(-180)
)

type latitude struct {
	decimal.Decimal
}

type longitude struct {
	decimal.Decimal
}

type Location struct {
	lat  latitude
	long longitude
}

func newLatitude(rawLat float64) (latitude, error) {
	lat := latitude{decimal.NewFromFloat(rawLat)}

	lat.Truncate(7)
	if lat.GreaterThan(maxLatitude) || lat.LessThan(minLatitude) {
		return lat, newErrInvalidLatitude(rawLat)
	}

	return lat, nil
}

func newLongitude(rawLong float64) (longitude, error) {
	long := longitude{decimal.NewFromFloat(rawLong)}

	long.Truncate(7)
	if long.GreaterThan(maxLatitude) || long.LessThan(minLatitude) {
		return long, newErrInvalidLongitude(rawLong)
	}

	return long, nil
}

func NewLocation(rawLat, rawLong float64) (Location, error) {
	var loc Location

	lat, err := newLatitude(rawLat)
	if err != nil {
		return loc, err
	}

	long, err := newLongitude(rawLong)
	if err != nil {
		return loc, err
	}

	loc.lat = lat
	loc.long = long

	return loc, nil
}
