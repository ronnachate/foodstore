package repository

import (
	"context"

	"github.com/ronnachate/foodstore/food-api/domain"
	"gorm.io/gorm"
)

type orderDiscountRepository struct {
	DB *gorm.DB
}

func NewOrderDiscountRepository(db *gorm.DB) domain.OrderDiscountRepository {
	return &orderDiscountRepository{
		DB: db,
	}
}

func (or *orderDiscountRepository) GetByType(c context.Context, typeID uint64) (domain.OrderDiscount, error) {
	var discount domain.OrderDiscount

	result := or.DB.Model(&domain.OrderDiscount{}).First(&discount, "type = ?", typeID)
	if result.Error != nil {
		return discount, result.Error
	}
	return discount, nil
}
