package main

import (
	"log"

	"github.com/nillocoelho/microservices/order/config"
	"github.com/nillocoelho/microservices/order/internal/adapter/db"
	grpcAdapter "github.com/nillocoelho/microservices/order/internal/adapter/grpc"
	"github.com/nillocoelho/microservices/order/internal/application/core/api"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("DB connection error: %v", err)
	}

	application := api.NewApplication(dbAdapter)

	grpcServer := grpcAdapter.NewAdapter(application, config.GetApplicationPort())

	grpcServer.Run()
}
