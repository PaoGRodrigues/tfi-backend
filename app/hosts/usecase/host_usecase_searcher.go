package usecase

import (
	"github.com/PaoGRodrigues/tfi-backend/app/hosts/domains"
)

type HostSearcher struct {
	hostService  domains.HostService
	currentHosts []domains.Host
}

func NewHostSearcher(service domains.HostService) *HostSearcher {

	return &HostSearcher{
		hostService: service,
	}
}

func (hs *HostSearcher) GetAllHosts() ([]domains.Host, error) {
	res, err := hs.hostService.GetAllHosts()
	if err != nil {
		return nil, err
	}
	hs.currentHosts = res
	return res, nil
}

func (hs *HostSearcher) GetHosts() []domains.Host {
	return hs.currentHosts
}
