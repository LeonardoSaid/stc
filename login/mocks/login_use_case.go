// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/leonardosaid/stc/accounts/internal/domain"
	mock "github.com/stretchr/testify/mock"
)

// LoginUseCase is an autogenerated mock type for the LoginUseCase type
type LoginUseCase struct {
	mock.Mock
}

// Login provides a mock function with given fields: _a0, _a1
func (_m *LoginUseCase) Login(_a0 context.Context, _a1 *domain.LoginCredentials) (string, error) {
	ret := _m.Called(_a0, _a1)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, *domain.LoginCredentials) string); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.LoginCredentials) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewLoginUseCase interface {
	mock.TestingT
	Cleanup(func())
}

// NewLoginUseCase creates a new instance of LoginUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewLoginUseCase(t mockConstructorTestingTNewLoginUseCase) *LoginUseCase {
	mock := &LoginUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
