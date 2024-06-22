package usecase_test

import (
	"fmt"
	"testing"

	"github.com/PaoGRodrigues/tfi-backend/app/traffic/domains"
	"github.com/PaoGRodrigues/tfi-backend/app/traffic/usecase"
	mocks "github.com/PaoGRodrigues/tfi-backend/mocks/traffic"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetBytesPerDestReturnsBytesSuccessfully(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := []domains.BytesPerDestination{
		{
			Bytes:       expectedFlowFromSearcher[0].Bytes,
			Destination: expectedFlowFromSearcher[0].Server.Name,
			Country:     expectedHosts[0].Country,
		},
	}

	mockFlowStorage := mocks.NewMockTrafficRepository(ctrl)
	mockFlowStorage.EXPECT().GetServers().Return([]domains.Server{server1}, nil)
	mockFlowStorage.EXPECT().GetFlowByKey(server1.Key).Return(expectedFlowFromSearcher[0], nil)

	parser := usecase.NewBytesParser(mockFlowStorage)
	got, err := parser.GetBytesPerDestination()

	if err != nil {
		t.Fail()
	}

	assert.ElementsMatch(t, expected, got)
}

func TestGetBytesPerDestReturnsBytesSuccessfullyWhenHaveMoreThanOneServer(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := []domains.BytesPerDestination{
		{
			Bytes:       secondExpectedFlowFromSearcher[0].Bytes + secondExpectedFlowFromSearcher[1].Bytes,
			Destination: secondExpectedFlowFromSearcher[0].Server.Name,
			Country:     expectedHosts[0].Country,
		},
		{
			Bytes:       expectedPerCountrySearcher[1].Bytes,
			Destination: expectedPerCountrySearcher[1].Server.Name,
			Country:     expectedHosts[2].Country,
		},
	}

	mockFlowStorage := mocks.NewMockTrafficRepository(ctrl)
	mockFlowStorage.EXPECT().GetServers().Return([]domains.Server{server1, server2, server3}, nil)
	mockFlowStorage.EXPECT().GetFlowByKey(server1.Key).Return(secondExpectedFlowFromSearcher[0], nil)
	mockFlowStorage.EXPECT().GetFlowByKey(server2.Key).Return(secondExpectedFlowFromSearcher[1], nil)
	mockFlowStorage.EXPECT().GetFlowByKey(server3.Key).Return(expectedPerCountrySearcher[1], nil)

	parser := usecase.NewBytesParser(mockFlowStorage)
	got, err := parser.GetBytesPerDestination()

	if err != nil {
		t.Fail()
	}

	assert.ElementsMatch(t, expected, got)
}

func TestGetBytesPerDestReturnsErrorWhenThereIsAnErrorInGetServersList(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFlowStorage := mocks.NewMockTrafficRepository(ctrl)
	mockFlowStorage.EXPECT().GetServers().Return([]domains.Server{}, fmt.Errorf("Test error"))

	parser := usecase.NewBytesParser(mockFlowStorage)
	_, err := parser.GetBytesPerDestination()

	if err == nil {
		t.Fail()
	}
}

func TestGetBytesPerDestReturnsErrorWhenThereIsAnErrorInGetFlowByKey(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFlowStorage := mocks.NewMockTrafficRepository(ctrl)
	mockFlowStorage.EXPECT().GetServers().Return([]domains.Server{server}, nil)
	mockFlowStorage.EXPECT().GetFlowByKey(server.Key).Return(domains.ActiveFlow{}, fmt.Errorf("Test error"))

	parser := usecase.NewBytesParser(mockFlowStorage)
	_, err := parser.GetBytesPerDestination()

	if err == nil {
		t.Fail()
	}
}

func TestGetBytesPerDestReturnsTheSumOfBytesSuccessfully(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := []domains.BytesPerDestination{
		{
			Bytes:       secondExpectedFlowFromSearcher[0].Bytes + secondExpectedFlowFromSearcher[1].Bytes,
			Destination: secondExpectedFlowFromSearcher[0].Server.Name,
			Country:     secondExpectedFlowFromSearcher[0].Server.Country,
		},
	}

	mockFlowStorage := mocks.NewMockTrafficRepository(ctrl)
	mockFlowStorage.EXPECT().GetServers().Return([]domains.Server{server1, server2}, nil)
	mockFlowStorage.EXPECT().GetFlowByKey(server1.Key).Return(secondExpectedFlowFromSearcher[0], nil)
	mockFlowStorage.EXPECT().GetFlowByKey(server2.Key).Return(secondExpectedFlowFromSearcher[1], nil)

	parser := usecase.NewBytesParser(mockFlowStorage)
	got, err := parser.GetBytesPerDestination()

	assert.ElementsMatch(t, expected, got)

	if err != nil {
		t.Fail()
	}
}

func TestGetBytesPerCountryReturnBytesSuccessfully(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := []domains.BytesPerCountry{
		{
			Bytes:   expectedPerCountrySearcher[0].Bytes + secondExpectedFlowFromSearcher[1].Bytes,
			Country: "US",
		},
		{
			Bytes:   expectedPerCountrySearcher[1].Bytes,
			Country: "RU",
		},
	}

	mockFlowStorage := mocks.NewMockTrafficRepository(ctrl)
	mockFlowStorage.EXPECT().GetServers().Return([]domains.Server{server1, server2, server3}, nil)
	mockFlowStorage.EXPECT().GetFlowByKey(server1.Key).Return(expectedPerCountrySearcher[0], nil)
	mockFlowStorage.EXPECT().GetFlowByKey(server2.Key).Return(secondExpectedFlowFromSearcher[1], nil)
	mockFlowStorage.EXPECT().GetFlowByKey(server3.Key).Return(expectedPerCountrySearcher[1], nil)

	parser := usecase.NewBytesParser(mockFlowStorage)
	got, err := parser.GetBytesPerCountry()

	if err != nil {
		t.Fail()
	}

	assert.ElementsMatch(t, expected, got)
}

func TestGetBytesPerCountryReturnsErrorWhenThereIsAnErrorInGetServersList(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFlowStorage := mocks.NewMockTrafficRepository(ctrl)
	mockFlowStorage.EXPECT().GetServers().Return([]domains.Server{}, fmt.Errorf("Test error"))

	parser := usecase.NewBytesParser(mockFlowStorage)
	_, err := parser.GetBytesPerCountry()

	if err == nil {
		t.Fail()
	}
}

func TestGetBytesPerCountryReturnsErrorWhenThereIsAnErrorInGetFlowByKey(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFlowStorage := mocks.NewMockTrafficRepository(ctrl)
	mockFlowStorage.EXPECT().GetServers().Return([]domains.Server{server}, nil)
	mockFlowStorage.EXPECT().GetFlowByKey(server.Key).Return(domains.ActiveFlow{}, fmt.Errorf("Test error"))

	parser := usecase.NewBytesParser(mockFlowStorage)
	_, err := parser.GetBytesPerCountry()

	if err == nil {
		t.Fail()
	}
}

func TestGetBytesPerDestReturnsBytesSuccessfullyWhenHaveMoreThanOneServerAndAServerWithoutName(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := []domains.BytesPerDestination{
		{
			Bytes:       secondExpectedFlowFromSearcher[0].Bytes + secondExpectedFlowFromSearcher[1].Bytes,
			Destination: secondExpectedFlowFromSearcher[0].Server.Name,
			Country:     expectedHosts[0].Country,
		},
		{
			Bytes:       expectedPerCountrySearcher[1].Bytes,
			Destination: expectedPerCountrySearcher[1].Server.Name,
			Country:     expectedHosts[2].Country,
		},
		{
			Bytes:       expectedFlowFromSearcherWithoutName[0].Bytes,
			Destination: expectedFlowFromSearcherWithoutName[0].Server.IP,
			Country:     expectedHosts[3].Country,
		},
	}

	mockFlowStorage := mocks.NewMockTrafficRepository(ctrl)
	mockFlowStorage.EXPECT().GetServers().Return([]domains.Server{server1, server2, server3, noNameServer}, nil)
	mockFlowStorage.EXPECT().GetFlowByKey(server1.Key).Return(secondExpectedFlowFromSearcher[0], nil)
	mockFlowStorage.EXPECT().GetFlowByKey(server2.Key).Return(secondExpectedFlowFromSearcher[1], nil)
	mockFlowStorage.EXPECT().GetFlowByKey(server3.Key).Return(expectedPerCountrySearcher[1], nil)
	mockFlowStorage.EXPECT().GetFlowByKey(noNameServer.Key).Return(expectedFlowFromSearcherWithoutName[0], nil)

	parser := usecase.NewBytesParser(mockFlowStorage)
	got, err := parser.GetBytesPerDestination()

	if err != nil {
		t.Fail()
	}

	assert.ElementsMatch(t, expected, got)
}

func TestGetBytesPerCountryReturnsBytesSuccessfullyWhenHaveMoreThanOneServerAndAServerWithoutName(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := []domains.BytesPerCountry{
		{
			Bytes: secondExpectedFlowFromSearcher[0].Bytes + secondExpectedFlowFromSearcher[1].Bytes +
				expectedFlowFromSearcherWithoutName[0].Bytes,
			Country: "US",
		},
		{
			Bytes:   expectedPerCountrySearcher[1].Bytes,
			Country: "RU",
		},
	}

	mockFlowStorage := mocks.NewMockTrafficRepository(ctrl)
	mockFlowStorage.EXPECT().GetServers().Return([]domains.Server{server1, server2, server3, noNameServer}, nil)
	mockFlowStorage.EXPECT().GetFlowByKey(server1.Key).Return(secondExpectedFlowFromSearcher[0], nil)
	mockFlowStorage.EXPECT().GetFlowByKey(server2.Key).Return(secondExpectedFlowFromSearcher[1], nil)
	mockFlowStorage.EXPECT().GetFlowByKey(server3.Key).Return(expectedPerCountrySearcher[1], nil)
	mockFlowStorage.EXPECT().GetFlowByKey(noNameServer.Key).Return(expectedFlowFromSearcherWithoutName[0], nil)

	parser := usecase.NewBytesParser(mockFlowStorage)
	got, err := parser.GetBytesPerCountry()

	if err != nil {
		t.Fail()
	}

	assert.ElementsMatch(t, expected, got)
}
