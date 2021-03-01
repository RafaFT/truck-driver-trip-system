package entity

import (
	"fmt"
)

type ErrInvalidAge struct {
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

func newErrInvalidAge(age int) ErrInvalidAge {
	return ErrInvalidAge{
		msg: fmt.Sprintf("Age invalid. age=[%d]", age),
	}
}

func (e ErrInvalidAge) Error() string {
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
