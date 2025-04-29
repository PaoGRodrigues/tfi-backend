package repository_test

import (
	"fmt"
	"testing"

	host_domains "github.com/PaoGRodrigues/tfi-backend/app/domain/host"
	domains "github.com/PaoGRodrigues/tfi-backend/app/domain/traffic"
	"github.com/PaoGRodrigues/tfi-backend/app/traffic/repository"
	services_mocks "github.com/PaoGRodrigues/tfi-backend/mocks/services"
	"github.com/go-playground/assert/v2"
	"go.uber.org/mock/gomock"
)

var host = host_domains.Host{
	Name:        "test",
	PrivateHost: false,
	IP:          "123.123.123.123",
	City:        "",
	Country:     "US",
}

var client = domains.Client{
	Name: "test",
	Port: 55672,
	IP:   "192.168.4.9",
}

var server = domains.Server{
	IP:                "123.123.123.123",
	IsBroadcastDomain: false,
	IsDHCP:            false,
	Port:              443,
	Name:              "lib.gen.rus",
	Country:           "US",
	Key:               "12344567",
}

var protocols = domains.Protocol{
	L4: "UDP.Youtube",
	L7: "TLS.GoogleServices",
}

func TestGetServersByAttrReturnServerSuccessfullyByIP(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDatabase := services_mocks.NewMockDatabase(ctrl)
	mockDatabase.EXPECT().GetServerByAttr("123.123.123.123").Return(server, nil)

	trafficStorage := repository.NewFlowsRepo(mockDatabase)
	got, err := trafficStorage.GetServerByAttr("123.123.123.123")

	if err != nil {
		t.Fail()
	}

	assert.Equal(t, server, got)
}

func TestGetServersByAttrReturnServerSuccessfullyByFQDN(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDatabase := services_mocks.NewMockDatabase(ctrl)
	mockDatabase.EXPECT().GetServerByAttr("lib.gen.rus").Return(server, nil)

	trafficStorage := repository.NewFlowsRepo(mockDatabase)
	got, err := trafficStorage.GetServerByAttr("lib.gen.rus")

	if err != nil {
		t.Fail()
	}

	assert.Equal(t, server, got)
}

func TestGetServersByAttrReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDatabase := services_mocks.NewMockDatabase(ctrl)
	mockDatabase.EXPECT().GetServerByAttr("lib.gen.rus").Return(domains.Server{}, fmt.Errorf("Test Error"))

	trafficStorage := repository.NewFlowsRepo(mockDatabase)
	_, err := trafficStorage.GetServerByAttr("lib.gen.rus")

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

	mockDatabase := services_mocks.NewMockDatabase(ctrl)
	mockDatabase.EXPECT().GetClients().Return([]domains.Client{expected}, nil)

	trafficStorage := repository.NewFlowsRepo(mockDatabase)
	got, err := trafficStorage.GetClients()

	if err != nil {
		t.Fail()
	}

	assert.Equal(t, []domains.Client{expected}, got)
}

func TestGetClientsReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDatabase := services_mocks.NewMockDatabase(ctrl)
	mockDatabase.EXPECT().GetClients().Return(nil, fmt.Errorf("Error test"))

	trafficStorage := repository.NewFlowsRepo(mockDatabase)
	_, err := trafficStorage.GetClients()

	if err == nil {
		t.Fail()
	}
}

func TestGetServersListReturnServersSuccessfully(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := domains.Server{
		IP:      "190.190.190.10",
		Name:    "Google.com",
		Port:    443,
		Country: "US",
	}

	mockDatabase := services_mocks.NewMockDatabase(ctrl)
	mockDatabase.EXPECT().GetServers().Return([]domains.Server{expected}, nil)

	trafficStorage := repository.NewFlowsRepo(mockDatabase)
	got, err := trafficStorage.GetServers()

	if err != nil {
		t.Fail()
	}

	assert.Equal(t, []domains.Server{expected}, got)
}

func TestGetServersReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDatabase := services_mocks.NewMockDatabase(ctrl)
	mockDatabase.EXPECT().GetServers().Return(nil, fmt.Errorf("Error test"))

	trafficStorage := repository.NewFlowsRepo(mockDatabase)
	_, err := trafficStorage.GetServers()

	if err == nil {
		t.Fail()
	}
}

func TestGetFlowByKeyReturnFlow(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	activeFlowExpected := domains.TrafficFlow{
		Key:      "12345",
		Client:   client,
		Server:   server,
		Bytes:    1000,
		Protocol: protocols,
	}

	mockDatabase := services_mocks.NewMockDatabase(ctrl)
	mockDatabase.EXPECT().GetFlowByKey(activeFlowExpected.Key).Return(activeFlowExpected, nil)

	trafficStorage := repository.NewFlowsRepo(mockDatabase)
	got, _ := trafficStorage.GetFlowByKey(activeFlowExpected.Key)

	assert.Equal(t, activeFlowExpected, got)
}

func TestGetFlowByKeyReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDatabase := services_mocks.NewMockDatabase(ctrl)
	mockDatabase.EXPECT().GetFlowByKey("1234").Return(domains.TrafficFlow{}, fmt.Errorf("Test error"))

	trafficStorage := repository.NewFlowsRepo(mockDatabase)
	_, err := trafficStorage.GetFlowByKey("1234")

	if err == nil {
		t.Fail()
	}
}

func TestStoreFlowsSuccessfully(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	activeFlows := []domains.TrafficFlow{
		{
			Key:      "12345",
			Client:   client,
			Server:   server,
			Bytes:    1000,
			Protocol: protocols,
		},
		{
			Key:      "12346",
			Client:   client,
			Server:   server,
			Bytes:    1000,
			Protocol: protocols,
		},
		{
			Key:      "123457",
			Client:   client,
			Server:   server,
			Bytes:    1000,
			Protocol: protocols,
		},
	}

	mockDatabase := services_mocks.NewMockDatabase(ctrl)
	mockDatabase.EXPECT().AddActiveFlows(activeFlows).Return(nil)

	trafficStorage := repository.NewFlowsRepo(mockDatabase)
	err := trafficStorage.StoreFlows(activeFlows)

	if err != nil {
		t.Fail()
	}
}

func TestStoreFlowsReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDatabase := services_mocks.NewMockDatabase(ctrl)
	mockDatabase.EXPECT().AddActiveFlows([]domains.TrafficFlow{}).Return(fmt.Errorf("Error Test"))

	trafficStorage := repository.NewFlowsRepo(mockDatabase)
	err := trafficStorage.StoreFlows([]domains.TrafficFlow{})

	if err == nil {
		t.Fail()
	}
}
