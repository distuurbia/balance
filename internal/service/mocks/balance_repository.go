// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// BalanceRepository is an autogenerated mock type for the BalanceRepository type
type BalanceRepository struct {
	mock.Mock
}

// AddBalanceChange provides a mock function with given fields: ctx, profileID, amount
func (_m *BalanceRepository) AddBalanceChange(ctx context.Context, profileID uuid.UUID, amount float64) error {
	ret := _m.Called(ctx, profileID, amount)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, float64) error); ok {
		r0 = rf(ctx, profileID, amount)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteProfilesBalance provides a mock function with given fields: ctx, profileID
func (_m *BalanceRepository) DeleteProfilesBalance(ctx context.Context, profileID uuid.UUID) error {
	ret := _m.Called(ctx, profileID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, profileID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetBalance provides a mock function with given fields: ctx, profileID
func (_m *BalanceRepository) GetBalance(ctx context.Context, profileID uuid.UUID) (float64, error) {
	ret := _m.Called(ctx, profileID)

	var r0 float64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (float64, error)); ok {
		return rf(ctx, profileID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) float64); ok {
		r0 = rf(ctx, profileID)
	} else {
		r0 = ret.Get(0).(float64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, profileID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewBalanceRepository creates a new instance of BalanceRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBalanceRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *BalanceRepository {
	mock := &BalanceRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
