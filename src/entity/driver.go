package entity

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

var (
	ErrInvalidBirthDate = errors.New("Invalid birth date")
	ErrInvalidGender    = errors.New("Invalid gender")
	ErrInvalidName      = errors.New("Invalid name")
)

type TruckDriver struct {
	birthDate  time.Time
	cnh        CNH
	cpf        CPF
	gender     string
	hasVehicle bool
	name       string
}

func NewTruckDriver(cpf, name, gender, cnh string, birthDate time.Time, hasVehicle bool) (*TruckDriver, error) {
	var err error

	newCPF, err := NewCPF(cpf)
	if err != nil {
		return nil, err
	}

	name, err = NewName(name)
	if err != nil {
		return nil, err
	}

	gender, err = NewGender(gender)
	if err != nil {
		return nil, err
	}

	newCNH, err := NewCNH(cnh)
	if err != nil {
		return nil, err
	}

	birthDate, err = NewBirthDate(birthDate)
	if err != nil {
		return nil, err
	}

	driver := TruckDriver{
		cpf:        newCPF,
		name:       name,
		gender:     gender,
		cnh:        newCNH,
		birthDate:  birthDate,
		hasVehicle: hasVehicle,
	}

	return &driver, nil
}

func (td *TruckDriver) Age() int {
	return calculateAge(time.Now(), time.Time(td.birthDate))
}

func (td *TruckDriver) BirthDate() time.Time {
	return td.birthDate
}

func (td *TruckDriver) CNHType() string {
	return string(td.cnh)
}

func (td *TruckDriver) CPF() string {
	return string(td.cpf)
}

func (td *TruckDriver) Gender() string {
	return td.gender
}

func (td *TruckDriver) HasVehicle() bool {
	return td.hasVehicle
}

func (td *TruckDriver) Name() string {
	return td.name
}

func calculateAge(baseDate, birthDate time.Time) int {
	years := baseDate.Year() - birthDate.Year()
	if years < 0 {
		return 0
	}

	birthMonthNDay, _ := strconv.Atoi(fmt.Sprintf("%d%d", int(birthDate.Month()), birthDate.Day()))
	baseDateMonthNDay, _ := strconv.Atoi(fmt.Sprintf("%d%d", int(baseDate.Month()), baseDate.Day()))

	if birthMonthNDay > baseDateMonthNDay {
		years--
	}

	return years
}

func NewGender(gender string) (string, error) {
	gender = strings.ToUpper(gender)
	genderOptions := "FMO"

	if len(gender) != 1 || !strings.Contains(genderOptions, gender) {
		return "", ErrInvalidGender
	}

	return gender, nil
}

func NewName(name string) (string, error) {
	if len(name) == 0 {
		return "", ErrInvalidName
	}

	return strings.ToLower(name), nil
}

func NewBirthDate(birthDate time.Time) (time.Time, error) {
	age := calculateAge(time.Now(), birthDate)

	// if system supported drivers from other countries,
	// the minimum age would depend on the location of the
	// CNH Type
	if age < 18 {
		return time.Time{}, ErrInvalidBirthDate
	}

	return birthDate, nil
}
