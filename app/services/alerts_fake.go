package services

import (
	alert "github.com/PaoGRodrigues/tfi-backend/app/domain/alert"
	flow "github.com/PaoGRodrigues/tfi-backend/app/traffic/domains"
)

func (d *FakeTool) GetAllAlerts(epoch_begin, epoch_end int) ([]alert.Alert, error) {

	alerts := []alert.Alert{
		{

			Name:     "test",
			Family:   "flow",
			Time:     "10/10/10 11:11:11",
			Severity: "Advertencia",
			AlertFlow: alert.AlertFlow{
				Client: flow.Client{
					Name: "192.168.4.14",
					IP:   "192.168.4.14",
					Port: 3550,
				},

				Server: flow.Server{
					IP:   "104.15.15.60",
					Port: 443,
					Name: "test2",
				},
			},
			AlertProtocol: flow.Protocol{

				L4:    "TCP",
				L7:    "TLS.Google",
				Label: "TCP:TLS.Google",
			},
		},
		{

			Name:     "test",
			Family:   "flow",
			Time:     "10/10/10 11:11:11",
			Severity: "Advertencia",
			AlertFlow: alert.AlertFlow{
				Client: flow.Client{
					IP:   "192.168.4.14",
					Name: "192.168.4.14",
					Port: 33566,
				},

				Server: flow.Server{
					IP:   "104.15.15.60",
					Port: 443,
					Name: "test2",
				},
			},
			AlertProtocol: flow.Protocol{
				L4:    "TCP",
				L7:    "TLS.Google",
				Label: "TCP:TLS.Google",
			},
		},
	}

	return alerts, nil
}
