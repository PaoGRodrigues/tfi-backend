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
	activeFlowStruct := []domains.ActiveFlow{
		domains.ActiveFlow{
			Client:   client,
			Server:   server,
			Bytes:    345,
			Protocol: protocols,
		},
		domains.ActiveFlow{
			Client:   client,
			Server:   server,
			Bytes:    10000,
			Protocol: protocols,
		},
		domains.ActiveFlow{
			Client:   client,
			Server:   server,
			Bytes:    1000,
			Protocol: protocols,
		},
	}

	return activeFlowStruct, nil
}
