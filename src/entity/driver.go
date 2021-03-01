package entity

import (
	"errors"
	"strings"
	"time"
)

const minimumDriverAge = 18

var (
	ErrInvalidName = errors.New("Invalid name")
)

type TruckDriver struct {
	birthDate  BirthDate
	cnh        CNH
	cpf        CPF
	gender     Gender
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

	newGender, err := NewGender(gender)
	if err != nil {
		return nil, err
	}

	newCNH, err := NewCNH(cnh)
	if err != nil {
		return nil, err
	}

	newBirthDate := NewBirthDate(birthDate)
	if age := newBirthDate.CalculateAge(); age < minimumDriverAge {
		return nil, newErrInvalidAge(age)
	}

	driver := TruckDriver{
		cpf:        newCPF,
		name:       name,
		gender:     newGender,
		cnh:        newCNH,
		birthDate:  newBirthDate,
		hasVehicle: hasVehicle,
	}

	return &driver, nil
}

func (td *TruckDriver) Age() int {
	return td.birthDate.CalculateAge()
}

func (td *TruckDriver) BirthDate() time.Time {
	return td.birthDate.Time
}

func (td *TruckDriver) CNHType() string {
	return string(td.cnh)
}

func (td *TruckDriver) CPF() string {
	return string(td.cpf)
}

func (td *TruckDriver) Gender() string {
	return string(td.gender)
}

func (td *TruckDriver) HasVehicle() bool {
	return td.hasVehicle
}

func (td *TruckDriver) Name() string {
	return td.name
}

func NewName(name string) (string, error) {
	if len(name) == 0 {
		return "", ErrInvalidName
	}

	return strings.ToLower(name), nil
}
