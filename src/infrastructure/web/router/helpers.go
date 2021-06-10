package router

import (
	"net/http"
	"strings"
)

func MethodNotAllowedHandler(allowedMethods ...string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Allow", strings.Join(allowedMethods, ", "))
		w.WriteHeader(http.StatusMethodNotAllowed)
	})
}

func UnsupportedMediaTypeJSON(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mediaTypes := r.Header["Content-Type"]
		if len(mediaTypes) != 1 || !strings.HasPrefix(strings.ToLower(mediaTypes[0]), "application/json") {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			return
		}

		h.ServeHTTP(w, r)
	})
}
