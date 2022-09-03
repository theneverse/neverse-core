// Code generated by MockGen. DO NOT EDIT.
// Source: types.go

// Package mock_appchainMgr is a generated GoMock package.
package mock_appchainMgr

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	appchain_mgr "github.com/theneverse/neverse-core/appchain-mgr"
	governance "github.com/theneverse/neverse-core/governance"
)

// MockAppchainMgr is a mock of AppchainMgr interface.
type MockAppchainMgr struct {
	ctrl     *gomock.Controller
	recorder *MockAppchainMgrMockRecorder
}

// MockAppchainMgrMockRecorder is the mock recorder for MockAppchainMgr.
type MockAppchainMgrMockRecorder struct {
	mock *MockAppchainMgr
}

// NewMockAppchainMgr creates a new mock instance.
func NewMockAppchainMgr(ctrl *gomock.Controller) *MockAppchainMgr {
	mock := &MockAppchainMgr{ctrl: ctrl}
	mock.recorder = &MockAppchainMgrMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppchainMgr) EXPECT() *MockAppchainMgrMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockAppchainMgr) All(extra []byte) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", extra)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockAppchainMgrMockRecorder) All(extra interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockAppchainMgr)(nil).All), extra)
}

// Audit mocks base method.
func (m *MockAppchainMgr) Audit(proposer string, isApproved int32, desc string) (bool, []byte) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Audit", proposer, isApproved, desc)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].([]byte)
	return ret0, ret1
}

// Audit indicates an expected call of Audit.
func (mr *MockAppchainMgrMockRecorder) Audit(proposer, isApproved, desc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Audit", reflect.TypeOf((*MockAppchainMgr)(nil).Audit), proposer, isApproved, desc)
}

// ChangeStatus mocks base method.
func (m *MockAppchainMgr) ChangeStatus(id, trigger, lastStatus string, extra []byte) (bool, []byte) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeStatus", id, trigger, lastStatus, extra)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].([]byte)
	return ret0, ret1
}

// ChangeStatus indicates an expected call of ChangeStatus.
func (mr *MockAppchainMgrMockRecorder) ChangeStatus(id, trigger, lastStatus, extra interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeStatus", reflect.TypeOf((*MockAppchainMgr)(nil).ChangeStatus), id, trigger, lastStatus, extra)
}

// CountAll mocks base method.
func (m *MockAppchainMgr) CountAll(extra []byte) (bool, []byte) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountAll", extra)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].([]byte)
	return ret0, ret1
}

// CountAll indicates an expected call of CountAll.
func (mr *MockAppchainMgrMockRecorder) CountAll(extra interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountAll", reflect.TypeOf((*MockAppchainMgr)(nil).CountAll), extra)
}

// CountAvailable mocks base method.
func (m *MockAppchainMgr) CountAvailable(extra []byte) (bool, []byte) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountAvailable", extra)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].([]byte)
	return ret0, ret1
}

// CountAvailable indicates an expected call of CountAvailable.
func (mr *MockAppchainMgrMockRecorder) CountAvailable(extra interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountAvailable", reflect.TypeOf((*MockAppchainMgr)(nil).CountAvailable), extra)
}

// DeleteAppchain mocks base method.
func (m *MockAppchainMgr) DeleteAppchain(id string) (bool, []byte) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAppchain", id)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].([]byte)
	return ret0, ret1
}

// DeleteAppchain indicates an expected call of DeleteAppchain.
func (mr *MockAppchainMgrMockRecorder) DeleteAppchain(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAppchain", reflect.TypeOf((*MockAppchainMgr)(nil).DeleteAppchain), id)
}

// FetchAuditRecords mocks base method.
func (m *MockAppchainMgr) FetchAuditRecords(id string) (bool, []byte) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchAuditRecords", id)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].([]byte)
	return ret0, ret1
}

// FetchAuditRecords indicates an expected call of FetchAuditRecords.
func (mr *MockAppchainMgrMockRecorder) FetchAuditRecords(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchAuditRecords", reflect.TypeOf((*MockAppchainMgr)(nil).FetchAuditRecords), id)
}

// GovernancePre mocks base method.
func (m *MockAppchainMgr) GovernancePre(id string, event governance.EventType, extra []byte) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GovernancePre", id, event, extra)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GovernancePre indicates an expected call of GovernancePre.
func (mr *MockAppchainMgrMockRecorder) GovernancePre(id, event, extra interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GovernancePre", reflect.TypeOf((*MockAppchainMgr)(nil).GovernancePre), id, event, extra)
}

// QueryById mocks base method.
func (m *MockAppchainMgr) QueryById(id string, extra []byte) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryById", id, extra)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryById indicates an expected call of QueryById.
func (mr *MockAppchainMgrMockRecorder) QueryById(id, extra interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryById", reflect.TypeOf((*MockAppchainMgr)(nil).QueryById), id, extra)
}

// Register mocks base method.
func (m *MockAppchainMgr) Register(chainInfo *appchain_mgr.Appchain) (bool, []byte) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", chainInfo)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].([]byte)
	return ret0, ret1
}

// Register indicates an expected call of Register.
func (mr *MockAppchainMgrMockRecorder) Register(chainInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockAppchainMgr)(nil).Register), chainInfo)
}

// Update mocks base method.
func (m *MockAppchainMgr) Update(updateInfo *appchain_mgr.Appchain) (bool, []byte) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", updateInfo)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].([]byte)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockAppchainMgrMockRecorder) Update(updateInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockAppchainMgr)(nil).Update), updateInfo)
}
