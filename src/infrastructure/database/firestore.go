package database

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

func NewFirestoreClient(ctx context.Context, projectID string) *firestore.Client {
	c, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatal(err)
	}

	return c
}
