package db

import (
	"github.com/nillocoelho/microservices/payment/internal/application/core/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Adapter struct {
	db *gorm.DB
}

func NewAdapter(databaseURL string) (*Adapter, error) {
	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	_ = db.AutoMigrate(&domain.Payment{}, &domain.Bill{})
	return &Adapter{db: db}, nil
}

func (a *Adapter) SavePayment(p *domain.Payment) error {
	return a.db.Create(p).Error
}

func (a *Adapter) SaveBill(b *domain.Bill) error {
	return a.db.Create(b).Error
}
