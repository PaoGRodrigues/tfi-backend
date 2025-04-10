package usecase_test

import (
	"github.com/PaoGRodrigues/tfi-backend/app/alerts/domains"
	flow "github.com/PaoGRodrigues/tfi-backend/app/traffic/domains"
)

var expected = []domains.Alert{
	{
		Name:     "test1",
		Family:   "flow",
		Time:     "10/10/10 11:11:11",
		Severity: "Advertencia",
		AlertFlow: domains.AlertFlow{
			Client: flow.Client{
				Name: "test1",
				Port: 33566,
				IP:   "192.168.4.14",
			},

			Server: flow.Server{
				IP:   "104.15.15.60",
				Port: 443,
				Name: "test1",
			},
		},
		AlertProtocol: flow.Protocol{
			L4: "TCP",
			L7: "TLS.Google",
		},
		Category: "Cybersecurity",
	},
	{
		Name:     "test2",
		Family:   "flow",
		Time:     "10/10/10 11:11:11",
		Severity: "Error",
		AlertFlow: domains.AlertFlow{
			Client: flow.Client{
				Name: "test2",
				Port: 33566,
				IP:   "192.168.4.15",
			},

			Server: flow.Server{
				IP:   "104.15.15.70",
				Port: 443,
				Name: "test2",
			},
		},
		AlertProtocol: flow.Protocol{
			L4: "TCP",
			L7: "TLS.YouTube",
		},
		Category: "Cybersecurity",
	},
	{
		Name:     "test3",
		Family:   "flow",
		Time:     "10/10/10 11:11:11",
		Severity: "Error",
		AlertFlow: domains.AlertFlow{
			Client: flow.Client{
				Name: "test3",
				Port: 33566,
				IP:   "192.168.4.15",
			},

			Server: flow.Server{
				IP:   "104.15.15.70",
				Port: 443,
				Name: "test3",
			},
		},
		AlertProtocol: flow.Protocol{
			L4: "TCP",
			L7: "TLS.YouTube",
		},
		Category: "Flow threshold",
	},
}
