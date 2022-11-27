package services

import (
	"github.com/PaoGRodrigues/tfi-backend/app/traffic/domains"
)

func (trff *FakeTool) GetAllActiveTraffic() ([]domains.ActiveFlow, error) {
	client := domains.Client{
		Name: "test",
		Port: 55672,
		IP:   "192.168.4.9",
	}
	protocols := domains.Protocol{
		L4: "UDP.Youtube",
		L7: "TLS.GoogleServices",
	}
	activeFlowStruct := []domains.ActiveFlow{
		{
			Key:    "345",
			Client: client,
			Server: domains.Server{
				IP:                "123.1.5.1",
				IsBroadcastDomain: false,
				IsDHCP:            false,
				Port:              443,
				Name:              "lib.gen.rus",
			},
			FirstSeen: 1589741868,
			LastSeen:  1589741868,
			Bytes:     345,
			Protocol:  protocols,
		},
		{
			Key:    "346",
			Client: client,
			Server: domains.Server{
				IP:                "123.123.123.123",
				IsBroadcastDomain: false,
				IsDHCP:            false,
				Port:              443,
				Name:              "lib.gen.rus",
			},
			Bytes:     10000,
			FirstSeen: 1589741868,
			LastSeen:  1589741868,
			Protocol:  protocols,
		},
		{
			Key:    "347",
			Client: client,
			Server: domains.Server{
				IP:                "172.98.98.109",
				IsBroadcastDomain: false,
				IsDHCP:            false,
				Port:              443,
				Name:              "lib.gen.rus",
			},
			FirstSeen: 1589741868,
			LastSeen:  1589741868,
			Bytes:     1000,
			Protocol:  protocols,
		},
	}

	return activeFlowStruct, nil
}
