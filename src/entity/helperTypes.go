package entity

import (
	"regexp"
	"strings"
	"time"
)

type CPF string
type Gender string
type CNHType string
type Name string
type BirthDate time.Time

var cpfPattern = regexp.MustCompile(`^\d{11}$`)

func NewCPF(cpf string) (CPF, error) {
	if !cpfPattern.MatchString(cpf) {
		return "", ErrInvalidCPF
	}

	return CPF(cpf), nil
}

func NewGender(gender string) (Gender, error) {
	gender = strings.ToUpper(gender)
	genderOptions := "FMO"

	if len(gender) != 1 || !strings.Contains(genderOptions, gender) {
		return "", ErrInvalidGender
	}

	return Gender(gender), nil
}

func NewCNHType(cnhType string) (CNHType, error) {
	cnhType = strings.ToUpper(cnhType)
	cnhTypeOptions := "ABCDE"

	if len(cnhType) != 1 || !strings.Contains(cnhTypeOptions, cnhType) {
		return "", ErrInvalidCNHType
	}

	return CNHType(cnhType), nil
}

func NewName(name string) (Name, error) {
	if len(name) == 0 {
		return "", ErrInvalidName
	}

	return Name(strings.ToLower(name)), nil
}

func NewBirthDate(birthDate time.Time) (BirthDate, error) {
	age := calculateAge(birthDate)

	// if system supported drivers from other countries,
	// the minimum age would depend on the location of the
	// CNH Type
	if age < 18 {
		return BirthDate(time.Time{}), ErrInvalidBirthDate
	}

	return BirthDate(birthDate), nil
}
