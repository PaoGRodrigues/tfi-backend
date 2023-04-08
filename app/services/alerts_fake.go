package services

import (
	"github.com/PaoGRodrigues/tfi-backend/app/alerts/domains"
	flow "github.com/PaoGRodrigues/tfi-backend/app/traffic/domains"
)

func (d *FakeTool) GetAllAlerts(epoch_begin, epoch_end int) ([]domains.Alert, error) {

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
					IP:   "192.168.4.4",
				},

				Server: flow.Server{
					IP:   "17.36.202.159",
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
					IP:   "192.168.4.4",
				},

				Server: flow.Server{
					IP:   "17.36.202.159",
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
