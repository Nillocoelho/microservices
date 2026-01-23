package db

import (
	"fmt"
	"time"

	"github.com/nillocoelho/microservices/shipping/internal/application/core/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Adapter struct {
	db *gorm.DB
}

type shippingRecord struct {
	ID           int64 `gorm:"primaryKey"`
	OrderID      int64
	DeliveryDays int32
	Status       string    `gorm:"default:Pending"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
}

func (shippingRecord) TableName() string { return "shipping" }

func NewAdapter(databaseURL string) (*Adapter, error) {
	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	if err := db.AutoMigrate(&shippingRecord{}); err != nil {
		return nil, fmt.Errorf("migration failed: %w", err)
	}

	return &Adapter{db: db}, nil
}

func (a *Adapter) Save(shipping *domain.Shipping) error {
	record := shippingRecord{
		OrderID:      shipping.OrderID,
		DeliveryDays: shipping.DeliveryDays,
		Status:       "Pending",
	}

	if err := a.db.Create(&record).Error; err != nil {
		return fmt.Errorf("failed to insert shipping: %w", err)
	}

	shipping.ID = record.ID
	return nil
}
