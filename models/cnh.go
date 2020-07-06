package models

import (
	"encoding/json"
	"fmt"
	"strings"
)

type CNHType string

func (cnh *CNHType) UnmarshalJSON(b []byte) error {
	var sCNH string
	json.Unmarshal(b, &sCNH)

	sCNH = strings.ToUpper(sCNH)
	validCNHTypes := "ABCDE"
	if len(sCNH) != 1 || !strings.Contains(validCNHTypes, sCNH) {
		return fmt.Errorf("'cnh_type' must be 'A', 'B', 'C', 'D' or 'E'")
	}

	*cnh = CNHType(sCNH)
	return nil
}
