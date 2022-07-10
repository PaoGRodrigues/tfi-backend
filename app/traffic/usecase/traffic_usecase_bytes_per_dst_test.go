package usecase_test

import (
	"fmt"
	"reflect"
	"testing"

	hosts "github.com/PaoGRodrigues/tfi-backend/app/hosts/domains"
	"github.com/PaoGRodrigues/tfi-backend/app/traffic/domains"
	"github.com/PaoGRodrigues/tfi-backend/app/traffic/usecase"
	mock_host "github.com/PaoGRodrigues/tfi-backend/mocks/host"
	mocks "github.com/PaoGRodrigues/tfi-backend/mocks/traffic"
	"github.com/golang/mock/gomock"
)

func TestGetBytesPerDestReturnsBytesSuccessfully(t *testing.T) {

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

	hosts := []hosts.Host{
		hosts.Host{
			Name:        "sarasa",
			PrivateHost: false,
			IP:          "8.8.8.8",
			Country:     "USA",
			City:        "California",
		},
		hosts.Host{
			Name:        "sarasa2",
			PrivateHost: false,
			IP:          "198.8.8.8",
		},
	}

	expected := []domains.BytesPerDestination{
		domains.BytesPerDestination{
			Bytes:       expectedFlowFromSearcher[0].Bytes,
			Destination: expectedFlowFromSearcher[0].Server.Name,
			City:        hosts[0].City,
			Country:     hosts[0].Country,
		},
	}

	mockSearcher := mocks.NewMockTrafficUseCase(ctrl)
	mockSearcher.EXPECT().GetActiveFlows().Return(expectedFlowFromSearcher)
	mockHostsSearcher := mock_host.NewMockHostsFilter(ctrl)
	mockHostsSearcher.EXPECT().GetRemoteHosts().Return(hosts, nil)

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

	expected := []domains.BytesPerDestination{
		domains.BytesPerDestination{
			Bytes:       expectedFlowFromSearcher[0].Bytes,
			Destination: expectedFlowFromSearcher[0].Server.Name,
		},
	}

	hosts := []hosts.Host{
		hosts.Host{
			Name:        "sarasa",
			PrivateHost: false,
			IP:          "8.8.8.8",
		},
		hosts.Host{
			Name:        "sarasa2",
			PrivateHost: false,
			IP:          "198.8.8.8",
		},
	}

	mockSearcher := mocks.NewMockTrafficUseCase(ctrl)
	mockSearcher.EXPECT().GetActiveFlows().Return([]domains.ActiveFlow{})
	mockSearcher.EXPECT().GetAllActiveTraffic().Return(expectedFlowFromSearcher, nil)
	mockHostsSearcher := mock_host.NewMockHostsFilter(ctrl)
	mockHostsSearcher.EXPECT().GetRemoteHosts().Return(hosts, nil)

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
