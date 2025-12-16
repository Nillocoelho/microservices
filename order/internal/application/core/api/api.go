package api

import (
	"github.com/nillocoelho/microservices/order/internal/application/core/domain"
	"github.com/nillocoelho/microservices/order/internal/ports"
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

	if err := a.payment.Charge(&order); err != nil {
		return domain.Order{}, err
	}

	return order, nil
}
