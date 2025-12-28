package api

import (
	"github.com/nillocoelho/microservices/order/internal/application/core/domain"
	"github.com/nillocoelho/microservices/order/internal/ports"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Application struct {
	db      ports.DBPort
	payment ports.PaymentPort
}

func NewApplication(db ports.DBPort, payment ports.PaymentPort) *Application {
	return &Application{db: db, payment: payment}
}

func (a Application) PlaceOrder(order domain.Order) (domain.Order, error) {
	if err := a.db.Save(&order); err != nil {
		return domain.Order{}, err
	}

	if err := validateItemLimit(order); err != nil {
		_ = a.db.UpdateStatus(order.ID, "Canceled")
		order.Status = "Canceled"
		return domain.Order{}, err
	}

	if err := a.payment.Charge(&order); err != nil {
		_ = a.db.UpdateStatus(order.ID, "Canceled")
		order.Status = "Canceled"
		return domain.Order{}, err
	}

	if err := a.db.UpdateStatus(order.ID, "Paid"); err != nil {
		return domain.Order{}, err
	}

	order.Status = "Paid"
	return order, nil
}

func validateItemLimit(order domain.Order) error {
	var totalItems int32
	for _, item := range order.OrderItems {
		totalItems += item.Quantity
	}

	if totalItems > 50 {
		return status.Errorf(codes.InvalidArgument, "Order cannot contain more than 50 items in total.")
	}

	return nil
}
