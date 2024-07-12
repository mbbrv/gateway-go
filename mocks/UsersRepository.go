// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	models "gateway-go/internal/models"

	mock "github.com/stretchr/testify/mock"
)

// UsersRepository is an autogenerated mock type for the UsersRepository type
type UsersRepository struct {
	mock.Mock
}

// Login provides a mock function with given fields: username, password
func (_m *UsersRepository) Login(username string, password string) (models.User, error) {
	ret := _m.Called(username, password)

	if len(ret) == 0 {
		panic("no return value specified for Login")
	}

	var r0 models.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (models.User, error)); ok {
		return rf(username, password)
	}
	if rf, ok := ret.Get(0).(func(string, string) models.User); ok {
		r0 = rf(username, password)
	} else {
		r0 = ret.Get(0).(models.User)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(username, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SignUp provides a mock function with given fields: username, password
func (_m *UsersRepository) SignUp(username string, password string) error {
	ret := _m.Called(username, password)

	if len(ret) == 0 {
		panic("no return value specified for SignUp")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(username, password)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewUsersRepository creates a new instance of UsersRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUsersRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *UsersRepository {
	mock := &UsersRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
