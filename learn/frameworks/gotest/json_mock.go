// Code generated by MockGen. DO NOT EDIT.
// Source: json.go

// Package gotest is a generated GoMock package.
package gotest

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockPersonEncoder is a mock of PersonEncoder interface
type MockPersonEncoder struct {
	ctrl     *gomock.Controller
	recorder *MockPersonEncoderMockRecorder
}

// MockPersonEncoderMockRecorder is the mock recorder for MockPersonEncoder
type MockPersonEncoderMockRecorder struct {
	mock *MockPersonEncoder
}

// NewMockPersonEncoder creates a new mock instance
func NewMockPersonEncoder(ctrl *gomock.Controller) *MockPersonEncoder {
	mock := &MockPersonEncoder{ctrl: ctrl}
	mock.recorder = &MockPersonEncoderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPersonEncoder) EXPECT() *MockPersonEncoderMockRecorder {
	return m.recorder
}

// encodePerson mocks base method
func (m *MockPersonEncoder) encodePerson(p person) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "encodePerson", p)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// encodePerson indicates an expected call of encodePerson
func (mr *MockPersonEncoderMockRecorder) encodePerson(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "encodePerson", reflect.TypeOf((*MockPersonEncoder)(nil).encodePerson), p)
}

// decodePerson mocks base method
func (m *MockPersonEncoder) decodePerson(encodedPerson string) (person, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "decodePerson", encodedPerson)
	ret0, _ := ret[0].(person)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// decodePerson indicates an expected call of decodePerson
func (mr *MockPersonEncoderMockRecorder) decodePerson(encodedPerson interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "decodePerson", reflect.TypeOf((*MockPersonEncoder)(nil).decodePerson), encodedPerson)
}
