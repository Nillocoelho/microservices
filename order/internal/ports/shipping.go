package ports

import "github.com/nillocoelho/microservices/order/internal/application/core/domain"

type ShippingPort interface {
	CreateShipping(order *domain.Order) error
}
