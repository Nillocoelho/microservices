package ports

import "github.com/nillocoelho/microservices/order/internal/application/core/domain"

type DBPort interface {
	Get(id string) (domain.Order, error)
	Save(order *domain.Order) error
	UpdateStatus(orderID int64, status string) error
}
