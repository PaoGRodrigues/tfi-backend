package services

import (
	alerts_domains "github.com/PaoGRodrigues/tfi-backend/app/alerts/domains"
	"github.com/PaoGRodrigues/tfi-backend/app/hosts/domains"
	hosts_domains "github.com/PaoGRodrigues/tfi-backend/app/hosts/domains"
	traffic_domains "github.com/PaoGRodrigues/tfi-backend/app/traffic/domains"
)

type Tool interface {
	GetAllHosts() ([]hosts_domains.Host, error)
	GetAllActiveTraffic() ([]traffic_domains.ActiveFlow, error)
	GetAllAlerts(epoch_begin, epoch_end int) ([]alerts_domains.Alert, error)
}

type Terminal interface {
	BlockHost(domains.Host) error
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
	AddHosts([]hosts_domains.Host) error
}
