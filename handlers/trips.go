package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/rafaft/truck-pad/models"
)

func GetAllTrips(client *firestore.Client) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		r.ParseForm()

		q := createTripsQuery(client, r)
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
}
