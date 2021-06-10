package router

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"github.com/rafaft/truck-driver-trip-system/adapter/controller/rest"
	"github.com/rafaft/truck-driver-trip-system/adapter/presenter"
	"github.com/rafaft/truck-driver-trip-system/entity"
	"github.com/rafaft/truck-driver-trip-system/infrastructure/log"
	"github.com/rafaft/truck-driver-trip-system/usecase"
)

type cloudRunRouter struct {
	port      string
	projectID string
	repo      entity.DriverRepository
	router    *mux.Router
}

func NewDriverCloudRun(port, projectID string, repo entity.DriverRepository) http.Handler {
	r := &cloudRunRouter{
		port:      port,
		projectID: projectID,
		repo:      repo,
		router:    mux.NewRouter(),
	}

	driverSubRoute := r.router.PathPrefix("/drivers").Subrouter()
	driverSubRoute.HandleFunc("", r.GetDriversRoute()).Methods(http.MethodGet)
	driverSubRoute.HandleFunc("", r.CreateDriverRoute()).Methods(http.MethodPost)
	driverSubRoute.MethodNotAllowedHandler = MethodNotAllowedHandler(http.MethodGet, http.MethodPost)

	driversCPFSubRoute := r.router.PathPrefix("/drivers/{cpf:[0-9]+}").Subrouter()
	driversCPFSubRoute.HandleFunc("", r.GetDriverByCPFRoute()).Methods(http.MethodGet)
	driversCPFSubRoute.HandleFunc("", r.DeleteDriverRoute()).Methods(http.MethodDelete)
	driversCPFSubRoute.HandleFunc("", r.UpdateDriverRoute()).Methods(http.MethodPatch)
	driversCPFSubRoute.MethodNotAllowedHandler = MethodNotAllowedHandler(http.MethodGet, http.MethodDelete, http.MethodPatch)

	return r
}

func (router *cloudRunRouter) CreateDriverRoute() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), rest.URLKey("url"), getURI(router.port, r))

		w.Header().Set("Content-Type", "application/json")

		l := log.NewCloudRunLogger(router.projectID, getGCPTrace(r))
		p := presenter.NewCreateDriver()
		uc := usecase.NewCreateDriver(l, router.repo)
		c := rest.NewCreateDriver(p, uc)

		c.ServeHTTP(w, r.WithContext(ctx))
	}
}

func (router *cloudRunRouter) DeleteDriverRoute() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), rest.CPFKey("cpf"), mux.Vars(r)["cpf"])

		w.Header().Set("Content-Type", "application/json")

		l := log.NewCloudRunLogger(router.projectID, getGCPTrace(r))
		p := presenter.NewDeleteDriver()
		uc := usecase.NewDeleteDriver(l, router.repo)
		c := rest.NewDeleteDriverByCPF(p, uc)

		c.ServeHTTP(w, r.WithContext(ctx))
	}
}

func (router *cloudRunRouter) GetDriverByCPFRoute() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), rest.CPFKey("cpf"), mux.Vars(r)["cpf"])

		w.Header().Set("Content-Type", "application/json")

		l := log.NewCloudRunLogger(router.projectID, getGCPTrace(r))
		p := presenter.NewGetDriverByCPF()
		uc := usecase.NewGetDriverByCPF(l, router.repo)
		c := rest.NewGetDriverByCPF(p, uc)

		c.ServeHTTP(w, r.WithContext(ctx))
	}
}

func (router *cloudRunRouter) GetDriversRoute() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		l := log.NewCloudRunLogger(router.projectID, getGCPTrace(r))
		p := presenter.NewGetDrivers()
		uc := usecase.NewGetDrivers(l, router.repo)
		c := rest.NewGetDrivers(p, uc)

		c.ServeHTTP(w, r)
	}
}

func (router *cloudRunRouter) UpdateDriverRoute() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), rest.CPFKey("cpf"), mux.Vars(r)["cpf"])

		w.Header().Set("Content-Type", "application/json")

		l := log.NewCloudRunLogger(router.projectID, getGCPTrace(r))
		p := presenter.NewUpdateDriver()
		uc := usecase.NewUpdateDriver(l, router.repo)
		c := rest.NewUpdateDriver(p, uc)

		c.ServeHTTP(w, r.WithContext(ctx))
	}
}

func (router *cloudRunRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router.router.ServeHTTP(w, r)
}

func getGCPTrace(r *http.Request) string {
	var trace string

	traceHeader := r.Header.Get("X-Cloud-Trace-Context")
	traceParts := strings.Split(traceHeader, "/")
	if len(traceParts) > 0 && len(traceParts[0]) > 0 {
		trace = traceParts[0]
	}

	return trace
}

func getURI(port string, r *http.Request) string {
	return fmt.Sprintf("https://%s:%s%s", r.Host, port, r.URL.Path)
}
