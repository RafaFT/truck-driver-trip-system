// https://pkg.go.dev/cloud.google.com/go/firestore

package repository

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/rafaft/truck-driver-trip-system/entity"
)

type driverDoc struct {
	BirthDate  time.Time `firestore:"birth_date"`
	CNH        string    `firestore:"cnh"`
	CPF        string    `firestore:"cpf"`
	Gender     string    `firestore:"gender"`
	HasVehicle bool      `firestore:"has_vehicle"`
	Name       string    `firestore:"name"`
}

// driver repository implementation
type driverFirestore struct {
	client *firestore.Client
	coll   string
}

func NewDriverFirestore(c *firestore.Client) entity.DriverRepository {
	return &driverFirestore{
		client: c,
		coll:   "drivers",
	}
}

func (df driverFirestore) DeleteDriverByCPF(ctx context.Context, cpf entity.CPF) error {
	doc := df.client.Doc(fmt.Sprintf("%s/%s", df.coll, cpf))

	if _, err := doc.Get(ctx); err != nil {
		if status.Code(err) == codes.NotFound {
			return entity.NewErrDriverNotFound(cpf)
		}
		return err
	}

	_, err := doc.Delete(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (df driverFirestore) FindDriverByCPF(ctx context.Context, cpf entity.CPF) (*entity.Driver, error) {
	docSnap, err := df.client.Doc(fmt.Sprintf("%s/%s", df.coll, cpf)).Get(ctx)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, entity.NewErrDriverNotFound(cpf)
		}
		return nil, err
	}

	var driverDocument driverDoc
	err = docSnap.DataTo(&driverDocument)
	if err != nil {
		return nil, err
	}

	driver, err := entity.NewDriver(
		driverDocument.CPF,
		driverDocument.Name,
		driverDocument.Gender,
		driverDocument.CNH,
		driverDocument.BirthDate,
		driverDocument.HasVehicle,
	)
	if err != nil {
		return nil, err
	}

	return driver, nil
}

func (df driverFirestore) FindDrivers(ctx context.Context, rawQ entity.FindDriversQuery) ([]*entity.Driver, error) {
	q := df.client.Collection(df.coll).Query

	if rawQ.CNH != nil {
		q = q.Where("cnh", "==", *rawQ.CNH)
	}
	if rawQ.Gender != nil {
		q = q.Where("gender", "==", *rawQ.Gender)
	}
	if rawQ.HasVehicle != nil {
		q = q.Where("has_vehicle", "==", *rawQ.HasVehicle)
	}
	if rawQ.Limit != nil {
		q = q.Limit(int(*rawQ.Limit))
	}

	docs, err := q.Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}

	drivers := make([]*entity.Driver, len(docs))
	for i, doc := range docs {
		var driverDocument driverDoc
		err := doc.DataTo(&driverDocument)
		if err != nil {
			return nil, err
		}

		driver, err := entity.NewDriver(
			driverDocument.CPF,
			driverDocument.Name,
			driverDocument.Gender,
			driverDocument.CNH,
			driverDocument.BirthDate,
			driverDocument.HasVehicle,
		)
		if err != nil {
			return nil, err
		}

		drivers[i] = driver
	}

	return drivers, nil
}

func (df driverFirestore) SaveDriver(ctx context.Context, driver *entity.Driver) error {
	newDriverDoc := driverDoc{
		BirthDate:  driver.BirthDate().Time,
		CNH:        string(driver.CNH()),
		CPF:        string(driver.CPF()),
		Gender:     string(driver.Gender()),
		HasVehicle: driver.HasVehicle(),
		Name:       string(driver.Name()),
	}

	path := fmt.Sprintf("%s/%s", df.coll, driver.CPF())
	_, err := df.client.Doc(path).Create(ctx, &newDriverDoc)
	if err != nil {
		if status.Code(err) == codes.AlreadyExists {
			return entity.NewErrDriverAlreadyExists(driver.CPF())
		}
		return err
	}

	return nil
}

func (df driverFirestore) UpdateDriver(ctx context.Context, driver *entity.Driver) error {
	driverDocument := driverDoc{
		BirthDate:  driver.BirthDate().Time,
		CNH:        string(driver.CNH()),
		CPF:        string(driver.CPF()),
		Gender:     string(driver.Gender()),
		HasVehicle: driver.HasVehicle(),
		Name:       string(driver.Name()),
	}

	// TODO: set behaves as an upsert (update/create)
	path := fmt.Sprintf("%s/%s", df.coll, driver.CPF())
	_, err := df.client.Doc(path).Set(ctx, &driverDocument)
	if err != nil {
		return err
	}

	return nil
}
