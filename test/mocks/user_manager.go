// Code generated by MockGen. DO NOT EDIT.
// Source: api/v1/users.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/readactedworks/go-http-server/api/model"
)

// MockUserDataManager is a mock of UserDataManager interface.
type MockUserDataManager struct {
	ctrl     *gomock.Controller
	recorder *MockUserDataManagerMockRecorder
}

// MockUserDataManagerMockRecorder is the mock recorder for MockUserDataManager.
type MockUserDataManagerMockRecorder struct {
	mock *MockUserDataManager
}

// NewMockUserDataManager creates a new mock instance.
func NewMockUserDataManager(ctrl *gomock.Controller) *MockUserDataManager {
	mock := &MockUserDataManager{ctrl: ctrl}
	mock.recorder = &MockUserDataManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserDataManager) EXPECT() *MockUserDataManagerMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUserDataManager) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, user)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserDataManagerMockRecorder) CreateUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserDataManager)(nil).CreateUser), ctx, user)
}

// DeleteUser mocks base method.
func (m *MockUserDataManager) DeleteUser(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockUserDataManagerMockRecorder) DeleteUser(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUserDataManager)(nil).DeleteUser), ctx, id)
}

// GetUser mocks base method.
func (m *MockUserDataManager) GetUser(ctx context.Context, id string) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", ctx, id)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockUserDataManagerMockRecorder) GetUser(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockUserDataManager)(nil).GetUser), ctx, id)
}

// UpdateUser mocks base method.
func (m *MockUserDataManager) UpdateUser(ctx context.Context, user *model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockUserDataManagerMockRecorder) UpdateUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUserDataManager)(nil).UpdateUser), ctx, user)
}