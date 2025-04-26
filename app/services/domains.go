package services

import (
	alerts_domains "github.com/PaoGRodrigues/tfi-backend/app/alerts/domains"
	"github.com/PaoGRodrigues/tfi-backend/app/domain/host"
	traffic_domains "github.com/PaoGRodrigues/tfi-backend/app/traffic/domains"
)

type Tool interface {
	SetInterfaceID() error
	GetAllHosts() ([]host.Host, error)
	GetAllActiveTraffic() ([]traffic_domains.ActiveFlow, error)
	GetAllAlerts(epoch_begin, epoch_end int) ([]alerts_domains.Alert, error)
	EnableChecks()
}

type Terminal interface {
	BlockHost(string) error
}

type NotificationChannel interface {
	Configure(string, string) error
	SendMessage(string) error
}

type Database interface {
	AddActiveFlows([]traffic_domains.ActiveFlow) error
	GetServerByAttr(attr string) (traffic_domains.Server, error)
	GetClients() ([]traffic_domains.Client, error)
	GetServers() ([]traffic_domains.Server, error)
	GetFlowByKey(key string) (traffic_domains.ActiveFlow, error)
	AddHosts([]host.Host) error
	GetHostByIp(string) (host.Host, error)
}
