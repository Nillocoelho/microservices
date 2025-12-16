package domain

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	UserID     int64
	OrderID    int64
	TotalPrice float32
}

type Bill struct {
	gorm.Model
	PaymentID uint
}
