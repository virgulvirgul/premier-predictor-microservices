// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cshep4/premier-predictor-microservices/src/predictionservice/internal/interfaces (interfaces: FixtureService)

// Package fixturemocks is a generated GoMock package.
package fixturemocks

import (
	model "github.com/cshep4/premier-predictor-microservices/src/common/model"
	model0 "github.com/cshep4/premier-predictor-microservices/src/predictionservice/internal/model"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockFixtureService is a mock of FixtureService interface
type MockFixtureService struct {
	ctrl     *gomock.Controller
	recorder *MockFixtureServiceMockRecorder
}

// MockFixtureServiceMockRecorder is the mock recorder for MockFixtureService
type MockFixtureServiceMockRecorder struct {
	mock *MockFixtureService
}

// NewMockFixtureService creates a new mock instance
func NewMockFixtureService(ctrl *gomock.Controller) *MockFixtureService {
	mock := &MockFixtureService{ctrl: ctrl}
	mock.recorder = &MockFixtureServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFixtureService) EXPECT() *MockFixtureServiceMockRecorder {
	return m.recorder
}

// GetMatches mocks base method
func (m *MockFixtureService) GetMatches() ([]model.Fixture, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMatches")
	ret0, _ := ret[0].([]model.Fixture)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMatches indicates an expected call of GetMatches
func (mr *MockFixtureServiceMockRecorder) GetMatches() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMatches", reflect.TypeOf((*MockFixtureService)(nil).GetMatches))
}

// GetTeamForm mocks base method
func (m *MockFixtureService) GetTeamForm() (map[string]model0.TeamForm, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTeamForm")
	ret0, _ := ret[0].(map[string]model0.TeamForm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTeamForm indicates an expected call of GetTeamForm
func (mr *MockFixtureServiceMockRecorder) GetTeamForm() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTeamForm", reflect.TypeOf((*MockFixtureService)(nil).GetTeamForm))
}