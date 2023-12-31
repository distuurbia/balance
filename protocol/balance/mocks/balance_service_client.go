// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	context "context"

	balance "github.com/distuurbia/balance/protocol/balance"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"
)

// BalanceServiceClient is an autogenerated mock type for the BalanceServiceClient type
type BalanceServiceClient struct {
	mock.Mock
}

// AddBalanceChange provides a mock function with given fields: ctx, in, opts
func (_m *BalanceServiceClient) AddBalanceChange(ctx context.Context, in *balance.AddBalanceChangeRequest, opts ...grpc.CallOption) (*balance.AddBalanceChangeResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *balance.AddBalanceChangeResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *balance.AddBalanceChangeRequest, ...grpc.CallOption) (*balance.AddBalanceChangeResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *balance.AddBalanceChangeRequest, ...grpc.CallOption) *balance.AddBalanceChangeResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*balance.AddBalanceChangeResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *balance.AddBalanceChangeRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteProfilesBalance provides a mock function with given fields: ctx, in, opts
func (_m *BalanceServiceClient) DeleteProfilesBalance(ctx context.Context, in *balance.DeleteProfilesBalanceRequest, opts ...grpc.CallOption) (*balance.DeleteProfilesBalanceResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *balance.DeleteProfilesBalanceResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *balance.DeleteProfilesBalanceRequest, ...grpc.CallOption) (*balance.DeleteProfilesBalanceResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *balance.DeleteProfilesBalanceRequest, ...grpc.CallOption) *balance.DeleteProfilesBalanceResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*balance.DeleteProfilesBalanceResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *balance.DeleteProfilesBalanceRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBalance provides a mock function with given fields: ctx, in, opts
func (_m *BalanceServiceClient) GetBalance(ctx context.Context, in *balance.GetBalanceRequest, opts ...grpc.CallOption) (*balance.GetBalanceResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *balance.GetBalanceResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *balance.GetBalanceRequest, ...grpc.CallOption) (*balance.GetBalanceResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *balance.GetBalanceRequest, ...grpc.CallOption) *balance.GetBalanceResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*balance.GetBalanceResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *balance.GetBalanceRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewBalanceServiceClient creates a new instance of BalanceServiceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBalanceServiceClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *BalanceServiceClient {
	mock := &BalanceServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
