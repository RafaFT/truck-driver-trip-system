package entity

import "fmt"

type ErrInvalidCPF struct {
	msg string
}

func newErrInvalidCPF(cpf string) ErrInvalidCPF {
	return ErrInvalidCPF{
		msg: fmt.Sprintf("CPF invalid. cpf=[%s]", cpf),
	}
}

func (e ErrInvalidCPF) Error() string {
	return e.msg
}
