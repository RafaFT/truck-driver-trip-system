package rest

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type DriversRouter interface {
	// http.Handler interface
	ServeHTTP(http.ResponseWriter, *http.Request)

	// Allows use of helper Setter functions for routers config
	MuxRouter() *mux.Router

	// Drivers related routes
	CreateDriverRoute() http.HandlerFunc
	DeleteDriverRoute() http.HandlerFunc
	GetDriverByCPFRoute() http.HandlerFunc
	GetDriversRoute() http.HandlerFunc
	UpdateDriverRoute() http.HandlerFunc
}

func SetDriversRoutes(drivers DriversRouter) {
	r := drivers.MuxRouter()

	// drivers sub-router
	driverSubRoute := r.PathPrefix("/drivers").Subrouter()
	driverSubRoute.MethodNotAllowedHandler = methodNotAllowedHandler(http.MethodGet, http.MethodPost)

	// drivers get
	driverSubRoute.HandleFunc("", drivers.GetDriversRoute()).Methods(http.MethodGet)

	// driver post
	driverSubRoutePost := driverSubRoute.Methods(http.MethodPost).Subrouter()
	driverSubRoutePost.HandleFunc("", drivers.CreateDriverRoute())
	driverSubRoutePost.Use(checkPayloadSize)
	driverSubRoutePost.Use(unsupportedMediaTypeJSON)

	// drivers by cpf sub-routers
	driversCPFSubRoute := r.PathPrefix("/drivers/{cpf:[0-9]+}").Subrouter()
	driversCPFSubRoute.MethodNotAllowedHandler = methodNotAllowedHandler(http.MethodGet, http.MethodDelete, http.MethodPatch)

	// drivers by cpf get, delete
	driversCPFSubRoute.HandleFunc("", drivers.GetDriverByCPFRoute()).Methods(http.MethodGet)
	driversCPFSubRoute.HandleFunc("", drivers.DeleteDriverRoute()).Methods(http.MethodDelete)

	// drivers by cpf patch
	driversCPFSubRoutePatch := driversCPFSubRoute.Methods(http.MethodPatch).Subrouter()
	driversCPFSubRoutePatch.HandleFunc("", drivers.UpdateDriverRoute())
	driversCPFSubRoutePatch.Use(checkPayloadSize)
	driversCPFSubRoutePatch.Use(unsupportedMediaTypeJSON)
}

func checkPayloadSize(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contentLength, exists := r.Header["Content-Length"]

		if !exists || len(contentLength) != 1 {
			w.WriteHeader(http.StatusLengthRequired)
			return
		}

		if length, err := strconv.Atoi(contentLength[0]); err != nil {
			w.WriteHeader(http.StatusLengthRequired)
			return
		} else if length > 1_000 {
			w.WriteHeader(http.StatusRequestEntityTooLarge)
			return
		}

		h.ServeHTTP(w, r)
	})
}

func methodNotAllowedHandler(allowedMethods ...string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Allow", strings.Join(allowedMethods, ", "))
		w.WriteHeader(http.StatusMethodNotAllowed)
	})
}

func unsupportedMediaTypeJSON(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mediaTypes := r.Header["Content-Type"]
		if len(mediaTypes) != 1 || !strings.HasPrefix(strings.ToLower(mediaTypes[0]), "application/json") {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			return
		}

		h.ServeHTTP(w, r)
	})
}
