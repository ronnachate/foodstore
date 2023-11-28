package usecase

import (
	"context"
	"database/sql"
	"math"
	"time"

	"github.com/google/uuid"
	"github.com/ronnachate/foodstore/food-api/domain"
	constant "github.com/ronnachate/foodstore/food-api/domain/constants"
	"github.com/ronnachate/foodstore/food-api/domain/dtos"
)

type orderUsecase struct {
	orderRepository         domain.OrderRepository
	orderDiscountRepository domain.OrderDiscountRepository
	productRepository       domain.ProductRepository
	contextTimeout          time.Duration
}

func NewOrderUsecase(
	orderRepository domain.OrderRepository,
	productRepository domain.ProductRepository,
	orderDiscountRepository domain.OrderDiscountRepository,
	timeout time.Duration) domain.OrderUsecase {
	return &orderUsecase{
		orderRepository:         orderRepository,
		orderDiscountRepository: orderDiscountRepository,
		productRepository:       productRepository,
		contextTimeout:          timeout,
	}
}

func (ou *orderUsecase) NewOrder(c context.Context, dto dtos.OrderDTO) (domain.Order, error) {
	ctx, cancel := context.WithTimeout(c, ou.contextTimeout)
	defer cancel()
	//converto to domain.Order

	order := domain.Order{
		Items: make([]domain.OrderItem, 0),
	}
	if dto.MemberID == uuid.Nil {
		order.MemberID = sql.NullString{}
	}
	productIDs := make([]uuid.UUID, 0)
	for _, item := range dto.Items {
		productIDs = append(productIDs, item.ProductID)
	}
	//get products list from db
	products, err := ou.productRepository.GetProducts(ctx, productIDs)
	if err != nil {
		return domain.Order{}, err
	}
	//append order items
	ou.CalculateOrder(&order, dto, products)

	//apply member discount
	ou.ApplyMemberDiscount(c, &order, dto)

	//create order
	return ou.orderRepository.NewOrder(ctx, &order)
}

func (ou *orderUsecase) GetByID(c context.Context, orderID string) (domain.Order, error) {
	ctx, cancel := context.WithTimeout(c, ou.contextTimeout)
	defer cancel()
	return ou.orderRepository.GetByID(ctx, orderID)
}

func (ou *orderUsecase) CalculateOrder(order *domain.Order, dto dtos.OrderDTO, products []domain.Product) {
	for _, item := range dto.Items {
		for _, product := range products {
			if item.ProductID == product.ID {
				//mapping order item
				orderItem := domain.OrderItem{
					ProductID:      product.ID,
					ItemDiscountID: sql.NullInt64{},
					Quantity:       item.Quantity,
					Price:          product.Price,
					TotalPrice:     (product.Price * float64(item.Quantity)),
				}
				if product.Discounts != nil {
					for _, discount := range product.Discounts {
						if discount.ItemDiscount.Type == constant.MIN_DISCOUNT_TYPE {
							if orderItem.Quantity >= discount.ItemDiscount.Min {
								//min discount found
								if discount.ItemDiscount.DiscountType == constant.PERCENTAGE_DISCOUNT_TYPE {
									orderItem.TotalPrice = orderItem.TotalPrice - ((orderItem.TotalPrice * discount.ItemDiscount.DiscountValue) / 100)
								} else if discount.ItemDiscount.DiscountType == constant.PRICE_DISCOUNT_TYPE {
									orderItem.TotalPrice = orderItem.TotalPrice - discount.ItemDiscount.DiscountValue
								}
								orderItem.ItemDiscountID.Int64 = int64(discount.ItemDiscount.ID)
							}
						} else if discount.ItemDiscount.Type == constant.PAIR_DISCOUNT_TYPE {
							if orderItem.Quantity >= 2 {
								//pair discount found
								//get pair quantity
								pairQuantity := float64(orderItem.Quantity / 2)

								//apply pair discount
								if discount.ItemDiscount.DiscountType == constant.PERCENTAGE_DISCOUNT_TYPE {
									orderItem.TotalPrice = orderItem.TotalPrice - (math.Floor(pairQuantity) * ((orderItem.Price * 2 * discount.ItemDiscount.DiscountValue) / 100))
								} else if discount.ItemDiscount.DiscountType == constant.PRICE_DISCOUNT_TYPE {
									orderItem.TotalPrice = orderItem.TotalPrice - (math.Floor(pairQuantity) * discount.ItemDiscount.DiscountValue)
								}
								orderItem.ItemDiscountID.Int64 = int64(discount.ItemDiscount.ID)
							}
						}
					}
				}
				order.Items = append(order.Items, orderItem)
				order.TotalPrice += orderItem.TotalPrice
			}
		}
	}
}

func (ou *orderUsecase) ApplyMemberDiscount(c context.Context, order *domain.Order, dto dtos.OrderDTO) {
	ctx, cancel := context.WithTimeout(c, ou.contextTimeout)
	defer cancel()
	if dto.MemberID != uuid.Nil {
		//find order discount
		orderDiscount, err := ou.orderDiscountRepository.GetByType(ctx, constant.MEMBER_DISCOUNT_TYPE)
		if err == nil {
			//apply discount
			discountAmount := 0.0
			if orderDiscount.DiscountType == constant.PERCENTAGE_DISCOUNT_TYPE {
				discountAmount = (order.TotalPrice * orderDiscount.DiscountValue) / 100
			} else if orderDiscount.DiscountType == constant.PRICE_DISCOUNT_TYPE {
				discountAmount = orderDiscount.DiscountValue
			}
			order.DiscountID.Int64 = int64(orderDiscount.ID)
			order.TotalPrice = order.TotalPrice - discountAmount
		}
	}
}
