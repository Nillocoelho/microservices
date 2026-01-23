package main

import (
	"log"

	"github.com/nillocoelho/microservices/shipping/config"
	"github.com/nillocoelho/microservices/shipping/internal/adapter/db"
	grpcAdapter "github.com/nillocoelho/microservices/shipping/internal/adapter/grpc"
	"github.com/nillocoelho/microservices/shipping/internal/application/core/api"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDatabaseURL())
	if err != nil {
		log.Fatalf("DB connection error: %v", err)
	}

	application := api.NewApplication(dbAdapter)

	grpcServer := grpcAdapter.NewAdapter(application, config.GetApplicationPort())

	grpcServer.Run()
}
