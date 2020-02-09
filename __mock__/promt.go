// Code generated by MockGen. DO NOT EDIT.
// Source: promt.go

// Package mock_promter is a generated GoMock package.
package mock_promter

import (
	gomock "github.com/golang/mock/gomock"
	promter "github.com/meinto/promter"
	reflect "reflect"
)

// MockPromter is a mock of Promter interface
type MockPromter struct {
	ctrl     *gomock.Controller
	recorder *MockPromterMockRecorder
}

// MockPromterMockRecorder is the mock recorder for MockPromter
type MockPromterMockRecorder struct {
	mock *MockPromter
}

// NewMockPromter creates a new mock instance
func NewMockPromter(ctrl *gomock.Controller) *MockPromter {
	mock := &MockPromter{ctrl: ctrl}
	mock.recorder = &MockPromterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPromter) EXPECT() *MockPromterMockRecorder {
	return m.recorder
}

// YesNo mocks base method
func (m *MockPromter) YesNo(label string) (int, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "YesNo", label)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// YesNo indicates an expected call of YesNo
func (mr *MockPromterMockRecorder) YesNo(label interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "YesNo", reflect.TypeOf((*MockPromter)(nil).YesNo), label)
}

// YesNoDefault mocks base method
func (m *MockPromter) YesNoDefault(label, defaultValue string) (int, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "YesNoDefault", label, defaultValue)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// YesNoDefault indicates an expected call of YesNoDefault
func (mr *MockPromterMockRecorder) YesNoDefault(label, defaultValue interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "YesNoDefault", reflect.TypeOf((*MockPromter)(nil).YesNoDefault), label, defaultValue)
}

// Select mocks base method
func (m *MockPromter) Select(label string, options []string) (int, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Select", label, options)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Select indicates an expected call of Select
func (mr *MockPromterMockRecorder) Select(label, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Select", reflect.TypeOf((*MockPromter)(nil).Select), label, options)
}

// SelectDefault mocks base method
func (m *MockPromter) SelectDefault(label, defaultValue string, options []string) (int, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectDefault", label, defaultValue, options)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SelectDefault indicates an expected call of SelectDefault
func (mr *MockPromterMockRecorder) SelectDefault(label, defaultValue, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectDefault", reflect.TypeOf((*MockPromter)(nil).SelectDefault), label, defaultValue, options)
}

// Text mocks base method
func (m *MockPromter) Text(label string, options ...promter.PromterOptions) (string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{label}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Text", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Text indicates an expected call of Text
func (mr *MockPromterMockRecorder) Text(label interface{}, options ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{label}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Text", reflect.TypeOf((*MockPromter)(nil).Text), varargs...)
}

// TextDefault mocks base method
func (m *MockPromter) TextDefault(label, defaultValue string, options ...promter.PromterOptions) (string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{label, defaultValue}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TextDefault", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TextDefault indicates an expected call of TextDefault
func (mr *MockPromterMockRecorder) TextDefault(label, defaultValue interface{}, options ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{label, defaultValue}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TextDefault", reflect.TypeOf((*MockPromter)(nil).TextDefault), varargs...)
}

// OptionalText mocks base method
func (m *MockPromter) OptionalText(label string, options ...promter.PromterOptions) (string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{label}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "OptionalText", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OptionalText indicates an expected call of OptionalText
func (mr *MockPromterMockRecorder) OptionalText(label interface{}, options ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{label}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OptionalText", reflect.TypeOf((*MockPromter)(nil).OptionalText), varargs...)
}

// OptionalTextDefault mocks base method
func (m *MockPromter) OptionalTextDefault(label, defaultValue string, options ...promter.PromterOptions) (string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{label, defaultValue}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "OptionalTextDefault", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OptionalTextDefault indicates an expected call of OptionalTextDefault
func (mr *MockPromterMockRecorder) OptionalTextDefault(label, defaultValue interface{}, options ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{label, defaultValue}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OptionalTextDefault", reflect.TypeOf((*MockPromter)(nil).OptionalTextDefault), varargs...)
}

// URL mocks base method
func (m *MockPromter) URL(label string, options ...promter.PromterOptions) (string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{label}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "URL", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// URL indicates an expected call of URL
func (mr *MockPromterMockRecorder) URL(label interface{}, options ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{label}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "URL", reflect.TypeOf((*MockPromter)(nil).URL), varargs...)
}

// URLDefault mocks base method
func (m *MockPromter) URLDefault(label, defaultValue string, options ...promter.PromterOptions) (string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{label, defaultValue}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "URLDefault", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// URLDefault indicates an expected call of URLDefault
func (mr *MockPromterMockRecorder) URLDefault(label, defaultValue interface{}, options ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{label, defaultValue}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "URLDefault", reflect.TypeOf((*MockPromter)(nil).URLDefault), varargs...)
}
