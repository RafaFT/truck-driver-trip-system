package repo

import (
	"context"

	"github.com/rafaft/truck-driver-trip-system/entity"
)

// driver repository mock implementation
type InMemoryDrivers struct {
	Drivers []*entity.Driver
}

func NewDriverInMemory() entity.DriverRepository {
	return &InMemoryDrivers{
		Drivers: make([]*entity.Driver, 0),
	}
}

func (d *InMemoryDrivers) DeleteDriverByCPF(ctx context.Context, cpf entity.CPF) error {
	for i, driver := range d.Drivers {
		if driver.CPF() == cpf {
			d.Drivers[i] = d.Drivers[len(d.Drivers)-1]
			d.Drivers = d.Drivers[:len(d.Drivers)-1]

			return nil
		}
	}

	return entity.NewErrDriverNotFound(cpf)
}

func (d *InMemoryDrivers) FindDriverByCPF(ctx context.Context, cpf entity.CPF) (*entity.Driver, error) {
	for _, driver := range d.Drivers {
		if driver.CPF() == cpf {
			return driver, nil
		}
	}

	return nil, entity.NewErrDriverNotFound(cpf)
}

func (d *InMemoryDrivers) FindDrivers(ctx context.Context) ([]*entity.Driver, error) {
	return d.Drivers, nil
}

func (d *InMemoryDrivers) SaveDriver(ctx context.Context, driver *entity.Driver) error {
	for _, storedDriver := range d.Drivers {
		if storedDriver.CPF() == driver.CPF() {
			return entity.NewErrDriverAlreadyExists(driver.CPF())
		}
	}

	d.Drivers = append(d.Drivers, driver)

	return nil
}

func (d *InMemoryDrivers) UpdateDriver(ctx context.Context, driver *entity.Driver) error {
	for i, storedDriver := range d.Drivers {
		if storedDriver.CPF() == driver.CPF() {
			updatedDriver := *driver
			d.Drivers[i] = &updatedDriver

			return nil
		}
	}

	return entity.NewErrDriverNotFound(driver.CPF())
}