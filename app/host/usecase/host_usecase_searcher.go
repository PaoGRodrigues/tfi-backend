package usecase

import (
	"github.com/PaoGRodrigues/tfi-backend/app/host/domains"
)

type HostSearcher struct {
	hostRepo     domains.HostRepository
	currentHosts []domains.Host
}

func NewHostSearcher(repo domains.HostRepository) *HostSearcher {

	return &HostSearcher{
		hostRepo: repo,
	}
}

func (hs *HostSearcher) GetAllHosts() ([]domains.Host, error) {
	res, err := hs.hostRepo.GetAll()
	if err != nil {
		return nil, err
	}
	hs.currentHosts = res
	return res, nil
}

func (hs *HostSearcher) GetHosts() []domains.Host {
	return hs.currentHosts
}
