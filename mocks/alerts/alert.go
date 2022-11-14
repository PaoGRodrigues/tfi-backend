// Code generated by MockGen. DO NOT EDIT.
// Source: app/alerts/domains/alert.go

// Package mock_domains is a generated GoMock package.
package mock_domains

import (
	reflect "reflect"

	domains "github.com/PaoGRodrigues/tfi-backend/app/alerts/domains"
	gomock "github.com/golang/mock/gomock"
)

// MockAlertUseCase is a mock of AlertUseCase interface.
type MockAlertUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockAlertUseCaseMockRecorder
}

// MockAlertUseCaseMockRecorder is the mock recorder for MockAlertUseCase.
type MockAlertUseCaseMockRecorder struct {
	mock *MockAlertUseCase
}

// NewMockAlertUseCase creates a new mock instance.
func NewMockAlertUseCase(ctrl *gomock.Controller) *MockAlertUseCase {
	mock := &MockAlertUseCase{ctrl: ctrl}
	mock.recorder = &MockAlertUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAlertUseCase) EXPECT() *MockAlertUseCaseMockRecorder {
	return m.recorder
}

// GetAllAlerts mocks base method.
func (m *MockAlertUseCase) GetAllAlerts() ([]domains.Alert, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllAlerts")
	ret0, _ := ret[0].([]domains.Alert)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllAlerts indicates an expected call of GetAllAlerts.
func (mr *MockAlertUseCaseMockRecorder) GetAllAlerts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllAlerts", reflect.TypeOf((*MockAlertUseCase)(nil).GetAllAlerts))
}

// MockAlertService is a mock of AlertService interface.
type MockAlertService struct {
	ctrl     *gomock.Controller
	recorder *MockAlertServiceMockRecorder
}

// MockAlertServiceMockRecorder is the mock recorder for MockAlertService.
type MockAlertServiceMockRecorder struct {
	mock *MockAlertService
}

// NewMockAlertService creates a new mock instance.
func NewMockAlertService(ctrl *gomock.Controller) *MockAlertService {
	mock := &MockAlertService{ctrl: ctrl}
	mock.recorder = &MockAlertServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAlertService) EXPECT() *MockAlertServiceMockRecorder {
	return m.recorder
}

// GetAllAlerts mocks base method.
func (m *MockAlertService) GetAllAlerts(arg0, arg1 int, arg2 string) ([]domains.Alert, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllAlerts", arg0, arg1, arg2)
	ret0, _ := ret[0].([]domains.Alert)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllAlerts indicates an expected call of GetAllAlerts.
func (mr *MockAlertServiceMockRecorder) GetAllAlerts(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllAlerts", reflect.TypeOf((*MockAlertService)(nil).GetAllAlerts), arg0, arg1, arg2)
}
