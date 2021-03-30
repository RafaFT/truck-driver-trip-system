package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	repo "github.com/rafaft/truck-driver-trip-system/adapter/gateway/database"
	"github.com/rafaft/truck-driver-trip-system/infrastructure/database"
	"github.com/rafaft/truck-driver-trip-system/infrastructure/web/router"
)

func main() {
	projectID := os.Getenv("projectID")
	port := os.Getenv("port")
	if port == "" {
		port = "8080"
	}

	fc := database.NewFirestoreClient(context.Background(), projectID)
	repo := repo.NewDriverFirestore(fc)
	router := router.NewDriverCloudRun(port, projectID, repo)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
