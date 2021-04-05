package rest

import (
	"fmt"
	"reflect"
)

type ErrInvalidParameterValue struct {
	msg string
}

func newErrInvalidParameterValue(p, v string, t reflect.Type) error {
	var tName string

	switch t.Name() {
	case "bool":
		tName = "BOOLEAN"
	case "int":
		tName = "INTEGER"
	case "uint":
		tName = "POSITIVE_INTEGER"
	case "CNH":
		tName = "CNH"
	case "Gender":
		tName = "GENDER"
	}

	return ErrInvalidParameterValue{
		msg: fmt.Sprintf("invalid value at '%s' (type: %s), got %s", p, tName, v),
	}
}

func (e ErrInvalidParameterValue) Error() string {
	return e.msg
}
