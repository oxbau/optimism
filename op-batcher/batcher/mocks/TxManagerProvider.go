// Code generated by mockery v2.22.1. DO NOT EDIT.

package mocks

import (
	context "context"
	big "math/big"

	common "github.com/ethereum/go-ethereum/common"

	mock "github.com/stretchr/testify/mock"

	types "github.com/ethereum/go-ethereum/core/types"
)

// TxManagerProvider is an autogenerated mock type for the TxManagerProvider type
type TxManagerProvider struct {
	mock.Mock
}

// HeaderByNumber provides a mock function with given fields: ctx, number
func (_m *TxManagerProvider) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	ret := _m.Called(ctx, number)

	var r0 *types.Header
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int) (*types.Header, error)); ok {
		return rf(ctx, number)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int) *types.Header); ok {
		r0 = rf(ctx, number)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Header)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *big.Int) error); ok {
		r1 = rf(ctx, number)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NonceAt provides a mock function with given fields: ctx, account, blockNumber
func (_m *TxManagerProvider) NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error) {
	ret := _m.Called(ctx, account, blockNumber)

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, *big.Int) (uint64, error)); ok {
		return rf(ctx, account, blockNumber)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, *big.Int) uint64); ok {
		r0 = rf(ctx, account, blockNumber)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Address, *big.Int) error); ok {
		r1 = rf(ctx, account, blockNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SuggestGasTipCap provides a mock function with given fields: ctx
func (_m *TxManagerProvider) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	ret := _m.Called(ctx)

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*big.Int, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *big.Int); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewTxManagerProvider interface {
	mock.TestingT
	Cleanup(func())
}

// NewTxManagerProvider creates a new instance of TxManagerProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTxManagerProvider(t mockConstructorTestingTNewTxManagerProvider) *TxManagerProvider {
	mock := &TxManagerProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
