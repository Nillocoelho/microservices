package ports

import "github.com/nillocoelho/microservices/shipping/internal/application/core/domain"

type DBPort interface {
	Save(shipping *domain.Shipping) error
}

type APIPort interface {
	CreateShipping(shipping domain.Shipping) (domain.Shipping, error)
}
