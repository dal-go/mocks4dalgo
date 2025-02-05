// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/dal-go/dalgo/dal (interfaces: TransactionCoordinator,ReadTransactionCoordinator,ReadwriteTransactionCoordinator)
//
// Generated by this command:
//
//	mockgen github.com/dal-go/dalgo/dal TransactionCoordinator,ReadTransactionCoordinator,ReadwriteTransactionCoordinator
//

// Package mock_dal is a generated GoMock package.
package mock_dal

import (
	context "context"
	reflect "reflect"

	dal "github.com/dal-go/dalgo/dal"
	gomock "go.uber.org/mock/gomock"
)

// MockTransactionCoordinator is a mock of TransactionCoordinator interface.
type MockTransactionCoordinator struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionCoordinatorMockRecorder
	isgomock struct{}
}

// MockTransactionCoordinatorMockRecorder is the mock recorder for MockTransactionCoordinator.
type MockTransactionCoordinatorMockRecorder struct {
	mock *MockTransactionCoordinator
}

// NewMockTransactionCoordinator creates a new mock instance.
func NewMockTransactionCoordinator(ctrl *gomock.Controller) *MockTransactionCoordinator {
	mock := &MockTransactionCoordinator{ctrl: ctrl}
	mock.recorder = &MockTransactionCoordinatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionCoordinator) EXPECT() *MockTransactionCoordinatorMockRecorder {
	return m.recorder
}

// RunReadonlyTransaction mocks base method.
func (m *MockTransactionCoordinator) RunReadonlyTransaction(ctx context.Context, f func(context.Context, dal.ReadTransaction) error, options ...dal.TransactionOption) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, f}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunReadonlyTransaction", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunReadonlyTransaction indicates an expected call of RunReadonlyTransaction.
func (mr *MockTransactionCoordinatorMockRecorder) RunReadonlyTransaction(ctx, f any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, f}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunReadonlyTransaction", reflect.TypeOf((*MockTransactionCoordinator)(nil).RunReadonlyTransaction), varargs...)
}

// RunReadwriteTransaction mocks base method.
func (m *MockTransactionCoordinator) RunReadwriteTransaction(ctx context.Context, f func(context.Context, dal.ReadwriteTransaction) error, options ...dal.TransactionOption) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, f}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunReadwriteTransaction", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunReadwriteTransaction indicates an expected call of RunReadwriteTransaction.
func (mr *MockTransactionCoordinatorMockRecorder) RunReadwriteTransaction(ctx, f any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, f}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunReadwriteTransaction", reflect.TypeOf((*MockTransactionCoordinator)(nil).RunReadwriteTransaction), varargs...)
}

// MockReadTransactionCoordinator is a mock of ReadTransactionCoordinator interface.
type MockReadTransactionCoordinator struct {
	ctrl     *gomock.Controller
	recorder *MockReadTransactionCoordinatorMockRecorder
	isgomock struct{}
}

// MockReadTransactionCoordinatorMockRecorder is the mock recorder for MockReadTransactionCoordinator.
type MockReadTransactionCoordinatorMockRecorder struct {
	mock *MockReadTransactionCoordinator
}

// NewMockReadTransactionCoordinator creates a new mock instance.
func NewMockReadTransactionCoordinator(ctrl *gomock.Controller) *MockReadTransactionCoordinator {
	mock := &MockReadTransactionCoordinator{ctrl: ctrl}
	mock.recorder = &MockReadTransactionCoordinatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReadTransactionCoordinator) EXPECT() *MockReadTransactionCoordinatorMockRecorder {
	return m.recorder
}

// RunReadonlyTransaction mocks base method.
func (m *MockReadTransactionCoordinator) RunReadonlyTransaction(ctx context.Context, f func(context.Context, dal.ReadTransaction) error, options ...dal.TransactionOption) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, f}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunReadonlyTransaction", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunReadonlyTransaction indicates an expected call of RunReadonlyTransaction.
func (mr *MockReadTransactionCoordinatorMockRecorder) RunReadonlyTransaction(ctx, f any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, f}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunReadonlyTransaction", reflect.TypeOf((*MockReadTransactionCoordinator)(nil).RunReadonlyTransaction), varargs...)
}

// MockReadwriteTransactionCoordinator is a mock of ReadwriteTransactionCoordinator interface.
type MockReadwriteTransactionCoordinator struct {
	ctrl     *gomock.Controller
	recorder *MockReadwriteTransactionCoordinatorMockRecorder
	isgomock struct{}
}

// MockReadwriteTransactionCoordinatorMockRecorder is the mock recorder for MockReadwriteTransactionCoordinator.
type MockReadwriteTransactionCoordinatorMockRecorder struct {
	mock *MockReadwriteTransactionCoordinator
}

// NewMockReadwriteTransactionCoordinator creates a new mock instance.
func NewMockReadwriteTransactionCoordinator(ctrl *gomock.Controller) *MockReadwriteTransactionCoordinator {
	mock := &MockReadwriteTransactionCoordinator{ctrl: ctrl}
	mock.recorder = &MockReadwriteTransactionCoordinatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReadwriteTransactionCoordinator) EXPECT() *MockReadwriteTransactionCoordinatorMockRecorder {
	return m.recorder
}

// RunReadwriteTransaction mocks base method.
func (m *MockReadwriteTransactionCoordinator) RunReadwriteTransaction(ctx context.Context, f func(context.Context, dal.ReadwriteTransaction) error, options ...dal.TransactionOption) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, f}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunReadwriteTransaction", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunReadwriteTransaction indicates an expected call of RunReadwriteTransaction.
func (mr *MockReadwriteTransactionCoordinatorMockRecorder) RunReadwriteTransaction(ctx, f any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, f}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunReadwriteTransaction", reflect.TypeOf((*MockReadwriteTransactionCoordinator)(nil).RunReadwriteTransaction), varargs...)
}
