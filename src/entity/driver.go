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

	if age := newBirthDate.age(); age < minimumDriverAge {
		return nil, NewErrInvalidAge(age)
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
func (d *Driver) Age() int {
	return d.birthDate.age()
}

func (d *Driver) BirthDate() BirthDate {
	return d.birthDate
}

func (d *Driver) CNH() CNH {
	return d.cnh
}

func (d *Driver) CPF() CPF {
	return d.cpf
}

func (d *Driver) Gender() Gender {
	return d.gender
}

func (d *Driver) HasVehicle() bool {
	return d.hasVehicle
}

func (d *Driver) Name() Name {
	return d.name
}

// setters
func (d *Driver) SetCNH(cnh string) error {
	newCNH, err := NewCNH(cnh)
	if err != nil {
		return err
	}

	d.cnh = newCNH

	return nil
}

func (d *Driver) SetGender(gender string) error {
	newGender, err := NewGender(gender)
	if err != nil {
		return err
	}

	d.gender = newGender

	return nil
}

func (d *Driver) SetHasVehicle(hasVehicle bool) {
	d.hasVehicle = hasVehicle
}

func (d *Driver) SetName(name string) error {
	newName, err := NewName(name)
	if err != nil {
		return err
	}

	d.name = newName

	return nil
}
