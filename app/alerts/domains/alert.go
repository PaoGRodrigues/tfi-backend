package domains

import (
	flow "github.com/PaoGRodrigues/tfi-backend/app/traffic/domains"
)

// *********** Entities
// Alerts
type Alert struct {
	Row_id string `json:"-"`
	Name   string
	Family string
	Time   struct {
		Label string
	} `json:"tstamp"`
	Score         string
	Severity      Severity
	AlertFlow     AlertFlow
	AlertProtocol AlertProtocol
}

type Severity struct {
	Label string
}

type AlertFlow struct {
	Client flow.Client
	Server flow.Server
}

type AlertProtocol struct {
	Protocol flow.Protocol
}

// *********** Use Cases
// AlertUseCase needs to be implemented in Alert use cases
type AlertUseCase interface {
	GetAllAlerts() ([]Alert, error)
}

// *********** Services
type AlertService interface {
	GetAllAlerts(int, int, string) ([]Alert, error)
}
