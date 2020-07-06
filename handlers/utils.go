package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/gorilla/mux"

	"github.com/rafaft/truck-pad/models"
)

const ISO8601 = "2006-01-02"

func createDriversQuery(r *http.Request) firestore.Query {
	q := client.Collection("drivers").Query

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
		// always get birth_date, because it's necessary for calculating
		// the age, remove it later if necessary
		fields := []string{
			"birth_date",
		}
		for _, field := range splitFields {
			if len(field) > 0 {
				fields = append(fields, field)
			}
		}

		q = q.Select(fields...)
	}

	return q
}

func createTripsQuery(r *http.Request) firestore.Query {
	q := client.CollectionGroup("trips").Query

	// add filters
	if cpf := r.Form.Get("cpf"); len(cpf) > 0 {
		q = q.Where("cpf", "==", cpf)
	}
	if str_has_load := r.Form.Get("has_load"); len(str_has_load) > 0 {
		has_load, err := strconv.ParseBool(str_has_load)
		if err == nil {
			q = q.Where("has_load", "==", has_load)
		}
	}
	if vehicle_type := r.Form.Get("vehicle_type"); len(vehicle_type) > 0 {
		q = q.Where("vehicle_type", "==", vehicle_type)
	}
	if str_start_date := r.Form.Get("start_date"); len(str_start_date) > 0 {
		start_date, err := time.Parse(ISO8601, str_start_date)
		if err == nil {
			q = q.Where("time", ">=", start_date)
		}
	}
	if str_end_date := r.Form.Get("end_date"); len(str_end_date) > 0 {
		end_date, err := time.Parse(ISO8601, str_end_date)
		if err == nil {
			q = q.Where("time", "<", end_date)
		}
	}
	if order := r.Form.Get("order"); strings.ToLower(order) == "asc" {
		q = q.OrderBy("time", firestore.Asc)
	} else {
		q = q.OrderBy("time", firestore.Desc)
	}
	if str_limit := r.Form.Get("limit"); len(str_limit) > 0 {
		limit, err := strconv.Atoi(str_limit)
		if err == nil {
			q = q.Limit(limit)
		}
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
	output := models.ErrorJSON{
		Error: e.Error(),
	}

	content, _ := json.Marshal(&output)
	return content
}

func calculateAge(birthDate, now time.Time) int {
	years := now.Year() - birthDate.Year()
	if years < 0 {
		return 0
	}

	birthMonthNDay, _ := strconv.Atoi(fmt.Sprintf("%d%d", int(birthDate.Month()), birthDate.Day()))
	nowMonthNDay, _ := strconv.Atoi(fmt.Sprintf("%d%d", int(now.Month()), now.Day()))

	if birthMonthNDay > nowMonthNDay {
		years--
	}

	return years
}
