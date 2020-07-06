package models

import (
	"encoding/json"
	"fmt"
	"regexp"
)

type CPF string

func (cpf *CPF) UnmarshalJSON(b []byte) error {
	var sCPF string
	err := json.Unmarshal(b, &sCPF)
	if err != nil {
		return err
	}

	matched, err := regexp.MatchString(`^\d{11}$`, sCPF)
	if err != nil {
		return err
	}
	if !matched {
		return fmt.Errorf("invalid value for 'CPF'")
	}

	*cpf = CPF(sCPF)
	return nil
}
