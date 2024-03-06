// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks_test

import mock "github.com/stretchr/testify/mock"

// MockValidator is an autogenerated mock type for the Validator type
type MockValidator struct {
	mock.Mock
}

type MockValidator_Expecter struct {
	mock *mock.Mock
}

func (_m *MockValidator) EXPECT() *MockValidator_Expecter {
	return &MockValidator_Expecter{mock: &_m.Mock}
}

// ValidateAndFormat provides a mock function with given fields: repo
func (_m *MockValidator) ValidateAndFormat(repo string) (string, error) {
	ret := _m.Called(repo)

	if len(ret) == 0 {
		panic("no return value specified for ValidateAndFormat")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(repo)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(repo)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(repo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockValidator_ValidateAndFormat_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ValidateAndFormat'
type MockValidator_ValidateAndFormat_Call struct {
	*mock.Call
}

// ValidateAndFormat is a helper method to define mock.On call
//   - repo string
func (_e *MockValidator_Expecter) ValidateAndFormat(repo interface{}) *MockValidator_ValidateAndFormat_Call {
	return &MockValidator_ValidateAndFormat_Call{Call: _e.mock.On("ValidateAndFormat", repo)}
}

func (_c *MockValidator_ValidateAndFormat_Call) Run(run func(repo string)) *MockValidator_ValidateAndFormat_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockValidator_ValidateAndFormat_Call) Return(_a0 string, _a1 error) *MockValidator_ValidateAndFormat_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockValidator_ValidateAndFormat_Call) RunAndReturn(run func(string) (string, error)) *MockValidator_ValidateAndFormat_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockValidator creates a new instance of MockValidator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockValidator(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockValidator {
	mock := &MockValidator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
