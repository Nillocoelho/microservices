package db

import (
	"fmt"

	"github.com/nillocoelho/microservices/order/internal/application/core/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	CustomerID int64
	Status     string
	OrderItems []OrderItem
}

type OrderItem struct {
	gorm.Model
	ProductCode string
	UnitPrice   float32
	Quantity    int32
	OrderID     uint
}

type Adapter struct {
	db *gorm.DB
}

func NewAdapter(dataSourceUrl string) (*Adapter, error) {
	db, err := gorm.Open(mysql.Open(dataSourceUrl), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("database connection error: %v", err)
	}

	if err := db.AutoMigrate(&Order{}, &OrderItem{}); err != nil {
		return nil, fmt.Errorf("migration error: %v", err)
	}

	return &Adapter{db: db}, nil
}

func (a *Adapter) Get(id string) (domain.Order, error) {
	var entity Order
	res := a.db.Preload("OrderItems").First(&entity, id)
	if res.Error != nil {
		return domain.Order{}, res.Error
	}

	var items []domain.OrderItem
	for _, it := range entity.OrderItems {
		items = append(items, domain.OrderItem{
			ProductCode: it.ProductCode,
			UnitPrice:   it.UnitPrice,
			Quantity:    it.Quantity,
		})
	}

	return domain.Order{
		ID:         int64(entity.ID),
		CustomerID: entity.CustomerID,
		Status:     entity.Status,
		OrderItems: items,
		CreatedAt:  entity.CreatedAt.Unix(),
	}, nil
}

func (a *Adapter) Save(order *domain.Order) error {
	var items []OrderItem
	for _, it := range order.OrderItems {
		items = append(items, OrderItem{
			ProductCode: it.ProductCode,
			UnitPrice:   it.UnitPrice,
			Quantity:    it.Quantity,
		})
	}

	entity := Order{
		CustomerID: order.CustomerID,
		Status:     order.Status,
		OrderItems: items,
	}

	res := a.db.Create(&entity)
	if res.Error == nil {
		order.ID = int64(entity.ID)
	}

	return res.Error
}

func (a *Adapter) UpdateStatus(orderID int64, status string) error {
	return a.db.Model(&Order{}).Where("id = ?", orderID).Update("status", status).Error
}
