// Copyright © 2024 Ory Corp
// SPDX-License-Identifier: Apache-2.0

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/yudaprama/fosite/handler/openid (interfaces: OpenIDConnectTokenStrategy)

// Package internal is a generated GoMock package.
package internal

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"

	fosite "github.com/yudaprama/fosite"
)

// MockOpenIDConnectTokenStrategy is a mock of OpenIDConnectTokenStrategy interface.
type MockOpenIDConnectTokenStrategy struct {
	ctrl     *gomock.Controller
	recorder *MockOpenIDConnectTokenStrategyMockRecorder
}

// MockOpenIDConnectTokenStrategyMockRecorder is the mock recorder for MockOpenIDConnectTokenStrategy.
type MockOpenIDConnectTokenStrategyMockRecorder struct {
	mock *MockOpenIDConnectTokenStrategy
}

// NewMockOpenIDConnectTokenStrategy creates a new mock instance.
func NewMockOpenIDConnectTokenStrategy(ctrl *gomock.Controller) *MockOpenIDConnectTokenStrategy {
	mock := &MockOpenIDConnectTokenStrategy{ctrl: ctrl}
	mock.recorder = &MockOpenIDConnectTokenStrategyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOpenIDConnectTokenStrategy) EXPECT() *MockOpenIDConnectTokenStrategyMockRecorder {
	return m.recorder
}

// GenerateIDToken mocks base method.
func (m *MockOpenIDConnectTokenStrategy) GenerateIDToken(arg0 context.Context, arg1 time.Duration, arg2 fosite.Requester) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateIDToken", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateIDToken indicates an expected call of GenerateIDToken.
func (mr *MockOpenIDConnectTokenStrategyMockRecorder) GenerateIDToken(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateIDToken", reflect.TypeOf((*MockOpenIDConnectTokenStrategy)(nil).GenerateIDToken), arg0, arg1, arg2)
}
