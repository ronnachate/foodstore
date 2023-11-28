package domain

import (
	"time"

	"github.com/google/uuid"
)

type ProductDiscount struct {
	ID             uint64 `gorm:"primaryKey"`
	ProductID      uuid.UUID
	Product        *Product `gorm:"foreignKey:ProductID;references:ID"`
	ItemDiscountID uint64
	ItemDiscount   *ItemDiscount `gorm:"foreignKey:ItemDiscountID;references:ID"`
	CreatedAt      time.Time     `gorm:"autoCreateTime"`
}
