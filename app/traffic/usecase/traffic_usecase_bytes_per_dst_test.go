package usecase_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/PaoGRodrigues/tfi-backend/app/traffic/domains"
	"github.com/PaoGRodrigues/tfi-backend/app/traffic/usecase"
	mock_host "github.com/PaoGRodrigues/tfi-backend/mocks/hosts"
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
			City:        expectedHosts[0].City,
			Country:     expectedHosts[0].Country,
		},
	}

	mockSearcher := mocks.NewMockTrafficUseCase(ctrl)
	mockSearcher.EXPECT().GetActiveFlows().Return(expectedFlowFromSearcher)
	mockHostsSearcher := mock_host.NewMockHostsFilter(ctrl)
	mockHostsSearcher.EXPECT().GetRemoteHosts().Return(expectedHosts, nil)

	parser := usecase.NewBytesDestinationParser(mockSearcher, mockHostsSearcher)
	got, err := parser.GetBytesPerDestination()

	if err != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("expected:\n%+v\ngot:\n%+v", expected, got)
	}
}

func TestGetBytesPerDestSearcherActiveFlowsIsEmptyReturnsBytesSuccessfully(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := []domains.BytesPerDestination{
		domains.BytesPerDestination{
			Bytes:       expectedFlowFromSearcher[0].Bytes,
			Destination: expectedFlowFromSearcher[0].Server.Name,
			City:        expectedHosts[0].City,
			Country:     expectedHosts[0].Country,
		},
	}

	mockSearcher := mocks.NewMockTrafficUseCase(ctrl)
	mockSearcher.EXPECT().GetActiveFlows().Return([]domains.ActiveFlow{})
	mockSearcher.EXPECT().GetAllActiveTraffic().Return(expectedFlowFromSearcher, nil)
	mockHostsSearcher := mock_host.NewMockHostsFilter(ctrl)
	mockHostsSearcher.EXPECT().GetRemoteHosts().Return(expectedHosts, nil)

	parser := usecase.NewBytesDestinationParser(mockSearcher, mockHostsSearcher)
	got, err := parser.GetBytesPerDestination()

	if err != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("expected:\n%+v\ngot:\n%+v", expected, got)
	}
}

func TestGetBytesPerDestSearcherActiveFlowsIsEmptyReturnsBytesFailedAndFailedTheEntireFunction(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSearcher := mocks.NewMockTrafficUseCase(ctrl)
	mockSearcher.EXPECT().GetActiveFlows().Return(nil)
	mockSearcher.EXPECT().GetAllActiveTraffic().Return(nil, fmt.Errorf("Test Error"))
	mockHostsSearcher := mock_host.NewMockHostsFilter(ctrl)

	parser := usecase.NewBytesDestinationParser(mockSearcher, mockHostsSearcher)
	_, err := parser.GetBytesPerDestination()

	if err == nil {
		t.Fail()
	}
}

func TestGetBytesPerDestSearcherHostFilterGetRemoteReturnsErrorAndThenReturnsBytesSuccessfully(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSearcher := mocks.NewMockTrafficUseCase(ctrl)
	mockSearcher.EXPECT().GetActiveFlows().Return([]domains.ActiveFlow{})
	mockSearcher.EXPECT().GetAllActiveTraffic().Return(expectedFlowFromSearcher, nil)
	mockHostsSearcher := mock_host.NewMockHostsFilter(ctrl)
	mockHostsSearcher.EXPECT().GetRemoteHosts().Return(nil, fmt.Errorf("Test error"))

	parser := usecase.NewBytesDestinationParser(mockSearcher, mockHostsSearcher)
	_, err := parser.GetBytesPerDestination()

	if err == nil {
		t.Fail()
	}
}

func TestGetBytesPerDestReturnsBytesSuccessfullyWhenCompareByIP(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := []domains.BytesPerDestination{
		domains.BytesPerDestination{
			Bytes:       expectedFlowFromSearcherWithoutName[0].Bytes,
			Destination: expectedFlowFromSearcherWithoutName[0].Server.IP,
			City:        expectedHosts[0].City,
			Country:     expectedHosts[0].Country,
		},
	}

	mockSearcher := mocks.NewMockTrafficUseCase(ctrl)
	mockSearcher.EXPECT().GetActiveFlows().Return(expectedFlowFromSearcherWithoutName)
	mockHostsSearcher := mock_host.NewMockHostsFilter(ctrl)
	mockHostsSearcher.EXPECT().GetRemoteHosts().Return(expectedHosts, nil)

	parser := usecase.NewBytesDestinationParser(mockSearcher, mockHostsSearcher)
	got, err := parser.GetBytesPerDestination()

	if err != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("expected:\n%+v\ngot:\n%+v", expected, got)
	}
}

func TestGetBytesPerDestReturnsTheSumOfBytesSuccessfully(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := []domains.BytesPerDestination{
		{
			Bytes:       secondExpectedFlowFromSearcher[0].Bytes + secondExpectedFlowFromSearcher[1].Bytes,
			Destination: secondExpectedFlowFromSearcher[0].Server.Name,
			City:        expectedHosts[0].City,
			Country:     expectedHosts[0].Country,
		},
	}

	mockSearcher := mocks.NewMockTrafficUseCase(ctrl)
	mockSearcher.EXPECT().GetActiveFlows().Return(secondExpectedFlowFromSearcher)
	mockHostsSearcher := mock_host.NewMockHostsFilter(ctrl)
	mockHostsSearcher.EXPECT().GetRemoteHosts().Return(expectedHosts, nil)

	parser := usecase.NewBytesDestinationParser(mockSearcher, mockHostsSearcher)
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

	mockSearcher := mocks.NewMockTrafficUseCase(ctrl)
	mockSearcher.EXPECT().GetActiveFlows().Return(expectedPerCountrySearcher)
	mockHostsSearcher := mock_host.NewMockHostsFilter(ctrl)
	mockHostsSearcher.EXPECT().GetRemoteHosts().Return(expectedHostsPerCountry, nil)

	parser := usecase.NewBytesDestinationParser(mockSearcher, mockHostsSearcher)
	got, err := parser.GetBytesPerCountry()

	if err != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("expected:\n%+v\ngot:\n%+v", expected, got)
	}
}

func TestGetBytesPerCountryReturnErrorWhenGetActiveTrafficReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSearcher := mocks.NewMockTrafficUseCase(ctrl)
	mockSearcher.EXPECT().GetActiveFlows().Return(nil)
	mockSearcher.EXPECT().GetAllActiveTraffic().Return(nil, fmt.Errorf("Test Error"))
	mockHostsSearcher := mock_host.NewMockHostsFilter(ctrl)

	parser := usecase.NewBytesDestinationParser(mockSearcher, mockHostsSearcher)
	_, err := parser.GetBytesPerCountry()

	if err == nil {
		t.Fail()
	}
}

func TestGetBytesPerCountryReturnBytesSuccessfullyWhenGetActiveFlowsReturn0(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := []domains.BytesPerCountry{
		{
			Bytes:   expectedPerCountrySearcher[0].Bytes + expectedPerCountrySearcher[1].Bytes,
			Country: "US",
		},
	}

	mockSearcher := mocks.NewMockTrafficUseCase(ctrl)
	mockSearcher.EXPECT().GetActiveFlows().Return(nil)
	mockSearcher.EXPECT().GetAllActiveTraffic().Return(expectedPerCountrySearcher, nil)
	mockHostsSearcher := mock_host.NewMockHostsFilter(ctrl)
	mockHostsSearcher.EXPECT().GetRemoteHosts().Return(expectedHostsPerCountry, nil)

	parser := usecase.NewBytesDestinationParser(mockSearcher, mockHostsSearcher)
	got, err := parser.GetBytesPerCountry()

	if err != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("expected:\n%+v\ngot:\n%+v", expected, got)
	}
}
