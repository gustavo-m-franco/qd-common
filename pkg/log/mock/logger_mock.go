// Code generated by MockGen. DO NOT EDIT.
// Source: ./logger.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockLoggerer is a mock of Loggerer interface.
type MockLoggerer struct {
	ctrl     *gomock.Controller
	recorder *MockLoggererMockRecorder
}

// MockLoggererMockRecorder is the mock recorder for MockLoggerer.
type MockLoggererMockRecorder struct {
	mock *MockLoggerer
}

// NewMockLoggerer creates a new mock instance.
func NewMockLoggerer(ctrl *gomock.Controller) *MockLoggerer {
	mock := &MockLoggerer{ctrl: ctrl}
	mock.recorder = &MockLoggererMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLoggerer) EXPECT() *MockLoggererMockRecorder {
	return m.recorder
}

// Error mocks base method.
func (m *MockLoggerer) Error(err error, message string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Error", err, message)
}

// Error indicates an expected call of Error.
func (mr *MockLoggererMockRecorder) Error(err, message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockLoggerer)(nil).Error), err, message)
}

// Info mocks base method.
func (m *MockLoggerer) Info(message string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Info", message)
}

// Info indicates an expected call of Info.
func (mr *MockLoggererMockRecorder) Info(message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockLoggerer)(nil).Info), message)
}

// Warn mocks base method.
func (m *MockLoggerer) Warn(message string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Warn", message)
}

// Warn indicates an expected call of Warn.
func (mr *MockLoggererMockRecorder) Warn(message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warn", reflect.TypeOf((*MockLoggerer)(nil).Warn), message)
}
