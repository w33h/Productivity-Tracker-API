// Code generated by mockery v2.12.2. DO NOT EDIT.

package mocks

import (
	//testing "testing"

	mock "github.com/stretchr/testify/mock"

	user "github.com/w33h/Productivity-Tracker-API/business/users"
)

// RepositoryUser is an autogenerated mock type for the RepositoryUser type
type RepositoryUser struct {
	mock.Mock
}

// DeleteUser provides a mock function with given fields: id
func (_m *RepositoryUser) DeleteUser(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindById provides a mock function with given fields: id
func (_m *RepositoryUser) FindById(id string) (*user.Users, error) {
	ret := _m.Called(id)

	var r0 *user.Users
	if rf, ok := ret.Get(0).(func(string) *user.Users); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.Users)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertUser provides a mock function with given fields: _a0
func (_m *RepositoryUser) InsertUser(_a0 user.Users) (user.Users, error) {
	ret := _m.Called(_a0)

	var r0 user.Users
	if rf, ok := ret.Get(0).(func(user.Users) user.Users); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(user.Users)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(user.Users) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: _a0
func (_m *RepositoryUser) UpdateUser(_a0 *user.Users) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*user.Users) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewRepositoryUser creates a new instance of RepositoryUser. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
//func NewRepositoryUser(t testing.TB) *RepositoryUser {
//	mock := &RepositoryUser{}
//	mock.Mock.Test(t)
//
//	t.Cleanup(func() { mock.AssertExpectations(t) })
//
//	return mock
//}
