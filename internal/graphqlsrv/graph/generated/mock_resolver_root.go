// Code generated by mockery v2.12.0. DO NOT EDIT.

package generated

import (
	testing "testing"

	mock "github.com/stretchr/testify/mock"
)

// MockResolverRoot is an autogenerated mock type for the ResolverRoot type
type MockResolverRoot struct {
	mock.Mock
}

// Query provides a mock function with given fields:
func (_m *MockResolverRoot) Query() QueryResolver {
	ret := _m.Called()

	var r0 QueryResolver
	if rf, ok := ret.Get(0).(func() QueryResolver); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(QueryResolver)
		}
	}

	return r0
}

// NewMockResolverRoot creates a new instance of MockResolverRoot. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockResolverRoot(t testing.TB) *MockResolverRoot {
	mock := &MockResolverRoot{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
