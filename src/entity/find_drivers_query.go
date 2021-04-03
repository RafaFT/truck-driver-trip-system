package entity

type FindDriversQuery struct {
	CNH        *CNH
	Gender     *Gender
	HasVehicle *bool
	Limit      *uint
}

func NewFindDriversQuery(cnh, gender string, has_vehicle *bool, limit *uint) FindDriversQuery {
	var q FindDriversQuery

	cnhT, err := NewCNH(cnh)
	if err == nil {
		q.CNH = &cnhT
	}

	genderT, err := NewGender(gender)
	if err == nil {
		q.Gender = &genderT
	}

	if has_vehicle != nil {
		q.HasVehicle = has_vehicle
	}

	if limit != nil {
		q.Limit = limit
	}

	return q
}
