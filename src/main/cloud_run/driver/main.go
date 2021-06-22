package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	repository "github.com/rafaft/truck-driver-trip-system/adapter/gateway/database"
	"github.com/rafaft/truck-driver-trip-system/infrastructure/database"
	"github.com/rafaft/truck-driver-trip-system/infrastructure/web/rest"
)

func main() {
	projectID := os.Getenv("projectID")
	port := os.Getenv("port")
	if port == "" {
		port = "8080"
	}

	fc := database.NewFirestoreClient(context.Background(), projectID)
	repo := repository.NewDriverFirestore(fc)
	router := rest.NewDriverCloudRun(port, projectID, repo)
	rest.SetDriversRoutes(router)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
