// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks_test

import (
	plugin "github.com/renatoaraujo/modular/internal/plugin"
	mock "github.com/stretchr/testify/mock"
)

// MockLoader is an autogenerated mock type for the Loader type
type MockLoader struct {
	mock.Mock
}

type MockLoader_Expecter struct {
	mock *mock.Mock
}

func (_m *MockLoader) EXPECT() *MockLoader_Expecter {
	return &MockLoader_Expecter{mock: &_m.Mock}
}

// Load provides a mock function with given fields: path
func (_m *MockLoader) Load(path string) (*plugin.Installation, error) {
	ret := _m.Called(path)

	if len(ret) == 0 {
		panic("no return value specified for Load")
	}

	var r0 *plugin.Installation
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*plugin.Installation, error)); ok {
		return rf(path)
	}
	if rf, ok := ret.Get(0).(func(string) *plugin.Installation); ok {
		r0 = rf(path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*plugin.Installation)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockLoader_Load_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Load'
type MockLoader_Load_Call struct {
	*mock.Call
}

// Load is a helper method to define mock.On call
//   - path string
func (_e *MockLoader_Expecter) Load(path interface{}) *MockLoader_Load_Call {
	return &MockLoader_Load_Call{Call: _e.mock.On("Load", path)}
}

func (_c *MockLoader_Load_Call) Run(run func(path string)) *MockLoader_Load_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockLoader_Load_Call) Return(_a0 *plugin.Installation, _a1 error) *MockLoader_Load_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockLoader_Load_Call) RunAndReturn(run func(string) (*plugin.Installation, error)) *MockLoader_Load_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockLoader creates a new instance of MockLoader. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockLoader(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockLoader {
	mock := &MockLoader{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
