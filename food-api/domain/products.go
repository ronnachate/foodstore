package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name      string    `gorm:"type:varchar(100);not null"`
	Price     float64   `gorm:"not null"`
	Discounts []ProductDiscount
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type ProductRepository interface {
	GetProducts(c context.Context, productIDs []int) ([]Product, error)
}
