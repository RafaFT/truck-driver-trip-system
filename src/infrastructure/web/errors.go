package web

import (
	"fmt"
	"net/http"
	"strings"
)

type HTTPError interface {
	Code() int
	Error() string
}

type ErrUnsupportedMediaType struct {
	code int
	msg  string
}

func newErrUnsupportedMediaType(supportedMediaTypes []string) HTTPError {
	return ErrUnsupportedMediaType{
		code: http.StatusUnsupportedMediaType,
		msg:  fmt.Sprintf("Supported Media Type(s): %s", strings.Join(supportedMediaTypes, ", ")),
	}
}

func (e ErrUnsupportedMediaType) Code() int {
	return e.code
}

func (e ErrUnsupportedMediaType) Error() string {
	return e.msg
}
