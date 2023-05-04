// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Azure/kubelogin/pkg/token (interfaces: TokenProvider)

// Package mock_token is a generated GoMock package.
package mock_token

import (
	adal "github.com/Azure/go-autorest/autorest/adal"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockTokenProvider is a mock of TokenProvider interface.
type MockTokenProvider struct {
	ctrl     *gomock.Controller
	recorder *MockTokenProviderMockRecorder
}

// MockTokenProviderMockRecorder is the mock recorder for MockTokenProvider.
type MockTokenProviderMockRecorder struct {
	mock *MockTokenProvider
}

// NewMockTokenProvider creates a new mock instance.
func NewMockTokenProvider(ctrl *gomock.Controller) *MockTokenProvider {
	mock := &MockTokenProvider{ctrl: ctrl}
	mock.recorder = &MockTokenProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTokenProvider) EXPECT() *MockTokenProviderMockRecorder {
	return m.recorder
}

// Token mocks base method.
func (m *MockTokenProvider) Token() (adal.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Token")
	ret0, _ := ret[0].(adal.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Token indicates an expected call of Token.
func (mr *MockTokenProviderMockRecorder) Token() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Token", reflect.TypeOf((*MockTokenProvider)(nil).Token))
}