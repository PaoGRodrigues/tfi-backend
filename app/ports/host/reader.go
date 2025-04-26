package host

import "github.com/PaoGRodrigues/tfi-backend/app/domain/host"

// HostReader needs to be implemented in Host use cases
type HostReader interface {
	GetAllHosts() ([]host.Host, error)
}
