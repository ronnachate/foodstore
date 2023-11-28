package repository

import (
	"context"

	"github.com/ronnachate/foodstore/food-api/domain"
	"gorm.io/gorm"
)

type orderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) domain.OrderRepository {
	return &orderRepository{
		DB: db,
	}
}

func (or *orderRepository) NewOrder(c context.Context, order *domain.Order) (domain.Order, error) {
	result := or.DB.Create(&order)
	if result.Error != nil {
		return domain.Order{}, result.Error
	}
	return *order, nil
}

func (or *orderRepository) GetByID(c context.Context, id string) (domain.Order, error) {
	var order domain.Order

	result := or.DB.Model(&domain.Order{}).First(&order, "id = ?", id)
	if result.Error != nil {
		return order, result.Error
	}
	return order, nil
}
