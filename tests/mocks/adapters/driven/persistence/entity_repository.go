// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	common "github.com/NicklasWallgren/go-template/domain/common"

	gorm "gorm.io/gorm"

	mock "github.com/stretchr/testify/mock"

	models "github.com/NicklasWallgren/go-template/adapters/driven/persistence/models"

	persistence "github.com/NicklasWallgren/go-template/adapters/driven/persistence"
)

// EntityRepository is an autogenerated mock type for the EntityRepository type
type EntityRepository[T common.EntityConstraint] struct {
	mock.Mock
}

type EntityRepository_Expecter[T common.EntityConstraint] struct {
	mock *mock.Mock
}

func (_m *EntityRepository[T]) EXPECT() *EntityRepository_Expecter[T] {
	return &EntityRepository_Expecter[T]{mock: &_m.Mock}
}

// Count provides a mock function with given fields: ctx
func (_m *EntityRepository[T]) Count(ctx context.Context) (int64, error) {
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

// EntityRepository_Count_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Count'
type EntityRepository_Count_Call[T common.EntityConstraint] struct {
	*mock.Call
}

// Count is a helper method to define mock.On call
//  - ctx context.Context
func (_e *EntityRepository_Expecter[T]) Count(ctx interface{}) *EntityRepository_Count_Call[T] {
	return &EntityRepository_Count_Call[T]{Call: _e.mock.On("Count", ctx)}
}

func (_c *EntityRepository_Count_Call[T]) Run(run func(ctx context.Context)) *EntityRepository_Count_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *EntityRepository_Count_Call[T]) Return(_a0 int64, _a1 error) *EntityRepository_Count_Call[T] {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Create provides a mock function with given fields: ctx, entity
func (_m *EntityRepository[T]) Create(ctx context.Context, entity *T) (*T, error) {
	ret := _m.Called(ctx, entity)

	var r0 *T
	if rf, ok := ret.Get(0).(func(context.Context, *T) *T); ok {
		r0 = rf(ctx, entity)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*T)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *T) error); ok {
		r1 = rf(ctx, entity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EntityRepository_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type EntityRepository_Create_Call[T common.EntityConstraint] struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//  - ctx context.Context
//  - entity *T
func (_e *EntityRepository_Expecter[T]) Create(ctx interface{}, entity interface{}) *EntityRepository_Create_Call[T] {
	return &EntityRepository_Create_Call[T]{Call: _e.mock.On("Create", ctx, entity)}
}

func (_c *EntityRepository_Create_Call[T]) Run(run func(ctx context.Context, entity *T)) *EntityRepository_Create_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*T))
	})
	return _c
}

func (_c *EntityRepository_Create_Call[T]) Return(_a0 *T, _a1 error) *EntityRepository_Create_Call[T] {
	_c.Call.Return(_a0, _a1)
	return _c
}

// DeleteByID provides a mock function with given fields: ctx, id
func (_m *EntityRepository[T]) DeleteByID(ctx context.Context, id uint) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EntityRepository_DeleteByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteByID'
type EntityRepository_DeleteByID_Call[T common.EntityConstraint] struct {
	*mock.Call
}

// DeleteByID is a helper method to define mock.On call
//  - ctx context.Context
//  - id uint
func (_e *EntityRepository_Expecter[T]) DeleteByID(ctx interface{}, id interface{}) *EntityRepository_DeleteByID_Call[T] {
	return &EntityRepository_DeleteByID_Call[T]{Call: _e.mock.On("DeleteByID", ctx, id)}
}

func (_c *EntityRepository_DeleteByID_Call[T]) Run(run func(ctx context.Context, id uint)) *EntityRepository_DeleteByID_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint))
	})
	return _c
}

func (_c *EntityRepository_DeleteByID_Call[T]) Return(_a0 error) *EntityRepository_DeleteByID_Call[T] {
	_c.Call.Return(_a0)
	return _c
}

// FindAll provides a mock function with given fields: ctx, pagination
func (_m *EntityRepository[T]) FindAll(ctx context.Context, pagination *models.Pagination) (*models.Page[*T], error) {
	ret := _m.Called(ctx, pagination)

	var r0 *models.Page[*T]
	if rf, ok := ret.Get(0).(func(context.Context, *models.Pagination) *models.Page[*T]); ok {
		r0 = rf(ctx, pagination)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Page[*T])
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

// EntityRepository_FindAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindAll'
type EntityRepository_FindAll_Call[T common.EntityConstraint] struct {
	*mock.Call
}

// FindAll is a helper method to define mock.On call
//  - ctx context.Context
//  - pagination *models.Pagination
func (_e *EntityRepository_Expecter[T]) FindAll(ctx interface{}, pagination interface{}) *EntityRepository_FindAll_Call[T] {
	return &EntityRepository_FindAll_Call[T]{Call: _e.mock.On("FindAll", ctx, pagination)}
}

func (_c *EntityRepository_FindAll_Call[T]) Run(run func(ctx context.Context, pagination *models.Pagination)) *EntityRepository_FindAll_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*models.Pagination))
	})
	return _c
}

func (_c *EntityRepository_FindAll_Call[T]) Return(page *models.Page[*T], err error) *EntityRepository_FindAll_Call[T] {
	_c.Call.Return(page, err)
	return _c
}

// FindOneByID provides a mock function with given fields: ctx, id
func (_m *EntityRepository[T]) FindOneByID(ctx context.Context, id uint) (*T, error) {
	ret := _m.Called(ctx, id)

	var r0 *T
	if rf, ok := ret.Get(0).(func(context.Context, uint) *T); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*T)
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

// EntityRepository_FindOneByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindOneByID'
type EntityRepository_FindOneByID_Call[T common.EntityConstraint] struct {
	*mock.Call
}

// FindOneByID is a helper method to define mock.On call
//  - ctx context.Context
//  - id uint
func (_e *EntityRepository_Expecter[T]) FindOneByID(ctx interface{}, id interface{}) *EntityRepository_FindOneByID_Call[T] {
	return &EntityRepository_FindOneByID_Call[T]{Call: _e.mock.On("FindOneByID", ctx, id)}
}

func (_c *EntityRepository_FindOneByID_Call[T]) Run(run func(ctx context.Context, id uint)) *EntityRepository_FindOneByID_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint))
	})
	return _c
}

func (_c *EntityRepository_FindOneByID_Call[T]) Return(entity *T, err error) *EntityRepository_FindOneByID_Call[T] {
	_c.Call.Return(entity, err)
	return _c
}

// FindOneByIDForUpdate provides a mock function with given fields: ctx, id
func (_m *EntityRepository[T]) FindOneByIDForUpdate(ctx context.Context, id uint) (*T, error) {
	ret := _m.Called(ctx, id)

	var r0 *T
	if rf, ok := ret.Get(0).(func(context.Context, uint) *T); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*T)
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

// EntityRepository_FindOneByIDForUpdate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindOneByIDForUpdate'
type EntityRepository_FindOneByIDForUpdate_Call[T common.EntityConstraint] struct {
	*mock.Call
}

// FindOneByIDForUpdate is a helper method to define mock.On call
//  - ctx context.Context
//  - id uint
func (_e *EntityRepository_Expecter[T]) FindOneByIDForUpdate(ctx interface{}, id interface{}) *EntityRepository_FindOneByIDForUpdate_Call[T] {
	return &EntityRepository_FindOneByIDForUpdate_Call[T]{Call: _e.mock.On("FindOneByIDForUpdate", ctx, id)}
}

func (_c *EntityRepository_FindOneByIDForUpdate_Call[T]) Run(run func(ctx context.Context, id uint)) *EntityRepository_FindOneByIDForUpdate_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint))
	})
	return _c
}

func (_c *EntityRepository_FindOneByIDForUpdate_Call[T]) Return(entity *T, err error) *EntityRepository_FindOneByIDForUpdate_Call[T] {
	_c.Call.Return(entity, err)
	return _c
}

// Gorm provides a mock function with given fields:
func (_m *EntityRepository[T]) Gorm() *gorm.DB {
	ret := _m.Called()

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func() *gorm.DB); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// EntityRepository_Gorm_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Gorm'
type EntityRepository_Gorm_Call[T common.EntityConstraint] struct {
	*mock.Call
}

// Gorm is a helper method to define mock.On call
func (_e *EntityRepository_Expecter[T]) Gorm() *EntityRepository_Gorm_Call[T] {
	return &EntityRepository_Gorm_Call[T]{Call: _e.mock.On("Gorm")}
}

func (_c *EntityRepository_Gorm_Call[T]) Run(run func()) *EntityRepository_Gorm_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *EntityRepository_Gorm_Call[T]) Return(_a0 *gorm.DB) *EntityRepository_Gorm_Call[T] {
	_c.Call.Return(_a0)
	return _c
}

// Save provides a mock function with given fields: ctx, entity
func (_m *EntityRepository[T]) Save(ctx context.Context, entity *T) (*T, error) {
	ret := _m.Called(ctx, entity)

	var r0 *T
	if rf, ok := ret.Get(0).(func(context.Context, *T) *T); ok {
		r0 = rf(ctx, entity)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*T)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *T) error); ok {
		r1 = rf(ctx, entity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EntityRepository_Save_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Save'
type EntityRepository_Save_Call[T common.EntityConstraint] struct {
	*mock.Call
}

// Save is a helper method to define mock.On call
//  - ctx context.Context
//  - entity *T
func (_e *EntityRepository_Expecter[T]) Save(ctx interface{}, entity interface{}) *EntityRepository_Save_Call[T] {
	return &EntityRepository_Save_Call[T]{Call: _e.mock.On("Save", ctx, entity)}
}

func (_c *EntityRepository_Save_Call[T]) Run(run func(ctx context.Context, entity *T)) *EntityRepository_Save_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*T))
	})
	return _c
}

func (_c *EntityRepository_Save_Call[T]) Return(_a0 *T, _a1 error) *EntityRepository_Save_Call[T] {
	_c.Call.Return(_a0, _a1)
	return _c
}

// TransactWithDefaultRetry provides a mock function with given fields: operation
func (_m *EntityRepository[T]) TransactWithDefaultRetry(operation func(*gorm.DB) error) error {
	ret := _m.Called(operation)

	var r0 error
	if rf, ok := ret.Get(0).(func(func(*gorm.DB) error) error); ok {
		r0 = rf(operation)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EntityRepository_TransactWithDefaultRetry_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TransactWithDefaultRetry'
type EntityRepository_TransactWithDefaultRetry_Call[T common.EntityConstraint] struct {
	*mock.Call
}

// TransactWithDefaultRetry is a helper method to define mock.On call
//  - operation func(*gorm.DB) error
func (_e *EntityRepository_Expecter[T]) TransactWithDefaultRetry(operation interface{}) *EntityRepository_TransactWithDefaultRetry_Call[T] {
	return &EntityRepository_TransactWithDefaultRetry_Call[T]{Call: _e.mock.On("TransactWithDefaultRetry", operation)}
}

func (_c *EntityRepository_TransactWithDefaultRetry_Call[T]) Run(run func(operation func(*gorm.DB) error)) *EntityRepository_TransactWithDefaultRetry_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(func(*gorm.DB) error))
	})
	return _c
}

func (_c *EntityRepository_TransactWithDefaultRetry_Call[T]) Return(_a0 error) *EntityRepository_TransactWithDefaultRetry_Call[T] {
	_c.Call.Return(_a0)
	return _c
}

// WithTx provides a mock function with given fields: tx
func (_m *EntityRepository[T]) WithTx(tx *gorm.DB) persistence.EntityRepository[T] {
	ret := _m.Called(tx)

	var r0 persistence.EntityRepository[T]
	if rf, ok := ret.Get(0).(func(*gorm.DB) persistence.EntityRepository[T]); ok {
		r0 = rf(tx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(persistence.EntityRepository[T])
		}
	}

	return r0
}

// EntityRepository_WithTx_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithTx'
type EntityRepository_WithTx_Call[T common.EntityConstraint] struct {
	*mock.Call
}

// WithTx is a helper method to define mock.On call
//  - tx *gorm.DB
func (_e *EntityRepository_Expecter[T]) WithTx(tx interface{}) *EntityRepository_WithTx_Call[T] {
	return &EntityRepository_WithTx_Call[T]{Call: _e.mock.On("WithTx", tx)}
}

func (_c *EntityRepository_WithTx_Call[T]) Run(run func(tx *gorm.DB)) *EntityRepository_WithTx_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gorm.DB))
	})
	return _c
}

func (_c *EntityRepository_WithTx_Call[T]) Return(_a0 persistence.EntityRepository[T]) *EntityRepository_WithTx_Call[T] {
	_c.Call.Return(_a0)
	return _c
}

// WrapError provides a mock function with given fields: err
func (_m *EntityRepository[T]) WrapError(err error) error {
	ret := _m.Called(err)

	var r0 error
	if rf, ok := ret.Get(0).(func(error) error); ok {
		r0 = rf(err)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EntityRepository_WrapError_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WrapError'
type EntityRepository_WrapError_Call[T common.EntityConstraint] struct {
	*mock.Call
}

// WrapError is a helper method to define mock.On call
//  - err error
func (_e *EntityRepository_Expecter[T]) WrapError(err interface{}) *EntityRepository_WrapError_Call[T] {
	return &EntityRepository_WrapError_Call[T]{Call: _e.mock.On("WrapError", err)}
}

func (_c *EntityRepository_WrapError_Call[T]) Run(run func(err error)) *EntityRepository_WrapError_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(error))
	})
	return _c
}

func (_c *EntityRepository_WrapError_Call[T]) Return(_a0 error) *EntityRepository_WrapError_Call[T] {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewEntityRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewEntityRepository creates a new instance of EntityRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewEntityRepository[T common.EntityConstraint](t mockConstructorTestingTNewEntityRepository) *EntityRepository[T] {
	mock := &EntityRepository[T]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
