package entity

type FindTripsQuery struct {
	CPF     *CPF
	HasLoad *bool
	Limit   *uint
	Vehicle *Vehicle
}

func NewFindTripsQuery(cpf *string, hasLoad *bool, limit *uint, vehicleCode *int) (FindTripsQuery, error) {
	var q FindTripsQuery

	if cpf != nil {
		CPF, err := NewCPF(*cpf)
		if err != nil {
			return q, err
		}

		q.CPF = &CPF
	}

	if vehicleCode != nil {
		Vehicle, err := NewVehicle(*vehicleCode)
		if err != nil {
			return q, err
		}

		q.Vehicle = &Vehicle
	}

	q.HasLoad = hasLoad
	q.Limit = limit

	return q, nil
}
