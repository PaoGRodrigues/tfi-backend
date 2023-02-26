package usecase_test

import (
	"fmt"
	"testing"

	"github.com/PaoGRodrigues/tfi-backend/app/traffic/domains"
	"github.com/PaoGRodrigues/tfi-backend/app/traffic/usecase"
	mocks "github.com/PaoGRodrigues/tfi-backend/mocks/traffic"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
)

func TestStoreTrafficSuccessfullyGettingTrafficFromSearcher(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := domains.Client{
		Name: "test",
		Port: 55672,
		IP:   "192.168.4.9",
	}
	server := domains.Server{
		IP:                "123.123.123.123",
		IsBroadcastDomain: false,
		IsDHCP:            false,
		Port:              443,
		Name:              "lib.gen.rus",
	}
	protocols := domains.Protocol{
		L4: "UDP.Youtube",
		L7: "TLS.GoogleServices",
	}

	activeFlowToStore := []domains.ActiveFlow{
		domains.ActiveFlow{
			Client:   client,
			Server:   server,
			Bytes:    1000,
			Protocol: protocols,
		},
	}

	mockSearcher := mocks.NewMockTrafficUseCase(ctrl)
	mockSearcher.EXPECT().GetActiveFlows().Return(activeFlowToStore)
	mockTrafficRepoStorage := mocks.NewMockTrafficRepository(ctrl)
	mockTrafficRepoStorage.EXPECT().AddActiveFlows(activeFlowToStore).Return(nil)

	trafficStorage := usecase.NewFlowsStorage(mockSearcher, mockTrafficRepoStorage)
	err := trafficStorage.StoreFlows()

	if err != nil {
		t.Fail()
	}
}

func TestStoreTrafficSuccessfullyGettingTrafficFromEmptySearcherFirstly(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := domains.Client{
		Name: "test",
		Port: 55672,
		IP:   "192.168.4.9",
	}
	server := domains.Server{
		IP:                "123.123.123.123",
		IsBroadcastDomain: false,
		IsDHCP:            false,
		Port:              443,
		Name:              "lib.gen.rus",
	}
	protocols := domains.Protocol{
		L4: "UDP.Youtube",
		L7: "TLS.GoogleServices",
	}

	activeFlowToStore := []domains.ActiveFlow{
		domains.ActiveFlow{
			Client:   client,
			Server:   server,
			Bytes:    1000,
			Protocol: protocols,
		},
	}

	mockSearcher := mocks.NewMockTrafficUseCase(ctrl)
	mockSearcher.EXPECT().GetActiveFlows().Return([]domains.ActiveFlow{})
	mockSearcher.EXPECT().GetAllActiveTraffic().Return(activeFlowToStore, nil)

	mockTrafficRepoStorage := mocks.NewMockTrafficRepository(ctrl)
	mockTrafficRepoStorage.EXPECT().AddActiveFlows(activeFlowToStore).Return(nil)

	trafficStorage := usecase.NewFlowsStorage(mockSearcher, mockTrafficRepoStorage)
	err := trafficStorage.StoreFlows()

	if err != nil {
		t.Fail()
	}
}

func TestStoreTrafficWithError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	activeFlowToStore := []domains.ActiveFlow{
		domains.ActiveFlow{
			Client:   domains.Client{},
			Server:   domains.Server{},
			Bytes:    1000,
			Protocol: domains.Protocol{},
		},
	}

	mockSearcher := mocks.NewMockTrafficUseCase(ctrl)
	mockSearcher.EXPECT().GetActiveFlows().Return(activeFlowToStore)
	mockTrafficRepoStorage := mocks.NewMockTrafficRepository(ctrl)
	mockTrafficRepoStorage.EXPECT().AddActiveFlows(activeFlowToStore).Return(fmt.Errorf("Testing Error"))

	trafficStorage := usecase.NewFlowsStorage(mockSearcher, mockTrafficRepoStorage)
	err := trafficStorage.StoreFlows()

	if err == nil {
		t.Fail()
	}
}

func TestGetServersByAttrReturnServerSuccessfullyByIP(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	server := domains.Server{
		IP:                "123.123.123.123",
		IsBroadcastDomain: false,
		IsDHCP:            false,
		Port:              443,
		Name:              "lib.gen.rus",
	}

	mockSearcher := mocks.NewMockTrafficUseCase(ctrl)
	mockTrafficRepoStorage := mocks.NewMockTrafficRepository(ctrl)
	mockTrafficRepoStorage.EXPECT().GetServerByAttr("123.123.123.123").Return(server, nil)

	trafficStorage := usecase.NewFlowsStorage(mockSearcher, mockTrafficRepoStorage)
	got, err := trafficStorage.GetFlows("123.123.123.123")

	if err != nil {
		t.Fail()
	}

	assert.Equal(t, server, got)
}

func TestGetServersByAttrReturnServerSuccessfullyByFQDN(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	server := domains.Server{
		IP:                "123.123.123.123",
		IsBroadcastDomain: false,
		IsDHCP:            false,
		Port:              443,
		Name:              "lib.gen.rus",
	}

	mockSearcher := mocks.NewMockTrafficUseCase(ctrl)
	mockTrafficRepoStorage := mocks.NewMockTrafficRepository(ctrl)
	mockTrafficRepoStorage.EXPECT().GetServerByAttr("lib.gen.rus").Return(server, nil)

	trafficStorage := usecase.NewFlowsStorage(mockSearcher, mockTrafficRepoStorage)
	got, err := trafficStorage.GetFlows("lib.gen.rus")

	if err != nil {
		t.Fail()
	}

	assert.Equal(t, server, got)
}

func TestGetServersByAttrReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSearcher := mocks.NewMockTrafficUseCase(ctrl)
	mockTrafficRepoStorage := mocks.NewMockTrafficRepository(ctrl)
	mockTrafficRepoStorage.EXPECT().GetServerByAttr("lib.gen.rus").Return(domains.Server{}, fmt.Errorf("Test Error"))

	trafficStorage := usecase.NewFlowsStorage(mockSearcher, mockTrafficRepoStorage)
	_, err := trafficStorage.GetFlows("lib.gen.rus")

	if err == nil {
		t.Fail()
	}
}

func TestGetClientsListReturnClientsSuccessfully(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := domains.Client{
		IP:   "10.10.10.10",
		Name: "host1",
		Port: 34667,
	}

	mockSearcher := mocks.NewMockTrafficUseCase(ctrl)
	mockTrafficRepoStorage := mocks.NewMockTrafficRepository(ctrl)
	mockTrafficRepoStorage.EXPECT().GetClients().Return([]domains.Client{expected}, nil)

	trafficStorage := usecase.NewFlowsStorage(mockSearcher, mockTrafficRepoStorage)
	got, err := trafficStorage.GetClientsList()

	if err != nil {
		t.Fail()
	}

	assert.Equal(t, []domains.Client{expected}, got)
}

func TestGetClientsListReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSearcher := mocks.NewMockTrafficUseCase(ctrl)
	mockTrafficRepoStorage := mocks.NewMockTrafficRepository(ctrl)
	mockTrafficRepoStorage.EXPECT().GetClients().Return(nil, fmt.Errorf("Error test"))

	trafficStorage := usecase.NewFlowsStorage(mockSearcher, mockTrafficRepoStorage)
	_, err := trafficStorage.GetClientsList()

	if err == nil {
		t.Fail()
	}
}

func TestGetServersListReturnServersSuccessfully(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := domains.Server{
		IP:   "190.190.190.10",
		Name: "Google.com",
		Port: 443,
	}

	mockSearcher := mocks.NewMockTrafficUseCase(ctrl)
	mockTrafficRepoStorage := mocks.NewMockTrafficRepository(ctrl)
	mockTrafficRepoStorage.EXPECT().GetServers().Return([]domains.Server{expected}, nil)

	trafficStorage := usecase.NewFlowsStorage(mockSearcher, mockTrafficRepoStorage)
	got, err := trafficStorage.GetServersList()

	if err != nil {
		t.Fail()
	}

	assert.Equal(t, []domains.Server{expected}, got)
}

func TestGetServersListReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSearcher := mocks.NewMockTrafficUseCase(ctrl)
	mockTrafficRepoStorage := mocks.NewMockTrafficRepository(ctrl)
	mockTrafficRepoStorage.EXPECT().GetServers().Return(nil, fmt.Errorf("Error test"))

	trafficStorage := usecase.NewFlowsStorage(mockSearcher, mockTrafficRepoStorage)
	_, err := trafficStorage.GetServersList()

	if err == nil {
		t.Fail()
	}
}
