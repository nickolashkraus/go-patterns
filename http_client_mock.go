// Code generated by MockGen. DO NOT EDIT.
// Source: http_client.go

// Package main is a generated GoMock package.
package main

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockFooAPI is a mock of FooAPI interface.
type MockFooAPI struct {
	ctrl     *gomock.Controller
	recorder *MockFooAPIMockRecorder
}

// MockFooAPIMockRecorder is the mock recorder for MockFooAPI.
type MockFooAPIMockRecorder struct {
	mock *MockFooAPI
}

// NewMockFooAPI creates a new mock instance.
func NewMockFooAPI(ctrl *gomock.Controller) *MockFooAPI {
	mock := &MockFooAPI{ctrl: ctrl}
	mock.recorder = &MockFooAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFooAPI) EXPECT() *MockFooAPIMockRecorder {
	return m.recorder
}

// GetFooData mocks base method.
func (m *MockFooAPI) GetFooData(arg0 string) (*FooResponseData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFooData", arg0)
	ret0, _ := ret[0].(*FooResponseData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFooData indicates an expected call of GetFooData.
func (mr *MockFooAPIMockRecorder) GetFooData(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFooData", reflect.TypeOf((*MockFooAPI)(nil).GetFooData), arg0)
}

// GetFooURL mocks base method.
func (m *MockFooAPI) GetFooURL(arg0, arg1, arg2 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFooURL", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetFooURL indicates an expected call of GetFooURL.
func (mr *MockFooAPIMockRecorder) GetFooURL(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFooURL", reflect.TypeOf((*MockFooAPI)(nil).GetFooURL), arg0, arg1, arg2)
}