package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/ronnachate/foodstore/food-api/domain"
	constant "github.com/ronnachate/foodstore/food-api/domain/constants"
	"github.com/ronnachate/foodstore/food-api/domain/dtos"
	"github.com/ronnachate/foodstore/food-api/domain/mocks"
	"github.com/ronnachate/foodstore/food-api/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewOrder(t *testing.T) {
	mockOrderRepository := new(mocks.OrderRepository)
	mockProductRepository := new(mocks.ProductRepository)
	mockOrderDiscountRepository := new(mocks.OrderDiscountRepository)
	productID1 := uuid.New()
	productID2 := uuid.New()

	t.Run("success", func(t *testing.T) {

		mockProducts := []domain.Product{
			{ID: productID1, Name: "test-product", Price: 100},
			{ID: productID2, Name: "test-product", Price: 50},
		}

		mockProductRepository.On("GetProducts", mock.Anything, []uuid.UUID{productID1, productID2}).Return(mockProducts, nil).Once()
		order := domain.Order{
			TotalPrice: 150,
		}
		mockOrderRepository.On("NewOrder", mock.Anything, mock.Anything).Return(order, nil).Once()

		u := usecase.NewOrderUsecase(mockOrderRepository, mockProductRepository, mockOrderDiscountRepository, time.Second*2)
		orderRequest := dtos.OrderDTO{
			MemberID: uuid.Nil,
			Items: []dtos.OrderItemDTO{
				{ProductID: productID1, Quantity: 1},
				{ProductID: productID2, Quantity: 1},
			},
		}
		_, err := u.NewOrder(context.Background(), orderRequest)

		assert.NoError(t, err)

		mockOrderRepository.AssertExpectations(t)
	})
}

func TestCalculateOrder(t *testing.T) {
	mockOrderRepository := new(mocks.OrderRepository)
	mockProductRepository := new(mocks.ProductRepository)
	mockOrderDiscountRepository := new(mocks.OrderDiscountRepository)
	productID1 := uuid.New()
	productID2 := uuid.New()

	t.Run("success", func(t *testing.T) {

		u := usecase.NewOrderUsecase(mockOrderRepository, mockProductRepository, mockOrderDiscountRepository, time.Second*2)

		mockProducts := []domain.Product{
			{ID: productID1, Name: "test-product", Price: 100},
			{ID: productID2, Name: "test-product", Price: 50},
		}

		orderRequest := dtos.OrderDTO{
			MemberID: uuid.Nil,
			Items: []dtos.OrderItemDTO{
				{ProductID: productID1, Quantity: 1},
				{ProductID: productID2, Quantity: 2},
			},
		}

		order := domain.Order{
			Items: make([]domain.OrderItem, 0),
		}

		u.CalculateOrder(&order, orderRequest, mockProducts)
		assert.NotNil(t, order)
		assert.Equal(t, order.TotalPrice, float64(200))

	})

	t.Run("success with item min discount - percentage discount", func(t *testing.T) {

		u := usecase.NewOrderUsecase(mockOrderRepository, mockProductRepository, mockOrderDiscountRepository, time.Second*2)

		mockProducts := []domain.Product{
			{
				ID: productID1, Name: "test-product", Price: 100,
				Discounts: []domain.ProductDiscount{
					{ItemDiscount: &domain.ItemDiscount{
						ID: 1, Type: constant.MIN_DISCOUNT_TYPE, Min: 2, DiscountType: constant.PERCENTAGE_DISCOUNT_TYPE, DiscountValue: 10}}},
			},
		}

		orderRequest := dtos.OrderDTO{
			MemberID: uuid.Nil,
			Items: []dtos.OrderItemDTO{
				{ProductID: productID1, Quantity: 2},
			},
		}

		order := domain.Order{
			Items: make([]domain.OrderItem, 0),
		}

		u.CalculateOrder(&order, orderRequest, mockProducts)
		assert.NotNil(t, order)
		assert.Equal(t, order.TotalPrice, float64(180))

	})

	t.Run("success with item min discount - value discount", func(t *testing.T) {

		u := usecase.NewOrderUsecase(mockOrderRepository, mockProductRepository, mockOrderDiscountRepository, time.Second*2)

		mockProducts := []domain.Product{
			{
				ID: productID1, Name: "test-product", Price: 100,
				Discounts: []domain.ProductDiscount{
					{ItemDiscount: &domain.ItemDiscount{
						ID: 1, Type: constant.MIN_DISCOUNT_TYPE, Min: 2, DiscountType: constant.PRICE_DISCOUNT_TYPE, DiscountValue: 10}}},
			},
		}

		orderRequest := dtos.OrderDTO{
			MemberID: uuid.Nil,
			Items: []dtos.OrderItemDTO{
				{ProductID: productID1, Quantity: 2},
			},
		}

		order := domain.Order{
			Items: make([]domain.OrderItem, 0),
		}

		u.CalculateOrder(&order, orderRequest, mockProducts)
		assert.NotNil(t, order)
		assert.Equal(t, order.TotalPrice, float64(190))

	})

	t.Run("success with item pair discount - percentage discount", func(t *testing.T) {

		u := usecase.NewOrderUsecase(mockOrderRepository, mockProductRepository, mockOrderDiscountRepository, time.Second*2)

		mockProducts := []domain.Product{
			{
				ID: productID1, Name: "test-product", Price: 100,
				Discounts: []domain.ProductDiscount{
					{ItemDiscount: &domain.ItemDiscount{
						ID: 1, Type: constant.PAIR_DISCOUNT_TYPE, Min: 2, DiscountType: constant.PERCENTAGE_DISCOUNT_TYPE, DiscountValue: 10}}},
			},
		}

		orderRequest := dtos.OrderDTO{
			MemberID: uuid.Nil,
			Items: []dtos.OrderItemDTO{
				{ProductID: productID1, Quantity: 3},
			},
		}

		order := domain.Order{
			Items: make([]domain.OrderItem, 0),
		}

		//1 pair discount
		u.CalculateOrder(&order, orderRequest, mockProducts)
		assert.NotNil(t, order)
		// discount 10% from 1 pair(100 * 2)
		assert.Equal(t, float64(280), order.TotalPrice)

	})

	t.Run("success with item pair discount - value discount", func(t *testing.T) {

		u := usecase.NewOrderUsecase(mockOrderRepository, mockProductRepository, mockOrderDiscountRepository, time.Second*2)

		mockProducts := []domain.Product{
			{
				ID: productID1, Name: "test-product", Price: 100,
				Discounts: []domain.ProductDiscount{
					{ItemDiscount: &domain.ItemDiscount{
						ID: 1, Type: constant.PAIR_DISCOUNT_TYPE, Min: 2, DiscountType: constant.PRICE_DISCOUNT_TYPE, DiscountValue: 10}}},
			},
		}

		orderRequest := dtos.OrderDTO{
			MemberID: uuid.Nil,
			Items: []dtos.OrderItemDTO{
				{ProductID: productID1, Quantity: 4},
			},
		}

		order := domain.Order{
			Items: make([]domain.OrderItem, 0),
		}

		//2 pair discount
		u.CalculateOrder(&order, orderRequest, mockProducts)
		assert.NotNil(t, order)
		assert.Equal(t, float64(380), order.TotalPrice)

	})
}

func TestGetByID(t *testing.T) {
	mockOrderRepository := new(mocks.OrderRepository)
	mockProductRepository := new(mocks.ProductRepository)
	mockOrderDiscountRepository := new(mocks.OrderDiscountRepository)
	orderID := uuid.UUID{}

	t.Run("success", func(t *testing.T) {

		mockOrder := domain.Order{
			ID: orderID,
		}

		mockOrderRepository.On("GetByID", mock.Anything, orderID.String()).Return(mockOrder, nil).Once()

		u := usecase.NewOrderUsecase(mockOrderRepository, mockProductRepository, mockOrderDiscountRepository, time.Second*2)

		order, err := u.GetByID(context.Background(), orderID.String())

		assert.NoError(t, err)
		assert.NotNil(t, order)
		assert.Equal(t, orderID, order.ID)

		mockOrderRepository.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {

		//need to mock return empty order due to 'ret.Get(0).(domain.Order)' error in generated file mocks/OrderRepository.go
		mockOrderRepository.On("GetByID", mock.Anything, orderID.String()).Return(domain.Order{}, errors.New("Unexpected")).Once()

		u := usecase.NewOrderUsecase(mockOrderRepository, mockProductRepository, mockOrderDiscountRepository, time.Second*2)

		_, err := u.GetByID(context.Background(), orderID.String())

		assert.Error(t, err)

		mockOrderRepository.AssertExpectations(t)
	})
}
