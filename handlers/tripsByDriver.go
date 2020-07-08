package handlers

import (
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/gorilla/mux"
)

func GetTripsByDriver(client *firestore.Client) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		r.ParseForm()

		cpf := mux.Vars(r)["cpf"]
		r.Form.Set("driver_id", cpf)

		q := createTripsQuery(client, r)

		getTrips(w, r, q)
	}
}

func GetTripsByDriverByYear(client *firestore.Client) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		r.ParseForm()

		cpf := mux.Vars(r)["cpf"]
		r.Form.Set("driver_id", cpf)
		setFilterByYear(r)

		q := createTripsQuery(client, r)

		getTrips(w, r, q)
	}
}
