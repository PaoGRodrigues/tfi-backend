package host

import "github.com/PaoGRodrigues/tfi-backend/app/domain/host"

type HostDBRepository interface {
	StoreHosts([]host.Host) error
	GetHostByIp(string) (host.Host, error)
}
