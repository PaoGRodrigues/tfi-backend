package host

import (
	"github.com/PaoGRodrigues/tfi-backend/app/domain/host"
	ports "github.com/PaoGRodrigues/tfi-backend/app/ports/host"
)

type GetLocalhostsUseCase struct {
	searcher ports.HostRepository
}

func NewGetLocalhostsUseCase(huc ports.HostRepository) *GetLocalhostsUseCase {
	return &GetLocalhostsUseCase{
		searcher: huc,
	}
}

func (l *GetLocalhostsUseCase) GetLocalHosts() ([]host.Host, error) {
	current, err := l.searcher.GetAllHosts()
	if err != nil {
		return []host.Host{}, err
	}
	localHosts := host.GetLocalHosts(current)
	return localHosts, nil
}
