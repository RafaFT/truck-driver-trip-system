package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func addDriver(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// get body's content
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(createErrorJSON(err))
		return
	}

	// load content into Driver instance
	var driver Driver
	err = json.Unmarshal(content, &driver)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(createErrorJSON(err))
		return
	}

	// get CPF (doc ID)
	cpf, _ := getCPF(r)
	driver.CPF = &cpf

	err = driver.ValidateDriver()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(createErrorJSON(err))
		return
	}

	collection := client.Collection("drivers")
	doc := collection.Doc(cpf)
	_, err = doc.Create(ctx, &driver)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		w.Write(createErrorJSON(fmt.Errorf("CPF=%s already exists", cpf)))
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func getAllDrivers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	r.ParseForm()

	returnAge := true       // age doesnt come from DB, it's calculated
	returnBirthDate := true // birth date is necessary to calculate age
	if fields := r.Form.Get("fields"); len(fields) > 0 {
		returnAge = strings.Contains(fields, "age")
		returnBirthDate = strings.Contains(fields, "birth_date")
	}

	q := createQuery(client.Collection("drivers"), r)
	docs, err := q.Documents(ctx).GetAll()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(createErrorJSON(fmt.Errorf("internal server error")))
		return
	}

	result := make([]*Driver, len(docs))
	for i, docSnapShot := range docs {
		var driver Driver
		err = docSnapShot.DataTo(&driver)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(createErrorJSON(fmt.Errorf("internal server error")))
			return
		}

		if returnAge {
			driver.Age = calculateAge(*driver.BirthDate, time.Now())
		}
		if !returnBirthDate {
			driver.BirthDate = nil
		}

		result[i] = &driver
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

func getDocs(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Docs..."))
}

func getDriver(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	r.ParseForm()

	returnAge := true       // age doesnt come from DB, it's calculated
	returnBirthDate := true // birth date is necessary to calculate age
	if fields := r.Form.Get("fields"); len(fields) > 0 {
		returnAge = strings.Contains(fields, "age")
		returnBirthDate = strings.Contains(fields, "birth_date")
	}

	q := createQuery(client.Collection("drivers"), r)
	docSnapshot, err := q.Documents(ctx).Next()
	if err != nil {
		if err == iterator.Done {
			cpf, _ := getCPF(r)
			w.WriteHeader(http.StatusNotFound)
			w.Write(createErrorJSON(fmt.Errorf("cpf=%s not found", cpf)))
		} else {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(createErrorJSON(fmt.Errorf("internal server error")))
		}
		return
	}

	var driver Driver
	err = docSnapshot.DataTo(&driver)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(createErrorJSON(fmt.Errorf("internal server error")))
		return
	}

	if returnAge {
		driver.Age = calculateAge(*driver.BirthDate, time.Now())
	}
	if !returnBirthDate {
		driver.BirthDate = nil
	}

	b, err := json.Marshal(&driver)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(createErrorJSON(fmt.Errorf("internal server error")))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func updateDriver(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// get body's content
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(createErrorJSON(err))
		return
	}

	// load content into Driver instance
	var driver Driver
	err = json.Unmarshal(content, &driver)
	if err != nil || driver.CPF != nil { // cannot update CPF
		w.WriteHeader(http.StatusBadRequest)
		w.Write(createErrorJSON(err))
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
	cpf, _ := getCPF(r)

	doc := client.Doc(fmt.Sprintf("drivers/%s", cpf))
	_, err = doc.Update(ctx, updates)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			w.WriteHeader(http.StatusNotFound)
			w.Write(createErrorJSON(fmt.Errorf("cpf=%s not found", cpf)))
		} else if err.Error() == "firestore: no paths to update" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(createErrorJSON(fmt.Errorf("empty update request")))
		} else {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(createErrorJSON(fmt.Errorf("internal server error")))
		}
		return
	}

	w.WriteHeader(http.StatusOK)
}
