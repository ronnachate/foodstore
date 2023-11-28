package domain

import (
	"time"
)

type ItemDiscount struct {
	ID            uint64 `gorm:"primaryKey"`
	Name          string `gorm:"type:varchar(100);not null"`
	Type          uint64 `gorm:"not null"`
	Min           uint64
	DiscountType  uint64    `gorm:"not null"`
	DiscountValue float64   `gorm:"not null"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
}
