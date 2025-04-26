package usecase

import "github.com/PaoGRodrigues/tfi-backend/app/domain/host"

type HostsStorage struct {
	hostSearcher host.HostUseCase
	hostRepo     host.HostsRepository
}

func NewHostsStorage(hostSearcher host.HostUseCase, hostRepo host.HostsRepository) *HostsStorage {
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

func (hs *HostsStorage) GetHostByIp(ip string) (host.Host, error) {
	current, err := hs.hostRepo.GetHostByIp(ip)
	if err != nil {
		return host.Host{}, err
	}

	return current, nil
}
