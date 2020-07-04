package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"

	"cloud.google.com/go/firestore"
	"github.com/gorilla/mux"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func addDriver(w http.ResponseWriter, r *http.Request) {
	// get body's content
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// load content into Driver instance
	var driver Driver
	err = json.Unmarshal(content, &driver)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// get CPF (doc ID)
	rawCPF := mux.Vars(r)["cpf"]
	cpf := strings.ReplaceAll(strings.ReplaceAll(rawCPF, ".", ""), "-", "")
	driver.CPF = &cpf

	collection := client.Collection("drivers")
	doc := collection.Doc(cpf)
	_, err = doc.Create(ctx, &driver)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

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

		age := 30
		driver.Age = &age // TODO: add age logic
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

	age := 30
	driver.Age = &age // TODO: add age logic
	driver.CPF = &cpf

	b, err := json.Marshal(&driver)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func updateDriver(w http.ResponseWriter, r *http.Request) {
	// get body's content
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// load content into Driver instance
	var driver Driver
	err = json.Unmarshal(content, &driver)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// explicitly convert Driver to map, because it's easier to iterate it
	mapDriver := map[string]interface{}{
		"name":        driver.Name,
		"birth_date":  driver.BirthDate,
		"gender":      driver.Gender,
		"has_vehicle": driver.HasVehicle,
		"cnh_type":    driver.CNHType,
	}

	// create slice of updates
	updates := make([]firestore.Update, 0)
	for fieldName, fieldValue := range mapDriver {
		if !reflect.ValueOf(fieldValue).IsNil() {
			update := firestore.Update{
				Path:  fieldName,
				Value: fieldValue,
			}
			updates = append(updates, update)
		}
	}

	// get CPF (doc ID)
	rawCPF := mux.Vars(r)["cpf"]
	cpf := strings.ReplaceAll(strings.ReplaceAll(rawCPF, ".", ""), "-", "")

	doc := client.Doc(fmt.Sprintf("drivers/%s", cpf))
	_, err = doc.Update(ctx, updates)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			w.WriteHeader(http.StatusNotFound)
		} else if err.Error() == "firestore: no paths to update" {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
}
