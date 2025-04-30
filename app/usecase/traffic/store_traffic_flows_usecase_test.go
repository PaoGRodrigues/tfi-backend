package traffic_test

import (
	"fmt"
	"testing"

	host_domains "github.com/PaoGRodrigues/tfi-backend/app/domain/host"
	domains "github.com/PaoGRodrigues/tfi-backend/app/domain/traffic"
	trafficUseCase "github.com/PaoGRodrigues/tfi-backend/app/usecase/traffic"
	hostPortsMock "github.com/PaoGRodrigues/tfi-backend/mocks/ports/host"
	trafficPortsMock "github.com/PaoGRodrigues/tfi-backend/mocks/ports/traffic"

	"go.uber.org/mock/gomock"
)

func TestStoreTrafficSuccessfullyGettingTrafficFromSearcher(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	activeFlowToStore := []domains.TrafficFlow{
		{
			Client:   client,
			Server:   server,
			Bytes:    1000,
			Protocol: protocols,
		},
	}

	mockTrafficRepository := trafficPortsMock.NewMockTrafficReader(ctrl)
	mockTrafficRepository.EXPECT().GetTrafficFlows().Return(activeFlowToStore, nil)

	mockHostsStorage := hostPortsMock.NewMockHostDBRepository(ctrl)
	mockHostsStorage.EXPECT().GetHostByIp(server.IP).Return(host, nil)

	mockTrafficRepoStorage := trafficPortsMock.NewMockTrafficDBRepository(ctrl)
	mockTrafficRepoStorage.EXPECT().StoreTrafficFlows(activeFlowToStore).Return(nil)

	trafficStorage := trafficUseCase.NewStoreTrafficFlowsUseCase(mockTrafficRepository, mockTrafficRepoStorage, mockHostsStorage)
	err := trafficStorage.StoreTrafficFlows()

	if err != nil {
		t.Fail()
	}
}

func TestStoreTrafficSuccessfullyGettingTrafficFromEmptySearcherFirstly(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	activeFlowToStore := []domains.TrafficFlow{
		{
			Client:   client,
			Server:   server,
			Bytes:    1000,
			Protocol: protocols,
		},
	}

	mockTrafficRepository := trafficPortsMock.NewMockTrafficReader(ctrl)
	mockTrafficRepository.EXPECT().GetTrafficFlows().Return(activeFlowToStore, nil)

	mockHostsStorage := hostPortsMock.NewMockHostDBRepository(ctrl)
	mockHostsStorage.EXPECT().GetHostByIp(server.IP).Return(host, nil)
	mockTrafficRepoStorage := trafficPortsMock.NewMockTrafficDBRepository(ctrl)
	mockTrafficRepoStorage.EXPECT().StoreTrafficFlows(activeFlowToStore).Return(nil)

	trafficStorage := trafficUseCase.NewStoreTrafficFlowsUseCase(mockTrafficRepository, mockTrafficRepoStorage, mockHostsStorage)
	err := trafficStorage.StoreTrafficFlows()

	if err != nil {
		t.Fail()
	}
}

func TestStoreTrafficWithError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	activeFlowToStore := []domains.TrafficFlow{
		{
			Client:   client,
			Server:   server,
			Bytes:    1000,
			Protocol: protocols,
		},
	}

	mockTrafficRepository := trafficPortsMock.NewMockTrafficReader(ctrl)
	mockTrafficRepository.EXPECT().GetTrafficFlows().Return(activeFlowToStore, nil)
	mockHostsStorage := hostPortsMock.NewMockHostDBRepository(ctrl)
	mockHostsStorage.EXPECT().GetHostByIp(server.IP).Return(host, nil)
	mockTrafficRepoStorage := trafficPortsMock.NewMockTrafficDBRepository(ctrl)
	mockTrafficRepoStorage.EXPECT().StoreTrafficFlows(activeFlowToStore).Return(fmt.Errorf("Testing Error"))

	trafficStorage := trafficUseCase.NewStoreTrafficFlowsUseCase(mockTrafficRepository, mockTrafficRepoStorage, mockHostsStorage)
	err := trafficStorage.StoreTrafficFlows()

	if err == nil {
		t.Fail()
	}
}

func TestStoreTrafficWithGetTrafficReturningError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTrafficRepository := trafficPortsMock.NewMockTrafficReader(ctrl)
	mockTrafficRepository.EXPECT().GetTrafficFlows().Return([]domains.TrafficFlow{}, fmt.Errorf("Test error"))
	mockHostsStorage := hostPortsMock.NewMockHostDBRepository(ctrl)
	mockTrafficRepoStorage := trafficPortsMock.NewMockTrafficDBRepository(ctrl)

	trafficStorage := trafficUseCase.NewStoreTrafficFlowsUseCase(mockTrafficRepository, mockTrafficRepoStorage, mockHostsStorage)
	err := trafficStorage.StoreTrafficFlows()

	if err == nil {
		t.Fail()
	}
}

func TestStoreTrafficWithErrorInEnrichData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	activeFlowToStore := []domains.TrafficFlow{
		{
			Client:   client,
			Server:   server,
			Bytes:    1000,
			Protocol: protocols,
		},
	}

	mockTrafficRepository := trafficPortsMock.NewMockTrafficReader(ctrl)
	mockTrafficRepository.EXPECT().GetTrafficFlows().Return(activeFlowToStore, nil)
	mockHostsStorage := hostPortsMock.NewMockHostDBRepository(ctrl)
	mockHostsStorage.EXPECT().GetHostByIp(server.IP).Return(host_domains.Host{}, fmt.Errorf("Test error"))
	mockTrafficRepoStorage := trafficPortsMock.NewMockTrafficDBRepository(ctrl)

	trafficStorage := trafficUseCase.NewStoreTrafficFlowsUseCase(mockTrafficRepository, mockTrafficRepoStorage, mockHostsStorage)
	err := trafficStorage.StoreTrafficFlows()

	if err == nil {
		t.Fail()
	}
}

func TestStoreBroadcastServerSuccessfully(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	got := []domains.TrafficFlow{
		{
			Client:   client,
			Server:   broadcastserver,
			Bytes:    1000,
			Protocol: protocols,
		},
	}

	expected := []domains.TrafficFlow{
		{
			Client:   client,
			Server:   broadcastserverchanged,
			Bytes:    1000,
			Protocol: protocols,
		},
	}

	mockTrafficRepository := trafficPortsMock.NewMockTrafficReader(ctrl)
	mockTrafficRepository.EXPECT().GetTrafficFlows().Return(got, nil)
	mockHostsStorage := hostPortsMock.NewMockHostDBRepository(ctrl)
	mockHostsStorage.EXPECT().GetHostByIp(broadcastserver.IP).Return(publichost, nil)
	mockTrafficRepoStorage := trafficPortsMock.NewMockTrafficDBRepository(ctrl)
	mockTrafficRepoStorage.EXPECT().StoreTrafficFlows(expected).Return(nil)

	trafficStorage := trafficUseCase.NewStoreTrafficFlowsUseCase(mockTrafficRepository, mockTrafficRepoStorage, mockHostsStorage)
	err := trafficStorage.StoreTrafficFlows()

	if err != nil {
		t.Fail()
	}
}
