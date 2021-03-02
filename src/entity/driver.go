package entity

import (
	"time"
)

const minimumDriverAge = 18

type TruckDriver struct {
	birthDate  BirthDate
	cnh        CNH
	cpf        CPF
	gender     Gender
	hasVehicle bool
	name       Name
}

func NewTruckDriver(cpf, name, gender, cnh string, birthDate time.Time, hasVehicle bool) (*TruckDriver, error) {
	newCPF, err := NewCPF(cpf)
	if err != nil {
		return nil, err
	}

	newName, err := NewName(name)
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

	newBirthDate, err := NewBirthDate(birthDate)
	if err != nil {
		return nil, err
	}

	if age := newBirthDate.CalculateAge(); age < minimumDriverAge {
		return nil, newErrInvalidAge(age)
	}

	driver := TruckDriver{
		cpf:        newCPF,
		name:       newName,
		gender:     newGender,
		cnh:        newCNH,
		birthDate:  newBirthDate,
		hasVehicle: hasVehicle,
	}

	return &driver, nil
}

// getters
func (td *TruckDriver) Age() int {
	return td.birthDate.CalculateAge()
}

func (td *TruckDriver) BirthDate() BirthDate {
	return td.birthDate
}

func (td *TruckDriver) CNHType() CNH {
	return td.cnh
}

func (td *TruckDriver) CPF() CPF {
	return td.cpf
}

func (td *TruckDriver) Gender() Gender {
	return td.gender
}

func (td *TruckDriver) HasVehicle() bool {
	return td.hasVehicle
}

func (td *TruckDriver) Name() Name {
	return td.name
}

// setters
func (td *TruckDriver) SetCNHType(cnh string) error {
	newCNH, err := NewCNH(cnh)
	if err != nil {
		return err
	}

	td.cnh = newCNH

	return nil
}

func (td *TruckDriver) SetGender(gender string) error {
	newGender, err := NewGender(gender)
	if err != nil {
		return err
	}

	td.gender = newGender

	return nil
}

func (td *TruckDriver) SetHasVehicle(hasVehicle bool) {
	td.hasVehicle = hasVehicle
}

func (td *TruckDriver) SetName(name string) error {
	newName, err := NewName(name)
	if err != nil {
		return err
	}

	td.name = newName

	return nil
}
