package host

import "github.com/PaoGRodrigues/tfi-backend/app/domain/host"

type HostWriter interface {
	StoreHosts() error
	GetHostByIp(string) (host.Host, error)
}
