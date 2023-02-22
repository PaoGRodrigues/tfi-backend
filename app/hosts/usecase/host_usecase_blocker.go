package usecase

import (
	"github.com/PaoGRodrigues/tfi-backend/app/hosts/domains"
	trafficDomains "github.com/PaoGRodrigues/tfi-backend/app/traffic/domains"
)

type Blocker struct {
	filter       trafficDomains.ActiveFlowsStorage
	blockService domains.HostBlockerService
}

func NewBlocker(srv domains.HostBlockerService, filter trafficDomains.ActiveFlowsStorage) *Blocker {
	return &Blocker{
		filter:       filter,
		blockService: srv,
	}
}

func (blocker *Blocker) Block(attr string) (domains.Host, error) {
	server, err := blocker.filter.GetFlows(attr)
	if err != nil {
		return domains.Host{}, err
	}
	host := convertServerToHost(server)
	err = blocker.blockService.BlockHost(host)
	if err != nil {
		return domains.Host{}, err
	}
	return host, nil
}

func convertServerToHost(server trafficDomains.Server) domains.Host {
	host := domains.Host{
		IP:   server.IP,
		Name: server.Name,
	}
	return host
}
