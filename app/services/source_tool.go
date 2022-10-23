package services

import (
	domains_alert "github.com/PaoGRodrigues/tfi-backend/app/alerts/domains"
	domains_host "github.com/PaoGRodrigues/tfi-backend/app/hosts/domains"
	domains_traffic "github.com/PaoGRodrigues/tfi-backend/app/traffic/domains"
)

type Tool interface {
	GetAllHosts() ([]domains_host.Host, error)
	GetAllActiveTraffic() ([]domains_traffic.ActiveFlow, error)
	GetAllAlerts(epoch_begin, epoch_end int) ([]domains_alert.Alert, error)
}

type NtopNG struct {
	UrlClient   string
	InterfaceId int
	Usr         string
	Pass        string
}

func NewTool(urlClient string, interfaceId int, usr string, pass string) *NtopNG {
	return &NtopNG{
		UrlClient:   urlClient,
		InterfaceId: interfaceId,
		Usr:         usr,
		Pass:        pass,
	}
}

type FakeTool struct{}

func NewFakeTool() *FakeTool {
	return &FakeTool{}
}
