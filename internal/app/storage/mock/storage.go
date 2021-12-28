// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/vabispklp/yap/internal/app/storage (interfaces: ShortenerExpected)

// Package storage_mock is a generated GoMock package.
package storage_mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/vabispklp/yap/internal/app/model"
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

// AddRedirectLink mocks base method.
func (m *MockShortenerExpected) AddRedirectLink(arg0 context.Context, arg1 *model.ShortURL) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRedirectLink", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddRedirectLink indicates an expected call of AddRedirectLink.
func (mr *MockShortenerExpectedMockRecorder) AddRedirectLink(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRedirectLink", reflect.TypeOf((*MockShortenerExpected)(nil).AddRedirectLink), arg0, arg1)
}

// GetRedirectLink mocks base method.
func (m *MockShortenerExpected) GetRedirectLink(arg0 context.Context, arg1 string) (*model.ShortURL, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRedirectLink", arg0, arg1)
	ret0, _ := ret[0].(*model.ShortURL)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRedirectLink indicates an expected call of GetRedirectLink.
func (mr *MockShortenerExpectedMockRecorder) GetRedirectLink(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRedirectLink", reflect.TypeOf((*MockShortenerExpected)(nil).GetRedirectLink), arg0, arg1)
}
