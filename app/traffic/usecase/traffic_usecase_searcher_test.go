package usecase_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/PaoGRodrigues/tfi-backend/app/traffic/domains"
	"github.com/PaoGRodrigues/tfi-backend/app/traffic/usecase"
	mocks "github.com/PaoGRodrigues/tfi-backend/mocks/traffic"
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

	expected := []domains.ActiveFlow{
		domains.ActiveFlow{
			Client:   client,
			Server:   server,
			Bytes:    1000,
			Protocol: protocols,
		},
	}

	mockTrafficService := mocks.NewMockTrafficService(ctrl)
	mockTrafficService.EXPECT().GetAllActiveTraffic().Return(expected, nil)

	trafficSearcher := usecase.NewTrafficSearcher(mockTrafficService)
	got, err := trafficSearcher.GetAllActiveTraffic()

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

	mockTrafficService := mocks.NewMockTrafficService(ctrl)
	mockTrafficService.EXPECT().GetAllActiveTraffic().Return(nil, fmt.Errorf("Testing Error"))

	trafficSearcher := usecase.NewTrafficSearcher(mockTrafficService)
	_, err := trafficSearcher.GetAllActiveTraffic()

	if err == nil {
		t.Errorf("We expected an error, but didn't get one.")
	}
}
