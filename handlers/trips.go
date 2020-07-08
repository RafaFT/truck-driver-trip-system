package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/rafaft/truck-pad/models"
	"google.golang.org/api/iterator"
)

func AddTrip(client *firestore.Client) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(createErrorJSON(err))
			return
		}

		trip, err := models.NewTrip(body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(createErrorJSON(err))
			return
		}

		ctx := r.Context()
		collection := client.Collection("drivers").Doc(string(*trip.DriverID)).Collection("trips")
		_, err = collection.Where("id", "==", trip.ID).Documents(ctx).Next()
		if err != iterator.Done {
			w.WriteHeader(http.StatusConflict)
			w.Write(createErrorJSON(fmt.Errorf(
				"there is already a trip with the same timestamp under driver=%s", *trip.DriverID),
			))
			return
		}

		_, _, err = collection.Add(ctx, &trip)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(createErrorJSON(fmt.Errorf("internal server error")))
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func getTrips(w http.ResponseWriter, r *http.Request, q firestore.Query) {
	docs, err := q.Documents(r.Context()).GetAll()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(createErrorJSON(fmt.Errorf("internal server error")))
		return
	}

	result := make([]*models.Trip, len(docs))
	for i, docSnapShot := range docs {
		var trip models.Trip
		err = docSnapShot.DataTo(&trip)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(createErrorJSON(fmt.Errorf("internal server error")))
			return
		}

		result[i] = &trip
	}

	b, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(createErrorJSON(fmt.Errorf("internal server error")))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func GetAllTrips(client *firestore.Client) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		r.ParseForm()

		q := createTripsQuery(client, r)

		getTrips(w, r, q)
	}
}

func GetTripsByDay(client *firestore.Client) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		r.ParseForm()

		err := setFilterByDay(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(createErrorJSON(err))
			return
		}

		q := createTripsQuery(client, r)

		getTrips(w, r, q)
	}
}

func GetTripsByMonth(client *firestore.Client) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		r.ParseForm()

		err := setFilterByMonth(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(createErrorJSON(err))
			return
		}

		q := createTripsQuery(client, r)

		getTrips(w, r, q)
	}
}

func GetTripsByYear(client *firestore.Client) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		r.ParseForm()

		setFilterByYear(r)

		q := createTripsQuery(client, r)

		getTrips(w, r, q)
	}
}
