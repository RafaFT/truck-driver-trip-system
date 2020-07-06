package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/gorilla/mux"

	"github.com/rafaft/truck-pad/handlers"
)

var ctx context.Context
var client *firestore.Client
var router *mux.Router

func init() {
	var err error

	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "firestore-credentials.json")

	ctx = context.Background()
	client, err = firestore.NewClient(ctx, "truck-pad")
	if err != nil {
		panic(err)
	}

	router = mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Docs..."))
	})

	// route for drivers
	router.HandleFunc("/drivers", handlers.getAllDrivers).Methods("GET")
	router.HandleFunc("/drivers", handlers.addDriver).Methods("POST")
	router.HandleFunc(`/drivers/{cpf:\d{11}}`, handlers.getDriver).Methods("GET")
	router.HandleFunc(`/drivers/{cpf:\d{11}}`, handlers.updateDriver).Methods("PATCH")

	// route for trips
	router.HandleFunc("/trips", handlers.getAllTrips).Methods("GET")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal(err)
	}
}
