package entity

import (
	"fmt"
	"time"
)

type ErrDriverAlreadyExists struct {
	msg string
}

type ErrDriverNotFound struct {
	msg string
}

type ErrInvalidAge struct {
	msg string
}

type ErrInvalidBirthDate struct {
	msg string
}

type ErrInvalidCNH struct {
	msg string
}

type ErrInvalidCPF struct {
	msg string
}

type ErrInvalidGender struct {
	msg string
}

type ErrInvalidLatitude struct {
	msg string
}

type ErrInvalidLongitude struct {
	msg string
}

type ErrInvalidName struct {
	msg string
}

type ErrInvalidTripEndDate struct {
	msg string
}

type ErrInvalidTripStartDate struct {
	msg string
}

type ErrInvalidVehicleCode struct {
	msg string
}

func NewErrDriverAlreadyExists(cpf CPF) ErrDriverAlreadyExists {
	return ErrDriverAlreadyExists{
		msg: fmt.Sprintf("Driver already exists. cpf=[%s]", cpf),
	}
}

func (e ErrDriverAlreadyExists) Error() string {
	return e.msg
}

func NewErrDriverNotFound(cpf CPF) ErrDriverNotFound {
	return ErrDriverNotFound{
		msg: fmt.Sprintf("Driver not found. cpf=[%s]", cpf),
	}
}

func (e ErrDriverNotFound) Error() string {
	return e.msg
}

func newErrInvalidAge(age int) ErrInvalidAge {
	return ErrInvalidAge{
		msg: fmt.Sprintf("Age invalid. age=[%d]", age),
	}
}

func (e ErrInvalidAge) Error() string {
	return e.msg
}

func newErrInvalidBirthDate(birthDate time.Time) ErrInvalidBirthDate {
	return ErrInvalidBirthDate{
		msg: fmt.Sprintf("Birth Date invalid. birthDate=[%v]", birthDate),
	}
}

func (e ErrInvalidBirthDate) Error() string {
	return e.msg
}

func newErrInvalidCNH(cnh string) ErrInvalidCNH {
	return ErrInvalidCNH{
		msg: fmt.Sprintf("CNH invalid. cnh=[%s]", cnh),
	}
}

func (e ErrInvalidCNH) Error() string {
	return e.msg
}

func newErrInvalidCPF(cpf string) ErrInvalidCPF {
	return ErrInvalidCPF{
		msg: fmt.Sprintf("CPF invalid. cpf=[%s]", cpf),
	}
}

func (e ErrInvalidCPF) Error() string {
	return e.msg
}

func newErrInvalidGender(gender string) ErrInvalidGender {
	return ErrInvalidGender{
		msg: fmt.Sprintf("Gender invalid. gender=[%s]", gender),
	}
}

func (e ErrInvalidGender) Error() string {
	return e.msg
}

func newErrInvalidLatitude(lat float64) ErrInvalidLatitude {
	return ErrInvalidLatitude{
		msg: fmt.Sprintf("Latitude out of range. latitude=[%f]", lat),
	}
}

func (e ErrInvalidLatitude) Error() string {
	return e.msg
}

func newErrInvalidLongitude(lat float64) ErrInvalidLongitude {
	return ErrInvalidLongitude{
		msg: fmt.Sprintf("Longitude out of range. longitude=[%f]", lat),
	}
}

func (e ErrInvalidLongitude) Error() string {
	return e.msg
}

func newErrInvalidName(name string) ErrInvalidName {
	return ErrInvalidName{
		msg: fmt.Sprintf("Name invalid. name=[%s]", name),
	}
}

func (e ErrInvalidName) Error() string {
	return e.msg
}

func newErrInvalidVehicleCode(vehicleCode int) ErrInvalidVehicleCode {
	return ErrInvalidVehicleCode{
		msg: fmt.Sprintf("Vehicle code invalid. code=[%d]", vehicleCode),
	}
}

func (e ErrInvalidVehicleCode) Error() string {
	return e.msg
}

func newErrInvalidTripEndDate(endDate time.Time) ErrInvalidTripEndDate {
	return ErrInvalidTripEndDate{
		msg: fmt.Sprintf("Trip end date invalid. date=[%s]", endDate),
	}
}

func (e ErrInvalidTripEndDate) Error() string {
	return e.msg
}

func newErrInvalidTripStartDate(startDate time.Time) ErrInvalidTripStartDate {
	return ErrInvalidTripStartDate{
		msg: fmt.Sprintf("Trip start date invalid. date=[%s]", startDate),
	}
}

func (e ErrInvalidTripStartDate) Error() string {
	return e.msg
}
