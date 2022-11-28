package usecase

import "github.com/PaoGRodrigues/tfi-backend/app/hosts/domains"

type Blocker struct {
	filter       domains.HostsFilter
	blockService domains.HostBlockerService
}

func NewBlocker(srv domains.HostBlockerService, filter domains.HostsFilter) domains.HostBlocker {
	return &Blocker{
		filter:       filter,
		blockService: srv,
	}
}

func (blocker *Blocker) Block(ip string) (domains.Host, error) {
	host, err := blocker.filter.GetHost(ip)
	if err != nil {
		return domains.Host{}, err
	}
	err = blocker.blockService.BlockHost(host)
	if err != nil {
		return domains.Host{}, err
	}
	return host, nil
}
