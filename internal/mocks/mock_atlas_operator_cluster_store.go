// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/internal/store (interfaces: AtlasOperatorClusterStore)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	mongodbatlas "go.mongodb.org/atlas/mongodbatlas"
	mongodbatlasv2 "go.mongodb.org/atlas/mongodbatlasv2"
)

// MockAtlasOperatorClusterStore is a mock of AtlasOperatorClusterStore interface.
type MockAtlasOperatorClusterStore struct {
	ctrl     *gomock.Controller
	recorder *MockAtlasOperatorClusterStoreMockRecorder
}

// MockAtlasOperatorClusterStoreMockRecorder is the mock recorder for MockAtlasOperatorClusterStore.
type MockAtlasOperatorClusterStoreMockRecorder struct {
	mock *MockAtlasOperatorClusterStore
}

// NewMockAtlasOperatorClusterStore creates a new mock instance.
func NewMockAtlasOperatorClusterStore(ctrl *gomock.Controller) *MockAtlasOperatorClusterStore {
	mock := &MockAtlasOperatorClusterStore{ctrl: ctrl}
	mock.recorder = &MockAtlasOperatorClusterStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAtlasOperatorClusterStore) EXPECT() *MockAtlasOperatorClusterStoreMockRecorder {
	return m.recorder
}

// AtlasCluster mocks base method.
func (m *MockAtlasOperatorClusterStore) AtlasCluster(arg0, arg1 string) (*mongodbatlas.AdvancedCluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AtlasCluster", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlas.AdvancedCluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AtlasCluster indicates an expected call of AtlasCluster.
func (mr *MockAtlasOperatorClusterStoreMockRecorder) AtlasCluster(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AtlasCluster", reflect.TypeOf((*MockAtlasOperatorClusterStore)(nil).AtlasCluster), arg0, arg1)
}

// AtlasClusterConfigurationOptions mocks base method.
func (m *MockAtlasOperatorClusterStore) AtlasClusterConfigurationOptions(arg0, arg1 string) (*mongodbatlasv2.ClusterDescriptionProcessArgs, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AtlasClusterConfigurationOptions", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlasv2.ClusterDescriptionProcessArgs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AtlasClusterConfigurationOptions indicates an expected call of AtlasClusterConfigurationOptions.
func (mr *MockAtlasOperatorClusterStoreMockRecorder) AtlasClusterConfigurationOptions(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AtlasClusterConfigurationOptions", reflect.TypeOf((*MockAtlasOperatorClusterStore)(nil).AtlasClusterConfigurationOptions), arg0, arg1)
}

// DescribeSchedule mocks base method.
func (m *MockAtlasOperatorClusterStore) DescribeSchedule(arg0, arg1 string) (*mongodbatlas.CloudProviderSnapshotBackupPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeSchedule", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlas.CloudProviderSnapshotBackupPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeSchedule indicates an expected call of DescribeSchedule.
func (mr *MockAtlasOperatorClusterStoreMockRecorder) DescribeSchedule(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeSchedule", reflect.TypeOf((*MockAtlasOperatorClusterStore)(nil).DescribeSchedule), arg0, arg1)
}

// GetServerlessInstance mocks base method.
func (m *MockAtlasOperatorClusterStore) GetServerlessInstance(arg0, arg1 string) (*mongodbatlasv2.ServerlessInstanceDescription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServerlessInstance", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlasv2.ServerlessInstanceDescription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServerlessInstance indicates an expected call of GetServerlessInstance.
func (mr *MockAtlasOperatorClusterStoreMockRecorder) GetServerlessInstance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServerlessInstance", reflect.TypeOf((*MockAtlasOperatorClusterStore)(nil).GetServerlessInstance), arg0, arg1)
}

// GlobalCluster mocks base method.
func (m *MockAtlasOperatorClusterStore) GlobalCluster(arg0, arg1 string) (*mongodbatlasv2.GeoSharding, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GlobalCluster", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlasv2.GeoSharding)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GlobalCluster indicates an expected call of GlobalCluster.
func (mr *MockAtlasOperatorClusterStoreMockRecorder) GlobalCluster(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GlobalCluster", reflect.TypeOf((*MockAtlasOperatorClusterStore)(nil).GlobalCluster), arg0, arg1)
}

// ProjectClusters mocks base method.
func (m *MockAtlasOperatorClusterStore) ProjectClusters(arg0 string, arg1 *mongodbatlas.ListOptions) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProjectClusters", arg0, arg1)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProjectClusters indicates an expected call of ProjectClusters.
func (mr *MockAtlasOperatorClusterStoreMockRecorder) ProjectClusters(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectClusters", reflect.TypeOf((*MockAtlasOperatorClusterStore)(nil).ProjectClusters), arg0, arg1)
}

// ServerlessInstance mocks base method.
func (m *MockAtlasOperatorClusterStore) ServerlessInstance(arg0, arg1 string) (*mongodbatlas.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerlessInstance", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlas.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServerlessInstance indicates an expected call of ServerlessInstance.
func (mr *MockAtlasOperatorClusterStoreMockRecorder) ServerlessInstance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerlessInstance", reflect.TypeOf((*MockAtlasOperatorClusterStore)(nil).ServerlessInstance), arg0, arg1)
}

// ServerlessInstances mocks base method.
func (m *MockAtlasOperatorClusterStore) ServerlessInstances(arg0 string, arg1 *mongodbatlas.ListOptions) (*mongodbatlasv2.PaginatedServerlessInstanceDescription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerlessInstances", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlasv2.PaginatedServerlessInstanceDescription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServerlessInstances indicates an expected call of ServerlessInstances.
func (mr *MockAtlasOperatorClusterStoreMockRecorder) ServerlessInstances(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerlessInstances", reflect.TypeOf((*MockAtlasOperatorClusterStore)(nil).ServerlessInstances), arg0, arg1)
}

// ServerlessPrivateEndpoints mocks base method.
func (m *MockAtlasOperatorClusterStore) ServerlessPrivateEndpoints(arg0, arg1 string) ([]mongodbatlasv2.ServerlessTenantEndpoint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerlessPrivateEndpoints", arg0, arg1)
	ret0, _ := ret[0].([]mongodbatlasv2.ServerlessTenantEndpoint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServerlessPrivateEndpoints indicates an expected call of ServerlessPrivateEndpoints.
func (mr *MockAtlasOperatorClusterStoreMockRecorder) ServerlessPrivateEndpoints(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerlessPrivateEndpoints", reflect.TypeOf((*MockAtlasOperatorClusterStore)(nil).ServerlessPrivateEndpoints), arg0, arg1)
}
