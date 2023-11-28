// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/ronnachate/foodstore/food-api/domain"
	mock "github.com/stretchr/testify/mock"
)

// OrderDiscountRepository is an autogenerated mock type for the OrderDiscountRepository type
type OrderDiscountRepository struct {
	mock.Mock
}

// GetByType provides a mock function with given fields: c, id
func (_m *OrderDiscountRepository) GetByType(c context.Context, id uint64) (domain.OrderDiscount, error) {
	ret := _m.Called(c, id)

	var r0 domain.OrderDiscount
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) (domain.OrderDiscount, error)); ok {
		return rf(c, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) domain.OrderDiscount); ok {
		r0 = rf(c, id)
	} else {
		r0 = ret.Get(0).(domain.OrderDiscount)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(c, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewOrderDiscountRepository creates a new instance of OrderDiscountRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOrderDiscountRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *OrderDiscountRepository {
	mock := &OrderDiscountRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
