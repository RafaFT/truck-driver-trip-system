package entity

import "fmt"

type FindDriversQuery struct {
	CNH        *CNH
	Gender     *Gender
	HasVehicle *bool
	Limit      *uint
}

func NewFindDriversQuery(cnh, gender *string, has_vehicle *bool, limit *uint) (FindDriversQuery, error) {
	errorMsg := "Invalid FindDriversQuery: %w"
	var q FindDriversQuery

	if cnh != nil {
		cnhT, err := NewCNH(*cnh)
		if err != nil {
			return q, fmt.Errorf(errorMsg, err)
		}
		q.CNH = &cnhT
	}

	if gender != nil {
		genderT, err := NewGender(*gender)
		if err != nil {
			return q, fmt.Errorf(errorMsg, err)
		}
		q.Gender = &genderT
	}

	if has_vehicle != nil {
		q.HasVehicle = has_vehicle
	}

	if limit != nil {
		q.Limit = limit
	}

	return q, nil
}
