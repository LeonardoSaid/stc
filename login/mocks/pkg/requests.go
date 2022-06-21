// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Requests is an autogenerated mock type for the Requests type
type Requests struct {
	mock.Mock
}

// Request provides a mock function with given fields: requestData, uri, method, expectedStatus
func (_m *Requests) Request(requestData interface{}, uri string, method string, expectedStatus int) ([]byte, error) {
	ret := _m.Called(requestData, uri, method, expectedStatus)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(interface{}, string, string, int) []byte); ok {
		r0 = rf(requestData, uri, method, expectedStatus)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}, string, string, int) error); ok {
		r1 = rf(requestData, uri, method, expectedStatus)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewRequests interface {
	mock.TestingT
	Cleanup(func())
}

// NewRequests creates a new instance of Requests. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRequests(t mockConstructorTestingTNewRequests) *Requests {
	mock := &Requests{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}