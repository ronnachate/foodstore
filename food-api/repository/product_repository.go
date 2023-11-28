package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/ronnachate/foodstore/food-api/domain"
	"gorm.io/gorm"
)

type productRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) domain.ProductRepository {
	return &productRepository{
		DB: db,
	}
}

func (pr *productRepository) GetProducts(c context.Context, productIDs []uuid.UUID) ([]domain.Product, error) {
	var products []domain.Product
	result := pr.DB.Model(&domain.Product{}).Preload("Discounts").Where("id IN ?", productIDs).Find(&products)
	if result.Error != nil {
		return products, result.Error
	}
	return products, nil
}
