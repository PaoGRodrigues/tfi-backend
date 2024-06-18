// Code generated by MockGen. DO NOT EDIT.
// Source: app/traffic/domains/traffic.go

// Package mock_domains is a generated GoMock package.
package mock_domains

import (
	reflect "reflect"

	domains "github.com/PaoGRodrigues/tfi-backend/app/traffic/domains"
	gomock "github.com/golang/mock/gomock"
)

// MockTrafficUseCase is a mock of TrafficUseCase interface.
type MockTrafficUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockTrafficUseCaseMockRecorder
}

// MockTrafficUseCaseMockRecorder is the mock recorder for MockTrafficUseCase.
type MockTrafficUseCaseMockRecorder struct {
	mock *MockTrafficUseCase
}

// NewMockTrafficUseCase creates a new mock instance.
func NewMockTrafficUseCase(ctrl *gomock.Controller) *MockTrafficUseCase {
	mock := &MockTrafficUseCase{ctrl: ctrl}
	mock.recorder = &MockTrafficUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTrafficUseCase) EXPECT() *MockTrafficUseCaseMockRecorder {
	return m.recorder
}

// GetActiveFlows mocks base method.
func (m *MockTrafficUseCase) GetActiveFlows() []domains.ActiveFlow {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActiveFlows")
	ret0, _ := ret[0].([]domains.ActiveFlow)
	return ret0
}

// GetActiveFlows indicates an expected call of GetActiveFlows.
func (mr *MockTrafficUseCaseMockRecorder) GetActiveFlows() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActiveFlows", reflect.TypeOf((*MockTrafficUseCase)(nil).GetActiveFlows))
}

// GetAllActiveTraffic mocks base method.
func (m *MockTrafficUseCase) GetAllActiveTraffic() ([]domains.ActiveFlow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllActiveTraffic")
	ret0, _ := ret[0].([]domains.ActiveFlow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllActiveTraffic indicates an expected call of GetAllActiveTraffic.
func (mr *MockTrafficUseCaseMockRecorder) GetAllActiveTraffic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllActiveTraffic", reflect.TypeOf((*MockTrafficUseCase)(nil).GetAllActiveTraffic))
}

// MockTrafficActiveFlowsSearcher is a mock of TrafficActiveFlowsSearcher interface.
type MockTrafficActiveFlowsSearcher struct {
	ctrl     *gomock.Controller
	recorder *MockTrafficActiveFlowsSearcherMockRecorder
}

// MockTrafficActiveFlowsSearcherMockRecorder is the mock recorder for MockTrafficActiveFlowsSearcher.
type MockTrafficActiveFlowsSearcherMockRecorder struct {
	mock *MockTrafficActiveFlowsSearcher
}

// NewMockTrafficActiveFlowsSearcher creates a new mock instance.
func NewMockTrafficActiveFlowsSearcher(ctrl *gomock.Controller) *MockTrafficActiveFlowsSearcher {
	mock := &MockTrafficActiveFlowsSearcher{ctrl: ctrl}
	mock.recorder = &MockTrafficActiveFlowsSearcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTrafficActiveFlowsSearcher) EXPECT() *MockTrafficActiveFlowsSearcherMockRecorder {
	return m.recorder
}

// GetBytesPerCountry mocks base method.
func (m *MockTrafficActiveFlowsSearcher) GetBytesPerCountry() ([]domains.BytesPerCountry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBytesPerCountry")
	ret0, _ := ret[0].([]domains.BytesPerCountry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBytesPerCountry indicates an expected call of GetBytesPerCountry.
func (mr *MockTrafficActiveFlowsSearcherMockRecorder) GetBytesPerCountry() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBytesPerCountry", reflect.TypeOf((*MockTrafficActiveFlowsSearcher)(nil).GetBytesPerCountry))
}

// GetBytesPerDestination mocks base method.
func (m *MockTrafficActiveFlowsSearcher) GetBytesPerDestination() ([]domains.BytesPerDestination, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBytesPerDestination")
	ret0, _ := ret[0].([]domains.BytesPerDestination)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBytesPerDestination indicates an expected call of GetBytesPerDestination.
func (mr *MockTrafficActiveFlowsSearcherMockRecorder) GetBytesPerDestination() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBytesPerDestination", reflect.TypeOf((*MockTrafficActiveFlowsSearcher)(nil).GetBytesPerDestination))
}

// MockTrafficStorage is a mock of TrafficStorage interface.
type MockTrafficStorage struct {
	ctrl     *gomock.Controller
	recorder *MockTrafficStorageMockRecorder
}

// MockTrafficStorageMockRecorder is the mock recorder for MockTrafficStorage.
type MockTrafficStorageMockRecorder struct {
	mock *MockTrafficStorage
}

// NewMockTrafficStorage creates a new mock instance.
func NewMockTrafficStorage(ctrl *gomock.Controller) *MockTrafficStorage {
	mock := &MockTrafficStorage{ctrl: ctrl}
	mock.recorder = &MockTrafficStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTrafficStorage) EXPECT() *MockTrafficStorageMockRecorder {
	return m.recorder
}

// StoreFlows mocks base method.
func (m *MockTrafficStorage) StoreFlows() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreFlows")
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreFlows indicates an expected call of StoreFlows.
func (mr *MockTrafficStorageMockRecorder) StoreFlows() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreFlows", reflect.TypeOf((*MockTrafficStorage)(nil).StoreFlows))
}

// MockTrafficService is a mock of TrafficService interface.
type MockTrafficService struct {
	ctrl     *gomock.Controller
	recorder *MockTrafficServiceMockRecorder
}

// MockTrafficServiceMockRecorder is the mock recorder for MockTrafficService.
type MockTrafficServiceMockRecorder struct {
	mock *MockTrafficService
}

// NewMockTrafficService creates a new mock instance.
func NewMockTrafficService(ctrl *gomock.Controller) *MockTrafficService {
	mock := &MockTrafficService{ctrl: ctrl}
	mock.recorder = &MockTrafficServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTrafficService) EXPECT() *MockTrafficServiceMockRecorder {
	return m.recorder
}

// GetAllActiveTraffic mocks base method.
func (m *MockTrafficService) GetAllActiveTraffic() ([]domains.ActiveFlow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllActiveTraffic")
	ret0, _ := ret[0].([]domains.ActiveFlow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllActiveTraffic indicates an expected call of GetAllActiveTraffic.
func (mr *MockTrafficServiceMockRecorder) GetAllActiveTraffic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllActiveTraffic", reflect.TypeOf((*MockTrafficService)(nil).GetAllActiveTraffic))
}

// MockTrafficRepository is a mock of TrafficRepository interface.
type MockTrafficRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTrafficRepositoryMockRecorder
}

// MockTrafficRepositoryMockRecorder is the mock recorder for MockTrafficRepository.
type MockTrafficRepositoryMockRecorder struct {
	mock *MockTrafficRepository
}

// NewMockTrafficRepository creates a new mock instance.
func NewMockTrafficRepository(ctrl *gomock.Controller) *MockTrafficRepository {
	mock := &MockTrafficRepository{ctrl: ctrl}
	mock.recorder = &MockTrafficRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTrafficRepository) EXPECT() *MockTrafficRepositoryMockRecorder {
	return m.recorder
}

// GetClients mocks base method.
func (m *MockTrafficRepository) GetClients() ([]domains.Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClients")
	ret0, _ := ret[0].([]domains.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClients indicates an expected call of GetClients.
func (mr *MockTrafficRepositoryMockRecorder) GetClients() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClients", reflect.TypeOf((*MockTrafficRepository)(nil).GetClients))
}

// GetFlowByKey mocks base method.
func (m *MockTrafficRepository) GetFlowByKey(arg0 string) (domains.ActiveFlow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFlowByKey", arg0)
	ret0, _ := ret[0].(domains.ActiveFlow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFlowByKey indicates an expected call of GetFlowByKey.
func (mr *MockTrafficRepositoryMockRecorder) GetFlowByKey(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFlowByKey", reflect.TypeOf((*MockTrafficRepository)(nil).GetFlowByKey), arg0)
}

// GetServerByAttr mocks base method.
func (m *MockTrafficRepository) GetServerByAttr(arg0 string) (domains.Server, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServerByAttr", arg0)
	ret0, _ := ret[0].(domains.Server)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServerByAttr indicates an expected call of GetServerByAttr.
func (mr *MockTrafficRepositoryMockRecorder) GetServerByAttr(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServerByAttr", reflect.TypeOf((*MockTrafficRepository)(nil).GetServerByAttr), arg0)
}

// GetServers mocks base method.
func (m *MockTrafficRepository) GetServers() ([]domains.Server, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServers")
	ret0, _ := ret[0].([]domains.Server)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServers indicates an expected call of GetServers.
func (mr *MockTrafficRepositoryMockRecorder) GetServers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServers", reflect.TypeOf((*MockTrafficRepository)(nil).GetServers))
}

// StoreFlows mocks base method.
func (m *MockTrafficRepository) StoreFlows(arg0 []domains.ActiveFlow) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreFlows", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreFlows indicates an expected call of StoreFlows.
func (mr *MockTrafficRepositoryMockRecorder) StoreFlows(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreFlows", reflect.TypeOf((*MockTrafficRepository)(nil).StoreFlows), arg0)
}
