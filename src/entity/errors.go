package entity

import (
	"fmt"
	"time"
)

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

type ErrInvalidName struct {
	msg string
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

func newErrInvalidName(name string) ErrInvalidName {
	return ErrInvalidName{
		msg: fmt.Sprintf("Name invalid. name=[%s]", name),
	}
}

func (e ErrInvalidName) Error() string {
	return e.msg
}
