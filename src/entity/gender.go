package entity

import "strings"

var genderValues = "FMO"

type Gender string

func NewGender(gender string) (Gender, error) {
	genderUpper := strings.ToUpper(gender)

	if len(genderUpper) != 1 || !strings.Contains(genderValues, genderUpper) {
		return "", newErrInvalidGender(gender)
	}

	return Gender(genderUpper), nil
}
