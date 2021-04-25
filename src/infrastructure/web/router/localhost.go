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

	r.router.HandleFunc("/drivers", r.GetDriversRoute()).Methods(http.MethodGet)
	r.router.HandleFunc("/drivers", r.CreateDriverRoute()).Methods(http.MethodPost)
	r.router.HandleFunc("/drivers/{cpf:[0-9]+}", r.GetDriverByCPFRoute()).Methods(http.MethodGet)
	r.router.HandleFunc("/drivers/{cpf:[0-9]+}", r.DeleteDriverRoute()).Methods(http.MethodDelete)
	r.router.HandleFunc("/drivers/{cpf:[0-9]+}", r.UpdateDriverRoute()).Methods(http.MethodPatch)

	return r
}

func (router *localHostRouter) CreateDriverRoute() http.HandlerFunc {
	url := fmt.Sprintf("%s:%s/%s", router.baseURL, router.port, "drivers")

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), rest.URLKey("url"), url)

		w.Header().Set("Content-Type", "application/json")

		l := log.NewPrintLogger()
		p := presenter.NewCreateDriverPresenter()
		uc := usecase.NewCreateDriver(l, router.repo)
		c := rest.NewCreateDriverController(p, uc)

		c.ServeHTTP(w, r.WithContext(ctx))
	}
}

func (router *localHostRouter) DeleteDriverRoute() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cpf := mux.Vars(r)["cpf"]
		ctx := context.WithValue(r.Context(), rest.CPFKey("cpf"), cpf)

		w.Header().Set("Content-Type", "application/json")

		l := log.NewPrintLogger()
		p := presenter.NewDeleteDriverPresenter()
		uc := usecase.NewDeleteDriver(l, router.repo)
		c := rest.NewDeleteDriverByCPFController(p, uc)

		c.ServeHTTP(w, r.WithContext(ctx))
	}
}

func (router *localHostRouter) GetDriverByCPFRoute() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cpf := mux.Vars(r)["cpf"]
		ctx := context.WithValue(r.Context(), rest.CPFKey("cpf"), cpf)

		w.Header().Set("Content-Type", "application/json")

		l := log.NewPrintLogger()
		p := presenter.NewGetDriverByCPFPresenter()
		uc := usecase.NewGetDriverByCPF(l, router.repo)
		c := rest.NewGetDriverByCPFController(p, uc)

		c.ServeHTTP(w, r.WithContext(ctx))
	}
}

func (router *localHostRouter) GetDriversRoute() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		l := log.NewPrintLogger()
		p := presenter.NewGetDriversPresenter()
		uc := usecase.NewGetDrivers(l, router.repo)
		c := rest.NewGetDriversController(p, uc)

		c.ServeHTTP(w, r)
	}
}

func (router *localHostRouter) UpdateDriverRoute() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cpf := mux.Vars(r)["cpf"]
		ctx := context.WithValue(r.Context(), rest.CPFKey("cpf"), cpf)

		w.Header().Set("Content-Type", "application/json")

		l := log.NewPrintLogger()
		p := presenter.NewUpdateDriverPresenter()
		uc := usecase.NewUpdateDriver(l, router.repo)
		c := rest.NewUpdateDriverController(p, uc)

		c.ServeHTTP(w, r.WithContext(ctx))
	}
}

func (router *localHostRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router.router.ServeHTTP(w, r)
}
