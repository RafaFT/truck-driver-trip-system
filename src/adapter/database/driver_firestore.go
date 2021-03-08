// https://pkg.go.dev/cloud.google.com/go/firestore

package database

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
type DriverFirestore struct {
	client     *firestore.Client
	collection string
}

func NewDriverFirestore(ctx context.Context, projectId string) (*DriverFirestore, error) {
	// TODO: NewClient uses credentials from environment variables, if not given explicitly
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		return nil, err
	}

	return &DriverFirestore{
		client:     client,
		collection: "drivers",
	}, nil
}

func (df DriverFirestore) DeleteDriverByCPF(ctx context.Context, cpf entity.CPF) error {
	doc := df.client.Doc(fmt.Sprintf("%s/%s", df.collection, cpf))

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

func (df DriverFirestore) FindDriverByCPF(ctx context.Context, cpf entity.CPF) (*entity.Driver, error) {
	docSnap, err := df.client.Doc(fmt.Sprintf("%s/%s", df.collection, cpf)).Get(ctx)
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

	driver, err := entity.NewTruckDriver(
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

func (df DriverFirestore) FindDrivers(ctx context.Context) ([]*entity.Driver, error) {
	docs, err := df.client.Collection(df.collection).Documents(ctx).GetAll()
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

		driver, err := entity.NewTruckDriver(
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

func (df DriverFirestore) SaveDriver(ctx context.Context, driver *entity.Driver) error {
	newDriverDoc := driverDoc{
		BirthDate:  driver.BirthDate().Time,
		CNH:        string(driver.CNHType()),
		CPF:        string(driver.CPF()),
		Gender:     string(driver.Gender()),
		HasVehicle: driver.HasVehicle(),
		Name:       string(driver.Name()),
	}

	_, err := df.client.Doc(string(driver.CPF())).Create(ctx, &newDriverDoc)
	if err != nil {
		if status.Code(err) == codes.AlreadyExists {
			return entity.NewErrDriverAlreadyExists(driver.CPF())
		}
		return err
	}

	return nil
}

func (df DriverFirestore) UpdateDriver(ctx context.Context, driver *entity.Driver) error {
	driverDocument := driverDoc{
		BirthDate:  driver.BirthDate().Time,
		CNH:        string(driver.CNHType()),
		CPF:        string(driver.CPF()),
		Gender:     string(driver.Gender()),
		HasVehicle: driver.HasVehicle(),
		Name:       string(driver.Name()),
	}

	// TODO: set behaves as an upsert (update/create)
	_, err := df.client.Doc(string(driver.CPF())).Set(ctx, &driverDocument)
	if err != nil {
		return err
	}

	return nil
}
