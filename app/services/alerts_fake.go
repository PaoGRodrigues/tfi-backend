package services

import (
	"github.com/PaoGRodrigues/tfi-backend/app/alerts/domains"
	flow "github.com/PaoGRodrigues/tfi-backend/app/traffic/domains"
)

func (d *FakeTool) GetAllAlerts(epoch_begin, epoch_end int, host string) ([]domains.Alert, error) {

	alerts := []domains.Alert{
		domains.Alert{

			Name:     "test",
			Family:   "flow",
			Time:     struct{ Label string }{"10/10/10 11:11:11"},
			Score:    "10",
			Severity: domains.Severity{Label: "2"},
			AlertFlow: domains.AlertFlow{
				Client: flow.Client{
					Name: "test1",
					Port: 33566,
					IP:   "192.168.4.14",
				},

				Server: flow.Server{
					IP:   "104.15.15.60",
					Port: 443,
					Name: "test2",
				},
			},
			AlertProtocol: domains.AlertProtocol{
				Protocol: flow.Protocol{
					L4: "TCP",
					L7: "TLS.Google",
				},
			},
		},
		domains.Alert{

			Name:     "test",
			Family:   "flow",
			Time:     struct{ Label string }{"10/10/10 11:11:11"},
			Score:    "10",
			Severity: domains.Severity{Label: "2"},
			AlertFlow: domains.AlertFlow{
				Client: flow.Client{
					Name: "test1",
					Port: 33566,
					IP:   "192.168.4.14",
				},

				Server: flow.Server{
					IP:   "104.15.15.60",
					Port: 443,
					Name: "test2",
				},
			},
			AlertProtocol: domains.AlertProtocol{
				Protocol: flow.Protocol{
					L4: "TCP",
					L7: "TLS.Google",
				},
			},
		},
	}

	return alerts, nil
}
