// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/vabispklp/yap/api/rest/handlers (interfaces: ShortenerExpected)

// Package service_mock is a generated GoMock package.
package service_mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/vabispklp/yap/internal/app/service/model"
	model0 "github.com/vabispklp/yap/internal/app/storage/model"
)

// MockShortenerExpected is a mock of ShortenerExpected interface.
type MockShortenerExpected struct {
	ctrl     *gomock.Controller
	recorder *MockShortenerExpectedMockRecorder
}

// MockShortenerExpectedMockRecorder is the mock recorder for MockShortenerExpected.
type MockShortenerExpectedMockRecorder struct {
	mock *MockShortenerExpected
}

// NewMockShortenerExpected creates a new mock instance.
func NewMockShortenerExpected(ctrl *gomock.Controller) *MockShortenerExpected {
	mock := &MockShortenerExpected{ctrl: ctrl}
	mock.recorder = &MockShortenerExpectedMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockShortenerExpected) EXPECT() *MockShortenerExpectedMockRecorder {
	return m.recorder
}

// AddManyRedirectLink mocks base method.
func (m *MockShortenerExpected) AddManyRedirectLink(arg0 context.Context, arg1 []model.ShortenBatchRequest, arg2 string) ([]model.ShortenBatchResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddManyRedirectLink", arg0, arg1, arg2)
	ret0, _ := ret[0].([]model.ShortenBatchResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddManyRedirectLink indicates an expected call of AddManyRedirectLink.
func (mr *MockShortenerExpectedMockRecorder) AddManyRedirectLink(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddManyRedirectLink", reflect.TypeOf((*MockShortenerExpected)(nil).AddManyRedirectLink), arg0, arg1, arg2)
}

// AddRedirectLink mocks base method.
func (m *MockShortenerExpected) AddRedirectLink(arg0 context.Context, arg1, arg2 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRedirectLink", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddRedirectLink indicates an expected call of AddRedirectLink.
func (mr *MockShortenerExpectedMockRecorder) AddRedirectLink(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRedirectLink", reflect.TypeOf((*MockShortenerExpected)(nil).AddRedirectLink), arg0, arg1, arg2)
}

// DeleteRedirectLinks mocks base method.
func (m *MockShortenerExpected) DeleteRedirectLinks(arg0 context.Context, arg1 []string, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRedirectLinks", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRedirectLinks indicates an expected call of DeleteRedirectLinks.
func (mr *MockShortenerExpectedMockRecorder) DeleteRedirectLinks(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRedirectLinks", reflect.TypeOf((*MockShortenerExpected)(nil).DeleteRedirectLinks), arg0, arg1, arg2)
}

// GetRedirectLink mocks base method.
func (m *MockShortenerExpected) GetRedirectLink(arg0 context.Context, arg1 string) (*model0.ShortURL, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRedirectLink", arg0, arg1)
	ret0, _ := ret[0].(*model0.ShortURL)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRedirectLink indicates an expected call of GetRedirectLink.
func (mr *MockShortenerExpectedMockRecorder) GetRedirectLink(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRedirectLink", reflect.TypeOf((*MockShortenerExpected)(nil).GetRedirectLink), arg0, arg1)
}

// GetUserURLs mocks base method.
func (m *MockShortenerExpected) GetUserURLs(arg0 context.Context, arg1 string) ([]model.ShortenItemResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserURLs", arg0, arg1)
	ret0, _ := ret[0].([]model.ShortenItemResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserURLs indicates an expected call of GetUserURLs.
func (mr *MockShortenerExpectedMockRecorder) GetUserURLs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserURLs", reflect.TypeOf((*MockShortenerExpected)(nil).GetUserURLs), arg0, arg1)
}

// Ping mocks base method.
func (m *MockShortenerExpected) Ping(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Ping indicates an expected call of Ping.
func (mr *MockShortenerExpectedMockRecorder) Ping(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockShortenerExpected)(nil).Ping), arg0)
}
