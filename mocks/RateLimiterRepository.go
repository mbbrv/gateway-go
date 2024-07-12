// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// RateLimiterRepository is an autogenerated mock type for the RateLimiterRepository type
type RateLimiterRepository struct {
	mock.Mock
}

// Get provides a mock function with given fields: identifier
func (_m *RateLimiterRepository) Get(identifier string) (int, error) {
	ret := _m.Called(identifier)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (int, error)); ok {
		return rf(identifier)
	}
	if rf, ok := ret.Get(0).(func(string) int); ok {
		r0 = rf(identifier)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(identifier)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Incr provides a mock function with given fields: identifier
func (_m *RateLimiterRepository) Incr(identifier string) error {
	ret := _m.Called(identifier)

	if len(ret) == 0 {
		panic("no return value specified for Incr")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(identifier)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewRateLimiterRepository creates a new instance of RateLimiterRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRateLimiterRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *RateLimiterRepository {
	mock := &RateLimiterRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}