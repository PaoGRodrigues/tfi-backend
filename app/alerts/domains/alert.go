package domains

import (
	flow "github.com/PaoGRodrigues/tfi-backend/app/traffic/domains"
)

// *********** Entities
// Alerts
type Alert struct {
	Row_id string `json:"-"`
	Name   string `json:"fullname"`
	Family string
	Time   struct {
		Label string
	} `json:"tstamp"`
	Severity      Severity
	AlertFlow     AlertFlow `json:"flow"`
	AlertProtocol AlertProtocol
}

type Severity struct {
	Label string
}

type AlertFlow struct {
	Client AlertClient `json:"cli_ip"`
	Server AlertServer `json:"srv_ip"`
}

type AlertClient struct {
	Value   string
	Contry  string
	CliPort int `json:"cli_port"`
}

type AlertServer struct {
	Name    string
	Value   string
	Country string
	SrvPort int `json:"srv_port"`
}

type AlertProtocol struct {
	Protocol flow.Protocol `json:"l7_proto"`
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
