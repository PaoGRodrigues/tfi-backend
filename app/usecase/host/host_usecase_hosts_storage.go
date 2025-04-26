package host

import (
	"github.com/PaoGRodrigues/tfi-backend/app/domain/host"
	ports "github.com/PaoGRodrigues/tfi-backend/app/ports/host"
)

type HostsStorage struct {
	hostRepositoryReader ports.HostReader
	hostRepo             host.HostsRepository
}

func NewHostsStorage(repository ports.HostReader, hostRepo host.HostsRepository) *HostsStorage {
	return &HostsStorage{
		hostRepositoryReader: repository,
		hostRepo:             hostRepo,
	}
}

func (hs *HostsStorage) StoreHosts() error {
	activeHosts, err := hs.hostRepositoryReader.GetAllHosts()
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
