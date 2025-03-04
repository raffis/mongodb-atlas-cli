// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/internal/store (interfaces: PeeringConnectionLister,PeeringConnectionDescriber,PeeringConnectionDeleter,AzurePeeringConnectionCreator,AWSPeeringConnectionCreator,GCPPeeringConnectionCreator,PeeringConnectionCreator)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	mongodbatlas "go.mongodb.org/atlas/mongodbatlas"
	mongodbatlasv2 "go.mongodb.org/atlas/mongodbatlasv2"
)

// MockPeeringConnectionLister is a mock of PeeringConnectionLister interface.
type MockPeeringConnectionLister struct {
	ctrl     *gomock.Controller
	recorder *MockPeeringConnectionListerMockRecorder
}

// MockPeeringConnectionListerMockRecorder is the mock recorder for MockPeeringConnectionLister.
type MockPeeringConnectionListerMockRecorder struct {
	mock *MockPeeringConnectionLister
}

// NewMockPeeringConnectionLister creates a new mock instance.
func NewMockPeeringConnectionLister(ctrl *gomock.Controller) *MockPeeringConnectionLister {
	mock := &MockPeeringConnectionLister{ctrl: ctrl}
	mock.recorder = &MockPeeringConnectionListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPeeringConnectionLister) EXPECT() *MockPeeringConnectionListerMockRecorder {
	return m.recorder
}

// PeeringConnections mocks base method.
func (m *MockPeeringConnectionLister) PeeringConnections(arg0 string, arg1 *mongodbatlas.ContainersListOptions) ([]mongodbatlas.Peer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PeeringConnections", arg0, arg1)
	ret0, _ := ret[0].([]mongodbatlas.Peer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PeeringConnections indicates an expected call of PeeringConnections.
func (mr *MockPeeringConnectionListerMockRecorder) PeeringConnections(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PeeringConnections", reflect.TypeOf((*MockPeeringConnectionLister)(nil).PeeringConnections), arg0, arg1)
}

// MockPeeringConnectionDescriber is a mock of PeeringConnectionDescriber interface.
type MockPeeringConnectionDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockPeeringConnectionDescriberMockRecorder
}

// MockPeeringConnectionDescriberMockRecorder is the mock recorder for MockPeeringConnectionDescriber.
type MockPeeringConnectionDescriberMockRecorder struct {
	mock *MockPeeringConnectionDescriber
}

// NewMockPeeringConnectionDescriber creates a new mock instance.
func NewMockPeeringConnectionDescriber(ctrl *gomock.Controller) *MockPeeringConnectionDescriber {
	mock := &MockPeeringConnectionDescriber{ctrl: ctrl}
	mock.recorder = &MockPeeringConnectionDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPeeringConnectionDescriber) EXPECT() *MockPeeringConnectionDescriberMockRecorder {
	return m.recorder
}

// PeeringConnection mocks base method.
func (m *MockPeeringConnectionDescriber) PeeringConnection(arg0, arg1 string) (*mongodbatlas.Peer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PeeringConnection", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlas.Peer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PeeringConnection indicates an expected call of PeeringConnection.
func (mr *MockPeeringConnectionDescriberMockRecorder) PeeringConnection(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PeeringConnection", reflect.TypeOf((*MockPeeringConnectionDescriber)(nil).PeeringConnection), arg0, arg1)
}

// MockPeeringConnectionDeleter is a mock of PeeringConnectionDeleter interface.
type MockPeeringConnectionDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockPeeringConnectionDeleterMockRecorder
}

// MockPeeringConnectionDeleterMockRecorder is the mock recorder for MockPeeringConnectionDeleter.
type MockPeeringConnectionDeleterMockRecorder struct {
	mock *MockPeeringConnectionDeleter
}

// NewMockPeeringConnectionDeleter creates a new mock instance.
func NewMockPeeringConnectionDeleter(ctrl *gomock.Controller) *MockPeeringConnectionDeleter {
	mock := &MockPeeringConnectionDeleter{ctrl: ctrl}
	mock.recorder = &MockPeeringConnectionDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPeeringConnectionDeleter) EXPECT() *MockPeeringConnectionDeleterMockRecorder {
	return m.recorder
}

// DeletePeeringConnection mocks base method.
func (m *MockPeeringConnectionDeleter) DeletePeeringConnection(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePeeringConnection", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePeeringConnection indicates an expected call of DeletePeeringConnection.
func (mr *MockPeeringConnectionDeleterMockRecorder) DeletePeeringConnection(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePeeringConnection", reflect.TypeOf((*MockPeeringConnectionDeleter)(nil).DeletePeeringConnection), arg0, arg1)
}

// MockAzurePeeringConnectionCreator is a mock of AzurePeeringConnectionCreator interface.
type MockAzurePeeringConnectionCreator struct {
	ctrl     *gomock.Controller
	recorder *MockAzurePeeringConnectionCreatorMockRecorder
}

// MockAzurePeeringConnectionCreatorMockRecorder is the mock recorder for MockAzurePeeringConnectionCreator.
type MockAzurePeeringConnectionCreatorMockRecorder struct {
	mock *MockAzurePeeringConnectionCreator
}

// NewMockAzurePeeringConnectionCreator creates a new mock instance.
func NewMockAzurePeeringConnectionCreator(ctrl *gomock.Controller) *MockAzurePeeringConnectionCreator {
	mock := &MockAzurePeeringConnectionCreator{ctrl: ctrl}
	mock.recorder = &MockAzurePeeringConnectionCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAzurePeeringConnectionCreator) EXPECT() *MockAzurePeeringConnectionCreatorMockRecorder {
	return m.recorder
}

// AzureContainers mocks base method.
func (m *MockAzurePeeringConnectionCreator) AzureContainers(arg0 string) ([]*mongodbatlasv2.AzureCloudProviderContainer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AzureContainers", arg0)
	ret0, _ := ret[0].([]*mongodbatlasv2.AzureCloudProviderContainer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AzureContainers indicates an expected call of AzureContainers.
func (mr *MockAzurePeeringConnectionCreatorMockRecorder) AzureContainers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AzureContainers", reflect.TypeOf((*MockAzurePeeringConnectionCreator)(nil).AzureContainers), arg0)
}

// CreateContainer mocks base method.
func (m *MockAzurePeeringConnectionCreator) CreateContainer(arg0 string, arg1 *mongodbatlasv2.CloudProviderContainer) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateContainer", arg0, arg1)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateContainer indicates an expected call of CreateContainer.
func (mr *MockAzurePeeringConnectionCreatorMockRecorder) CreateContainer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateContainer", reflect.TypeOf((*MockAzurePeeringConnectionCreator)(nil).CreateContainer), arg0, arg1)
}

// CreatePeeringConnection mocks base method.
func (m *MockAzurePeeringConnectionCreator) CreatePeeringConnection(arg0 string, arg1 *mongodbatlas.Peer) (*mongodbatlas.Peer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePeeringConnection", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlas.Peer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePeeringConnection indicates an expected call of CreatePeeringConnection.
func (mr *MockAzurePeeringConnectionCreatorMockRecorder) CreatePeeringConnection(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePeeringConnection", reflect.TypeOf((*MockAzurePeeringConnectionCreator)(nil).CreatePeeringConnection), arg0, arg1)
}

// MockAWSPeeringConnectionCreator is a mock of AWSPeeringConnectionCreator interface.
type MockAWSPeeringConnectionCreator struct {
	ctrl     *gomock.Controller
	recorder *MockAWSPeeringConnectionCreatorMockRecorder
}

// MockAWSPeeringConnectionCreatorMockRecorder is the mock recorder for MockAWSPeeringConnectionCreator.
type MockAWSPeeringConnectionCreatorMockRecorder struct {
	mock *MockAWSPeeringConnectionCreator
}

// NewMockAWSPeeringConnectionCreator creates a new mock instance.
func NewMockAWSPeeringConnectionCreator(ctrl *gomock.Controller) *MockAWSPeeringConnectionCreator {
	mock := &MockAWSPeeringConnectionCreator{ctrl: ctrl}
	mock.recorder = &MockAWSPeeringConnectionCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAWSPeeringConnectionCreator) EXPECT() *MockAWSPeeringConnectionCreatorMockRecorder {
	return m.recorder
}

// AWSContainers mocks base method.
func (m *MockAWSPeeringConnectionCreator) AWSContainers(arg0 string) ([]*mongodbatlasv2.AWSCloudProviderContainer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AWSContainers", arg0)
	ret0, _ := ret[0].([]*mongodbatlasv2.AWSCloudProviderContainer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AWSContainers indicates an expected call of AWSContainers.
func (mr *MockAWSPeeringConnectionCreatorMockRecorder) AWSContainers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AWSContainers", reflect.TypeOf((*MockAWSPeeringConnectionCreator)(nil).AWSContainers), arg0)
}

// CreateContainer mocks base method.
func (m *MockAWSPeeringConnectionCreator) CreateContainer(arg0 string, arg1 *mongodbatlasv2.CloudProviderContainer) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateContainer", arg0, arg1)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateContainer indicates an expected call of CreateContainer.
func (mr *MockAWSPeeringConnectionCreatorMockRecorder) CreateContainer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateContainer", reflect.TypeOf((*MockAWSPeeringConnectionCreator)(nil).CreateContainer), arg0, arg1)
}

// CreatePeeringConnection mocks base method.
func (m *MockAWSPeeringConnectionCreator) CreatePeeringConnection(arg0 string, arg1 *mongodbatlas.Peer) (*mongodbatlas.Peer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePeeringConnection", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlas.Peer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePeeringConnection indicates an expected call of CreatePeeringConnection.
func (mr *MockAWSPeeringConnectionCreatorMockRecorder) CreatePeeringConnection(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePeeringConnection", reflect.TypeOf((*MockAWSPeeringConnectionCreator)(nil).CreatePeeringConnection), arg0, arg1)
}

// MockGCPPeeringConnectionCreator is a mock of GCPPeeringConnectionCreator interface.
type MockGCPPeeringConnectionCreator struct {
	ctrl     *gomock.Controller
	recorder *MockGCPPeeringConnectionCreatorMockRecorder
}

// MockGCPPeeringConnectionCreatorMockRecorder is the mock recorder for MockGCPPeeringConnectionCreator.
type MockGCPPeeringConnectionCreatorMockRecorder struct {
	mock *MockGCPPeeringConnectionCreator
}

// NewMockGCPPeeringConnectionCreator creates a new mock instance.
func NewMockGCPPeeringConnectionCreator(ctrl *gomock.Controller) *MockGCPPeeringConnectionCreator {
	mock := &MockGCPPeeringConnectionCreator{ctrl: ctrl}
	mock.recorder = &MockGCPPeeringConnectionCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGCPPeeringConnectionCreator) EXPECT() *MockGCPPeeringConnectionCreatorMockRecorder {
	return m.recorder
}

// CreateContainer mocks base method.
func (m *MockGCPPeeringConnectionCreator) CreateContainer(arg0 string, arg1 *mongodbatlasv2.CloudProviderContainer) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateContainer", arg0, arg1)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateContainer indicates an expected call of CreateContainer.
func (mr *MockGCPPeeringConnectionCreatorMockRecorder) CreateContainer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateContainer", reflect.TypeOf((*MockGCPPeeringConnectionCreator)(nil).CreateContainer), arg0, arg1)
}

// CreatePeeringConnection mocks base method.
func (m *MockGCPPeeringConnectionCreator) CreatePeeringConnection(arg0 string, arg1 *mongodbatlas.Peer) (*mongodbatlas.Peer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePeeringConnection", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlas.Peer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePeeringConnection indicates an expected call of CreatePeeringConnection.
func (mr *MockGCPPeeringConnectionCreatorMockRecorder) CreatePeeringConnection(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePeeringConnection", reflect.TypeOf((*MockGCPPeeringConnectionCreator)(nil).CreatePeeringConnection), arg0, arg1)
}

// GCPContainers mocks base method.
func (m *MockGCPPeeringConnectionCreator) GCPContainers(arg0 string) ([]*mongodbatlasv2.GCPCloudProviderContainer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GCPContainers", arg0)
	ret0, _ := ret[0].([]*mongodbatlasv2.GCPCloudProviderContainer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GCPContainers indicates an expected call of GCPContainers.
func (mr *MockGCPPeeringConnectionCreatorMockRecorder) GCPContainers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GCPContainers", reflect.TypeOf((*MockGCPPeeringConnectionCreator)(nil).GCPContainers), arg0)
}

// MockPeeringConnectionCreator is a mock of PeeringConnectionCreator interface.
type MockPeeringConnectionCreator struct {
	ctrl     *gomock.Controller
	recorder *MockPeeringConnectionCreatorMockRecorder
}

// MockPeeringConnectionCreatorMockRecorder is the mock recorder for MockPeeringConnectionCreator.
type MockPeeringConnectionCreatorMockRecorder struct {
	mock *MockPeeringConnectionCreator
}

// NewMockPeeringConnectionCreator creates a new mock instance.
func NewMockPeeringConnectionCreator(ctrl *gomock.Controller) *MockPeeringConnectionCreator {
	mock := &MockPeeringConnectionCreator{ctrl: ctrl}
	mock.recorder = &MockPeeringConnectionCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPeeringConnectionCreator) EXPECT() *MockPeeringConnectionCreatorMockRecorder {
	return m.recorder
}

// CreateContainer mocks base method.
func (m *MockPeeringConnectionCreator) CreateContainer(arg0 string, arg1 *mongodbatlasv2.CloudProviderContainer) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateContainer", arg0, arg1)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateContainer indicates an expected call of CreateContainer.
func (mr *MockPeeringConnectionCreatorMockRecorder) CreateContainer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateContainer", reflect.TypeOf((*MockPeeringConnectionCreator)(nil).CreateContainer), arg0, arg1)
}

// CreatePeeringConnection mocks base method.
func (m *MockPeeringConnectionCreator) CreatePeeringConnection(arg0 string, arg1 *mongodbatlas.Peer) (*mongodbatlas.Peer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePeeringConnection", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlas.Peer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePeeringConnection indicates an expected call of CreatePeeringConnection.
func (mr *MockPeeringConnectionCreatorMockRecorder) CreatePeeringConnection(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePeeringConnection", reflect.TypeOf((*MockPeeringConnectionCreator)(nil).CreatePeeringConnection), arg0, arg1)
}
