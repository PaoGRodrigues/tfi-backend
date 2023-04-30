package services

import (
	domains_alert "github.com/PaoGRodrigues/tfi-backend/app/alerts/domains"
	"github.com/PaoGRodrigues/tfi-backend/app/hosts/domains"
	domains_host "github.com/PaoGRodrigues/tfi-backend/app/hosts/domains"
	domains_traffic "github.com/PaoGRodrigues/tfi-backend/app/traffic/domains"
)

type Tool interface {
	GetAllHosts() ([]domains_host.Host, error)
	GetAllActiveTraffic() ([]domains_traffic.ActiveFlow, error)
	GetAllAlerts(epoch_begin, epoch_end int) ([]domains_alert.Alert, error)
}

type Terminal interface {
	BlockHost(domains.Host) error
}

type NotificationChannel interface {
	Configure(string, string) error
	SendMessage(string) error
}
