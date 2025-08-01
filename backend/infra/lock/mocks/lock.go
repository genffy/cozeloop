// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/coze-dev/coze-loop/backend/infra/lock (interfaces: ILocker)
//
// Generated by this command:
//
//	mockgen -destination=mocks/lock.go -package=mocks . ILocker
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	time "time"

	lock "github.com/coze-dev/coze-loop/backend/infra/lock"
	gomock "go.uber.org/mock/gomock"
)

// MockILocker is a mock of ILocker interface.
type MockILocker struct {
	ctrl     *gomock.Controller
	recorder *MockILockerMockRecorder
	isgomock struct{}
}

// MockILockerMockRecorder is the mock recorder for MockILocker.
type MockILockerMockRecorder struct {
	mock *MockILocker
}

// NewMockILocker creates a new mock instance.
func NewMockILocker(ctrl *gomock.Controller) *MockILocker {
	mock := &MockILocker{ctrl: ctrl}
	mock.recorder = &MockILockerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockILocker) EXPECT() *MockILockerMockRecorder {
	return m.recorder
}

// ExpireLockIn mocks base method.
func (m *MockILocker) ExpireLockIn(key string, expiresIn time.Duration) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExpireLockIn", key, expiresIn)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExpireLockIn indicates an expected call of ExpireLockIn.
func (mr *MockILockerMockRecorder) ExpireLockIn(key, expiresIn any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExpireLockIn", reflect.TypeOf((*MockILocker)(nil).ExpireLockIn), key, expiresIn)
}

// Lock mocks base method.
func (m *MockILocker) Lock(ctx context.Context, key string, expiresIn time.Duration) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Lock", ctx, key, expiresIn)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Lock indicates an expected call of Lock.
func (mr *MockILockerMockRecorder) Lock(ctx, key, expiresIn any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Lock", reflect.TypeOf((*MockILocker)(nil).Lock), ctx, key, expiresIn)
}

// LockBackoff mocks base method.
func (m *MockILocker) LockBackoff(ctx context.Context, key string, expiresIn, maxWait time.Duration) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LockBackoff", ctx, key, expiresIn, maxWait)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LockBackoff indicates an expected call of LockBackoff.
func (mr *MockILockerMockRecorder) LockBackoff(ctx, key, expiresIn, maxWait any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LockBackoff", reflect.TypeOf((*MockILocker)(nil).LockBackoff), ctx, key, expiresIn, maxWait)
}

// LockBackoffWithRenew mocks base method.
func (m *MockILocker) LockBackoffWithRenew(parent context.Context, key string, ttl, maxHold time.Duration) (bool, context.Context, func(), error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LockBackoffWithRenew", parent, key, ttl, maxHold)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(context.Context)
	ret2, _ := ret[2].(func())
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// LockBackoffWithRenew indicates an expected call of LockBackoffWithRenew.
func (mr *MockILockerMockRecorder) LockBackoffWithRenew(parent, key, ttl, maxHold any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LockBackoffWithRenew", reflect.TypeOf((*MockILocker)(nil).LockBackoffWithRenew), parent, key, ttl, maxHold)
}

// LockWithRenew mocks base method.
func (m *MockILocker) LockWithRenew(parent context.Context, key string, ttl, maxHold time.Duration) (bool, context.Context, func(), error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LockWithRenew", parent, key, ttl, maxHold)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(context.Context)
	ret2, _ := ret[2].(func())
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// LockWithRenew indicates an expected call of LockWithRenew.
func (mr *MockILockerMockRecorder) LockWithRenew(parent, key, ttl, maxHold any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LockWithRenew", reflect.TypeOf((*MockILocker)(nil).LockWithRenew), parent, key, ttl, maxHold)
}

// Unlock mocks base method.
func (m *MockILocker) Unlock(key string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unlock", key)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Unlock indicates an expected call of Unlock.
func (mr *MockILockerMockRecorder) Unlock(key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unlock", reflect.TypeOf((*MockILocker)(nil).Unlock), key)
}

// WithHolder mocks base method.
func (m *MockILocker) WithHolder(holder string) lock.ILocker {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithHolder", holder)
	ret0, _ := ret[0].(lock.ILocker)
	return ret0
}

// WithHolder indicates an expected call of WithHolder.
func (mr *MockILockerMockRecorder) WithHolder(holder any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithHolder", reflect.TypeOf((*MockILocker)(nil).WithHolder), holder)
}
