// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/BelyaevEI/GophKeeper/server/internal/user/userrepository (interfaces: Store)

// Package usermocks is a generated GoMock package.
package usermocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CheckUniqueLogin mocks base method.
func (m *MockStore) CheckUniqueLogin(arg0 context.Context, arg1 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUniqueLogin", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CheckUniqueLogin indicates an expected call of CheckUniqueLogin.
func (mr *MockStoreMockRecorder) CheckUniqueLogin(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUniqueLogin", reflect.TypeOf((*MockStore)(nil).CheckUniqueLogin), arg0, arg1)
}

// GenerateRandomString mocks base method.
func (m *MockStore) GenerateRandomString(arg0 int) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateRandomString", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// GenerateRandomString indicates an expected call of GenerateRandomString.
func (mr *MockStoreMockRecorder) GenerateRandomString(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateRandomString", reflect.TypeOf((*MockStore)(nil).GenerateRandomString), arg0)
}

// GenerateUniqueUserID mocks base method.
func (m *MockStore) GenerateUniqueUserID() (uint32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateUniqueUserID")
	ret0, _ := ret[0].(uint32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateUniqueUserID indicates an expected call of GenerateUniqueUserID.
func (mr *MockStoreMockRecorder) GenerateUniqueUserID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateUniqueUserID", reflect.TypeOf((*MockStore)(nil).GenerateUniqueUserID))
}

// SaveDataNewUser mocks base method.
func (m *MockStore) SaveDataNewUser(arg0 context.Context, arg1, arg2, arg3 string, arg4 uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveDataNewUser", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveDataNewUser indicates an expected call of SaveDataNewUser.
func (mr *MockStoreMockRecorder) SaveDataNewUser(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveDataNewUser", reflect.TypeOf((*MockStore)(nil).SaveDataNewUser), arg0, arg1, arg2, arg3, arg4)
}
