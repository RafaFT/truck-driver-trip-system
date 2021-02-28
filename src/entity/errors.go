package entity

import "fmt"

type ErrInvalidCNH struct {
	msg string
}

type ErrInvalidCPF struct {
	msg string
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
