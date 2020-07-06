package models

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Gender string

func (gender *Gender) UnmarshalJSON(b []byte) error {
	var sGender string
	json.Unmarshal(b, &sGender)

	sGender = strings.ToUpper(sGender)
	validGenders := "FMO" // "O" stands for other
	if len(sGender) != 1 || !strings.Contains(validGenders, sGender) {
		return fmt.Errorf("'gender' must be 'F', 'M' or 'O'")
	}

	*gender = Gender(sGender)
	return nil
}
