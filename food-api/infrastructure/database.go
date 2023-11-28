package infrastructure

import (
	"errors"
	"fmt"
	"log"

	"github.com/ronnachate/foodstore/food-api/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupDatabase(config *Config) {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok", config.DBHost, config.DBUser, config.DBPassword, config.DBName, config.DBPort)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database", err)
	}

	MigrateDB()
}

func CloseDBConnection() {
	dbInstance, _ := DB.DB()
	_ = dbInstance.Close()
}

// Need to be refractor later
// https://gorm.io/docs/migration.html
func MigrateDB() {
	//create extension upport uuid;
	DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	DB.AutoMigrate(
		&domain.Product{}, &domain.ItemDiscount{}, &domain.ProductDiscount{}, &domain.Member{}, &domain.Order{}, &domain.OrderItem{}, &domain.OrderDiscount{})

	// itemsDiscounts := []domain.ItemDiscount{
	// 	{Name: "Double", Type: constant.PAIR_DISCOUNT_TYPE, DiscountType: constant.PERCENTAGE_DISCOUNT_TYPE, DiscountValue: 5},
	// }

	// if err := DB.First(&domain.ItemDiscount{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
	// 	DB.Create(itemsDiscounts)
	// }

	products := []domain.Product{
		{Name: "Red set", Price: 50},
		{Name: "Green set", Price: 40},
		{Name: "Blue set", Price: 30},
		{Name: "Yellow set", Price: 50},
		{Name: "Pink set", Price: 80},
		{Name: "Purple set", Price: 90},
		{Name: "Orange set", Price: 120},
	}

	if err := DB.First(&domain.Product{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		DB.Create(products)
	}

	// productDiscounts := []domain.ProductDiscount{
	// 	{Product: products[1], ItemDiscount: itemsDiscounts[0]},
	// 	{Product: products[4], ItemDiscount: itemsDiscounts[0]},
	// 	{Product: products[6], ItemDiscount: itemsDiscounts[0]},
	// }

	// if err := DB.First(&domain.ProductDiscount{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
	// 	DB.Create(productDiscounts)
	// }
}
