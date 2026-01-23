package api

import (
	"github.com/nillocoelho/microservices/shipping/internal/application/core/domain"
	"github.com/nillocoelho/microservices/shipping/internal/ports"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Application struct {
	db ports.DBPort
}

func NewApplication(db ports.DBPort) *Application {
	return &Application{db: db}
}

func (a Application) CreateShipping(shipping domain.Shipping) (domain.Shipping, error) {
	if shipping.OrderID == 0 {
		return domain.Shipping{}, status.Errorf(codes.InvalidArgument, "Order ID is required")
	}

	if len(shipping.Items) == 0 {
		return domain.Shipping{}, status.Errorf(codes.InvalidArgument, "At least one item is required")
	}

	// Calcula os dias de entrega
	shipping.CalculateDeliveryDays()

	// Salva no banco de dados
	if err := a.db.Save(&shipping); err != nil {
		return domain.Shipping{}, status.Errorf(codes.Internal, "Failed to save shipping: %v", err)
	}

	return shipping, nil
}
