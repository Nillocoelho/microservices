package grpc

import (
	"context"
	"fmt"
	"log"
	"net"

	orderpb "github.com/nillocoelho/microservices-proto/golang/order"
	"github.com/nillocoelho/microservices/order/config"
	"github.com/nillocoelho/microservices/order/internal/application/core/domain"
	"github.com/nillocoelho/microservices/order/internal/ports"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Adapter struct {
	api  ports.APIPort
	port int
	orderpb.UnimplementedOrderServer
}

func NewAdapter(api ports.APIPort, port int) *Adapter {
	return &Adapter{
		api:  api,
		port: port,
	}
}

func (a *Adapter) Create(ctx context.Context, request *orderpb.CreateOrderRequest) (*orderpb.CreateOrderResponse, error) {
	var items []domain.OrderItem
	for _, item := range request.OrderItems {
		items = append(items, domain.OrderItem{
			ProductCode: item.ProductCode,
			UnitPrice:   item.UnitPrice,
			Quantity:    item.Quantity,
		})
	}

	newOrder := domain.NewOrder(int64(request.CostumerId), items)

	result, err := a.api.PlaceOrder(newOrder)
	if err != nil {
		return nil, err
	}

	return &orderpb.CreateOrderResponse{
		OrderId: int32(result.ID),
	}, nil
}

func (a *Adapter) Run() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("failed to listen on port %d: %v", a.port, err)
	}

	grpcServer := grpc.NewServer()
	orderpb.RegisterOrderServer(grpcServer, a)

	if config.GetEnv() == "development" {
		reflection.Register(grpcServer)
	}

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve grpc: %v", err)
	}
}
