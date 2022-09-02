// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	"github.com/NicklasWallgren/go-template/adapters/driver/api/request"
	common "github.com/NicklasWallgren/go-template/domain/common"

	mock "github.com/stretchr/testify/mock"
)

// APIRequestCreateConverter is an autogenerated mock type for the APIRequestCreateConverter type
type APIRequestCreateConverter[T request.APIRequest, R common.EntityConstraint] struct {
	mock.Mock
}

type APIRequestCreateConverter_Expecter[T request.APIRequest, R common.EntityConstraint] struct {
	mock *mock.Mock
}

func (_m *APIRequestCreateConverter[T, R]) EXPECT() *APIRequestCreateConverter_Expecter[T, R] {
	return &APIRequestCreateConverter_Expecter[T, R]{mock: &_m.Mock}
}

// EntityOf provides a mock function with given fields: _a0
func (_m *APIRequestCreateConverter[T, R]) EntityOf(_a0 *T) R {
	ret := _m.Called(_a0)

	var r0 R
	if rf, ok := ret.Get(0).(func(*T) R); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(R)
	}

	return r0
}

// APIRequestCreateConverter_EntityOf_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EntityOf'
type APIRequestCreateConverter_EntityOf_Call[T request.APIRequest, R common.EntityConstraint] struct {
	*mock.Call
}

// EntityOf is a helper method to define mock.On call
//   - _a0 *T
func (_e *APIRequestCreateConverter_Expecter[T, R]) EntityOf(_a0 interface{}) *APIRequestCreateConverter_EntityOf_Call[T, R] {
	return &APIRequestCreateConverter_EntityOf_Call[T, R]{Call: _e.mock.On("EntityOf", _a0)}
}

func (_c *APIRequestCreateConverter_EntityOf_Call[T, R]) Run(run func(_a0 *T)) *APIRequestCreateConverter_EntityOf_Call[T, R] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*T))
	})
	return _c
}

func (_c *APIRequestCreateConverter_EntityOf_Call[T, R]) Return(_a0 R) *APIRequestCreateConverter_EntityOf_Call[T, R] {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewAPIRequestCreateConverter interface {
	mock.TestingT
	Cleanup(func())
}

// NewAPIRequestCreateConverter creates a new instance of APIRequestCreateConverter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAPIRequestCreateConverter[T request.APIRequest, R common.EntityConstraint](t mockConstructorTestingTNewAPIRequestCreateConverter) *APIRequestCreateConverter[T, R] {
	mock := &APIRequestCreateConverter[T, R]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
