package domain

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	domain "github.com/ronnachate/foodstore/food-api/domain/dtos"
)

type Order struct {
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	MemberID   sql.NullString
	Member     Member `gorm:"foreignKey:MemberID;references:ID"`
	Items      []OrderItem
	DiscountID sql.NullInt64
	Discount   OrderDiscount `gorm:"foreignKey:DiscountID;references:ID"`
	TotalPrice float64
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}

type OrderRepository interface {
	NewOrder(c context.Context, order *Order) error
	GetByID(c context.Context, id string) (Order, error)
}

type OrderUsecase interface {
	NewOrder(c context.Context, order domain.OrderDTO) error
	GetByID(c context.Context, orderID string) (Order, error)
}
