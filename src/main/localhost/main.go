package main

import (
	"fmt"
	"log"
	"net/http"

	repository "github.com/rafaft/truck-driver-trip-system/adapter/gateway/database"
	"github.com/rafaft/truck-driver-trip-system/infrastructure/web/router"
)

func main() {
	port := "8080"
	repo := repository.NewDriverInMemory(nil)
	router := router.NewDriverLocalHost(port, repo)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
