package usecase

import (
	"errors"

	"github.com/PaoGRodrigues/tfi-backend/app/domain/host"
)

type HostsFilter struct {
	searcher host.HostUseCase
}

func NewHostsFilter(huc host.HostUseCase) *HostsFilter {
	return &HostsFilter{
		searcher: huc,
	}
}

func (l *HostsFilter) checkHosts() ([]host.Host, error) {
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

func (l *HostsFilter) GetLocalHosts() ([]host.Host, error) {
	current, err := l.checkHosts()
	if err != nil {
		return []host.Host{}, err
	}
	localHosts := []host.Host{}
	if len(current) != 0 {
		for _, host := range current {
			if host.PrivateHost {
				localHosts = append(localHosts, host)
			}
		}
	}
	return localHosts, nil
}

func (l *HostsFilter) GetRemoteHosts() ([]host.Host, error) {
	current, err := l.checkHosts()
	if err != nil {
		return []host.Host{}, err
	}
	remoteHosts := []host.Host{}
	if len(current) != 0 {
		for _, host := range current {
			if !host.PrivateHost {
				remoteHosts = append(remoteHosts, host)
			}
		}
	}
	return remoteHosts, nil
}

func (l *HostsFilter) GetHost(attr string) (host.Host, error) {
	current, err := l.checkHosts()
	if err != nil {
		return host.Host{}, err
	}
	var ip string
	for _, host := range current {
		if host.IP == attr || host.Name == attr {
			return host, nil
		}
		ip = host.IP
	}

	return host.Host{}, errors.New("There's no host with this IP " + ip)
}
