package usecase

import "github.com/PaoGRodrigues/tfi-backend/app/hosts/domains"

type HostsStorage struct {
	hostSearcher domains.HostUseCase
	hostRepo     domains.HostsRepository
}

func NewHostsStorage(hostSearcher domains.HostUseCase, hostRepo domains.HostsRepository) *HostsStorage {
	return &HostsStorage{
		hostSearcher: hostSearcher,
		hostRepo:     hostRepo,
	}
}

func (hs *HostsStorage) StoreHosts() error {
	activeHosts, err := hs.hostSearcher.GetAllHosts()
	if err != nil {
		return err
	}

	if activeHosts != nil {
		err = hs.hostRepo.StoreHosts(activeHosts)
		if err != nil {
			return err
		}
	}
	return nil
}

func (hs *HostsStorage) GetHostByIp(ip string) (domains.Host, error) {
	host, err := hs.hostRepo.GetHostByIp(ip)
	if err != nil {
		return domains.Host{}, err
	}

	return host, nil
}
