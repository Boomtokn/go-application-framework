// Code generated by MockGen. DO NOT EDIT.
// Source: networking.go

// Package mocks is a generated GoMock package.
package mocks

import (
	http "net/http"
	url "net/url"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
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

// GetDefaultHeader mocks base method.
func (m *MockNetworkAccess) GetDefaultHeader(url *url.URL) http.Header {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDefaultHeader", url)
	ret0, _ := ret[0].(http.Header)
	return ret0
}

// GetDefaultHeader indicates an expected call of GetDefaultHeader.
func (mr *MockNetworkAccessMockRecorder) GetDefaultHeader(url interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDefaultHeader", reflect.TypeOf((*MockNetworkAccess)(nil).GetDefaultHeader), url)
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

// GetRoundtripper mocks base method.
func (m *MockNetworkAccess) GetRoundtripper() http.RoundTripper {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoundtripper")
	ret0, _ := ret[0].(http.RoundTripper)
	return ret0
}

// GetRoundtripper indicates an expected call of GetRoundtripper.
func (mr *MockNetworkAccessMockRecorder) GetRoundtripper() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoundtripper", reflect.TypeOf((*MockNetworkAccess)(nil).GetRoundtripper))
}
