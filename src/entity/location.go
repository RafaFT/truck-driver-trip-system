package entity

import (
	"github.com/shopspring/decimal"
)

var (
	minLatitude = decimal.NewFromInt(-90)
	maxLatitude = decimal.NewFromInt(90)

	minLongitude = decimal.NewFromInt(-180)
	maxLongitude = decimal.NewFromInt(180)
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
	lat := latitude{decimal.NewFromFloat(rawLat).Truncate(7)}

	if lat.GreaterThan(maxLatitude) || lat.LessThan(minLatitude) {
		return lat, NewErrInvalidLatitude(rawLat)
	}

	return lat, nil
}

func newLongitude(rawLong float64) (longitude, error) {
	long := longitude{decimal.NewFromFloat(rawLong).Truncate(7)}

	if long.GreaterThan(maxLongitude) || long.LessThan(minLongitude) {
		return long, NewErrInvalidLongitude(rawLong)
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

func (l Location) Latitude() float64 {
	f, _ := l.lat.Float64()
	return f
}

func (l Location) Longitude() float64 {
	f, _ := l.long.Float64()
	return f
}
