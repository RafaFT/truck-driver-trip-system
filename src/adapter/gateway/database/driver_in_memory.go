package repository

import (
	"context"

	"github.com/rafaft/truck-driver-trip-system/entity"
	"github.com/rafaft/truck-driver-trip-system/usecase"
)

// driver repository mock implementation
type InMemoryDrivers struct {
	Drivers []*entity.Driver
}

func NewDriverInMemory(drivers []*entity.Driver) usecase.DriverRepository {
	return &InMemoryDrivers{
		Drivers: drivers,
	}
}

func (d *InMemoryDrivers) DeleteByCPF(ctx context.Context, cpf entity.CPF) error {
	for i, driver := range d.Drivers {
		if driver.CPF() == cpf {
			d.Drivers[i] = d.Drivers[len(d.Drivers)-1]
			d.Drivers = d.Drivers[:len(d.Drivers)-1]

			return nil
		}
	}

	return entity.NewErrDriverNotFound(cpf)
}

func (d *InMemoryDrivers) FindByCPF(ctx context.Context, cpf entity.CPF) (*entity.Driver, error) {
	for _, driver := range d.Drivers {
		if driver.CPF() == cpf {
			return driver, nil
		}
	}

	return nil, entity.NewErrDriverNotFound(cpf)
}

func (d *InMemoryDrivers) Find(ctx context.Context, q usecase.FindDriversQuery) ([]*entity.Driver, error) {
	limit := len(d.Drivers)
	if q.Limit != nil {
		limit = int(*q.Limit)
	}

	drivers := make([]*entity.Driver, 0, limit)
	for _, driver := range d.Drivers {
		if len(drivers) == limit {
			break
		}

		if q.CNH != nil && driver.CNH() != *q.CNH {
			continue
		}
		if q.Gender != nil && driver.Gender() != *q.Gender {
			continue
		}
		if q.HasVehicle != nil && driver.HasVehicle() != *q.HasVehicle {
			continue
		}

		drivers = append(drivers, driver)
	}

	return drivers, nil
}

func (d *InMemoryDrivers) Save(ctx context.Context, driver *entity.Driver) error {
	for _, storedDriver := range d.Drivers {
		if storedDriver.CPF() == driver.CPF() {
			return entity.NewErrDriverAlreadyExists(driver.CPF())
		}
	}

	d.Drivers = append(d.Drivers, driver)

	return nil
}

func (d *InMemoryDrivers) Update(ctx context.Context, driver *entity.Driver) error {
	for i, storedDriver := range d.Drivers {
		if storedDriver.CPF() == driver.CPF() {
			updatedDriver := *driver
			d.Drivers[i] = &updatedDriver

			return nil
		}
	}

	return entity.NewErrDriverNotFound(driver.CPF())
}
