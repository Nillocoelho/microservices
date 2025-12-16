package api

import (
	"github.com/nillocoelho/microservices/payment/internal/application/core/domain"
	"github.com/nillocoelho/microservices/payment/internal/ports"
)

type Application struct {
	db ports.DBPort
}

func NewApplication(db ports.DBPort) *Application {
	return &Application{db: db}
}

func (a *Application) CreatePayment(userID, orderID int64, totalPrice float32) (domain.Payment, domain.Bill, error) {
	p := domain.Payment{
		UserID:     userID,
		OrderID:    orderID,
		TotalPrice: totalPrice,
	}

	if err := a.db.SavePayment(&p); err != nil {
		return domain.Payment{}, domain.Bill{}, err
	}

	b := domain.Bill{PaymentID: p.ID}
	if err := a.db.SaveBill(&b); err != nil {
		return domain.Payment{}, domain.Bill{}, err
	}

	return p, b, nil
}
