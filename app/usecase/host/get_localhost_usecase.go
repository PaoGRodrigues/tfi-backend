package host

import (
	"github.com/PaoGRodrigues/tfi-backend/app/domain/host"
)

type GetLocalhostsUseCase struct {
	searcher host.HostUseCase
}

func NewGetLocalhostsUseCase(huc host.HostUseCase) *GetLocalhostsUseCase {
	return &GetLocalhostsUseCase{
		searcher: huc,
	}
}

func (l *GetLocalhostsUseCase) GetLocalHosts() ([]host.Host, error) {
	current, err := l.checkHosts()
	if err != nil {
		return []host.Host{}, err
	}
	localHosts := host.GetLocalHosts(current)
	return localHosts, nil
}

func (l *GetLocalhostsUseCase) checkHosts() ([]host.Host, error) {
	current := l.searcher.GetHosts()
	if len(current) == 0 {
		hosts, err := l.searcher.GetAllHosts()
		if err != nil {
			return []host.Host{}, err
		}
		current = hosts
	}
	return current, nil
}
