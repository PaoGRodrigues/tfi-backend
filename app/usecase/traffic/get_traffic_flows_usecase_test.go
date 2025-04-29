package traffic_test

import (
	"fmt"
	"reflect"
	"testing"

	domains "github.com/PaoGRodrigues/tfi-backend/app/domain/traffic"
	usecase "github.com/PaoGRodrigues/tfi-backend/app/usecase/traffic"
	trafficPortsMock "github.com/PaoGRodrigues/tfi-backend/mocks/ports/traffic"

	"go.uber.org/mock/gomock"
)

func TestGetAllTrafficReturnAListOfTrafficJsons(t *testing.T) {

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

	expected := []domains.TrafficFlow{
		domains.TrafficFlow{
			Client:   client,
			Server:   server,
			Bytes:    1000,
			Protocol: protocols,
		},
	}

	mockRepository := trafficPortsMock.NewMockTrafficReader(ctrl)
	mockRepository.EXPECT().GetTrafficFlows().Return(expected, nil)

	usecase := usecase.NewTrafficFlowsUseCase(mockRepository)
	got, err := usecase.GetTrafficFlows()

	if err != nil {
		t.Fail()
	}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("expected:\n%+v\ngot:\n%+v", expected, got)
	}
}

func TestGetAllTrafficReturnAnError(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := trafficPortsMock.NewMockTrafficReader(ctrl)
	mockRepository.EXPECT().GetTrafficFlows().Return(nil, fmt.Errorf("Testing Error"))

	usecase := usecase.NewTrafficFlowsUseCase(mockRepository)
	_, err := usecase.GetTrafficFlows()

	if err == nil {
		t.Errorf("We expected an error, but didn't get one.")
	}
}
