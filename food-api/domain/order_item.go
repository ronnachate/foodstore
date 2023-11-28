package domain

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type OrderItem struct {
	ID             uint64 `gorm:"primaryKey"`
	OrderID        uuid.UUID
	Order          Order `gorm:"foreignKey:OrderID;references:ID"`
	ProductID      uuid.UUID
	Product        Product `gorm:"foreignKey:ProductID;references:ID"`
	Quantity       uint64
	Price          float64
	ItemDiscountID sql.NullInt64
	ItemDiscount   ItemDiscount `gorm:"foreignKey:ItemDiscountID;references:ID"`
	TotalPrice     float64
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
}
