package main

import (
	"encoding/json"
	"net/http"
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
