// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks_test

import mock "github.com/stretchr/testify/mock"

// MockRunner is an autogenerated mock type for the Runner type
type MockRunner struct {
	mock.Mock
}

type MockRunner_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRunner) EXPECT() *MockRunner_Expecter {
	return &MockRunner_Expecter{mock: &_m.Mock}
}

// Run provides a mock function with given fields: name, args
func (_m *MockRunner) Run(name string, args ...string) error {
	_va := make([]interface{}, len(args))
	for _i := range args {
		_va[_i] = args[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, name)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Run")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, ...string) error); ok {
		r0 = rf(name, args...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRunner_Run_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Run'
type MockRunner_Run_Call struct {
	*mock.Call
}

// Run is a helper method to define mock.On call
//   - name string
//   - args ...string
func (_e *MockRunner_Expecter) Run(name interface{}, args ...interface{}) *MockRunner_Run_Call {
	return &MockRunner_Run_Call{Call: _e.mock.On("Run",
		append([]interface{}{name}, args...)...)}
}

func (_c *MockRunner_Run_Call) Run(run func(name string, args ...string)) *MockRunner_Run_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]string, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(string)
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockRunner_Run_Call) Return(_a0 error) *MockRunner_Run_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRunner_Run_Call) RunAndReturn(run func(string, ...string) error) *MockRunner_Run_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockRunner creates a new instance of MockRunner. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRunner(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRunner {
	mock := &MockRunner{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}