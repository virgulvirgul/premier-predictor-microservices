// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cshep4/premier-predictor-microservices/src/common/interfaces (interfaces: Authenticator)

// Package authmocks is a generated GoMock package.
package authmocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	http "net/http"
	reflect "reflect"
)

// MockAuthenticator is a mock of Authenticator interface
type MockAuthenticator struct {
	ctrl     *gomock.Controller
	recorder *MockAuthenticatorMockRecorder
}

// MockAuthenticatorMockRecorder is the mock recorder for MockAuthenticator
type MockAuthenticatorMockRecorder struct {
	mock *MockAuthenticator
}

// NewMockAuthenticator creates a new mock instance
func NewMockAuthenticator(ctrl *gomock.Controller) *MockAuthenticator {
	mock := &MockAuthenticator{ctrl: ctrl}
	mock.recorder = &MockAuthenticatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuthenticator) EXPECT() *MockAuthenticatorMockRecorder {
	return m.recorder
}

// GrpcStreamInterceptor mocks base method
func (m *MockAuthenticator) GrpcStreamInterceptor(arg0 interface{}, arg1 grpc.ServerStream, arg2 *grpc.StreamServerInfo, arg3 grpc.StreamHandler) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GrpcStreamInterceptor", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// GrpcStreamInterceptor indicates an expected call of GrpcStreamInterceptor
func (mr *MockAuthenticatorMockRecorder) GrpcStreamInterceptor(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GrpcStreamInterceptor", reflect.TypeOf((*MockAuthenticator)(nil).GrpcStreamInterceptor), arg0, arg1, arg2, arg3)
}

// GrpcUnaryInterceptor mocks base method
func (m *MockAuthenticator) GrpcUnaryInterceptor(arg0 context.Context, arg1 interface{}, arg2 *grpc.UnaryServerInfo, arg3 grpc.UnaryHandler) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GrpcUnaryInterceptor", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GrpcUnaryInterceptor indicates an expected call of GrpcUnaryInterceptor
func (mr *MockAuthenticatorMockRecorder) GrpcUnaryInterceptor(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GrpcUnaryInterceptor", reflect.TypeOf((*MockAuthenticator)(nil).GrpcUnaryInterceptor), arg0, arg1, arg2, arg3)
}

// HttpMiddleware mocks base method
func (m *MockAuthenticator) HttpMiddleware(arg0 http.Handler) http.Handler {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HttpMiddleware", arg0)
	ret0, _ := ret[0].(http.Handler)
	return ret0
}

// HttpMiddleware indicates an expected call of HttpMiddleware
func (mr *MockAuthenticatorMockRecorder) HttpMiddleware(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HttpMiddleware", reflect.TypeOf((*MockAuthenticator)(nil).HttpMiddleware), arg0)
}