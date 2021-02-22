package entity

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	cpfPattern = regexp.MustCompile(`^\d{11}$`)

	ErrInvalidBirthDate = errors.New("Invalid birth date")
	ErrInvalidCPF       = errors.New("Invalid CPF")
	ErrInvalidCNHType   = errors.New("Invalid cnhType")
	ErrInvalidGender    = errors.New("Invalid gender")
	ErrInvalidName      = errors.New("Invalid name")
)

type TruckDriver struct {
	birthDate  time.Time
	cnhType    string
	cpf        string
	gender     string
	hasVehicle bool
	name       string
}

func NewTruckDriver(cpf, name, gender, cnhType string, birthDate time.Time, hasVehicle bool) (*TruckDriver, error) {
	var err error

	if !cpfPattern.MatchString(cpf) {
		return nil, ErrInvalidCPF
	}

	name, err = NewName(name)
	if err != nil {
		return nil, err
	}

	gender, err = NewGender(gender)
	if err != nil {
		return nil, err
	}

	cnhType, err = NewCNHType(cnhType)
	if err != nil {
		return nil, err
	}

	birthDate, err = NewBirthDate(birthDate)
	if err != nil {
		return nil, err
	}

	driver := TruckDriver{
		cpf:        cpf,
		name:       name,
		gender:     gender,
		cnhType:    cnhType,
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
	return td.cnhType
}

func (td *TruckDriver) CPF() string {
	return td.cpf
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

func NewCNHType(cnhType string) (string, error) {
	cnhType = strings.ToUpper(cnhType)
	cnhTypeOptions := "ABCDE"

	if len(cnhType) != 1 || !strings.Contains(cnhTypeOptions, cnhType) {
		return "", ErrInvalidCNHType
	}

	return cnhType, nil
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
