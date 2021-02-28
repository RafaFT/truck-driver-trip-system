package entity

import "strings"

var cnhValues = "ABCDE"

type CNH string

func NewCNH(cnh string) (CNH, error) {
	cnhUpper := strings.ToUpper(cnh)

	if len(cnhUpper) != 1 || !strings.Contains(cnhValues, cnhUpper) {
		return "", newErrInvalidCNH(cnh)
	}

	return CNH(cnhUpper), nil
}
