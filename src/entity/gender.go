package entity

import "strings"

var genderValues = map[string]string{
	"F": "female",
	"M": "male",
	"O": "other",
}

type Gender string

func NewGender(gender string) (Gender, error) {
	genderUpper := strings.ToUpper(gender)

	if _, ok := genderValues[genderUpper]; !ok {
		return "", NewErrInvalidGender(gender)
	}

	return Gender(genderUpper), nil
}
