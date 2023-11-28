package domain

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	MemberID   uuid.UUID
	Member     Member `gorm:"foreignKey:MemberID;references:ID"`
	Items      []OrderItem
	DiscountID uint64
	Discount   OrderDiscount `gorm:"foreignKey:DiscountID;references:ID"`
	TotalPrice float64
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}
