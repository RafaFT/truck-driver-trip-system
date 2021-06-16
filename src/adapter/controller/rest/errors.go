package rest

import (
	"fmt"
	"reflect"
)

var ErrExpectedJSONObject = fmt.Errorf("Expected JSON Object.")
var ErrInternalServerError = fmt.Errorf("Internal Server Error.")
var ErrInvalidJSON = fmt.Errorf("Invalid JSON.")

type ErrInvalidBody struct {
	msg string
}

type ErrInvalidJSONFieldType struct {
	msg string
}

type ErrInvalidParameterValue struct {
	msg string
}

type ErrParseQueryString struct {
	msg string
}

type ErrUnknownParameter struct {
	msg string
}

func newErrInvalidBody() error {
	return ErrInvalidBody{
		msg: "Could not read HTTP request body.",
	}
}

func (e ErrInvalidBody) Error() string {
	return e.msg
}

func newErrInvalidJSONFieldType(jsonField, expectedType, gotType string) error {
	switch expectedType {
	case "ISO8601Date":
		expectedType = "string (ISO-8601 date as 'YYYY-MM-DD')"
	}

	switch gotType {
	case "int", "float64":
		gotType = "number"
	}

	return ErrInvalidJSONFieldType{
		msg: fmt.Sprintf("Invalid value at '%s'. Expected %s, got %s.", jsonField, expectedType, gotType),
	}
}

func (e ErrInvalidJSONFieldType) Error() string {
	return e.msg
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
		msg: fmt.Sprintf("Invalid value at '%s'. Expected %s, got %s.", p, tName, v),
	}
}

func (e ErrInvalidParameterValue) Error() string {
	return e.msg
}

func newErrParseQueryString(queryString string) error {
	return ErrParseQueryString{
		msg: fmt.Sprintf("Could not parse query string: %s.", queryString),
	}
}

func (e ErrParseQueryString) Error() string {
	return e.msg
}

func newErrUnknownParameter(p string) error {
	return ErrUnknownParameter{
		msg: fmt.Sprintf("Unknown query parameter: %s.", p),
	}
}

func (e ErrUnknownParameter) Error() string {
	return e.msg
}
