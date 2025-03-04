// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/internal/store (interfaces: EncryptionAtRestDescriber)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	mongodbatlasv2 "go.mongodb.org/atlas/mongodbatlasv2"
)

// MockEncryptionAtRestDescriber is a mock of EncryptionAtRestDescriber interface.
type MockEncryptionAtRestDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockEncryptionAtRestDescriberMockRecorder
}

// MockEncryptionAtRestDescriberMockRecorder is the mock recorder for MockEncryptionAtRestDescriber.
type MockEncryptionAtRestDescriberMockRecorder struct {
	mock *MockEncryptionAtRestDescriber
}

// NewMockEncryptionAtRestDescriber creates a new mock instance.
func NewMockEncryptionAtRestDescriber(ctrl *gomock.Controller) *MockEncryptionAtRestDescriber {
	mock := &MockEncryptionAtRestDescriber{ctrl: ctrl}
	mock.recorder = &MockEncryptionAtRestDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEncryptionAtRestDescriber) EXPECT() *MockEncryptionAtRestDescriberMockRecorder {
	return m.recorder
}

// EncryptionAtRest mocks base method.
func (m *MockEncryptionAtRestDescriber) EncryptionAtRest(arg0 string) (*mongodbatlasv2.EncryptionAtRest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EncryptionAtRest", arg0)
	ret0, _ := ret[0].(*mongodbatlasv2.EncryptionAtRest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EncryptionAtRest indicates an expected call of EncryptionAtRest.
func (mr *MockEncryptionAtRestDescriberMockRecorder) EncryptionAtRest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EncryptionAtRest", reflect.TypeOf((*MockEncryptionAtRestDescriber)(nil).EncryptionAtRest), arg0)
}
