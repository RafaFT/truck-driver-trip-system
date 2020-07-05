package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/gorilla/mux"
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

	router.HandleFunc("/", getDocs)
	router.HandleFunc("/drivers", getAllDrivers).Methods("GET")
	router.HandleFunc(`/drivers/{cpf:(?:\d{11}|\d{3}\.\d{3}\.\d{3}-\d{2})}`,
		getDriver).Methods("GET")
	router.HandleFunc(`/drivers/{cpf:(?:\d{11}|\d{3}\.\d{3}\.\d{3}-\d{2})}`,
		addDriver).Methods("POST")
	router.HandleFunc(`/drivers/{cpf:(?:\d{11}|\d{3}\.\d{3}\.\d{3}-\d{2})}`,
		updateDriver).Methods("PATCH")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal(err)
	}
}
