package shipping

import (
	"context"
	"fmt"
	"log"
	"time"

	shippingpb "github.com/nillocoelho/microservices-proto/golang/shipping"
	"github.com/nillocoelho/microservices/order/internal/application/core/domain"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type Adapter struct {
	shippingClient shippingpb.ShippingClient
}

func NewAdapter(shippingServiceUrl string) (*Adapter, error) {
	conn, err := grpc.Dial(
		shippingServiceUrl,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to dial shipping service: %w", err)
	}

	client := shippingpb.NewShippingClient(conn)

	return &Adapter{shippingClient: client}, nil
}

func (a *Adapter) CreateShipping(order *domain.Order) error {
	// Converte os itens do dominio para o protobuf
	var items []*shippingpb.ShippingItem
	for _, item := range order.OrderItems {
		items = append(items, &shippingpb.ShippingItem{
			ProductCode: item.ProductCode,
			Quantity:    item.Quantity,
		})
	}

	// Cria a requisicao
	req := &shippingpb.CreateShippingRequest{
		OrderId: int64(order.ID),
		Items:   items,
	}

	// Chamada ao servico de shipping com timeout de 2s
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	resp, err := a.shippingClient.Create(ctx, req)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			if st.Code() == codes.DeadlineExceeded {
				log.Printf("Shipping request for order %d exceeded deadline (2 seconds)", order.ID)
				return status.Errorf(codes.DeadlineExceeded, "shipping service timeout")
			}
			return status.Errorf(st.Code(), "Shipping service error: %s", st.Message())
		}
		return fmt.Errorf("failed to create shipping: %w", err)
	}

	// Atribui o ID de envio ao pedido
	order.ShippingID = resp.ShippingId
	order.DeliveryDays = resp.DeliveryDays

	return nil
}
