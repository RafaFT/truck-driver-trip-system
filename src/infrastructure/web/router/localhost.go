package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/rafaft/truck-driver-trip-system/adapter/controller/rest"
	"github.com/rafaft/truck-driver-trip-system/adapter/presenter"
	"github.com/rafaft/truck-driver-trip-system/entity"
	"github.com/rafaft/truck-driver-trip-system/infrastructure/logger"
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

	return r
}

func (router *localHostRouter) CreateDriverRoute() http.HandlerFunc {
	url := fmt.Sprintf("%s:%s/%s", router.baseURL, router.port, "drivers")

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		l := logger.NewPrintLogger()
		p := presenter.NewCreateDriverPresenter()
		uc := usecase.NewCreateDriverInteractor(l, router.repo)
		c := rest.NewCreateDriverController(p, url, uc)

		c.ServeHTTP(w, r)
	}
}

func (router *localHostRouter) GetDriversRoute() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		l := logger.NewPrintLogger()
		p := presenter.NewGetDriversPresenter()
		uc := usecase.NewGetDriversInteractor(l, router.repo)
		c := rest.NewGetDriversController(p, uc)

		c.ServeHTTP(w, r)
	}
}

func (router *localHostRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router.router.ServeHTTP(w, r)
}
