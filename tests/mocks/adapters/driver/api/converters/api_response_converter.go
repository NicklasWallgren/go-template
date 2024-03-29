// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	common "github.com/NicklasWallgren/go-template/domain/common"

	mock "github.com/stretchr/testify/mock"

	response "github.com/NicklasWallgren/go-template/adapters/driver/api/response"
)

// APIResponseConverter is an autogenerated mock type for the APIResponseConverter type
type APIResponseConverter[T common.EntityConstraint, R response.APIResponse] struct {
	mock.Mock
}

type APIResponseConverter_Expecter[T common.EntityConstraint, R response.APIResponse] struct {
	mock *mock.Mock
}

func (_m *APIResponseConverter[T, R]) EXPECT() *APIResponseConverter_Expecter[T, R] {
	return &APIResponseConverter_Expecter[T, R]{mock: &_m.Mock}
}

// ResponseOf provides a mock function with given fields: _a0
func (_m *APIResponseConverter[T, R]) ResponseOf(_a0 T) R {
	ret := _m.Called(_a0)

	var r0 R
	if rf, ok := ret.Get(0).(func(T) R); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(R)
	}

	return r0
}

// APIResponseConverter_ResponseOf_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ResponseOf'
type APIResponseConverter_ResponseOf_Call[T common.EntityConstraint, R response.APIResponse] struct {
	*mock.Call
}

// ResponseOf is a helper method to define mock.On call
//  - _a0 T
func (_e *APIResponseConverter_Expecter[T, R]) ResponseOf(_a0 interface{}) *APIResponseConverter_ResponseOf_Call[T, R] {
	return &APIResponseConverter_ResponseOf_Call[T, R]{Call: _e.mock.On("ResponseOf", _a0)}
}

func (_c *APIResponseConverter_ResponseOf_Call[T, R]) Run(run func(_a0 T)) *APIResponseConverter_ResponseOf_Call[T, R] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(T))
	})
	return _c
}

func (_c *APIResponseConverter_ResponseOf_Call[T, R]) Return(_a0 R) *APIResponseConverter_ResponseOf_Call[T, R] {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewAPIResponseConverter interface {
	mock.TestingT
	Cleanup(func())
}

// NewAPIResponseConverter creates a new instance of APIResponseConverter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAPIResponseConverter[T common.EntityConstraint, R response.APIResponse](t mockConstructorTestingTNewAPIResponseConverter) *APIResponseConverter[T, R] {
	mock := &APIResponseConverter[T, R]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
