package usecase_test

import (
	"reflect"
	"testing"

	"github.com/PaoGRodrigues/tfi-backend/app/traffic/domains"
	"github.com/PaoGRodrigues/tfi-backend/app/traffic/usecase"
	mocks "github.com/PaoGRodrigues/tfi-backend/mocks/traffic"
	"github.com/golang/mock/gomock"
)

func TestGetBytesPerDestReturnsBytes(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expectedFlowFromSearcher := []domains.ActiveFlow{
		domains.ActiveFlow{
			Client: domains.Client{
				Name: "Local",
				Port: 12345,
				IP:   "192.168.4.1",
			},
			Server: domains.Server{
				IP:                "8.8.8.8",
				IsBroadcastDomain: false,
				IsDHCP:            false,
				Port:              443,
				Name:              "google.com.ar",
			},
			Protocol: domains.Protocol{
				L4: "TCP",
				L7: "TLS.Google",
			},
			Bytes: 5566778,
		},
	}

	expected := []usecase.BytesPerDestination{
		usecase.BytesPerDestination{
			Bytes:       expectedFlowFromSearcher[0].Bytes,
			Destination: expectedFlowFromSearcher[0].Server.Name,
		},
	}

	mockSearcher := mocks.NewMockTrafficUseCase(ctrl)
	mockSearcher.EXPECT().GetActiveFlows().Return(expectedFlowFromSearcher)

	parser := usecase.NewBytesDestinationParser(mockSearcher)
	got, err := parser.GetBytesPerDestination()

	if err != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("expected:\n%+v\ngot:\n%+v", expected, got)
	}
}
