// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/rimvydascivilis/book-tracker/backend/domain"
	mock "github.com/stretchr/testify/mock"
)

// BookService is an autogenerated mock type for the BookService type
type BookService struct {
	mock.Mock
}

type BookService_Expecter struct {
	mock *mock.Mock
}

func (_m *BookService) EXPECT() *BookService_Expecter {
	return &BookService_Expecter{mock: &_m.Mock}
}

// CreateBook provides a mock function with given fields: ctx, userID, book
func (_m *BookService) CreateBook(ctx context.Context, userID int64, book domain.Book) (domain.Book, error) {
	ret := _m.Called(ctx, userID, book)

	if len(ret) == 0 {
		panic("no return value specified for CreateBook")
	}

	var r0 domain.Book
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, domain.Book) (domain.Book, error)); ok {
		return rf(ctx, userID, book)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, domain.Book) domain.Book); ok {
		r0 = rf(ctx, userID, book)
	} else {
		r0 = ret.Get(0).(domain.Book)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, domain.Book) error); ok {
		r1 = rf(ctx, userID, book)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BookService_CreateBook_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateBook'
type BookService_CreateBook_Call struct {
	*mock.Call
}

// CreateBook is a helper method to define mock.On call
//   - ctx context.Context
//   - userID int64
//   - book domain.Book
func (_e *BookService_Expecter) CreateBook(ctx interface{}, userID interface{}, book interface{}) *BookService_CreateBook_Call {
	return &BookService_CreateBook_Call{Call: _e.mock.On("CreateBook", ctx, userID, book)}
}

func (_c *BookService_CreateBook_Call) Run(run func(ctx context.Context, userID int64, book domain.Book)) *BookService_CreateBook_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].(domain.Book))
	})
	return _c
}

func (_c *BookService_CreateBook_Call) Return(_a0 domain.Book, _a1 error) *BookService_CreateBook_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *BookService_CreateBook_Call) RunAndReturn(run func(context.Context, int64, domain.Book) (domain.Book, error)) *BookService_CreateBook_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteBook provides a mock function with given fields: ctx, userID, bookID
func (_m *BookService) DeleteBook(ctx context.Context, userID int64, bookID int64) error {
	ret := _m.Called(ctx, userID, bookID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteBook")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64) error); ok {
		r0 = rf(ctx, userID, bookID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BookService_DeleteBook_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteBook'
type BookService_DeleteBook_Call struct {
	*mock.Call
}

// DeleteBook is a helper method to define mock.On call
//   - ctx context.Context
//   - userID int64
//   - bookID int64
func (_e *BookService_Expecter) DeleteBook(ctx interface{}, userID interface{}, bookID interface{}) *BookService_DeleteBook_Call {
	return &BookService_DeleteBook_Call{Call: _e.mock.On("DeleteBook", ctx, userID, bookID)}
}

func (_c *BookService_DeleteBook_Call) Run(run func(ctx context.Context, userID int64, bookID int64)) *BookService_DeleteBook_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].(int64))
	})
	return _c
}

func (_c *BookService_DeleteBook_Call) Return(_a0 error) *BookService_DeleteBook_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BookService_DeleteBook_Call) RunAndReturn(run func(context.Context, int64, int64) error) *BookService_DeleteBook_Call {
	_c.Call.Return(run)
	return _c
}

// GetBooks provides a mock function with given fields: ctx, userID, page, limit
func (_m *BookService) GetBooks(ctx context.Context, userID int64, page int64, limit int64) ([]domain.Book, bool, error) {
	ret := _m.Called(ctx, userID, page, limit)

	if len(ret) == 0 {
		panic("no return value specified for GetBooks")
	}

	var r0 []domain.Book
	var r1 bool
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64, int64) ([]domain.Book, bool, error)); ok {
		return rf(ctx, userID, page, limit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64, int64) []domain.Book); ok {
		r0 = rf(ctx, userID, page, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Book)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, int64, int64) bool); ok {
		r1 = rf(ctx, userID, page, limit)
	} else {
		r1 = ret.Get(1).(bool)
	}

	if rf, ok := ret.Get(2).(func(context.Context, int64, int64, int64) error); ok {
		r2 = rf(ctx, userID, page, limit)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// BookService_GetBooks_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBooks'
type BookService_GetBooks_Call struct {
	*mock.Call
}

// GetBooks is a helper method to define mock.On call
//   - ctx context.Context
//   - userID int64
//   - page int64
//   - limit int64
func (_e *BookService_Expecter) GetBooks(ctx interface{}, userID interface{}, page interface{}, limit interface{}) *BookService_GetBooks_Call {
	return &BookService_GetBooks_Call{Call: _e.mock.On("GetBooks", ctx, userID, page, limit)}
}

func (_c *BookService_GetBooks_Call) Run(run func(ctx context.Context, userID int64, page int64, limit int64)) *BookService_GetBooks_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].(int64), args[3].(int64))
	})
	return _c
}

func (_c *BookService_GetBooks_Call) Return(_a0 []domain.Book, _a1 bool, _a2 error) *BookService_GetBooks_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *BookService_GetBooks_Call) RunAndReturn(run func(context.Context, int64, int64, int64) ([]domain.Book, bool, error)) *BookService_GetBooks_Call {
	_c.Call.Return(run)
	return _c
}

// SearchBooks provides a mock function with given fields: ctx, userID, title, limit
func (_m *BookService) SearchBooks(ctx context.Context, userID int64, title string, limit int64) ([]domain.Book, error) {
	ret := _m.Called(ctx, userID, title, limit)

	if len(ret) == 0 {
		panic("no return value specified for SearchBooks")
	}

	var r0 []domain.Book
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, string, int64) ([]domain.Book, error)); ok {
		return rf(ctx, userID, title, limit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, string, int64) []domain.Book); ok {
		r0 = rf(ctx, userID, title, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Book)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, string, int64) error); ok {
		r1 = rf(ctx, userID, title, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BookService_SearchBooks_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SearchBooks'
type BookService_SearchBooks_Call struct {
	*mock.Call
}

// SearchBooks is a helper method to define mock.On call
//   - ctx context.Context
//   - userID int64
//   - title string
//   - limit int64
func (_e *BookService_Expecter) SearchBooks(ctx interface{}, userID interface{}, title interface{}, limit interface{}) *BookService_SearchBooks_Call {
	return &BookService_SearchBooks_Call{Call: _e.mock.On("SearchBooks", ctx, userID, title, limit)}
}

func (_c *BookService_SearchBooks_Call) Run(run func(ctx context.Context, userID int64, title string, limit int64)) *BookService_SearchBooks_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].(string), args[3].(int64))
	})
	return _c
}

func (_c *BookService_SearchBooks_Call) Return(_a0 []domain.Book, _a1 error) *BookService_SearchBooks_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *BookService_SearchBooks_Call) RunAndReturn(run func(context.Context, int64, string, int64) ([]domain.Book, error)) *BookService_SearchBooks_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateBook provides a mock function with given fields: ctx, userID, book
func (_m *BookService) UpdateBook(ctx context.Context, userID int64, book domain.Book) (domain.Book, error) {
	ret := _m.Called(ctx, userID, book)

	if len(ret) == 0 {
		panic("no return value specified for UpdateBook")
	}

	var r0 domain.Book
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, domain.Book) (domain.Book, error)); ok {
		return rf(ctx, userID, book)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, domain.Book) domain.Book); ok {
		r0 = rf(ctx, userID, book)
	} else {
		r0 = ret.Get(0).(domain.Book)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, domain.Book) error); ok {
		r1 = rf(ctx, userID, book)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BookService_UpdateBook_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateBook'
type BookService_UpdateBook_Call struct {
	*mock.Call
}

// UpdateBook is a helper method to define mock.On call
//   - ctx context.Context
//   - userID int64
//   - book domain.Book
func (_e *BookService_Expecter) UpdateBook(ctx interface{}, userID interface{}, book interface{}) *BookService_UpdateBook_Call {
	return &BookService_UpdateBook_Call{Call: _e.mock.On("UpdateBook", ctx, userID, book)}
}

func (_c *BookService_UpdateBook_Call) Run(run func(ctx context.Context, userID int64, book domain.Book)) *BookService_UpdateBook_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].(domain.Book))
	})
	return _c
}

func (_c *BookService_UpdateBook_Call) Return(_a0 domain.Book, _a1 error) *BookService_UpdateBook_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *BookService_UpdateBook_Call) RunAndReturn(run func(context.Context, int64, domain.Book) (domain.Book, error)) *BookService_UpdateBook_Call {
	_c.Call.Return(run)
	return _c
}

// NewBookService creates a new instance of BookService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBookService(t interface {
	mock.TestingT
	Cleanup(func())
}) *BookService {
	mock := &BookService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
