// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/ronnachate/foodstore/food-api/domain"
	mock "github.com/stretchr/testify/mock"
)

// OrderRepository is an autogenerated mock type for the OrderRepository type
type OrderRepository struct {
	mock.Mock
}

// GetByID provides a mock function with given fields: c, id
func (_m *OrderRepository) GetByID(c context.Context, id string) (domain.Order, error) {
	ret := _m.Called(c, id)

	var r0 domain.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (domain.Order, error)); ok {
		return rf(c, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.Order); ok {
		r0 = rf(c, id)
	} else {
		r0 = ret.Get(0).(domain.Order)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(c, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewOrder provides a mock function with given fields: c, order
func (_m *OrderRepository) NewOrder(c context.Context, order *domain.Order) error {
	ret := _m.Called(c, order)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Order) error); ok {
		r0 = rf(c, order)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewOrderRepository creates a new instance of OrderRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOrderRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *OrderRepository {
	mock := &OrderRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
