// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/leonardosaid/stc/accounts/internal/domain"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// TransferRepository is an autogenerated mock type for the TransferRepository type
type TransferRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0, _a1
func (_m *TransferRepository) Create(_a0 context.Context, _a1 *domain.Transfer) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Transfer) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListByAccountID provides a mock function with given fields: _a0, _a1
func (_m *TransferRepository) ListByAccountID(_a0 context.Context, _a1 uuid.UUID) ([]domain.Transfer, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []domain.Transfer
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) []domain.Transfer); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Transfer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewTransferRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewTransferRepository creates a new instance of TransferRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTransferRepository(t mockConstructorTestingTNewTransferRepository) *TransferRepository {
	mock := &TransferRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}