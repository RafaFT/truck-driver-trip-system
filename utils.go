package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"cloud.google.com/go/firestore"
	"github.com/gorilla/mux"
)

func createQuery(collection *firestore.CollectionRef, r *http.Request) firestore.Query {
	q := collection.Query

	// cpf will only exists if it's the getDriver`s route
	if rawCPF, exist := mux.Vars(r)["cpf"]; exist {
		cpf := strings.ReplaceAll(strings.ReplaceAll(rawCPF, ".", ""), "-", "")
		q = q.Where("cpf", "==", cpf) // TODO: query for the document ID
	}
	if gender := r.Form.Get("gender"); len(gender) > 0 {
		q = q.Where("gender", "==", strings.ToUpper(gender))
	}
	if str_has_vehicle := r.Form.Get("has_vehicle"); len(str_has_vehicle) > 0 {
		has_vehicle, err := strconv.ParseBool(str_has_vehicle)
		if err == nil {
			q = q.Where("has_vehicle", "==", has_vehicle)
		}
	}
	if cnh_type := r.Form.Get("cnh_type"); len(cnh_type) > 0 {
		q = q.Where("cnh_type", "==", strings.ToUpper(cnh_type))
	}

	// get only requested fields
	if rawFields := r.Form.Get("fields"); len(rawFields) > 0 {
		splitFields := strings.Split(rawFields, ",")
		fields := make([]string, 0)
		for _, field := range splitFields {
			if len(field) > 0 {
				fields = append(fields, field)
			}
		}

		q = q.Select(fields...)
	}

	return q
}

func createErrorJSON(e error) []byte {
	output := errorJSON{
		Error: e.Error(),
	}

	content, _ := json.Marshal(&output)
	return content
}
