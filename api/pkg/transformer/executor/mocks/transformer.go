// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	types "github.com/gojek/merlin/pkg/transformer/types"
)

// Transformer is an autogenerated mock type for the Transformer type
type Transformer struct {
	mock.Mock
}

// Execute provides a mock function with given fields: ctx, requestBody, requestHeaders
func (_m *Transformer) Execute(ctx context.Context, requestBody types.Payload, requestHeaders map[string]string) (*types.PredictResponse, error) {
	ret := _m.Called(ctx, requestBody, requestHeaders)

	var r0 *types.PredictResponse
	if rf, ok := ret.Get(0).(func(context.Context, types.Payload, map[string]string) *types.PredictResponse); ok {
		r0 = rf(ctx, requestBody, requestHeaders)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.PredictResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, types.Payload, map[string]string) error); ok {
		r1 = rf(ctx, requestBody, requestHeaders)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewTransformer interface {
	mock.TestingT
	Cleanup(func())
}

// NewTransformer creates a new instance of Transformer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTransformer(t mockConstructorTestingTNewTransformer) *Transformer {
	mock := &Transformer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
