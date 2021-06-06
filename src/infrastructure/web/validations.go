package web

import (
	"net/http"
	"strings"
)

func checkJSONContentType(headers http.Header) HTTPError {
	contentType := headers["Content-Type"]

	if len(contentType) != 1 || strings.ToLower(contentType[0]) != "application/json" {
		return newErrUnsupportedMediaType([]string{"application/json"})
	}

	return nil
}
