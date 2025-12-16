package main

import (
	"log"
	"net"

	"github.com/nillocoelho/microservices/payment/config"
	dbadapter "github.com/nillocoelho/microservices/payment/internal/adapter/db"
	grpcadapter "github.com/nillocoelho/microservices/payment/internal/adapter/grpc"
	"github.com/nillocoelho/microservices/payment/internal/application/core/api"
	paymentpb "github.com/nillocoelho/microservices-proto/golang/payment"
	"google.golang.org/grpc"
)

func main() {
	db, err := dbadapter.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("db init error: %v", err)
	}

	app := api.NewApplication(db)
	srv := grpcadapter.NewServer(app)

	lis, err := net.Listen("tcp", ":"+config.GetApplicationPort())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	paymentpb.RegisterPaymentServer(grpcServer, srv)

	log.Printf("Payment gRPC listening on :%s", config.GetApplicationPort())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("grpc serve error: %v", err)
	}
}
