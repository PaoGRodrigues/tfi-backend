package usecase

import "github.com/PaoGRodrigues/tfi-backend/app/domain/host"

type HostSearcher struct {
	hostService  host.HostService
	currentHosts []host.Host
}

func NewHostSearcher(service host.HostService) *HostSearcher {
	return &HostSearcher{
		hostService: service,
	}
}

func (hs *HostSearcher) GetAllHosts() ([]host.Host, error) {
	res, err := hs.hostService.GetAllHosts()
	if err != nil {
		return nil, err
	}
	hs.currentHosts = res
	return res, nil
}

func (hs *HostSearcher) GetHosts() []host.Host {
	return hs.currentHosts
}
