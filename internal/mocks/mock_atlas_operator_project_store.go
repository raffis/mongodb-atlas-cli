// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/internal/store (interfaces: AtlasOperatorProjectStore)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	mongodbatlas "go.mongodb.org/atlas/mongodbatlas"
	mongodbatlasv2 "go.mongodb.org/atlas/mongodbatlasv2"
)

// MockAtlasOperatorProjectStore is a mock of AtlasOperatorProjectStore interface.
type MockAtlasOperatorProjectStore struct {
	ctrl     *gomock.Controller
	recorder *MockAtlasOperatorProjectStoreMockRecorder
}

// MockAtlasOperatorProjectStoreMockRecorder is the mock recorder for MockAtlasOperatorProjectStore.
type MockAtlasOperatorProjectStoreMockRecorder struct {
	mock *MockAtlasOperatorProjectStore
}

// NewMockAtlasOperatorProjectStore creates a new mock instance.
func NewMockAtlasOperatorProjectStore(ctrl *gomock.Controller) *MockAtlasOperatorProjectStore {
	mock := &MockAtlasOperatorProjectStore{ctrl: ctrl}
	mock.recorder = &MockAtlasOperatorProjectStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAtlasOperatorProjectStore) EXPECT() *MockAtlasOperatorProjectStoreMockRecorder {
	return m.recorder
}

// AlertConfigurations mocks base method.
func (m *MockAtlasOperatorProjectStore) AlertConfigurations(arg0 string, arg1 *mongodbatlas.ListOptions) ([]mongodbatlas.AlertConfiguration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AlertConfigurations", arg0, arg1)
	ret0, _ := ret[0].([]mongodbatlas.AlertConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AlertConfigurations indicates an expected call of AlertConfigurations.
func (mr *MockAtlasOperatorProjectStoreMockRecorder) AlertConfigurations(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AlertConfigurations", reflect.TypeOf((*MockAtlasOperatorProjectStore)(nil).AlertConfigurations), arg0, arg1)
}

// Auditing mocks base method.
func (m *MockAtlasOperatorProjectStore) Auditing(arg0 string) (*mongodbatlasv2.AuditLog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Auditing", arg0)
	ret0, _ := ret[0].(*mongodbatlasv2.AuditLog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Auditing indicates an expected call of Auditing.
func (mr *MockAtlasOperatorProjectStoreMockRecorder) Auditing(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Auditing", reflect.TypeOf((*MockAtlasOperatorProjectStore)(nil).Auditing), arg0)
}

// CloudProviderAccessRoles mocks base method.
func (m *MockAtlasOperatorProjectStore) CloudProviderAccessRoles(arg0 string) (*mongodbatlasv2.CloudProviderAccess, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloudProviderAccessRoles", arg0)
	ret0, _ := ret[0].(*mongodbatlasv2.CloudProviderAccess)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloudProviderAccessRoles indicates an expected call of CloudProviderAccessRoles.
func (mr *MockAtlasOperatorProjectStoreMockRecorder) CloudProviderAccessRoles(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudProviderAccessRoles", reflect.TypeOf((*MockAtlasOperatorProjectStore)(nil).CloudProviderAccessRoles), arg0)
}

// DatabaseRoles mocks base method.
func (m *MockAtlasOperatorProjectStore) DatabaseRoles(arg0 string) ([]mongodbatlasv2.CustomDBRole, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DatabaseRoles", arg0)
	ret0, _ := ret[0].([]mongodbatlasv2.CustomDBRole)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DatabaseRoles indicates an expected call of DatabaseRoles.
func (mr *MockAtlasOperatorProjectStoreMockRecorder) DatabaseRoles(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DatabaseRoles", reflect.TypeOf((*MockAtlasOperatorProjectStore)(nil).DatabaseRoles), arg0)
}

// EncryptionAtRest mocks base method.
func (m *MockAtlasOperatorProjectStore) EncryptionAtRest(arg0 string) (*mongodbatlasv2.EncryptionAtRest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EncryptionAtRest", arg0)
	ret0, _ := ret[0].(*mongodbatlasv2.EncryptionAtRest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EncryptionAtRest indicates an expected call of EncryptionAtRest.
func (mr *MockAtlasOperatorProjectStoreMockRecorder) EncryptionAtRest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EncryptionAtRest", reflect.TypeOf((*MockAtlasOperatorProjectStore)(nil).EncryptionAtRest), arg0)
}

// Integrations mocks base method.
func (m *MockAtlasOperatorProjectStore) Integrations(arg0 string) (*mongodbatlas.ThirdPartyIntegrations, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Integrations", arg0)
	ret0, _ := ret[0].(*mongodbatlas.ThirdPartyIntegrations)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Integrations indicates an expected call of Integrations.
func (mr *MockAtlasOperatorProjectStoreMockRecorder) Integrations(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Integrations", reflect.TypeOf((*MockAtlasOperatorProjectStore)(nil).Integrations), arg0)
}

// MaintenanceWindow mocks base method.
func (m *MockAtlasOperatorProjectStore) MaintenanceWindow(arg0 string) (*mongodbatlasv2.GroupMaintenanceWindow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MaintenanceWindow", arg0)
	ret0, _ := ret[0].(*mongodbatlasv2.GroupMaintenanceWindow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MaintenanceWindow indicates an expected call of MaintenanceWindow.
func (mr *MockAtlasOperatorProjectStoreMockRecorder) MaintenanceWindow(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaintenanceWindow", reflect.TypeOf((*MockAtlasOperatorProjectStore)(nil).MaintenanceWindow), arg0)
}

// PeeringConnections mocks base method.
func (m *MockAtlasOperatorProjectStore) PeeringConnections(arg0 string, arg1 *mongodbatlas.ContainersListOptions) ([]mongodbatlas.Peer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PeeringConnections", arg0, arg1)
	ret0, _ := ret[0].([]mongodbatlas.Peer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PeeringConnections indicates an expected call of PeeringConnections.
func (mr *MockAtlasOperatorProjectStoreMockRecorder) PeeringConnections(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PeeringConnections", reflect.TypeOf((*MockAtlasOperatorProjectStore)(nil).PeeringConnections), arg0, arg1)
}

// PrivateEndpoints mocks base method.
func (m *MockAtlasOperatorProjectStore) PrivateEndpoints(arg0, arg1 string, arg2 *mongodbatlas.ListOptions) ([]mongodbatlas.PrivateEndpointConnection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrivateEndpoints", arg0, arg1, arg2)
	ret0, _ := ret[0].([]mongodbatlas.PrivateEndpointConnection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrivateEndpoints indicates an expected call of PrivateEndpoints.
func (mr *MockAtlasOperatorProjectStoreMockRecorder) PrivateEndpoints(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrivateEndpoints", reflect.TypeOf((*MockAtlasOperatorProjectStore)(nil).PrivateEndpoints), arg0, arg1, arg2)
}

// Project mocks base method.
func (m *MockAtlasOperatorProjectStore) Project(arg0 string) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Project", arg0)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Project indicates an expected call of Project.
func (mr *MockAtlasOperatorProjectStoreMockRecorder) Project(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Project", reflect.TypeOf((*MockAtlasOperatorProjectStore)(nil).Project), arg0)
}

// ProjectIPAccessLists mocks base method.
func (m *MockAtlasOperatorProjectStore) ProjectIPAccessLists(arg0 string, arg1 *mongodbatlas.ListOptions) (*mongodbatlasv2.PaginatedNetworkAccess, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProjectIPAccessLists", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlasv2.PaginatedNetworkAccess)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProjectIPAccessLists indicates an expected call of ProjectIPAccessLists.
func (mr *MockAtlasOperatorProjectStoreMockRecorder) ProjectIPAccessLists(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectIPAccessLists", reflect.TypeOf((*MockAtlasOperatorProjectStore)(nil).ProjectIPAccessLists), arg0, arg1)
}

// ProjectSettings mocks base method.
func (m *MockAtlasOperatorProjectStore) ProjectSettings(arg0 string) (*mongodbatlas.ProjectSettings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProjectSettings", arg0)
	ret0, _ := ret[0].(*mongodbatlas.ProjectSettings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProjectSettings indicates an expected call of ProjectSettings.
func (mr *MockAtlasOperatorProjectStoreMockRecorder) ProjectSettings(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectSettings", reflect.TypeOf((*MockAtlasOperatorProjectStore)(nil).ProjectSettings), arg0)
}

// ProjectTeams mocks base method.
func (m *MockAtlasOperatorProjectStore) ProjectTeams(arg0 string) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProjectTeams", arg0)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProjectTeams indicates an expected call of ProjectTeams.
func (mr *MockAtlasOperatorProjectStoreMockRecorder) ProjectTeams(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectTeams", reflect.TypeOf((*MockAtlasOperatorProjectStore)(nil).ProjectTeams), arg0)
}

// TeamByID mocks base method.
func (m *MockAtlasOperatorProjectStore) TeamByID(arg0, arg1 string) (*mongodbatlas.Team, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TeamByID", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlas.Team)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TeamByID indicates an expected call of TeamByID.
func (mr *MockAtlasOperatorProjectStoreMockRecorder) TeamByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TeamByID", reflect.TypeOf((*MockAtlasOperatorProjectStore)(nil).TeamByID), arg0, arg1)
}

// TeamByName mocks base method.
func (m *MockAtlasOperatorProjectStore) TeamByName(arg0, arg1 string) (*mongodbatlas.Team, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TeamByName", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlas.Team)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TeamByName indicates an expected call of TeamByName.
func (mr *MockAtlasOperatorProjectStoreMockRecorder) TeamByName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TeamByName", reflect.TypeOf((*MockAtlasOperatorProjectStore)(nil).TeamByName), arg0, arg1)
}

// TeamUsers mocks base method.
func (m *MockAtlasOperatorProjectStore) TeamUsers(arg0, arg1 string) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TeamUsers", arg0, arg1)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TeamUsers indicates an expected call of TeamUsers.
func (mr *MockAtlasOperatorProjectStoreMockRecorder) TeamUsers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TeamUsers", reflect.TypeOf((*MockAtlasOperatorProjectStore)(nil).TeamUsers), arg0, arg1)
}
