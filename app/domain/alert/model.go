package alert

import (
	flow "github.com/PaoGRodrigues/tfi-backend/app/traffic/domains"
)

// *********** Entities
// Alerts
type Alert struct {
	Name          string
	Family        string
	Category      string
	Time          string
	Severity      string
	AlertFlow     AlertFlow
	AlertProtocol flow.Protocol `json:",omitempty"`
}

type AlertFlow struct {
	Client flow.Client
	Server flow.Server
}
