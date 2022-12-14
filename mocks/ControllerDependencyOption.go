// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	controller "github.com/rakhmadbudiono/code-scanner/internal/controller"
	mock "github.com/stretchr/testify/mock"
)

// ControllerDependencyOption is an autogenerated mock type for the ControllerDependencyOption type
type ControllerDependencyOption struct {
	mock.Mock
}

// Execute provides a mock function with given fields: _a0
func (_m *ControllerDependencyOption) Execute(_a0 *controller.Controller) {
	_m.Called(_a0)
}

type mockConstructorTestingTNewControllerDependencyOption interface {
	mock.TestingT
	Cleanup(func())
}

// NewControllerDependencyOption creates a new instance of ControllerDependencyOption. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewControllerDependencyOption(t mockConstructorTestingTNewControllerDependencyOption) *ControllerDependencyOption {
	mock := &ControllerDependencyOption{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
