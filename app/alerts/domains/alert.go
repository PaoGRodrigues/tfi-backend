package domains

import (
	flow "github.com/PaoGRodrigues/tfi-backend/app/traffic/domains"
)

// *********** Entities
// Alerts
type Alert struct {
	Name   string
	Family string
	Time   struct {
		Label string
	} `json:"tstamp"`
	Score         string `json:"score,omitempty"`
	Severity      string
	AlertFlow     AlertFlow
	AlertProtocol flow.Protocol
}

type AlertFlow struct {
	Client flow.Client
	Server flow.Server
}

// *********** Use Cases
// AlertUseCase needs to be implemented in Alert use cases
type AlertUseCase interface {
	GetAllAlerts() ([]Alert, error)
	GetAllAlertsByTime(int, int) ([]Alert, error)
}

type AlertsSender interface {
	SendLastAlertMessages() error
}

// *********** Services
type AlertService interface {
	GetAllAlerts(int, int) ([]Alert, error)
}

type Notifier interface {
	SendMessage(string) error
}
