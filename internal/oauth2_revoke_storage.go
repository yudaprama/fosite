// Copyright © 2024 Ory Corp
// SPDX-License-Identifier: Apache-2.0

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/yudaprama/fosite/handler/oauth2 (interfaces: TokenRevocationStorage)

// Package internal is a generated GoMock package.
package internal

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	fosite "github.com/yudaprama/fosite"
)

// MockTokenRevocationStorage is a mock of TokenRevocationStorage interface.
type MockTokenRevocationStorage struct {
	ctrl     *gomock.Controller
	recorder *MockTokenRevocationStorageMockRecorder
}

// MockTokenRevocationStorageMockRecorder is the mock recorder for MockTokenRevocationStorage.
type MockTokenRevocationStorageMockRecorder struct {
	mock *MockTokenRevocationStorage
}

// NewMockTokenRevocationStorage creates a new mock instance.
func NewMockTokenRevocationStorage(ctrl *gomock.Controller) *MockTokenRevocationStorage {
	mock := &MockTokenRevocationStorage{ctrl: ctrl}
	mock.recorder = &MockTokenRevocationStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTokenRevocationStorage) EXPECT() *MockTokenRevocationStorageMockRecorder {
	return m.recorder
}

// CreateAccessTokenSession mocks base method.
func (m *MockTokenRevocationStorage) CreateAccessTokenSession(arg0 context.Context, arg1 string, arg2 fosite.Requester) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccessTokenSession", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAccessTokenSession indicates an expected call of CreateAccessTokenSession.
func (mr *MockTokenRevocationStorageMockRecorder) CreateAccessTokenSession(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccessTokenSession", reflect.TypeOf((*MockTokenRevocationStorage)(nil).CreateAccessTokenSession), arg0, arg1, arg2)
}

// CreateRefreshTokenSession mocks base method.
func (m *MockTokenRevocationStorage) CreateRefreshTokenSession(arg0 context.Context, arg1 string, arg2 fosite.Requester) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRefreshTokenSession", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateRefreshTokenSession indicates an expected call of CreateRefreshTokenSession.
func (mr *MockTokenRevocationStorageMockRecorder) CreateRefreshTokenSession(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRefreshTokenSession", reflect.TypeOf((*MockTokenRevocationStorage)(nil).CreateRefreshTokenSession), arg0, arg1, arg2)
}

// DeleteAccessTokenSession mocks base method.
func (m *MockTokenRevocationStorage) DeleteAccessTokenSession(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAccessTokenSession", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAccessTokenSession indicates an expected call of DeleteAccessTokenSession.
func (mr *MockTokenRevocationStorageMockRecorder) DeleteAccessTokenSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAccessTokenSession", reflect.TypeOf((*MockTokenRevocationStorage)(nil).DeleteAccessTokenSession), arg0, arg1)
}

// DeleteRefreshTokenSession mocks base method.
func (m *MockTokenRevocationStorage) DeleteRefreshTokenSession(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRefreshTokenSession", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRefreshTokenSession indicates an expected call of DeleteRefreshTokenSession.
func (mr *MockTokenRevocationStorageMockRecorder) DeleteRefreshTokenSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRefreshTokenSession", reflect.TypeOf((*MockTokenRevocationStorage)(nil).DeleteRefreshTokenSession), arg0, arg1)
}

// GetAccessTokenSession mocks base method.
func (m *MockTokenRevocationStorage) GetAccessTokenSession(arg0 context.Context, arg1 string, arg2 fosite.Session) (fosite.Requester, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccessTokenSession", arg0, arg1, arg2)
	ret0, _ := ret[0].(fosite.Requester)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccessTokenSession indicates an expected call of GetAccessTokenSession.
func (mr *MockTokenRevocationStorageMockRecorder) GetAccessTokenSession(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccessTokenSession", reflect.TypeOf((*MockTokenRevocationStorage)(nil).GetAccessTokenSession), arg0, arg1, arg2)
}

// GetRefreshTokenSession mocks base method.
func (m *MockTokenRevocationStorage) GetRefreshTokenSession(arg0 context.Context, arg1 string, arg2 fosite.Session) (fosite.Requester, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRefreshTokenSession", arg0, arg1, arg2)
	ret0, _ := ret[0].(fosite.Requester)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRefreshTokenSession indicates an expected call of GetRefreshTokenSession.
func (mr *MockTokenRevocationStorageMockRecorder) GetRefreshTokenSession(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRefreshTokenSession", reflect.TypeOf((*MockTokenRevocationStorage)(nil).GetRefreshTokenSession), arg0, arg1, arg2)
}

// RevokeAccessToken mocks base method.
func (m *MockTokenRevocationStorage) RevokeAccessToken(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RevokeAccessToken", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RevokeAccessToken indicates an expected call of RevokeAccessToken.
func (mr *MockTokenRevocationStorageMockRecorder) RevokeAccessToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokeAccessToken", reflect.TypeOf((*MockTokenRevocationStorage)(nil).RevokeAccessToken), arg0, arg1)
}

// RevokeRefreshToken mocks base method.
func (m *MockTokenRevocationStorage) RevokeRefreshToken(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RevokeRefreshToken", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RevokeRefreshToken indicates an expected call of RevokeRefreshToken.
func (mr *MockTokenRevocationStorageMockRecorder) RevokeRefreshToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokeRefreshToken", reflect.TypeOf((*MockTokenRevocationStorage)(nil).RevokeRefreshToken), arg0, arg1)
}

// RevokeRefreshTokenMaybeGracePeriod mocks base method.
func (m *MockTokenRevocationStorage) RevokeRefreshTokenMaybeGracePeriod(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RevokeRefreshTokenMaybeGracePeriod", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RevokeRefreshTokenMaybeGracePeriod indicates an expected call of RevokeRefreshTokenMaybeGracePeriod.
func (mr *MockTokenRevocationStorageMockRecorder) RevokeRefreshTokenMaybeGracePeriod(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokeRefreshTokenMaybeGracePeriod", reflect.TypeOf((*MockTokenRevocationStorage)(nil).RevokeRefreshTokenMaybeGracePeriod), arg0, arg1, arg2)
}
