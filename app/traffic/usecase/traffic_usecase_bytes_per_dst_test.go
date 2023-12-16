package usecase_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/PaoGRodrigues/tfi-backend/app/traffic/domains"
	"github.com/PaoGRodrigues/tfi-backend/app/traffic/usecase"
	mocks "github.com/PaoGRodrigues/tfi-backend/mocks/traffic"
	"github.com/golang/mock/gomock"
)

func TestGetBytesPerDestReturnsBytesSuccessfully(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := []domains.BytesPerDestination{
		domains.BytesPerDestination{
			Bytes:       expectedFlowFromSearcher[0].Bytes,
			Destination: expectedFlowFromSearcher[0].Server.Name,
			Country:     expectedHosts[0].Country,
		},
	}

	mockFlowStorage := mocks.NewMockActiveFlowsStorage(ctrl)
	mockFlowStorage.EXPECT().GetServersList().Return([]domains.Server{server}, nil)
	mockFlowStorage.EXPECT().GetFlowByKey(server.Key).Return(expectedFlowFromSearcher[0], nil)

	parser := usecase.NewBytesDestinationParser(mockFlowStorage)
	got, err := parser.GetBytesPerDestination()

	if err != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("expected:\n%+v\ngot:\n%+v", expected, got)
	}
}

func TestGetBytesPerDestReturnsBytesSuccessfullyWhenCompareByIP(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := []domains.BytesPerDestination{
		domains.BytesPerDestination{
			Bytes:       expectedFlowFromSearcherWithoutName[0].Bytes,
			Destination: expectedFlowFromSearcherWithoutName[0].Server.IP,
			Country:     expectedHosts[0].Country,
		},
	}

	mockFlowStorage := mocks.NewMockActiveFlowsStorage(ctrl)
	mockFlowStorage.EXPECT().GetServersList().Return([]domains.Server{server}, nil)
	mockFlowStorage.EXPECT().GetFlowByKey(server.Key).Return(expectedFlowFromSearcherWithoutName[0], nil)

	parser := usecase.NewBytesDestinationParser(mockFlowStorage)
	got, err := parser.GetBytesPerDestination()

	if err != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("expected:\n%+v\ngot:\n%+v", expected, got)
	}
}

func TestGetBytesPerDestReturnsErrorWhenThereIsAnErrorInGetServersList(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFlowStorage := mocks.NewMockActiveFlowsStorage(ctrl)
	mockFlowStorage.EXPECT().GetServersList().Return([]domains.Server{}, fmt.Errorf("Test error"))

	parser := usecase.NewBytesDestinationParser(mockFlowStorage)
	_, err := parser.GetBytesPerDestination()

	if err == nil {
		t.Fail()
	}
}

func TestGetBytesPerDestReturnsErrorWhenThereIsAnErrorInGetFlowByKey(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFlowStorage := mocks.NewMockActiveFlowsStorage(ctrl)
	mockFlowStorage.EXPECT().GetServersList().Return([]domains.Server{server}, nil)
	mockFlowStorage.EXPECT().GetFlowByKey(server.Key).Return(domains.ActiveFlow{}, fmt.Errorf("Test error"))

	parser := usecase.NewBytesDestinationParser(mockFlowStorage)
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

	mockFlowStorage := mocks.NewMockActiveFlowsStorage(ctrl)
	mockFlowStorage.EXPECT().GetServersList().Return([]domains.Server{server1, server2}, nil)
	mockFlowStorage.EXPECT().GetFlowByKey(server1.Key).Return(secondExpectedFlowFromSearcher[0], nil)
	mockFlowStorage.EXPECT().GetFlowByKey(server2.Key).Return(secondExpectedFlowFromSearcher[1], nil)

	parser := usecase.NewBytesDestinationParser(mockFlowStorage)
	got, err := parser.GetBytesPerDestination()

	if err != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("expected:\n%+v\ngot:\n%+v", expected, got)
	}
}

func TestGetBytesPerCountryReturnBytesSuccessfully(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := []domains.BytesPerCountry{
		{
			Bytes:   expectedPerCountrySearcher[0].Bytes + expectedPerCountrySearcher[1].Bytes,
			Country: "US",
		},
	}

	mockFlowStorage := mocks.NewMockActiveFlowsStorage(ctrl)
	mockFlowStorage.EXPECT().GetServersList().Return([]domains.Server{server1, server3}, nil)
	mockFlowStorage.EXPECT().GetFlowByKey(server1.Key).Return(expectedPerCountrySearcher[0], nil)
	mockFlowStorage.EXPECT().GetFlowByKey(server3.Key).Return(expectedPerCountrySearcher[1], nil)

	parser := usecase.NewBytesDestinationParser(mockFlowStorage)
	got, err := parser.GetBytesPerCountry()

	if err != nil {
		t.Fail()
	}

	if got[0].Bytes != expected[0].Bytes {
		t.Errorf("expected:\n%+v\ngot:\n%+v", expected, got)
	}
}
