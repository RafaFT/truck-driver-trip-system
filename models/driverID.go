package models

import (
	"encoding/json"
	"fmt"
	"regexp"
)

type DriverID string

func (id *DriverID) UnmarshalJSON(b []byte) error {
	var sID string
	err := json.Unmarshal(b, &sID)
	if err != nil {
		return err
	}

	matched, err := regexp.MatchString(`^\d{11}$`, sID)
	if err != nil {
		return err
	}
	if !matched {
		return fmt.Errorf("invalid value for 'driver_id' (must be valid CPF) ")
	}

	*id = DriverID(sID)
	return nil
}
