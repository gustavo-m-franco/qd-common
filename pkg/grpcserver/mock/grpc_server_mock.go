// Code generated by MockGen. DO NOT EDIT.
// Source: ./grpc_server.go

// Package mock is a generated GoMock package.
package mock

import (
	net "net"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockGRPCServerer is a mock of GRPCServerer interface.
type MockGRPCServerer struct {
	ctrl     *gomock.Controller
	recorder *MockGRPCServererMockRecorder
}

// MockGRPCServererMockRecorder is the mock recorder for MockGRPCServerer.
type MockGRPCServererMockRecorder struct {
	mock *MockGRPCServerer
}

// NewMockGRPCServerer creates a new mock instance.
func NewMockGRPCServerer(ctrl *gomock.Controller) *MockGRPCServerer {
	mock := &MockGRPCServerer{ctrl: ctrl}
	mock.recorder = &MockGRPCServererMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGRPCServerer) EXPECT() *MockGRPCServererMockRecorder {
	return m.recorder
}

// Serve mocks base method.
func (m *MockGRPCServerer) Serve(arg0 net.Listener) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Serve", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Serve indicates an expected call of Serve.
func (mr *MockGRPCServererMockRecorder) Serve(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Serve", reflect.TypeOf((*MockGRPCServerer)(nil).Serve), arg0)
}

// Stop mocks base method.
func (m *MockGRPCServerer) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop.
func (mr *MockGRPCServererMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockGRPCServerer)(nil).Stop))
}

// MockGRPCServicer is a mock of GRPCServicer interface.
type MockGRPCServicer struct {
	ctrl     *gomock.Controller
	recorder *MockGRPCServicerMockRecorder
}

// MockGRPCServicerMockRecorder is the mock recorder for MockGRPCServicer.
type MockGRPCServicerMockRecorder struct {
	mock *MockGRPCServicer
}

// NewMockGRPCServicer creates a new mock instance.
func NewMockGRPCServicer(ctrl *gomock.Controller) *MockGRPCServicer {
	mock := &MockGRPCServicer{ctrl: ctrl}
	mock.recorder = &MockGRPCServicerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGRPCServicer) EXPECT() *MockGRPCServicerMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockGRPCServicer) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockGRPCServicerMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockGRPCServicer)(nil).Close))
}

// Serve mocks base method.
func (m *MockGRPCServicer) Serve() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Serve")
	ret0, _ := ret[0].(error)
	return ret0
}

// Serve indicates an expected call of Serve.
func (mr *MockGRPCServicerMockRecorder) Serve() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Serve", reflect.TypeOf((*MockGRPCServicer)(nil).Serve))
}