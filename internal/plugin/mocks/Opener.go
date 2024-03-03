// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks_test

import (
	plugin "github.com/renatoaraujo/modular/internal/plugin"
	mock "github.com/stretchr/testify/mock"
)

// MockOpener is an autogenerated mock type for the Opener type
type MockOpener struct {
	mock.Mock
}

type MockOpener_Expecter struct {
	mock *mock.Mock
}

func (_m *MockOpener) EXPECT() *MockOpener_Expecter {
	return &MockOpener_Expecter{mock: &_m.Mock}
}

// Open provides a mock function with given fields: path
func (_m *MockOpener) Open(path string) (plugin.SymbolLoader, error) {
	ret := _m.Called(path)

	if len(ret) == 0 {
		panic("no return value specified for Open")
	}

	var r0 plugin.SymbolLoader
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (plugin.SymbolLoader, error)); ok {
		return rf(path)
	}
	if rf, ok := ret.Get(0).(func(string) plugin.SymbolLoader); ok {
		r0 = rf(path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(plugin.SymbolLoader)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockOpener_Open_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Open'
type MockOpener_Open_Call struct {
	*mock.Call
}

// Open is a helper method to define mock.On call
//   - path string
func (_e *MockOpener_Expecter) Open(path interface{}) *MockOpener_Open_Call {
	return &MockOpener_Open_Call{Call: _e.mock.On("Open", path)}
}

func (_c *MockOpener_Open_Call) Run(run func(path string)) *MockOpener_Open_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockOpener_Open_Call) Return(_a0 plugin.SymbolLoader, _a1 error) *MockOpener_Open_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockOpener_Open_Call) RunAndReturn(run func(string) (plugin.SymbolLoader, error)) *MockOpener_Open_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockOpener creates a new instance of MockOpener. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockOpener(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockOpener {
	mock := &MockOpener{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
