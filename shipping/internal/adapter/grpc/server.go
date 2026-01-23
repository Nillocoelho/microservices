package grpc

import (
	"context"
	"fmt"
	"log"
	"net"

	shippingpb "github.com/nillocoelho/microservices-proto/golang/shipping"
	"github.com/nillocoelho/microservices/shipping/internal/application/core/domain"
	"github.com/nillocoelho/microservices/shipping/internal/ports"
	"google.golang.org/grpc"
)

type Adapter struct {
	shippingpb.UnimplementedShippingServer
	application ports.APIPort
	port        string
}

func NewAdapter(application ports.APIPort, port string) *Adapter {
	return &Adapter{
		application: application,
		port:        port,
	}
}

func (a *Adapter) Create(ctx context.Context, req *shippingpb.CreateShippingRequest) (*shippingpb.CreateShippingResponse, error) {
	// Converter os itens do proto para o domain
	var items []domain.ShippingItem
	for _, item := range req.Items {
		items = append(items, domain.ShippingItem{
			ProductCode: item.ProductCode,
			Quantity:    item.Quantity,
		})
	}

	// Criar o domínio
	shipping := domain.Shipping{
		OrderID: req.OrderId,
		Items:   items,
	}

	// Usar a aplicação para criar a entrega
	result, err := a.application.CreateShipping(shipping)
	if err != nil {
		return nil, err
	}

	// Retornar a resposta
	return &shippingpb.CreateShippingResponse{
		ShippingId:   result.ID,
		DeliveryDays: result.DeliveryDays,
	}, nil
}

func (a *Adapter) Run() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", a.port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	shippingpb.RegisterShippingServer(s, a)

	log.Printf("Shipping server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
