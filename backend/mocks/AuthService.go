// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// AuthService is an autogenerated mock type for the AuthService type
type AuthService struct {
	mock.Mock
}

type AuthService_Expecter struct {
	mock *mock.Mock
}

func (_m *AuthService) EXPECT() *AuthService_Expecter {
	return &AuthService_Expecter{mock: &_m.Mock}
}

// Login provides a mock function with given fields: ctx, googleOauthToken
func (_m *AuthService) Login(ctx context.Context, googleOauthToken string) (string, error) {
	ret := _m.Called(ctx, googleOauthToken)

	if len(ret) == 0 {
		panic("no return value specified for Login")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (string, error)); ok {
		return rf(ctx, googleOauthToken)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, googleOauthToken)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, googleOauthToken)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AuthService_Login_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Login'
type AuthService_Login_Call struct {
	*mock.Call
}

// Login is a helper method to define mock.On call
//   - ctx context.Context
//   - googleOauthToken string
func (_e *AuthService_Expecter) Login(ctx interface{}, googleOauthToken interface{}) *AuthService_Login_Call {
	return &AuthService_Login_Call{Call: _e.mock.On("Login", ctx, googleOauthToken)}
}

func (_c *AuthService_Login_Call) Run(run func(ctx context.Context, googleOauthToken string)) *AuthService_Login_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *AuthService_Login_Call) Return(_a0 string, _a1 error) *AuthService_Login_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AuthService_Login_Call) RunAndReturn(run func(context.Context, string) (string, error)) *AuthService_Login_Call {
	_c.Call.Return(run)
	return _c
}

// NewAuthService creates a new instance of AuthService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAuthService(t interface {
	mock.TestingT
	Cleanup(func())
}) *AuthService {
	mock := &AuthService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
