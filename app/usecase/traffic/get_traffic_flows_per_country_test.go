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

func TestGetBytesPerCountryReturnBytesSuccessfully(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := []traffic.BytesPerCountry{
		{
			Country: "US",
			Bytes:   expectedPerCountrySearcher[0].Bytes + secondExpectedFlowFromSearcher[1].Bytes,
		},
		{
			Country: "RU",
			Bytes:   expectedPerCountrySearcher[1].Bytes,
		},
	}

	mockFlowStorage := trafficPortsMock.NewMockTrafficDBRepository(ctrl)
	mockFlowStorage.EXPECT().GetServers().Return([]trafficDomains.Server{server1, server2, server3}, nil)
	mockFlowStorage.EXPECT().GetFlowByKey(server1.Key).Return(expectedPerCountrySearcher[0], nil)
	mockFlowStorage.EXPECT().GetFlowByKey(server2.Key).Return(secondExpectedFlowFromSearcher[1], nil)
	mockFlowStorage.EXPECT().GetFlowByKey(server3.Key).Return(expectedPerCountrySearcher[1], nil)

	parser := traffic.NewGetTrafficFlowsPerCountryUseCase(mockFlowStorage)
	got, err := parser.GetBytesPerCountry()

	if err != nil {
		t.Fail()
	}

	assert.ElementsMatch(t, expected, got)
}

func TestGetBytesPerCountryReturnsErrorWhenThereIsAnErrorInGetServersList(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFlowStorage := trafficPortsMock.NewMockTrafficDBRepository(ctrl)
	mockFlowStorage.EXPECT().GetServers().Return([]trafficDomains.Server{}, fmt.Errorf("Test error"))

	parser := traffic.NewGetTrafficFlowsPerCountryUseCase(mockFlowStorage)
	_, err := parser.GetBytesPerCountry()

	if err == nil {
		t.Fail()
	}
}

func TestGetBytesPerCountryReturnsErrorWhenThereIsAnErrorInGetFlowByKey(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFlowStorage := trafficPortsMock.NewMockTrafficDBRepository(ctrl)
	mockFlowStorage.EXPECT().GetServers().Return([]trafficDomains.Server{server}, nil)
	mockFlowStorage.EXPECT().GetFlowByKey(server.Key).Return(trafficDomains.TrafficFlow{}, fmt.Errorf("Test error"))

	parser := traffic.NewGetTrafficFlowsPerCountryUseCase(mockFlowStorage)
	_, err := parser.GetBytesPerCountry()

	if err == nil {
		t.Fail()
	}
}

func TestGetBytesPerCountryReturnsBytesSuccessfullyWhenHaveMoreThanOneServerAndAServerWithoutName(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := []traffic.BytesPerCountry{
		{
			Country: "US",
			Bytes: secondExpectedFlowFromSearcher[0].Bytes + secondExpectedFlowFromSearcher[1].Bytes +
				expectedFlowFromSearcherWithoutName[0].Bytes,
		},
		{
			Country: "RU",
			Bytes:   expectedPerCountrySearcher[1].Bytes,
		},
	}

	mockFlowStorage := trafficPortsMock.NewMockTrafficDBRepository(ctrl)
	mockFlowStorage.EXPECT().GetServers().Return([]trafficDomains.Server{server1, server2, server3, noNameServer}, nil)
	mockFlowStorage.EXPECT().GetFlowByKey(server1.Key).Return(secondExpectedFlowFromSearcher[0], nil)
	mockFlowStorage.EXPECT().GetFlowByKey(server2.Key).Return(secondExpectedFlowFromSearcher[1], nil)
	mockFlowStorage.EXPECT().GetFlowByKey(server3.Key).Return(expectedPerCountrySearcher[1], nil)
	mockFlowStorage.EXPECT().GetFlowByKey(noNameServer.Key).Return(expectedFlowFromSearcherWithoutName[0], nil)

	parser := traffic.NewGetTrafficFlowsPerCountryUseCase(mockFlowStorage)
	got, err := parser.GetBytesPerCountry()

	if err != nil {
		t.Fail()
	}

	assert.ElementsMatch(t, expected, got)
}
