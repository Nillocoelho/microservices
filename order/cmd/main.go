package main

import (
	"log"

	"github.com/nillocoelho/microservices/order/config"
	"github.com/nillocoelho/microservices/order/internal/adapter/db"
	grpcAdapter "github.com/nillocoelho/microservices/order/internal/adapter/grpc"
	payment_adapter "github.com/nillocoelho/microservices/order/internal/adapter/payment"
	shipping_adapter "github.com/nillocoelho/microservices/order/internal/adapter/shipping"
	"github.com/nillocoelho/microservices/order/internal/application/core/api"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDatabaseURL())
	if err != nil {
		log.Fatalf("DB connection error: %v", err)
	}

	paymentAdapter, err := payment_adapter.NewAdapter(config.GetPaymentServiceUrl())
	if err != nil {
		log.Fatalf("Failed to initialize payment stub. Error: %v", err)
	}

	shippingAdapter, err := shipping_adapter.NewAdapter(config.GetShippingServiceUrl())
	if err != nil {
		log.Fatalf("Failed to initialize shipping stub. Error: %v", err)
	}

	application := api.NewApplication(dbAdapter, paymentAdapter, shippingAdapter)

	grpcServer := grpcAdapter.NewAdapter(application, config.GetApplicationPort())

	grpcServer.Run()
}
