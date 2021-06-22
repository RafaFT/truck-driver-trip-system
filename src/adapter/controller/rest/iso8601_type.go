package rest

import (
	"encoding/json"
	"reflect"
	"time"
)

const ISO8601 = "2006-01-02"

type ISO8601Date struct {
	time.Time
}

func (t *ISO8601Date) UnmarshalJSON(b []byte) error {
	var dateInput interface{}
	err := json.Unmarshal(b, &dateInput)

	jsonTypeError := &json.UnmarshalTypeError{
		Value: reflect.TypeOf(dateInput).Name(),
		Type:  reflect.TypeOf(*t),
	}

	if err != nil || jsonTypeError.Value != "string" {
		return jsonTypeError
	}

	dateString := dateInput.(string)
	date, err := time.Parse(ISO8601, dateString)
	if err != nil {
		date, err = time.Parse(time.RFC3339, dateString)
		if err != nil {
			return jsonTypeError
		}
	}

	*t = ISO8601Date{
		date,
	}

	return nil
}
