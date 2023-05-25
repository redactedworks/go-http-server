// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/firebase/references/references.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	db "firebase.google.com/go/db"
	gomock "github.com/golang/mock/gomock"
	references "github.com/readactedworks/go-http-server/pkg/firebase/references"
)

// MockCreator is a mock of Creator interface.
type MockCreator struct {
	ctrl     *gomock.Controller
	recorder *MockCreatorMockRecorder
}

// MockCreatorMockRecorder is the mock recorder for MockCreator.
type MockCreatorMockRecorder struct {
	mock *MockCreator
}

// NewMockCreator creates a new mock instance.
func NewMockCreator(ctrl *gomock.Controller) *MockCreator {
	mock := &MockCreator{ctrl: ctrl}
	mock.recorder = &MockCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCreator) EXPECT() *MockCreatorMockRecorder {
	return m.recorder
}

// NewRef mocks base method.
func (m *MockCreator) NewRef(path string) *db.Ref {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewRef", path)
	ret0, _ := ret[0].(*db.Ref)
	return ret0
}

// NewRef indicates an expected call of NewRef.
func (mr *MockCreatorMockRecorder) NewRef(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewRef", reflect.TypeOf((*MockCreator)(nil).NewRef), path)
}

// MockOperatorCreator is a mock of OperatorCreator interface.
type MockOperatorCreator struct {
	ctrl     *gomock.Controller
	recorder *MockOperatorCreatorMockRecorder
}

// MockOperatorCreatorMockRecorder is the mock recorder for MockOperatorCreator.
type MockOperatorCreatorMockRecorder struct {
	mock *MockOperatorCreator
}

// NewMockOperatorCreator creates a new mock instance.
func NewMockOperatorCreator(ctrl *gomock.Controller) *MockOperatorCreator {
	mock := &MockOperatorCreator{ctrl: ctrl}
	mock.recorder = &MockOperatorCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOperatorCreator) EXPECT() *MockOperatorCreatorMockRecorder {
	return m.recorder
}

// NewOperator mocks base method.
func (m *MockOperatorCreator) NewOperator(ref references.Operator) references.Operator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewOperator", ref)
	ret0, _ := ret[0].(references.Operator)
	return ret0
}

// NewOperator indicates an expected call of NewOperator.
func (mr *MockOperatorCreatorMockRecorder) NewOperator(ref interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewOperator", reflect.TypeOf((*MockOperatorCreator)(nil).NewOperator), ref)
}

// MockOperator is a mock of Operator interface.
type MockOperator struct {
	ctrl     *gomock.Controller
	recorder *MockOperatorMockRecorder
}

// MockOperatorMockRecorder is the mock recorder for MockOperator.
type MockOperatorMockRecorder struct {
	mock *MockOperator
}

// NewMockOperator creates a new mock instance.
func NewMockOperator(ctrl *gomock.Controller) *MockOperator {
	mock := &MockOperator{ctrl: ctrl}
	mock.recorder = &MockOperatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOperator) EXPECT() *MockOperatorMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockOperator) Delete(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockOperatorMockRecorder) Delete(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockOperator)(nil).Delete), ctx)
}

// Get mocks base method.
func (m *MockOperator) Get(ctx context.Context, v any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, v)
	ret0, _ := ret[0].(error)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockOperatorMockRecorder) Get(ctx, v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockOperator)(nil).Get), ctx, v)
}

// Set mocks base method.
func (m *MockOperator) Set(ctx context.Context, v any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", ctx, v)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockOperatorMockRecorder) Set(ctx, v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockOperator)(nil).Set), ctx, v)
}
