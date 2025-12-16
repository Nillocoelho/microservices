package ports

import "github.com/nillocoelho/microservices/payment/internal/application/core/domain"

type DBPort interface {
	SavePayment(p *domain.Payment) error
	SaveBill(b *domain.Bill) error
}
