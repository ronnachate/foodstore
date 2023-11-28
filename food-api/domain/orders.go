package domain

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	dtos "github.com/ronnachate/foodstore/food-api/domain/dtos"
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
	NewOrder(c context.Context, order *Order) (Order, error)
	GetByID(c context.Context, id string) (Order, error)
}

type OrderUsecase interface {
	NewOrder(c context.Context, order dtos.OrderDTO) (Order, error)
	GetByID(c context.Context, orderID string) (Order, error)
	CalculateOrder(order *Order, dto dtos.OrderDTO, products []Product)
	ApplyMemberDiscount(ctx context.Context, order *Order, dto dtos.OrderDTO)
}
