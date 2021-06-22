package entity

import "strings"

var cnhValues = map[string]bool{
	"A": true,
	"B": true,
	"C": true,
	"D": true,
	"E": true,
}

type CNH string

func NewCNH(cnh string) (CNH, error) {
	cnhUpper := strings.ToUpper(cnh)

	if _, ok := cnhValues[cnhUpper]; !ok {
		return "", newErrInvalidCNH(cnh)
	}

	return CNH(cnhUpper), nil
}
