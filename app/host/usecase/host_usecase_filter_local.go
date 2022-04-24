package usecase

import "github.com/PaoGRodrigues/tfi-backend/app/host/domains"

type LocalHosts struct {
	searcher domains.HostUseCase
}

func NewLocalHosts(huc domains.HostUseCase) *LocalHosts {
	return &LocalHosts{
		searcher: huc,
	}
}

func (l *LocalHosts) GetLocalHosts() ([]domains.Host, error) {
	current := l.searcher.GetHosts()
	if len(current) == 0 {
		hosts, err := l.searcher.GetAllHosts()
		if err != nil {
			return nil, err
		}
		current = hosts
	}
	localHosts := l.getLocalHosts(current)
	return localHosts, nil
}

func (l *LocalHosts) getLocalHosts(hosts []domains.Host) []domains.Host {
	localHosts := []domains.Host{}

	for _, host := range hosts {
		if host.PrivateHost {
			localHosts = append(localHosts, host)
		}
	}
	return localHosts
}
