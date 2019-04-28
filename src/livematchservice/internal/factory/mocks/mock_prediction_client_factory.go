// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cshep4/premier-predictor-microservices/src/livematchservice/internal/interfaces (interfaces: PredictionClientFactory)

// Package factorymocks is a generated GoMock package.
package factorymocks

import (
	gen "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockPredictionClientFactory is a mock of PredictionClientFactory interface
type MockPredictionClientFactory struct {
	ctrl     *gomock.Controller
	recorder *MockPredictionClientFactoryMockRecorder
}

// MockPredictionClientFactoryMockRecorder is the mock recorder for MockPredictionClientFactory
type MockPredictionClientFactoryMockRecorder struct {
	mock *MockPredictionClientFactory
}

// NewMockPredictionClientFactory creates a new mock instance
func NewMockPredictionClientFactory(ctrl *gomock.Controller) *MockPredictionClientFactory {
	mock := &MockPredictionClientFactory{ctrl: ctrl}
	mock.recorder = &MockPredictionClientFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPredictionClientFactory) EXPECT() *MockPredictionClientFactoryMockRecorder {
	return m.recorder
}

// CloseConnection mocks base method
func (m *MockPredictionClientFactory) CloseConnection() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseConnection")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseConnection indicates an expected call of CloseConnection
func (mr *MockPredictionClientFactoryMockRecorder) CloseConnection() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseConnection", reflect.TypeOf((*MockPredictionClientFactory)(nil).CloseConnection))
}

// NewPredictionClient mocks base method
func (m *MockPredictionClientFactory) NewPredictionClient() (gen.PredictionServiceClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewPredictionClient")
	ret0, _ := ret[0].(gen.PredictionServiceClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewPredictionClient indicates an expected call of NewPredictionClient
func (mr *MockPredictionClientFactoryMockRecorder) NewPredictionClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewPredictionClient", reflect.TypeOf((*MockPredictionClientFactory)(nil).NewPredictionClient))
}