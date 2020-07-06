package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rafaft/truck-pad/models"
)

func getAllTrips(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	r.ParseForm()

	q := createTripsQuery(r)
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
