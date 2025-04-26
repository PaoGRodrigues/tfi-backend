package host

import "github.com/PaoGRodrigues/tfi-backend/app/domain/host"

// HostRepository needs to be implemented in Host use cases
type HostRepository interface {
	GetAllHosts() ([]host.Host, error)
}
