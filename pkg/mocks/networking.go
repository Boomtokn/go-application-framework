// Code generated by MockGen. DO NOT EDIT.
// Source: networking.go

// Package mocks is a generated GoMock package.
package mocks

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	auth "github.com/snyk/go-application-framework/pkg/auth"
)

// MockNetworkAccess is a mock of NetworkAccess interface.
type MockNetworkAccess struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkAccessMockRecorder
}

// MockNetworkAccessMockRecorder is the mock recorder for MockNetworkAccess.
type MockNetworkAccessMockRecorder struct {
	mock *MockNetworkAccess
}

// NewMockNetworkAccess creates a new mock instance.
func NewMockNetworkAccess(ctrl *gomock.Controller) *MockNetworkAccess {
	mock := &MockNetworkAccess{ctrl: ctrl}
	mock.recorder = &MockNetworkAccessMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNetworkAccess) EXPECT() *MockNetworkAccessMockRecorder {
	return m.recorder
}

// AddHeaderField mocks base method.
func (m *MockNetworkAccess) AddHeaderField(key, value string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddHeaderField", key, value)
}

// AddHeaderField indicates an expected call of AddHeaderField.
func (mr *MockNetworkAccessMockRecorder) AddHeaderField(key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddHeaderField", reflect.TypeOf((*MockNetworkAccess)(nil).AddHeaderField), key, value)
}

// AddHeaders mocks base method.
func (m *MockNetworkAccess) AddHeaders(request *http.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddHeaders", request)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddHeaders indicates an expected call of AddHeaders.
func (mr *MockNetworkAccessMockRecorder) AddHeaders(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddHeaders", reflect.TypeOf((*MockNetworkAccess)(nil).AddHeaders), request)
}

// AddRootCAs mocks base method.
func (m *MockNetworkAccess) AddRootCAs(pemFileLocation string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRootCAs", pemFileLocation)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddRootCAs indicates an expected call of AddRootCAs.
func (mr *MockNetworkAccessMockRecorder) AddRootCAs(pemFileLocation interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRootCAs", reflect.TypeOf((*MockNetworkAccess)(nil).AddRootCAs), pemFileLocation)
}

// GetAuthenticator mocks base method.
func (m *MockNetworkAccess) GetAuthenticator() auth.Authenticator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAuthenticator")
	ret0, _ := ret[0].(auth.Authenticator)
	return ret0
}

// GetAuthenticator indicates an expected call of GetAuthenticator.
func (mr *MockNetworkAccessMockRecorder) GetAuthenticator() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuthenticator", reflect.TypeOf((*MockNetworkAccess)(nil).GetAuthenticator))
}

// GetHttpClient mocks base method.
func (m *MockNetworkAccess) GetHttpClient() *http.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHttpClient")
	ret0, _ := ret[0].(*http.Client)
	return ret0
}

// GetHttpClient indicates an expected call of GetHttpClient.
func (mr *MockNetworkAccessMockRecorder) GetHttpClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHttpClient", reflect.TypeOf((*MockNetworkAccess)(nil).GetHttpClient))
}

// GetRoundTripper mocks base method.
func (m *MockNetworkAccess) GetRoundTripper() http.RoundTripper {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoundTripper")
	ret0, _ := ret[0].(http.RoundTripper)
	return ret0
}

// GetRoundTripper indicates an expected call of GetRoundTripper.
func (mr *MockNetworkAccessMockRecorder) GetRoundTripper() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoundTripper", reflect.TypeOf((*MockNetworkAccess)(nil).GetRoundTripper))
}

// GetUnauthorizedHttpClient mocks base method.
func (m *MockNetworkAccess) GetUnauthorizedHttpClient() *http.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUnauthorizedHttpClient")
	ret0, _ := ret[0].(*http.Client)
	return ret0
}

// GetUnauthorizedHttpClient indicates an expected call of GetUnauthorizedHttpClient.
func (mr *MockNetworkAccessMockRecorder) GetUnauthorizedHttpClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUnauthorizedHttpClient", reflect.TypeOf((*MockNetworkAccess)(nil).GetUnauthorizedHttpClient))
}
