// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/vabispklp/yap/internal/app/storage (interfaces: StorageExpected)

// Package storage_mock is a generated GoMock package.
package storage_mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/vabispklp/yap/internal/app/storage/model"
)

// MockStorageExpected is a mock of StorageExpected interface.
type MockStorageExpected struct {
	ctrl     *gomock.Controller
	recorder *MockStorageExpectedMockRecorder
}

// MockStorageExpectedMockRecorder is the mock recorder for MockStorageExpected.
type MockStorageExpectedMockRecorder struct {
	mock *MockStorageExpected
}

// NewMockStorageExpected creates a new mock instance.
func NewMockStorageExpected(ctrl *gomock.Controller) *MockStorageExpected {
	mock := &MockStorageExpected{ctrl: ctrl}
	mock.recorder = &MockStorageExpectedMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorageExpected) EXPECT() *MockStorageExpectedMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockStorageExpected) Add(arg0 context.Context, arg1 model.ShortURL) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockStorageExpectedMockRecorder) Add(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockStorageExpected)(nil).Add), arg0, arg1)
}

// AddMany mocks base method.
func (m *MockStorageExpected) AddMany(arg0 context.Context, arg1 []model.ShortURL) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddMany", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddMany indicates an expected call of AddMany.
func (mr *MockStorageExpectedMockRecorder) AddMany(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMany", reflect.TypeOf((*MockStorageExpected)(nil).AddMany), arg0, arg1)
}

// Close mocks base method.
func (m *MockStorageExpected) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockStorageExpectedMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockStorageExpected)(nil).Close))
}

// Get mocks base method.
func (m *MockStorageExpected) Get(arg0 context.Context, arg1 string) (*model.ShortURL, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*model.ShortURL)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockStorageExpectedMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockStorageExpected)(nil).Get), arg0, arg1)
}

// GetByUser mocks base method.
func (m *MockStorageExpected) GetByUser(arg0 context.Context, arg1 string) ([]model.ShortURL, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUser", arg0, arg1)
	ret0, _ := ret[0].([]model.ShortURL)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByUser indicates an expected call of GetByUser.
func (mr *MockStorageExpectedMockRecorder) GetByUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUser", reflect.TypeOf((*MockStorageExpected)(nil).GetByUser), arg0, arg1)
}

// Ping mocks base method.
func (m *MockStorageExpected) Ping(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Ping indicates an expected call of Ping.
func (mr *MockStorageExpectedMockRecorder) Ping(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockStorageExpected)(nil).Ping), arg0)
}
