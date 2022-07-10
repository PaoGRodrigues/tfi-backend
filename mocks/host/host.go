// Code generated by MockGen. DO NOT EDIT.
// Source: app/host/domains/host.go

// Package mock_domains is a generated GoMock package.
package mock_domains

import (
	reflect "reflect"

	domains "github.com/PaoGRodrigues/tfi-backend/app/host/domains"
	gomock "github.com/golang/mock/gomock"
)

// MockHostUseCase is a mock of HostUseCase interface.
type MockHostUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockHostUseCaseMockRecorder
}

// MockHostUseCaseMockRecorder is the mock recorder for MockHostUseCase.
type MockHostUseCaseMockRecorder struct {
	mock *MockHostUseCase
}

// NewMockHostUseCase creates a new mock instance.
func NewMockHostUseCase(ctrl *gomock.Controller) *MockHostUseCase {
	mock := &MockHostUseCase{ctrl: ctrl}
	mock.recorder = &MockHostUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHostUseCase) EXPECT() *MockHostUseCaseMockRecorder {
	return m.recorder
}

// GetAllHosts mocks base method.
func (m *MockHostUseCase) GetAllHosts() ([]domains.Host, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllHosts")
	ret0, _ := ret[0].([]domains.Host)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllHosts indicates an expected call of GetAllHosts.
func (mr *MockHostUseCaseMockRecorder) GetAllHosts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllHosts", reflect.TypeOf((*MockHostUseCase)(nil).GetAllHosts))
}

// GetHosts mocks base method.
func (m *MockHostUseCase) GetHosts() []domains.Host {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHosts")
	ret0, _ := ret[0].([]domains.Host)
	return ret0
}

// GetHosts indicates an expected call of GetHosts.
func (mr *MockHostUseCaseMockRecorder) GetHosts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHosts", reflect.TypeOf((*MockHostUseCase)(nil).GetHosts))
}

// MockHostsFilter is a mock of HostsFilter interface.
type MockHostsFilter struct {
	ctrl     *gomock.Controller
	recorder *MockHostsFilterMockRecorder
}

// MockHostsFilterMockRecorder is the mock recorder for MockHostsFilter.
type MockHostsFilterMockRecorder struct {
	mock *MockHostsFilter
}

// NewMockHostsFilter creates a new mock instance.
func NewMockHostsFilter(ctrl *gomock.Controller) *MockHostsFilter {
	mock := &MockHostsFilter{ctrl: ctrl}
	mock.recorder = &MockHostsFilterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHostsFilter) EXPECT() *MockHostsFilterMockRecorder {
	return m.recorder
}

// GetLocalHosts mocks base method.
func (m *MockHostsFilter) GetLocalHosts() ([]domains.Host, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLocalHosts")
	ret0, _ := ret[0].([]domains.Host)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLocalHosts indicates an expected call of GetLocalHosts.
func (mr *MockHostsFilterMockRecorder) GetLocalHosts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLocalHosts", reflect.TypeOf((*MockHostsFilter)(nil).GetLocalHosts))
}

// GetRemoteHosts mocks base method.
func (m *MockHostsFilter) GetRemoteHosts() ([]domains.Host, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRemoteHosts")
	ret0, _ := ret[0].([]domains.Host)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRemoteHosts indicates an expected call of GetRemoteHosts.
func (mr *MockHostsFilterMockRecorder) GetRemoteHosts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRemoteHosts", reflect.TypeOf((*MockHostsFilter)(nil).GetRemoteHosts))
}

// MockHostRepository is a mock of HostRepository interface.
type MockHostRepository struct {
	ctrl     *gomock.Controller
	recorder *MockHostRepositoryMockRecorder
}

// MockHostRepositoryMockRecorder is the mock recorder for MockHostRepository.
type MockHostRepositoryMockRecorder struct {
	mock *MockHostRepository
}

// NewMockHostRepository creates a new mock instance.
func NewMockHostRepository(ctrl *gomock.Controller) *MockHostRepository {
	mock := &MockHostRepository{ctrl: ctrl}
	mock.recorder = &MockHostRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHostRepository) EXPECT() *MockHostRepositoryMockRecorder {
	return m.recorder
}

// GetAllHosts mocks base method.
func (m *MockHostRepository) GetAllHosts() ([]domains.Host, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllHosts")
	ret0, _ := ret[0].([]domains.Host)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllHosts indicates an expected call of GetAllHosts.
func (mr *MockHostRepositoryMockRecorder) GetAllHosts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllHosts", reflect.TypeOf((*MockHostRepository)(nil).GetAllHosts))
}
