package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/gorilla/mux"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/rafaft/truck-pad/models"
)

func AddDriver(client *firestore.Client) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// get body's content
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(createErrorJSON(err))
			return
		}

		// load content into Driver instance
		var driver models.Driver
		err = json.Unmarshal(content, &driver)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(createErrorJSON(err))
			return
		}

		err = driver.ValidateDriver()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(createErrorJSON(err))
			return
		}

		collection := client.Collection("drivers")
		doc := collection.Doc(string(*driver.CPF))
		_, err = doc.Create(r.Context(), &driver)
		if err != nil {
			w.WriteHeader(http.StatusConflict)
			w.Write(createErrorJSON(fmt.Errorf("CPF=%s already registered", *driver.CPF)))
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func GetAllDrivers(client *firestore.Client) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		r.ParseForm()

		returnAge := true       // age doesnt come from DB, it's calculated
		returnBirthDate := true // birth date is necessary to calculate age
		if fields := r.Form.Get("fields"); len(fields) > 0 {
			returnAge = strings.Contains(fields, "age")
			returnBirthDate = strings.Contains(fields, "birth_date")
		}

		q := createDriversQuery(client, r)
		docs, err := q.Documents(r.Context()).GetAll()
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(createErrorJSON(fmt.Errorf("internal server error")))
			return
		}

		result := make([]*models.Driver, len(docs))
		for i, docSnapShot := range docs {
			var driver models.Driver
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
}

func GetDriver(client *firestore.Client) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		r.ParseForm()

		returnAge := true       // age doesnt come from DB, it's calculated
		returnBirthDate := true // birth date is necessary to calculate age
		if fields := r.Form.Get("fields"); len(fields) > 0 {
			returnAge = strings.Contains(fields, "age")
			returnBirthDate = strings.Contains(fields, "birth_date")
		}

		q := createDriversQuery(client, r)
		docSnapshot, err := q.Documents(r.Context()).Next()
		if err != nil {
			if err == iterator.Done {
				cpf := mux.Vars(r)["cpf"]
				w.WriteHeader(http.StatusNotFound)
				w.Write(createErrorJSON(fmt.Errorf("cpf=%s not found", cpf)))
			} else {
				fmt.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(createErrorJSON(fmt.Errorf("internal server error")))
			}
			return
		}

		var driver models.Driver
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
}

func UpdateDriver(client *firestore.Client) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// get body's content
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(createErrorJSON(err))
			return
		}

		// load content into Driver instance
		var driver models.Driver
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
		cpf := mux.Vars(r)["cpf"]

		doc := client.Doc(fmt.Sprintf("drivers/%s", cpf))
		_, err = doc.Update(r.Context(), updates)
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
}
