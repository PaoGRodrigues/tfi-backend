package services

import (
	alert "github.com/PaoGRodrigues/tfi-backend/app/domain/alert"
	host "github.com/PaoGRodrigues/tfi-backend/app/domain/host"
	traffic "github.com/PaoGRodrigues/tfi-backend/app/domain/traffic"
)

type Tool interface {
	SetInterfaceID() error
	GetAllHosts() ([]host.Host, error)
	GetAllActiveTraffic() ([]traffic.TrafficFlow, error)
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
	AddActiveFlows([]traffic.TrafficFlow) error
	GetServerByAttr(attr string) (traffic.Server, error)
	GetClients() ([]traffic.Client, error)
	GetServers() ([]traffic.Server, error)
	GetFlowByKey(key string) (traffic.TrafficFlow, error)
	StoreHosts([]host.Host) error
	GetHostByIp(string) (host.Host, error)
}
