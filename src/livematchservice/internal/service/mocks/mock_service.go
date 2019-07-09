// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cshep4/premier-predictor-microservices/src/livematchservice/internal/interfaces (interfaces: Service)

// Package livemocks is a generated GoMock package.
package livemocks

import (
	context "context"
	model "github.com/cshep4/premier-predictor-microservices/src/common/model"
	model0 "github.com/cshep4/premier-predictor-microservices/src/livematchservice/internal/model"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockService is a mock of Service interface
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// GetMatchFacts mocks base method
func (m *MockService) GetMatchFacts(arg0 string) (*model.MatchFacts, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMatchFacts", arg0)
	ret0, _ := ret[0].(*model.MatchFacts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMatchFacts indicates an expected call of GetMatchFacts
func (mr *MockServiceMockRecorder) GetMatchFacts(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMatchFacts", reflect.TypeOf((*MockService)(nil).GetMatchFacts), arg0)
}

// GetMatchSummary mocks base method
func (m *MockService) GetMatchSummary(arg0 context.Context, arg1 model0.PredictionRequest) (*model0.MatchSummary, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMatchSummary", arg0, arg1)
	ret0, _ := ret[0].(*model0.MatchSummary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMatchSummary indicates an expected call of GetMatchSummary
func (mr *MockServiceMockRecorder) GetMatchSummary(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMatchSummary", reflect.TypeOf((*MockService)(nil).GetMatchSummary), arg0, arg1)
}

// GetUpcomingMatches mocks base method
func (m *MockService) GetUpcomingMatches() (map[time.Time][]model.MatchFacts, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUpcomingMatches")
	ret0, _ := ret[0].(map[time.Time][]model.MatchFacts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUpcomingMatches indicates an expected call of GetUpcomingMatches
func (mr *MockServiceMockRecorder) GetUpcomingMatches() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUpcomingMatches", reflect.TypeOf((*MockService)(nil).GetUpcomingMatches))
}
