package rest

import (
	"fmt"
	"reflect"
)

type ErrInvalidParameterValue struct {
	msg string
}

type ErrParseQueryString struct {
	msg string
}

type ErrUnknownParameter struct {
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

func newErrParseQueryString(queryString string) error {
	return ErrParseQueryString{
		msg: fmt.Sprintf("could not parse query string: %s", queryString),
	}
}

func newErrUnknownParameter(p string) error {
	return ErrUnknownParameter{
		msg: fmt.Sprintf("unknown query parameter: %s", p),
	}
}

func (e ErrInvalidParameterValue) Error() string {
	return e.msg
}

func (e ErrParseQueryString) Error() string {
	return e.msg
}

func (e ErrUnknownParameter) Error() string {
	return e.msg
}
