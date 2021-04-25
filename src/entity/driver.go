package entity

import (
	"time"
)

const minimumDriverAge = 18

type Driver struct {
	birthDate  BirthDate
	cnh        CNH
	cpf        CPF
	gender     Gender
	hasVehicle bool
	name       Name
}

func NewDriver(cpf, name, gender, cnh string, birthDate time.Time, hasVehicle bool) (*Driver, error) {
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

	driver := Driver{
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
func (td *Driver) Age() int {
	return td.birthDate.CalculateAge()
}

func (td *Driver) BirthDate() BirthDate {
	return td.birthDate
}

func (td *Driver) CNHType() CNH {
	return td.cnh
}

func (td *Driver) CPF() CPF {
	return td.cpf
}

func (td *Driver) Gender() Gender {
	return td.gender
}

func (td *Driver) HasVehicle() bool {
	return td.hasVehicle
}

func (td *Driver) Name() Name {
	return td.name
}

// setters
func (td *Driver) SetCNHType(cnh string) error {
	newCNH, err := NewCNH(cnh)
	if err != nil {
		return err
	}

	td.cnh = newCNH

	return nil
}

func (td *Driver) SetGender(gender string) error {
	newGender, err := NewGender(gender)
	if err != nil {
		return err
	}

	td.gender = newGender

	return nil
}

func (td *Driver) SetHasVehicle(hasVehicle bool) {
	td.hasVehicle = hasVehicle
}

func (td *Driver) SetName(name string) error {
	newName, err := NewName(name)
	if err != nil {
		return err
	}

	td.name = newName

	return nil
}
