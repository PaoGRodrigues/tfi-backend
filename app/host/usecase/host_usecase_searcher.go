package usecase

import (
	"github.com/PaoGRodrigues/tfi-backend/app/host/domains"
)

type HostSearcher struct {
	hostRepo domains.HostRepository
}

func NewHostSearcher(repo domains.HostRepository) *HostSearcher {

	return &HostSearcher{
		hostRepo: repo,
	}
}

func (gw *HostSearcher) GetAllHosts() ([]domains.Host, error) {
	res, err := gw.hostRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return res, nil
}
