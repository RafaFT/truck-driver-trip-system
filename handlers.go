package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func getAllDrivers(w http.ResponseWriter, r *http.Request) {
	collection := client.Collection("drivers")
	docs, err := collection.DocumentRefs(ctx).GetAll()
	if err != nil {
		panic(err)
	}

	result := make(map[string]*Driver)
	for _, doc := range docs {
		cpf := doc.ID
		var driver Driver

		docSnapShot, err := doc.Get(ctx)
		if err != nil {
			panic(err)
		}

		err = docSnapShot.DataTo(&driver)
		if err != nil {
			panic(err)
		}

		driver.Age = 30 // TODO: add age logic
		result[cpf] = &driver
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func getDocs(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Docs..."))
}

func getDriver(w http.ResponseWriter, r *http.Request) {
	inputCPF := mux.Vars(r)["cpf"]
	cpf := strings.ReplaceAll(strings.ReplaceAll(inputCPF, "-", ""), ".", "")

	docSnapShot, err := client.Doc(fmt.Sprintf("drivers/%s", cpf)).Get(ctx)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	var driver Driver

	err = docSnapShot.DataTo(&driver)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	driver.Age = 30 // TODO: add age logic
	driver.CPF = cpf

	b, err := json.Marshal(&driver)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
