package traffic_test

import (
	"fmt"
	"testing"

	trafficDomains "github.com/PaoGRodrigues/tfi-backend/app/domain/traffic"
	"github.com/PaoGRodrigues/tfi-backend/app/usecase/traffic"
	trafficPortsMock "github.com/PaoGRodrigues/tfi-backend/mocks/ports/traffic"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetBytesPerDestReturnsBytesSuccessfully(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := []traffic.BytesPerDestination{
		{
			Bytes:       expectedFlowFromSearcher[0].Bytes,
			Destination: expectedFlowFromSearcher[0].Server.Name,
		},
	}

	mockFlowStorage := trafficPortsMock.NewMockTrafficDBRepository(ctrl)
	mockFlowStorage.EXPECT().GetServers().Return([]trafficDomains.Server{server1}, nil)
	mockFlowStorage.EXPECT().GetFlowByKey(server1.Key).Return(expectedFlowFromSearcher[0], nil)

	perDestinationUseCase := traffic.NewGetTrafficFlowsPerDestinationUseCase(mockFlowStorage)
	got, err := perDestinationUseCase.GetTrafficFlowsPerDestinations()

	if err != nil {
		t.Fail()
	}

	assert.ElementsMatch(t, expected, got)
}

func TestGetBytesPerDestReturnsBytesSuccessfullyWhenHaveMoreThanOneServer(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := []traffic.BytesPerDestination{
		{
			Bytes:       secondExpectedFlowFromSearcher[0].Bytes + secondExpectedFlowFromSearcher[1].Bytes,
			Destination: secondExpectedFlowFromSearcher[0].Server.Name,
		},
		{
			Bytes:       expectedPerCountrySearcher[1].Bytes,
			Destination: expectedPerCountrySearcher[1].Server.Name,
		},
	}

	mockFlowStorage := trafficPortsMock.NewMockTrafficDBRepository(ctrl)
	mockFlowStorage.EXPECT().GetServers().Return([]trafficDomains.Server{server1, server2, server3}, nil)
	mockFlowStorage.EXPECT().GetFlowByKey(server1.Key).Return(secondExpectedFlowFromSearcher[0], nil)
	mockFlowStorage.EXPECT().GetFlowByKey(server2.Key).Return(secondExpectedFlowFromSearcher[1], nil)
	mockFlowStorage.EXPECT().GetFlowByKey(server3.Key).Return(expectedPerCountrySearcher[1], nil)

	perDestinationUseCase := traffic.NewGetTrafficFlowsPerDestinationUseCase(mockFlowStorage)
	got, err := perDestinationUseCase.GetTrafficFlowsPerDestinations()

	if err != nil {
		t.Fail()
	}

	assert.ElementsMatch(t, expected, got)
}

func TestGetBytesPerDestReturnsErrorWhenThereIsAnErrorInGetServersList(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFlowStorage := trafficPortsMock.NewMockTrafficDBRepository(ctrl)
	mockFlowStorage.EXPECT().GetServers().Return([]trafficDomains.Server{}, fmt.Errorf("Test error"))

	perDestinationUseCase := traffic.NewGetTrafficFlowsPerDestinationUseCase(mockFlowStorage)
	_, err := perDestinationUseCase.GetTrafficFlowsPerDestinations()

	if err == nil {
		t.Fail()
	}
}

func TestGetBytesPerDestReturnsErrorWhenThereIsAnErrorInGetFlowByKey(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFlowStorage := trafficPortsMock.NewMockTrafficDBRepository(ctrl)
	mockFlowStorage.EXPECT().GetServers().Return([]trafficDomains.Server{server}, nil)
	mockFlowStorage.EXPECT().GetFlowByKey(server.Key).Return(trafficDomains.TrafficFlow{}, fmt.Errorf("Test error"))

	perDestinationUseCase := traffic.NewGetTrafficFlowsPerDestinationUseCase(mockFlowStorage)
	_, err := perDestinationUseCase.GetTrafficFlowsPerDestinations()

	if err == nil {
		t.Fail()
	}
}

func TestGetBytesPerDestReturnsTheSumOfBytesSuccessfully(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := []traffic.BytesPerDestination{
		{
			Bytes:       secondExpectedFlowFromSearcher[0].Bytes + secondExpectedFlowFromSearcher[1].Bytes,
			Destination: secondExpectedFlowFromSearcher[0].Server.Name,
		},
	}

	mockFlowStorage := trafficPortsMock.NewMockTrafficDBRepository(ctrl)
	mockFlowStorage.EXPECT().GetServers().Return([]trafficDomains.Server{server1, server2}, nil)
	mockFlowStorage.EXPECT().GetFlowByKey(server1.Key).Return(secondExpectedFlowFromSearcher[0], nil)
	mockFlowStorage.EXPECT().GetFlowByKey(server2.Key).Return(secondExpectedFlowFromSearcher[1], nil)

	perDestinationUseCase := traffic.NewGetTrafficFlowsPerDestinationUseCase(mockFlowStorage)
	got, err := perDestinationUseCase.GetTrafficFlowsPerDestinations()

	assert.ElementsMatch(t, expected, got)

	if err != nil {
		t.Fail()
	}
}

func TestGetBytesPerDestReturnsBytesSuccessfullyWhenHaveMoreThanOneServerAndAServerWithoutName(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := []traffic.BytesPerDestination{
		{
			Bytes:       secondExpectedFlowFromSearcher[0].Bytes + secondExpectedFlowFromSearcher[1].Bytes,
			Destination: secondExpectedFlowFromSearcher[0].Server.Name,
		},
		{
			Bytes:       expectedPerCountrySearcher[1].Bytes,
			Destination: expectedPerCountrySearcher[1].Server.Name,
		},
		{
			Bytes:       expectedFlowFromSearcherWithoutName[0].Bytes,
			Destination: expectedFlowFromSearcherWithoutName[0].Server.IP,
		},
	}

	mockFlowStorage := trafficPortsMock.NewMockTrafficDBRepository(ctrl)
	mockFlowStorage.EXPECT().GetServers().Return([]trafficDomains.Server{server1, server2, server3, noNameServer}, nil)
	mockFlowStorage.EXPECT().GetFlowByKey(server1.Key).Return(secondExpectedFlowFromSearcher[0], nil)
	mockFlowStorage.EXPECT().GetFlowByKey(server2.Key).Return(secondExpectedFlowFromSearcher[1], nil)
	mockFlowStorage.EXPECT().GetFlowByKey(server3.Key).Return(expectedPerCountrySearcher[1], nil)
	mockFlowStorage.EXPECT().GetFlowByKey(noNameServer.Key).Return(expectedFlowFromSearcherWithoutName[0], nil)

	perDestinationUseCase := traffic.NewGetTrafficFlowsPerDestinationUseCase(mockFlowStorage)
	got, err := perDestinationUseCase.GetTrafficFlowsPerDestinations()

	if err != nil {
		t.Fail()
	}

	assert.ElementsMatch(t, expected, got)
}
