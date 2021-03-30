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
	"github.com/rafaft/truck-driver-trip-system/infrastructure/logger"
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

	r.router.HandleFunc("/drivers", r.GetDriversRoute()).Methods(http.MethodGet)
	r.router.HandleFunc("/drivers", r.CreateDriverRoute()).Methods(http.MethodPost)
	r.router.HandleFunc("/drivers/{cpf:[0-9]+}", r.GetDriverByCPFRoute()).Methods(http.MethodGet)
	r.router.HandleFunc("/drivers/{cpf:[0-9]+}", r.DeleteDriverRoute()).Methods(http.MethodDelete)
	r.router.HandleFunc("/drivers/{cpf:[0-9]+}", r.UpdateDriverRoute()).Methods(http.MethodPatch)

	return r
}

func (router *cloudRunRouter) CreateDriverRoute() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), rest.URLKey("url"), getURI(router.port, r))

		w.Header().Set("Content-Type", "application/json")

		l := logger.NewCloudRunLogger(router.projectID, getGCPTrace(r))
		p := presenter.NewCreateDriverPresenter()
		uc := usecase.NewCreateDriverInteractor(l, router.repo)
		c := rest.NewCreateDriverController(p, uc)

		c.ServeHTTP(w, r.WithContext(ctx))
	}
}

func (router *cloudRunRouter) DeleteDriverRoute() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), rest.CPFKey("cpf"), mux.Vars(r)["cpf"])

		w.Header().Set("Content-Type", "application/json")

		l := logger.NewCloudRunLogger(router.projectID, getGCPTrace(r))
		p := presenter.NewDeleteDriverPresenter()
		uc := usecase.NewDeleteDriverInteractor(l, router.repo)
		c := rest.NewDeleteDriverByCPFController(p, uc)

		c.ServeHTTP(w, r.WithContext(ctx))
	}
}

func (router *cloudRunRouter) GetDriverByCPFRoute() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), rest.CPFKey("cpf"), mux.Vars(r)["cpf"])

		w.Header().Set("Content-Type", "application/json")

		l := logger.NewCloudRunLogger(router.projectID, getGCPTrace(r))
		p := presenter.NewGetDriverByCPFPresenter()
		uc := usecase.NewGetDriverByCPFInteractor(l, router.repo)
		c := rest.NewGetDriverByCPFController(p, uc)

		c.ServeHTTP(w, r.WithContext(ctx))
	}
}

func (router *cloudRunRouter) GetDriversRoute() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		l := logger.NewCloudRunLogger(router.projectID, getGCPTrace(r))
		p := presenter.NewGetDriversPresenter()
		uc := usecase.NewGetDriversInteractor(l, router.repo)
		c := rest.NewGetDriversController(p, uc)

		c.ServeHTTP(w, r)
	}
}

func (router *cloudRunRouter) UpdateDriverRoute() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), rest.CPFKey("cpf"), mux.Vars(r)["cpf"])

		w.Header().Set("Content-Type", "application/json")

		l := logger.NewCloudRunLogger(router.projectID, getGCPTrace(r))
		p := presenter.NewUpdateDriverPresenter()
		uc := usecase.NewUpdateDriverInteractor(l, router.repo)
		c := rest.NewUpdateDriverController(p, uc)

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
