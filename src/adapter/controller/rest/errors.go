package rest

import (
	"fmt"
	"strings"
)

var ErrExpectedJSONObject = fmt.Errorf("Expected JSON Object.")
var ErrInternalServerError = fmt.Errorf("Internal Server Error.")
var ErrInvalidBody = fmt.Errorf("Could not read HTTP request body.")
var ErrInvalidJSON = fmt.Errorf("Invalid JSON.")
var ErrInvalidQueryString = fmt.Errorf("Could not read URL query string.")

type ErrInvalidJSONFieldType struct {
	msg string
}

type ErrInvalidParameterType struct {
	msg string
}

type ErrMissingJSONFields struct {
	msg string
}

type ErrUnexpectedJSONField struct {
	msg string
}

type ErrUnexpectedParameter struct {
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

func newErrInvalidParameterType(param, expectedType string) error {
	switch expectedType {
	case "uint":
		expectedType = "positive integer"
	}

	return ErrInvalidParameterType{
		msg: fmt.Sprintf("Invalid value at '%s'. Expected %s.", param, expectedType),
	}
}

func (e ErrInvalidParameterType) Error() string {
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

func newErrUnexpectedJSONField(field string) error {
	return ErrUnexpectedJSONField{
		msg: fmt.Sprintf("Unexpected JSON field: '%s'.", field),
	}
}

func (e ErrUnexpectedJSONField) Error() string {
	return e.msg
}

func newErrUnexpectedParameter(param string) error {
	return ErrUnexpectedParameter{
		msg: fmt.Sprintf("Unexpected query parameter: %s.", param),
	}
}

func (e ErrUnexpectedParameter) Error() string {
	return e.msg
}
