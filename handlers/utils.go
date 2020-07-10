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

func createDriversQuery(client *firestore.Client, r *http.Request) firestore.Query {
	q := client.Collection("drivers").Query

	// cpf will only exists if it's the getDriver`s route
	if cpf, exist := mux.Vars(r)["cpf"]; exist {
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

func createTripsQuery(client *firestore.Client, r *http.Request) firestore.Query {
	// TODO: Since I query "trips" always by using Collection Group, I
	// 	should probably organize trips as a top level collection
	//  https://firebase.googleblog.com/2019/06/understanding-collection-group-queries.html
	q := client.CollectionGroup("trips").Query

	// add filters
	if driver_id := r.Form.Get("driver_id"); len(driver_id) > 0 {
		q = q.Where("driver_id", "==", driver_id)
	}
	if id := r.Form.Get("id"); len(id) > 0 {
		q = q.Where("id", "==", id)
	}
	if str_has_load := r.Form.Get("has_load"); len(str_has_load) > 0 {
		has_load, err := strconv.ParseBool(str_has_load)
		if err == nil {
			q = q.Where("has_load", "==", has_load)
		}
	}
	if str_vehicle_type := r.Form.Get("vehicle_type"); len(str_vehicle_type) > 0 {
		vehicle_type, err := strconv.Atoi(str_vehicle_type)
		if err == nil {
			q = q.Where("vehicle_type", "==", vehicle_type)
		}
	}
	if strFrom := r.Form.Get("from"); len(strFrom) > 0 {
		from, err := time.Parse(ISO8601, strFrom)
		if err == nil {
			q = q.Where("time", ">=", from)
		}
	}
	if strTo := r.Form.Get("to"); len(strTo) > 0 {
		to, err := time.Parse(ISO8601, strTo)
		if err == nil {
			q = q.Where("time", "<", to)
		}
	}
	// TODO: add query by origin and destination on lat and lng values
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
