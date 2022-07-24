package usecase

import "github.com/PaoGRodrigues/tfi-backend/app/hosts/domains"

type HostsFilter struct {
	searcher domains.HostUseCase
}

func NewHostsFilter(huc domains.HostUseCase) *HostsFilter {
	return &HostsFilter{
		searcher: huc,
	}
}

func (l *HostsFilter) checkHosts() ([]domains.Host, error) {
	current := l.searcher.GetHosts()
	if len(current) == 0 {
		hosts, err := l.searcher.GetAllHosts()
		if err != nil {
			return []domains.Host{}, err
		}
		current = hosts
	}
	return current, nil
}

func (l *HostsFilter) GetLocalHosts() ([]domains.Host, error) {
	current, err := l.checkHosts()
	if err != nil {
		return []domains.Host{}, err
	}
	localHosts := []domains.Host{}
	if len(current) != 0 {
		for _, host := range current {
			if host.PrivateHost {
				localHosts = append(localHosts, host)
			}
		}
	}
	return localHosts, nil
}

func (l *HostsFilter) GetRemoteHosts() ([]domains.Host, error) {
	current, err := l.checkHosts()
	if err != nil {
		return []domains.Host{}, err
	}
	remoteHosts := []domains.Host{}
	if len(current) != 0 {
		for _, host := range current {
			if !host.PrivateHost {
				remoteHosts = append(remoteHosts, host)
			}
		}
	}
	return remoteHosts, nil
}
