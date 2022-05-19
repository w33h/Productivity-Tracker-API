// Code generated by mockery v2.12.2. DO NOT EDIT.

package mocks

import (
	//testing "testing"

	mock "github.com/stretchr/testify/mock"

	todos "github.com/w33h/Productivity-Tracker-API/business/todos"
)

// RepositoryTodos is an autogenerated mock type for the RepositoryTodos type
type RepositoryTodos struct {
	mock.Mock
}

// DeleteTodo provides a mock function with given fields: id
func (_m *RepositoryTodos) DeleteTodo(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllTodo provides a mock function with given fields:
func (_m *RepositoryTodos) FindAllTodo() ([]todos.Todo, error) {
	ret := _m.Called()

	var r0 []todos.Todo
	if rf, ok := ret.Get(0).(func() []todos.Todo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]todos.Todo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: id
func (_m *RepositoryTodos) FindById(id string) (*todos.Todo, error) {
	ret := _m.Called(id)

	var r0 *todos.Todo
	if rf, ok := ret.Get(0).(func(string) *todos.Todo); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*todos.Todo)
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

// GetByStatus provides a mock function with given fields: status
func (_m *RepositoryTodos) FindByStatus(status string) ([]todos.Todo, error) {
	ret := _m.Called(status)

	var r0 []todos.Todo
	if rf, ok := ret.Get(0).(func(string) []todos.Todo); ok {
		r0 = rf(status)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]todos.Todo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(status)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertTodo provides a mock function with given fields: todo
func (_m *RepositoryTodos) InsertTodo(todo todos.Todo) (id string, err error) {
	ret := _m.Called(todo)

	var r0 string
	if rf, ok := ret.Get(0).(func(todos.Todo) string); ok {
		r0 = rf(todo)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(todos.Todo) error); ok {
		r1 = rf(todo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTodo provides a mock function with given fields: todo
func (_m *RepositoryTodos) UpdateTodo(todo todos.Todo) error {
	ret := _m.Called(todo)

	var r0 error
	if rf, ok := ret.Get(0).(func(todos.Todo) error); ok {
		r0 = rf(todo)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewRepositoryTodos creates a new instance of RepositoryTodos. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
//func NewRepositoryTodos(t testing.TB) *RepositoryTodos {
//	mock := &RepositoryTodos{}
//	mock.Mock.Test(t)
//
//	t.Cleanup(func() { mock.AssertExpectations(t) })
//
//	return mock
//}