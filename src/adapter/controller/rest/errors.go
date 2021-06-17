package rest

import (
	"fmt"
	"reflect"
	"strings"
)

var ErrExpectedJSONObject = fmt.Errorf("Expected JSON Object.")
var ErrInternalServerError = fmt.Errorf("Internal Server Error.")
var ErrInvalidBody = fmt.Errorf("Could not read HTTP request body.")
var ErrInvalidJSON = fmt.Errorf("Invalid JSON.")

type ErrInvalidJSONFieldType struct {
	msg string
}

type ErrInvalidParameterValue struct {
	msg string
}

type ErrMissingJSONFields struct {
	msg string
}

type ErrParseQueryString struct {
	msg string
}

type ErrUnknownParameter struct {
	msg string
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

func newErrMissingJSONFields(fieldTypes [][2]string) error {
	s := make([]string, 0, len(fieldTypes))

	for _, fieldAndType := range fieldTypes {
		field := fieldAndType[0]
		type_ := fieldAndType[1]

		switch type_ {
		case "ISO8601Date":
			type_ = "string (ISO-8601 date as 'YYYY-MM-DD')"
		}

		s = append(s, fmt.Sprintf("[%s: %s]", field, type_))
	}

	return ErrMissingJSONFields{
		msg: fmt.Sprintf("%s.", fmt.Sprintf("Missing JSON fields. %s", strings.Join(s, ", "))),
	}
}

func (e ErrMissingJSONFields) Error() string {
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
