// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entities "github.com/NicklasWallgren/go-template/domain/users/entities"
	gorm "gorm.io/gorm"

	mock "github.com/stretchr/testify/mock"

	models "github.com/NicklasWallgren/go-template/adapters/driven/persistence/models"

	users "github.com/NicklasWallgren/go-template/adapters/driven/persistence/users"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

type UserRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *UserRepository) EXPECT() *UserRepository_Expecter {
	return &UserRepository_Expecter{mock: &_m.Mock}
}

// Count provides a mock function with given fields: ctx
func (_m *UserRepository) Count(ctx context.Context) (int64, error) {
	ret := _m.Called(ctx)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context) int64); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserRepository_Count_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Count'
type UserRepository_Count_Call struct {
	*mock.Call
}

// Count is a helper method to define mock.On call
//  - ctx context.Context
func (_e *UserRepository_Expecter) Count(ctx interface{}) *UserRepository_Count_Call {
	return &UserRepository_Count_Call{Call: _e.mock.On("Count", ctx)}
}

func (_c *UserRepository_Count_Call) Run(run func(ctx context.Context)) *UserRepository_Count_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *UserRepository_Count_Call) Return(_a0 int64, _a1 error) *UserRepository_Count_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Create provides a mock function with given fields: ctx, user
func (_m *UserRepository) Create(ctx context.Context, user *entities.User) (*entities.User, error) {
	ret := _m.Called(ctx, user)

	var r0 *entities.User
	if rf, ok := ret.Get(0).(func(context.Context, *entities.User) *entities.User); ok {
		r0 = rf(ctx, user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *entities.User) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserRepository_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type UserRepository_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//  - ctx context.Context
//  - user *entities.User
func (_e *UserRepository_Expecter) Create(ctx interface{}, user interface{}) *UserRepository_Create_Call {
	return &UserRepository_Create_Call{Call: _e.mock.On("Create", ctx, user)}
}

func (_c *UserRepository_Create_Call) Run(run func(ctx context.Context, user *entities.User)) *UserRepository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*entities.User))
	})
	return _c
}

func (_c *UserRepository_Create_Call) Return(_a0 *entities.User, _a1 error) *UserRepository_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// DeleteByID provides a mock function with given fields: ctx, id
func (_m *UserRepository) DeleteByID(ctx context.Context, id uint) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserRepository_DeleteByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteByID'
type UserRepository_DeleteByID_Call struct {
	*mock.Call
}

// DeleteByID is a helper method to define mock.On call
//  - ctx context.Context
//  - id uint
func (_e *UserRepository_Expecter) DeleteByID(ctx interface{}, id interface{}) *UserRepository_DeleteByID_Call {
	return &UserRepository_DeleteByID_Call{Call: _e.mock.On("DeleteByID", ctx, id)}
}

func (_c *UserRepository_DeleteByID_Call) Run(run func(ctx context.Context, id uint)) *UserRepository_DeleteByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint))
	})
	return _c
}

func (_c *UserRepository_DeleteByID_Call) Return(_a0 error) *UserRepository_DeleteByID_Call {
	_c.Call.Return(_a0)
	return _c
}

// FindAll provides a mock function with given fields: ctx, pagination
func (_m *UserRepository) FindAll(ctx context.Context, pagination *models.Pagination) (*models.Page[*entities.User], error) {
	ret := _m.Called(ctx, pagination)

	var r0 *models.Page[*entities.User]
	if rf, ok := ret.Get(0).(func(context.Context, *models.Pagination) *models.Page[*entities.User]); ok {
		r0 = rf(ctx, pagination)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Page[*entities.User])
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.Pagination) error); ok {
		r1 = rf(ctx, pagination)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserRepository_FindAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindAll'
type UserRepository_FindAll_Call struct {
	*mock.Call
}

// FindAll is a helper method to define mock.On call
//  - ctx context.Context
//  - pagination *models.Pagination
func (_e *UserRepository_Expecter) FindAll(ctx interface{}, pagination interface{}) *UserRepository_FindAll_Call {
	return &UserRepository_FindAll_Call{Call: _e.mock.On("FindAll", ctx, pagination)}
}

func (_c *UserRepository_FindAll_Call) Run(run func(ctx context.Context, pagination *models.Pagination)) *UserRepository_FindAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*models.Pagination))
	})
	return _c
}

func (_c *UserRepository_FindAll_Call) Return(page *models.Page[*entities.User], err error) *UserRepository_FindAll_Call {
	_c.Call.Return(page, err)
	return _c
}

// FindOneByEmailWithExclusiveLock provides a mock function with given fields: ctx, email
func (_m *UserRepository) FindOneByEmailWithExclusiveLock(ctx context.Context, email string) (*entities.User, error) {
	ret := _m.Called(ctx, email)

	var r0 *entities.User
	if rf, ok := ret.Get(0).(func(context.Context, string) *entities.User); ok {
		r0 = rf(ctx, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserRepository_FindOneByEmailWithExclusiveLock_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindOneByEmailWithExclusiveLock'
type UserRepository_FindOneByEmailWithExclusiveLock_Call struct {
	*mock.Call
}

// FindOneByEmailWithExclusiveLock is a helper method to define mock.On call
//  - ctx context.Context
//  - email string
func (_e *UserRepository_Expecter) FindOneByEmailWithExclusiveLock(ctx interface{}, email interface{}) *UserRepository_FindOneByEmailWithExclusiveLock_Call {
	return &UserRepository_FindOneByEmailWithExclusiveLock_Call{Call: _e.mock.On("FindOneByEmailWithExclusiveLock", ctx, email)}
}

func (_c *UserRepository_FindOneByEmailWithExclusiveLock_Call) Run(run func(ctx context.Context, email string)) *UserRepository_FindOneByEmailWithExclusiveLock_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *UserRepository_FindOneByEmailWithExclusiveLock_Call) Return(_a0 *entities.User, _a1 error) *UserRepository_FindOneByEmailWithExclusiveLock_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// FindOneByID provides a mock function with given fields: ctx, id
func (_m *UserRepository) FindOneByID(ctx context.Context, id uint) (*entities.User, error) {
	ret := _m.Called(ctx, id)

	var r0 *entities.User
	if rf, ok := ret.Get(0).(func(context.Context, uint) *entities.User); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserRepository_FindOneByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindOneByID'
type UserRepository_FindOneByID_Call struct {
	*mock.Call
}

// FindOneByID is a helper method to define mock.On call
//  - ctx context.Context
//  - id uint
func (_e *UserRepository_Expecter) FindOneByID(ctx interface{}, id interface{}) *UserRepository_FindOneByID_Call {
	return &UserRepository_FindOneByID_Call{Call: _e.mock.On("FindOneByID", ctx, id)}
}

func (_c *UserRepository_FindOneByID_Call) Run(run func(ctx context.Context, id uint)) *UserRepository_FindOneByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint))
	})
	return _c
}

func (_c *UserRepository_FindOneByID_Call) Return(user *entities.User, err error) *UserRepository_FindOneByID_Call {
	_c.Call.Return(user, err)
	return _c
}

// FindOneByIDForUpdate provides a mock function with given fields: ctx, id
func (_m *UserRepository) FindOneByIDForUpdate(ctx context.Context, id uint) (*entities.User, error) {
	ret := _m.Called(ctx, id)

	var r0 *entities.User
	if rf, ok := ret.Get(0).(func(context.Context, uint) *entities.User); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserRepository_FindOneByIDForUpdate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindOneByIDForUpdate'
type UserRepository_FindOneByIDForUpdate_Call struct {
	*mock.Call
}

// FindOneByIDForUpdate is a helper method to define mock.On call
//  - ctx context.Context
//  - id uint
func (_e *UserRepository_Expecter) FindOneByIDForUpdate(ctx interface{}, id interface{}) *UserRepository_FindOneByIDForUpdate_Call {
	return &UserRepository_FindOneByIDForUpdate_Call{Call: _e.mock.On("FindOneByIDForUpdate", ctx, id)}
}

func (_c *UserRepository_FindOneByIDForUpdate_Call) Run(run func(ctx context.Context, id uint)) *UserRepository_FindOneByIDForUpdate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint))
	})
	return _c
}

func (_c *UserRepository_FindOneByIDForUpdate_Call) Return(_a0 *entities.User, _a1 error) *UserRepository_FindOneByIDForUpdate_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Save provides a mock function with given fields: ctx, user
func (_m *UserRepository) Save(ctx context.Context, user *entities.User) (*entities.User, error) {
	ret := _m.Called(ctx, user)

	var r0 *entities.User
	if rf, ok := ret.Get(0).(func(context.Context, *entities.User) *entities.User); ok {
		r0 = rf(ctx, user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *entities.User) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserRepository_Save_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Save'
type UserRepository_Save_Call struct {
	*mock.Call
}

// Save is a helper method to define mock.On call
//  - ctx context.Context
//  - user *entities.User
func (_e *UserRepository_Expecter) Save(ctx interface{}, user interface{}) *UserRepository_Save_Call {
	return &UserRepository_Save_Call{Call: _e.mock.On("Save", ctx, user)}
}

func (_c *UserRepository_Save_Call) Run(run func(ctx context.Context, user *entities.User)) *UserRepository_Save_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*entities.User))
	})
	return _c
}

func (_c *UserRepository_Save_Call) Return(_a0 *entities.User, _a1 error) *UserRepository_Save_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// TransactWithDefaultRetry provides a mock function with given fields: ctx, operation
func (_m *UserRepository) TransactWithDefaultRetry(ctx context.Context, operation func(*gorm.DB) error) error {
	ret := _m.Called(ctx, operation)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, func(*gorm.DB) error) error); ok {
		r0 = rf(ctx, operation)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserRepository_TransactWithDefaultRetry_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TransactWithDefaultRetry'
type UserRepository_TransactWithDefaultRetry_Call struct {
	*mock.Call
}

// TransactWithDefaultRetry is a helper method to define mock.On call
//  - ctx context.Context
//  - operation func(*gorm.DB) error
func (_e *UserRepository_Expecter) TransactWithDefaultRetry(ctx interface{}, operation interface{}) *UserRepository_TransactWithDefaultRetry_Call {
	return &UserRepository_TransactWithDefaultRetry_Call{Call: _e.mock.On("TransactWithDefaultRetry", ctx, operation)}
}

func (_c *UserRepository_TransactWithDefaultRetry_Call) Run(run func(ctx context.Context, operation func(*gorm.DB) error)) *UserRepository_TransactWithDefaultRetry_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(func(*gorm.DB) error))
	})
	return _c
}

func (_c *UserRepository_TransactWithDefaultRetry_Call) Return(_a0 error) *UserRepository_TransactWithDefaultRetry_Call {
	_c.Call.Return(_a0)
	return _c
}

// WithTx provides a mock function with given fields: tx
func (_m *UserRepository) WithTx(tx *gorm.DB) users.UserRepository {
	ret := _m.Called(tx)

	var r0 users.UserRepository
	if rf, ok := ret.Get(0).(func(*gorm.DB) users.UserRepository); ok {
		r0 = rf(tx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(users.UserRepository)
		}
	}

	return r0
}

// UserRepository_WithTx_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithTx'
type UserRepository_WithTx_Call struct {
	*mock.Call
}

// WithTx is a helper method to define mock.On call
//  - tx *gorm.DB
func (_e *UserRepository_Expecter) WithTx(tx interface{}) *UserRepository_WithTx_Call {
	return &UserRepository_WithTx_Call{Call: _e.mock.On("WithTx", tx)}
}

func (_c *UserRepository_WithTx_Call) Run(run func(tx *gorm.DB)) *UserRepository_WithTx_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gorm.DB))
	})
	return _c
}

func (_c *UserRepository_WithTx_Call) Return(_a0 users.UserRepository) *UserRepository_WithTx_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewUserRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserRepository creates a new instance of UserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserRepository(t mockConstructorTestingTNewUserRepository) *UserRepository {
	mock := &UserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
