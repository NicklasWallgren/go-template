// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	common "github.com/NicklasWallgren/go-template/domain/common"

	mock "github.com/stretchr/testify/mock"
)

// EntityUpdateValidator is an autogenerated mock type for the EntityUpdateValidator type
type EntityUpdateValidator[T common.EntityConstraint] struct {
	mock.Mock
}

type EntityUpdateValidator_Expecter[T common.EntityConstraint] struct {
	mock *mock.Mock
}

func (_m *EntityUpdateValidator[T]) EXPECT() *EntityUpdateValidator_Expecter[T] {
	return &EntityUpdateValidator_Expecter[T]{mock: &_m.Mock}
}

// ValidateToUpdate provides a mock function with given fields: ctx, origin, updated
func (_m *EntityUpdateValidator[T]) ValidateToUpdate(ctx context.Context, origin *T, updated *T) error {
	ret := _m.Called(ctx, origin, updated)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *T, *T) error); ok {
		r0 = rf(ctx, origin, updated)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EntityUpdateValidator_ValidateToUpdate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ValidateToUpdate'
type EntityUpdateValidator_ValidateToUpdate_Call[T common.EntityConstraint] struct {
	*mock.Call
}

// ValidateToUpdate is a helper method to define mock.On call
//  - ctx context.Context
//  - origin *T
//  - updated *T
func (_e *EntityUpdateValidator_Expecter[T]) ValidateToUpdate(ctx interface{}, origin interface{}, updated interface{}) *EntityUpdateValidator_ValidateToUpdate_Call[T] {
	return &EntityUpdateValidator_ValidateToUpdate_Call[T]{Call: _e.mock.On("ValidateToUpdate", ctx, origin, updated)}
}

func (_c *EntityUpdateValidator_ValidateToUpdate_Call[T]) Run(run func(ctx context.Context, origin *T, updated *T)) *EntityUpdateValidator_ValidateToUpdate_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*T), args[2].(*T))
	})
	return _c
}

func (_c *EntityUpdateValidator_ValidateToUpdate_Call[T]) Return(_a0 error) *EntityUpdateValidator_ValidateToUpdate_Call[T] {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewEntityUpdateValidator interface {
	mock.TestingT
	Cleanup(func())
}

// NewEntityUpdateValidator creates a new instance of EntityUpdateValidator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewEntityUpdateValidator[T common.EntityConstraint](t mockConstructorTestingTNewEntityUpdateValidator) *EntityUpdateValidator[T] {
	mock := &EntityUpdateValidator[T]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
