// Code generated by MockGen. DO NOT EDIT.
// Source: sample.go

// Package mock_sample is a generated GoMock package.
package mock_sample

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockSample is a mock of Sample interface.
type MockSample struct {
	ctrl     *gomock.Controller
	recorder *MockSampleMockRecorder
}

// MockSampleMockRecorder is the mock recorder for MockSample.
type MockSampleMockRecorder struct {
	mock *MockSample
}

// NewMockSample creates a new mock instance.
func NewMockSample(ctrl *gomock.Controller) *MockSample {
	mock := &MockSample{ctrl: ctrl}
	mock.recorder = &MockSampleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSample) EXPECT() *MockSampleMockRecorder {
	return m.recorder
}

// Method mocks base method.
// func (m *MockSample) Method(s string) int {
// 	m.ctrl.T.Helper()
// 	ret := m.ctrl.Call(m, "Method", s)
// 	ret0, _ := ret[0].(int)
// 	return ret0
// }において、mocksample構造体のメソッドを呼び出すように内部でCall関数で処理が記述されているため、結果が等しくなる!
func (m *MockSample) Method(s string) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Method", s)
	ret0, _ := ret[0].(int)
	return ret0
}

// Method indicates an expected call of Method.
func (mr *MockSampleMockRecorder) Method(s interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Method", reflect.TypeOf((*MockSample)(nil).Method), s)
}
