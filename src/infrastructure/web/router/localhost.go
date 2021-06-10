package router

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/rafaft/truck-driver-trip-system/adapter/controller/rest"
	"github.com/rafaft/truck-driver-trip-system/adapter/presenter"
	"github.com/rafaft/truck-driver-trip-system/entity"
	"github.com/rafaft/truck-driver-trip-system/infrastructure/log"
	"github.com/rafaft/truck-driver-trip-system/usecase"
)

type localHostRouter struct {
	baseURL string
	port    string
	repo    entity.DriverRepository
	router  *mux.Router
}

func NewDriverLocalHost(port string, repo entity.DriverRepository) http.Handler {
	r := &localHostRouter{
		baseURL: "http://localhost",
		port:    port,
		repo:    repo,
		router:  mux.NewRouter(),
	}

	// drivers sub-router
	driverSubRoute := r.router.PathPrefix("/drivers").Subrouter()
	driverSubRoute.MethodNotAllowedHandler = MethodNotAllowedHandler(http.MethodGet, http.MethodPost)

	// drivers get
	driverSubRoute.HandleFunc("", r.GetDriversRoute()).Methods(http.MethodGet)

	// driver post
	driverSubRoutePost := driverSubRoute.Methods(http.MethodPost).Subrouter()
	driverSubRoutePost.HandleFunc("", r.CreateDriverRoute())
	driverSubRoutePost.Use(UnsupportedMediaTypeJSON)

	// drivers by cpf sub-routers
	driversCPFSubRoute := r.router.PathPrefix("/drivers/{cpf:[0-9]+}").Subrouter()
	driversCPFSubRoute.MethodNotAllowedHandler = MethodNotAllowedHandler(http.MethodGet, http.MethodDelete, http.MethodPatch)

	// drivers by cpf get, delete
	driversCPFSubRoute.HandleFunc("", r.GetDriverByCPFRoute()).Methods(http.MethodGet)
	driversCPFSubRoute.HandleFunc("", r.DeleteDriverRoute()).Methods(http.MethodDelete)

	// drivers by cpf patch
	driversCPFSubRoutePatch := driversCPFSubRoute.Methods(http.MethodPatch).Subrouter()
	driversCPFSubRoutePatch.HandleFunc("", r.UpdateDriverRoute())
	driversCPFSubRoutePatch.Use(UnsupportedMediaTypeJSON)

	return r
}

func (router *localHostRouter) CreateDriverRoute() http.HandlerFunc {
	url := fmt.Sprintf("%s:%s/%s", router.baseURL, router.port, "drivers")

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), rest.URLKey("url"), url)

		w.Header().Set("Content-Type", "application/json")

		l := log.NewPrintLogger()
		p := presenter.NewCreateDriver()
		uc := usecase.NewCreateDriver(l, router.repo)
		c := rest.NewCreateDriver(p, uc)

		c.ServeHTTP(w, r.WithContext(ctx))
	}
}

func (router *localHostRouter) DeleteDriverRoute() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cpf := mux.Vars(r)["cpf"]
		ctx := context.WithValue(r.Context(), rest.CPFKey("cpf"), cpf)

		w.Header().Set("Content-Type", "application/json")

		l := log.NewPrintLogger()
		p := presenter.NewDeleteDriver()
		uc := usecase.NewDeleteDriver(l, router.repo)
		c := rest.NewDeleteDriverByCPF(p, uc)

		c.ServeHTTP(w, r.WithContext(ctx))
	}
}

func (router *localHostRouter) GetDriverByCPFRoute() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cpf := mux.Vars(r)["cpf"]
		ctx := context.WithValue(r.Context(), rest.CPFKey("cpf"), cpf)

		w.Header().Set("Content-Type", "application/json")

		l := log.NewPrintLogger()
		p := presenter.NewGetDriverByCPF()
		uc := usecase.NewGetDriverByCPF(l, router.repo)
		c := rest.NewGetDriverByCPF(p, uc)

		c.ServeHTTP(w, r.WithContext(ctx))
	}
}

func (router *localHostRouter) GetDriversRoute() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		l := log.NewPrintLogger()
		p := presenter.NewGetDrivers()
		uc := usecase.NewGetDrivers(l, router.repo)
		c := rest.NewGetDrivers(p, uc)

		c.ServeHTTP(w, r)
	}
}

func (router *localHostRouter) UpdateDriverRoute() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cpf := mux.Vars(r)["cpf"]
		ctx := context.WithValue(r.Context(), rest.CPFKey("cpf"), cpf)

		w.Header().Set("Content-Type", "application/json")

		l := log.NewPrintLogger()
		p := presenter.NewUpdateDriver()
		uc := usecase.NewUpdateDriver(l, router.repo)
		c := rest.NewUpdateDriver(p, uc)

		c.ServeHTTP(w, r.WithContext(ctx))
	}
}

func (router *localHostRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router.router.ServeHTTP(w, r)
}
