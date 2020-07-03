package main

import (
	"context"
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
	router.HandleFunc("/drivers", getAllDrivers)
	router.HandleFunc(`/drivers/{cpf:(?:\d{11}|\d{3}\.\d{3}\.\d{3}-\d{2})}`, getDriver)
}

func main() {
	http.ListenAndServe(":3000", router)
}
