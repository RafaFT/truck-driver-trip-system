package entity

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

var (
	ErrInvalidBirthDate = errors.New("Invalid birth date")
	ErrInvalidCPF       = errors.New("Invalid CPF")
	ErrInvalidCNHType   = errors.New("Invalid cnhType")
	ErrInvalidGender    = errors.New("Invalid gender")
	ErrInvalidName      = errors.New("Invalid name")
)

type TruckDriver struct {
	birthDate  BirthDate
	cnhType    CNHType
	cpf        CPF
	gender     Gender
	hasVehicle bool
	name       Name
}

func NewTruckDriver(cpf, name, gender, cnhType string, birthDate time.Time, hasVehicle bool) (*TruckDriver, error) {
	cpfValue, err := NewCPF(cpf)
	if err != nil {
		return nil, err
	}

	nameValue, err := NewName(name)
	if err != nil {
		return nil, err
	}

	genderValue, err := NewGender(gender)
	if err != nil {
		return nil, err
	}

	cnhTypeValue, err := NewCNHType(cnhType)
	if err != nil {
		return nil, err
	}

	birthDateValue, err := NewBirthDate(birthDate)
	if err != nil {
		return nil, err
	}

	driver := TruckDriver{
		cpf:        cpfValue,
		name:       nameValue,
		gender:     genderValue,
		cnhType:    cnhTypeValue,
		birthDate:  birthDateValue,
		hasVehicle: hasVehicle,
	}

	return &driver, nil
}

func (td *TruckDriver) Age() int {
	return calculateAge(time.Now(), time.Time(td.birthDate))
}

func (td *TruckDriver) BirthDate() BirthDate {
	return td.birthDate
}

func (td *TruckDriver) CNHType() CNHType {
	return td.cnhType
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
