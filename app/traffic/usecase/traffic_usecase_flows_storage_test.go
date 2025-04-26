package usecase_test

import (
	"fmt"
	"testing"

	host_domains "github.com/PaoGRodrigues/tfi-backend/app/domain/host"
	"github.com/PaoGRodrigues/tfi-backend/app/traffic/domains"
	"github.com/PaoGRodrigues/tfi-backend/app/traffic/usecase"
	host_mocks "github.com/PaoGRodrigues/tfi-backend/mocks/hosts"
	mocks "github.com/PaoGRodrigues/tfi-backend/mocks/traffic"
	"github.com/golang/mock/gomock"
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

var broadcastserver = domains.Server{
	IP:                "1.1.1.1",
	IsBroadcastDomain: true,
	IsDHCP:            false,
	Port:              443,
	Name:              "SARASA",
	Country:           "US",
	Key:               "12344569",
}

var broadcastserverchanged = domains.Server{
	IP:                "1.1.1.1",
	IsBroadcastDomain: true,
	IsDHCP:            false,
	Port:              443,
	Name:              "1.1.1.1",
	Country:           "US",
	Key:               "12344569",
}
var publichost = host_domains.Host{
	Name:        "SARASA",
	PrivateHost: false,
	IP:          "1.1.1.1",
	City:        "",
	Country:     "US",
}

func TestStoreTrafficSuccessfullyGettingTrafficFromSearcher(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	activeFlowToStore := []domains.ActiveFlow{
		{
			Client:   client,
			Server:   server,
			Bytes:    1000,
			Protocol: protocols,
		},
	}

	mockSearcher := mocks.NewMockTrafficUseCase(ctrl)
	mockSearcher.EXPECT().GetActiveFlows().Return(activeFlowToStore)
	mockHostsStorage := host_mocks.NewMockHostsStorage(ctrl)
	mockHostsStorage.EXPECT().GetHostByIp(server.IP).Return(host, nil)
	mockTrafficRepoStorage := mocks.NewMockTrafficRepository(ctrl)
	mockTrafficRepoStorage.EXPECT().StoreFlows(activeFlowToStore).Return(nil)

	trafficStorage := usecase.NewFlowsStorage(mockSearcher, mockTrafficRepoStorage, mockHostsStorage)
	err := trafficStorage.StoreFlows()

	if err != nil {
		t.Fail()
	}
}

func TestStoreTrafficSuccessfullyGettingTrafficFromEmptySearcherFirstly(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	activeFlowToStore := []domains.ActiveFlow{
		{
			Client:   client,
			Server:   server,
			Bytes:    1000,
			Protocol: protocols,
		},
	}

	mockSearcher := mocks.NewMockTrafficUseCase(ctrl)
	mockSearcher.EXPECT().GetActiveFlows().Return([]domains.ActiveFlow{})
	mockSearcher.EXPECT().GetAllActiveTraffic().Return(activeFlowToStore, nil)
	mockHostsStorage := host_mocks.NewMockHostsStorage(ctrl)
	mockHostsStorage.EXPECT().GetHostByIp(server.IP).Return(host, nil)
	mockTrafficRepoStorage := mocks.NewMockTrafficRepository(ctrl)
	mockTrafficRepoStorage.EXPECT().StoreFlows(activeFlowToStore).Return(nil)

	trafficStorage := usecase.NewFlowsStorage(mockSearcher, mockTrafficRepoStorage, mockHostsStorage)
	err := trafficStorage.StoreFlows()

	if err != nil {
		t.Fail()
	}
}

func TestStoreTrafficWithError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	activeFlowToStore := []domains.ActiveFlow{
		{
			Client:   client,
			Server:   server,
			Bytes:    1000,
			Protocol: protocols,
		},
	}

	mockSearcher := mocks.NewMockTrafficUseCase(ctrl)
	mockSearcher.EXPECT().GetActiveFlows().Return(activeFlowToStore)
	mockHostsStorage := host_mocks.NewMockHostsStorage(ctrl)
	mockHostsStorage.EXPECT().GetHostByIp(server.IP).Return(host, nil)
	mockTrafficRepoStorage := mocks.NewMockTrafficRepository(ctrl)
	mockTrafficRepoStorage.EXPECT().StoreFlows(activeFlowToStore).Return(fmt.Errorf("Testing Error"))

	trafficStorage := usecase.NewFlowsStorage(mockSearcher, mockTrafficRepoStorage, mockHostsStorage)
	err := trafficStorage.StoreFlows()

	if err == nil {
		t.Fail()
	}
}

func TestStoreTrafficWithGetTrafficReturningError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSearcher := mocks.NewMockTrafficUseCase(ctrl)
	mockSearcher.EXPECT().GetActiveFlows().Return([]domains.ActiveFlow{})
	mockSearcher.EXPECT().GetAllActiveTraffic().Return([]domains.ActiveFlow{}, fmt.Errorf("Test error"))
	mockHostsStorage := host_mocks.NewMockHostsStorage(ctrl)
	mockTrafficRepoStorage := mocks.NewMockTrafficRepository(ctrl)

	trafficStorage := usecase.NewFlowsStorage(mockSearcher, mockTrafficRepoStorage, mockHostsStorage)
	err := trafficStorage.StoreFlows()

	if err == nil {
		t.Fail()
	}
}

func TestStoreTrafficWithErrorInEnrichData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	activeFlowToStore := []domains.ActiveFlow{
		{
			Client:   client,
			Server:   server,
			Bytes:    1000,
			Protocol: protocols,
		},
	}

	mockSearcher := mocks.NewMockTrafficUseCase(ctrl)
	mockSearcher.EXPECT().GetActiveFlows().Return(activeFlowToStore)
	mockHostsStorage := host_mocks.NewMockHostsStorage(ctrl)
	mockHostsStorage.EXPECT().GetHostByIp(server.IP).Return(host_domains.Host{}, fmt.Errorf("Test error"))
	mockTrafficRepoStorage := mocks.NewMockTrafficRepository(ctrl)

	trafficStorage := usecase.NewFlowsStorage(mockSearcher, mockTrafficRepoStorage, mockHostsStorage)
	err := trafficStorage.StoreFlows()

	if err == nil {
		t.Fail()
	}
}

func TestStoreBroadcastServerSuccessfully(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	got := []domains.ActiveFlow{
		{
			Client:   client,
			Server:   broadcastserver,
			Bytes:    1000,
			Protocol: protocols,
		},
	}

	expected := []domains.ActiveFlow{
		{
			Client:   client,
			Server:   broadcastserverchanged,
			Bytes:    1000,
			Protocol: protocols,
		},
	}

	mockSearcher := mocks.NewMockTrafficUseCase(ctrl)
	mockSearcher.EXPECT().GetActiveFlows().Return(got)
	mockHostsStorage := host_mocks.NewMockHostsStorage(ctrl)
	mockHostsStorage.EXPECT().GetHostByIp(broadcastserver.IP).Return(publichost, nil)
	mockTrafficRepoStorage := mocks.NewMockTrafficRepository(ctrl)
	mockTrafficRepoStorage.EXPECT().StoreFlows(expected).Return(nil)

	trafficStorage := usecase.NewFlowsStorage(mockSearcher, mockTrafficRepoStorage, mockHostsStorage)
	err := trafficStorage.StoreFlows()

	if err != nil {
		t.Fail()
	}
}
