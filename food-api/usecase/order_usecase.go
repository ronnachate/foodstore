package usecase

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/ronnachate/foodstore/food-api/domain"
	"github.com/ronnachate/foodstore/food-api/domain/dtos"
)

type orderUsecase struct {
	orderRepository domain.OrderRepository
	contextTimeout  time.Duration
}

func NewOrderUsecase(orderRepository domain.OrderRepository, timeout time.Duration) domain.OrderUsecase {
	return &orderUsecase{
		orderRepository: orderRepository,
		contextTimeout:  timeout,
	}
}

func (ou *orderUsecase) NewOrder(c context.Context, dto dtos.OrderDTO) error {
	ctx, cancel := context.WithTimeout(c, ou.contextTimeout)
	defer cancel()
	//converto to domain.Order

	order := domain.Order{}
	if dto.MemberID == uuid.Nil {
		order.MemberID = sql.NullString{}
	}
	//find order discount
	return ou.orderRepository.NewOrder(ctx, &order)
}

func (ou *orderUsecase) GetByID(c context.Context, orderID string) (domain.Order, error) {
	ctx, cancel := context.WithTimeout(c, ou.contextTimeout)
	defer cancel()
	return ou.orderRepository.GetByID(ctx, orderID)
}
