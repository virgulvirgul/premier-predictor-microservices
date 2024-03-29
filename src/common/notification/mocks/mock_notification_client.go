// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cshep4/premier-predictor-microservices/src/common/interfaces (interfaces: Notifier)

// Package notificationmocks is a generated GoMock package.
package notificationmocks

import (
	context "context"
	model "github.com/cshep4/premier-predictor-microservices/src/common/model"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockNotifier is a mock of Notifier interface
type MockNotifier struct {
	ctrl     *gomock.Controller
	recorder *MockNotifierMockRecorder
}

// MockNotifierMockRecorder is the mock recorder for MockNotifier
type MockNotifierMockRecorder struct {
	mock *MockNotifier
}

// NewMockNotifier creates a new mock instance
func NewMockNotifier(ctrl *gomock.Controller) *MockNotifier {
	mock := &MockNotifier{ctrl: ctrl}
	mock.recorder = &MockNotifierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNotifier) EXPECT() *MockNotifierMockRecorder {
	return m.recorder
}

// Send mocks base method
func (m *MockNotifier) Send(arg0 context.Context, arg1 model.Notification, arg2 ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Send", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockNotifierMockRecorder) Send(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockNotifier)(nil).Send), varargs...)
}
