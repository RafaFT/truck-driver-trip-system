package entity

import (
	"time"
)

type BirthDate struct {
	time.Time
}

// arbitrary lowest valid birth date
var minBirthDate = time.Date(1950, 1, 1, 0, 0, 0, 0, time.UTC)

func NewBirthDate(birthDate time.Time) (BirthDate, error) {
	if minBirthDate.After(birthDate) {
		return BirthDate{}, newErrInvalidBirthDate(birthDate)
	}

	return BirthDate{birthDate.UTC()}, nil
}

func (bd BirthDate) age() int {
	now := time.Now().UTC()

	years := now.Year() - bd.Year()
	if years < 0 {
		return 0
	}

	if bd.Month() >= now.Month() && bd.Day() > now.Day() {
		years--
	}

	return years
}
