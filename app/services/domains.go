package services

import (
	alert "github.com/PaoGRodrigues/tfi-backend/app/domain/alert"
	host "github.com/PaoGRodrigues/tfi-backend/app/domain/host"
	traffic_domains "github.com/PaoGRodrigues/tfi-backend/app/traffic/domains"
)

type Tool interface {
	SetInterfaceID() error
	GetAllHosts() ([]host.Host, error)
	GetAllActiveTraffic() ([]traffic_domains.ActiveFlow, error)
	GetAllAlerts(epoch_begin, epoch_end int) ([]alert.Alert, error)
	EnableChecks()
}

type Terminal interface {
	Block(string) (*string, error)
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
	StoreHosts([]host.Host) error
	GetHostByIp(string) (host.Host, error)
}
