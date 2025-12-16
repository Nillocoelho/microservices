package main

import (
	"log"

	"github.com/nillocoelho/microservices/order/config"
	"github.com/nillocoelho/microservices/order/internal/adapter/db"
	grpcAdapter "github.com/nillocoelho/microservices/order/internal/adapter/grpc"
	payment_adapter "github.com/nillocoelho/microservices/order/internal/adapter/payment"
	"github.com/nillocoelho/microservices/order/internal/application/core/api"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("DB connection error: %v", err)
	}

	paymentAdapter, err := payment_adapter.NewAdapter(config.GetPaymentServiceUrl())
	if err != nil {
		log.Fatalf("Failed to initialize payment stub. Error: %v", err)
	}

	application := api.NewApplication(dbAdapter, paymentAdapter)

	grpcServer := grpcAdapter.NewAdapter(application, config.GetApplicationPort())

	grpcServer.Run()
}
