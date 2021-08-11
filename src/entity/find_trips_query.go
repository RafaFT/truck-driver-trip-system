package entity

type FindTripsQuery interface {
	CPF() *CPF
	HasLoad() *bool
	Limit() *uint
	Vehicle() *Vehicle
}

type findTripsQuery struct {
	cpf     *CPF
	hasLoad *bool
	limit   *uint
	vehicle *Vehicle
}

func NewFindTripsQuery(cpf *string, hasLoad *bool, limit *uint, vehicleCode *int) (FindTripsQuery, error) {
	var q findTripsQuery

	if cpf != nil {
		CPF, err := NewCPF(*cpf)
		if err != nil {
			return nil, err
		}

		q.cpf = &CPF
	}

	if vehicleCode != nil {
		Vehicle, err := NewVehicle(*vehicleCode)
		if err != nil {
			return nil, err
		}

		q.vehicle = &Vehicle
	}

	q.hasLoad = hasLoad
	q.limit = limit

	return q, nil
}

func (q findTripsQuery) CPF() *CPF {
	return q.cpf
}

func (q findTripsQuery) HasLoad() *bool {
	return q.hasLoad
}

func (q findTripsQuery) Limit() *uint {
	return q.limit
}

func (q findTripsQuery) Vehicle() *Vehicle {
	return q.vehicle
}
