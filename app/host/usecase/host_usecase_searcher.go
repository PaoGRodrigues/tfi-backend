package usecase

import (
	"github.com/PaoGRodrigues/tfi-backend/app/host/domains"
)

type HostSearcher struct {
	hostService  domains.HostService
	currentHosts []domains.Host
}

func NewHostSearcher(repo domains.HostService) *HostSearcher {

	return &HostSearcher{
		hostService: repo,
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
