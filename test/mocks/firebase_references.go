// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/firebase/firebase.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	db "firebase.google.com/go/db"
	gomock "github.com/golang/mock/gomock"
	firebase "github.com/readactedworks/go-http-server/pkg/firebase"
)

// MockReferenceCreator is a mock of ReferenceCreator interface.
type MockReferenceCreator struct {
	ctrl     *gomock.Controller
	recorder *MockReferenceCreatorMockRecorder
}

// MockReferenceCreatorMockRecorder is the mock recorder for MockReferenceCreator.
type MockReferenceCreatorMockRecorder struct {
	mock *MockReferenceCreator
}

// NewMockReferenceCreator creates a new mock instance.
func NewMockReferenceCreator(ctrl *gomock.Controller) *MockReferenceCreator {
	mock := &MockReferenceCreator{ctrl: ctrl}
	mock.recorder = &MockReferenceCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReferenceCreator) EXPECT() *MockReferenceCreatorMockRecorder {
	return m.recorder
}

// NewRef mocks base method.
func (m *MockReferenceCreator) NewRef(path string) *db.Ref {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewRef", path)
	ret0, _ := ret[0].(*db.Ref)
	return ret0
}

// NewRef indicates an expected call of NewRef.
func (mr *MockReferenceCreatorMockRecorder) NewRef(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewRef", reflect.TypeOf((*MockReferenceCreator)(nil).NewRef), path)
}

// MockReferenceOperator is a mock of ReferenceOperator interface.
type MockReferenceOperator struct {
	ctrl     *gomock.Controller
	recorder *MockReferenceOperatorMockRecorder
}

// MockReferenceOperatorMockRecorder is the mock recorder for MockReferenceOperator.
type MockReferenceOperatorMockRecorder struct {
	mock *MockReferenceOperator
}

// NewMockReferenceOperator creates a new mock instance.
func NewMockReferenceOperator(ctrl *gomock.Controller) *MockReferenceOperator {
	mock := &MockReferenceOperator{ctrl: ctrl}
	mock.recorder = &MockReferenceOperatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReferenceOperator) EXPECT() *MockReferenceOperatorMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockReferenceOperator) Delete(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockReferenceOperatorMockRecorder) Delete(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockReferenceOperator)(nil).Delete), ctx)
}

// Get mocks base method.
func (m *MockReferenceOperator) Get(ctx context.Context, v any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, v)
	ret0, _ := ret[0].(error)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockReferenceOperatorMockRecorder) Get(ctx, v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockReferenceOperator)(nil).Get), ctx, v)
}

// Set mocks base method.
func (m *MockReferenceOperator) Set(ctx context.Context, v any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", ctx, v)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockReferenceOperatorMockRecorder) Set(ctx, v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockReferenceOperator)(nil).Set), ctx, v)
}

// MockReferenceManager is a mock of ReferenceManager interface.
type MockReferenceManager struct {
	ctrl     *gomock.Controller
	recorder *MockReferenceManagerMockRecorder
}

// MockReferenceManagerMockRecorder is the mock recorder for MockReferenceManager.
type MockReferenceManagerMockRecorder struct {
	mock *MockReferenceManager
}

// NewMockReferenceManager creates a new mock instance.
func NewMockReferenceManager(ctrl *gomock.Controller) *MockReferenceManager {
	mock := &MockReferenceManager{ctrl: ctrl}
	mock.recorder = &MockReferenceManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReferenceManager) EXPECT() *MockReferenceManagerMockRecorder {
	return m.recorder
}

// NewReference mocks base method.
func (m *MockReferenceManager) NewReference(ref *db.Ref) firebase.ReferenceOperator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewReference", ref)
	ret0, _ := ret[0].(firebase.ReferenceOperator)
	return ret0
}

// NewReference indicates an expected call of NewReference.
func (mr *MockReferenceManagerMockRecorder) NewReference(ref interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewReference", reflect.TypeOf((*MockReferenceManager)(nil).NewReference), ref)
}
