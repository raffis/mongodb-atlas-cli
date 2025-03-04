// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/internal/cli/atlas/quickstart (interfaces: Flow)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockFlow is a mock of Flow interface.
type MockFlow struct {
	ctrl     *gomock.Controller
	recorder *MockFlowMockRecorder
}

// MockFlowMockRecorder is the mock recorder for MockFlow.
type MockFlowMockRecorder struct {
	mock *MockFlow
}

// NewMockFlow creates a new mock instance.
func NewMockFlow(ctrl *gomock.Controller) *MockFlow {
	mock := &MockFlow{ctrl: ctrl}
	mock.recorder = &MockFlowMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFlow) EXPECT() *MockFlowMockRecorder {
	return m.recorder
}

// PreRun mocks base method.
func (m *MockFlow) PreRun(arg0 context.Context, arg1 io.Writer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PreRun", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// PreRun indicates an expected call of PreRun.
func (mr *MockFlowMockRecorder) PreRun(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PreRun", reflect.TypeOf((*MockFlow)(nil).PreRun), arg0, arg1)
}

// Run mocks base method.
func (m *MockFlow) Run() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run")
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run.
func (mr *MockFlowMockRecorder) Run() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockFlow)(nil).Run))
}
