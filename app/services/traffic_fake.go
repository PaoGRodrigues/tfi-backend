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
		domains.ActiveFlow{
			Client: client,
			Server: domains.Server{
				IP:                "123.1.5.1",
				IsBroadcastDomain: false,
				IsDHCP:            false,
				Port:              443,
				Name:              "lib.gen.rus",
			},
			Bytes:    345,
			Protocol: protocols,
		},
		domains.ActiveFlow{
			Client: client,
			Server: domains.Server{
				IP:                "123.123.123.123",
				IsBroadcastDomain: false,
				IsDHCP:            false,
				Port:              443,
				Name:              "lib.gen.rus",
			},
			Bytes:    10000,
			Protocol: protocols,
		},
		domains.ActiveFlow{
			Client: client,
			Server: domains.Server{
				IP:                "172.98.98.109",
				IsBroadcastDomain: false,
				IsDHCP:            false,
				Port:              443,
				Name:              "lib.gen.rus",
			},
			Bytes:    1000,
			Protocol: protocols,
		},
	}

	return activeFlowStruct, nil
}
