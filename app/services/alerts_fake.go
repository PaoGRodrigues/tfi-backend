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
			Severity: domains.Severity{Label: "2"},
			AlertFlow: domains.AlertFlow{
				Client: domains.AlertClient{
					CliPort: 33566,
					Value:   "192.168.4.14",
				},

				Server: domains.AlertServer{
					Value:   "104.15.15.60",
					SrvPort: 443,
					Name:    "test2",
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
			Severity: domains.Severity{Label: "2"},
			AlertFlow: domains.AlertFlow{
				Client: domains.AlertClient{
					Value:   "192.168.4.14",
					CliPort: 33566,
				},

				Server: domains.AlertServer{
					Value:   "104.15.15.60",
					SrvPort: 443,
					Name:    "test2",
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
