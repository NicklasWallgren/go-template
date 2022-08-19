// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entities "github.com/NicklasWallgren/go-template/domain/users/entities"
	mock "github.com/stretchr/testify/mock"

	response "github.com/NicklasWallgren/go-template/adapters/driven/api/users/response"

	users "github.com/NicklasWallgren/go-template/adapters/driven/api/users"
)

// UserApiConverter is an autogenerated mock type for the UserApiConverter type
type UserApiConverter struct {
	mock.Mock
}

type UserApiConverter_Expecter struct {
	mock *mock.Mock
}

func (_m *UserApiConverter) EXPECT() *UserApiConverter_Expecter {
	return &UserApiConverter_Expecter{mock: &_m.Mock}
}

// EntityOf provides a mock function with given fields: _a0
func (_m *UserApiConverter) EntityOf(_a0 *users.CreateUserRequest) entities.User {
	ret := _m.Called(_a0)

	var r0 entities.User
	if rf, ok := ret.Get(0).(func(*users.CreateUserRequest) entities.User); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(entities.User)
	}

	return r0
}

// UserApiConverter_EntityOf_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EntityOf'
type UserApiConverter_EntityOf_Call struct {
	*mock.Call
}

// EntityOf is a helper method to define mock.On call
//  - _a0 *users.CreateUserRequest
func (_e *UserApiConverter_Expecter) EntityOf(_a0 interface{}) *UserApiConverter_EntityOf_Call {
	return &UserApiConverter_EntityOf_Call{Call: _e.mock.On("EntityOf", _a0)}
}

func (_c *UserApiConverter_EntityOf_Call) Run(run func(_a0 *users.CreateUserRequest)) *UserApiConverter_EntityOf_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*users.CreateUserRequest))
	})
	return _c
}

func (_c *UserApiConverter_EntityOf_Call) Return(_a0 entities.User) *UserApiConverter_EntityOf_Call {
	_c.Call.Return(_a0)
	return _c
}

// ResponseOf provides a mock function with given fields: _a0
func (_m *UserApiConverter) ResponseOf(_a0 *entities.User) response.UserResponse {
	ret := _m.Called(_a0)

	var r0 response.UserResponse
	if rf, ok := ret.Get(0).(func(*entities.User) response.UserResponse); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(response.UserResponse)
	}

	return r0
}

// UserApiConverter_ResponseOf_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ResponseOf'
type UserApiConverter_ResponseOf_Call struct {
	*mock.Call
}

// ResponseOf is a helper method to define mock.On call
//  - _a0 *entities.User
func (_e *UserApiConverter_Expecter) ResponseOf(_a0 interface{}) *UserApiConverter_ResponseOf_Call {
	return &UserApiConverter_ResponseOf_Call{Call: _e.mock.On("ResponseOf", _a0)}
}

func (_c *UserApiConverter_ResponseOf_Call) Run(run func(_a0 *entities.User)) *UserApiConverter_ResponseOf_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*entities.User))
	})
	return _c
}

func (_c *UserApiConverter_ResponseOf_Call) Return(_a0 response.UserResponse) *UserApiConverter_ResponseOf_Call {
	_c.Call.Return(_a0)
	return _c
}

// UpdatedEntityOf provides a mock function with given fields: ctx, request
func (_m *UserApiConverter) UpdatedEntityOf(ctx context.Context, request *users.UpdateUserRequest) (*entities.User, error) {
	ret := _m.Called(ctx, request)

	var r0 *entities.User
	if rf, ok := ret.Get(0).(func(context.Context, *users.UpdateUserRequest) *entities.User); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *users.UpdateUserRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserApiConverter_UpdatedEntityOf_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdatedEntityOf'
type UserApiConverter_UpdatedEntityOf_Call struct {
	*mock.Call
}

// UpdatedEntityOf is a helper method to define mock.On call
//  - ctx context.Context
//  - request *users.UpdateUserRequest
func (_e *UserApiConverter_Expecter) UpdatedEntityOf(ctx interface{}, request interface{}) *UserApiConverter_UpdatedEntityOf_Call {
	return &UserApiConverter_UpdatedEntityOf_Call{Call: _e.mock.On("UpdatedEntityOf", ctx, request)}
}

func (_c *UserApiConverter_UpdatedEntityOf_Call) Run(run func(ctx context.Context, request *users.UpdateUserRequest)) *UserApiConverter_UpdatedEntityOf_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*users.UpdateUserRequest))
	})
	return _c
}

func (_c *UserApiConverter_UpdatedEntityOf_Call) Return(_a0 *entities.User, _a1 error) *UserApiConverter_UpdatedEntityOf_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewUserApiConverter interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserApiConverter creates a new instance of UserApiConverter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserApiConverter(t mockConstructorTestingTNewUserApiConverter) *UserApiConverter {
	mock := &UserApiConverter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
