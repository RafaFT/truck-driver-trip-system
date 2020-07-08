package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/gorilla/mux"
	"github.com/rafaft/truck-pad/models"
	"google.golang.org/api/iterator"
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

func GetTripsByDriverByMonth(client *firestore.Client) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		r.ParseForm()

		cpf := mux.Vars(r)["cpf"]
		r.Form.Set("driver_id", cpf)
		setFilterByMonth(r)

		q := createTripsQuery(client, r)

		getTrips(w, r, q)
	}
}

func GetTripsByDriverByDay(client *firestore.Client) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		r.ParseForm()

		cpf := mux.Vars(r)["cpf"]
		r.Form.Set("driver_id", cpf)
		setFilterByDay(r)

		q := createTripsQuery(client, r)

		getTrips(w, r, q)
	}
}

func GetLatestTrip(client *firestore.Client) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		r.ParseForm()

		r.Form.Del("id")
		r.Form.Del("has_load")
		r.Form.Del("vehicle_type")
		r.Form.Del("start_date")
		r.Form.Del("end_date")

		cpf := mux.Vars(r)["cpf"]
		r.Form.Set("driver_id", cpf)
		r.Form.Set("order", "desc")
		r.Form.Set("limit", "1")

		q := createTripsQuery(client, r)
		doc, err := q.Documents(r.Context()).Next()
		if err == iterator.Done {
			w.WriteHeader(http.StatusNotFound)
			w.Write(createErrorJSON(fmt.Errorf("no trip found for driver=%s", cpf)))
		}

		var trip models.Trip
		err = doc.DataTo(&trip)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(createErrorJSON(fmt.Errorf("internal server error")))
			return
		}

		b, err := json.Marshal(&trip)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(createErrorJSON(fmt.Errorf("internal server error")))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}
