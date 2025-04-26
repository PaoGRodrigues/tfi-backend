package host

import (
	"github.com/PaoGRodrigues/tfi-backend/app/domain/host"
	ports "github.com/PaoGRodrigues/tfi-backend/app/ports/host"
)

type StoreHostUseCase struct {
	hostRepositoryReader ports.HostReader
	hostDBRepository     ports.HostDBRepository
}

func NewHostsStorage(repository ports.HostReader, hostDBRepository host.HostsRepository) *StoreHostUseCase {
	return &StoreHostUseCase{
		hostRepositoryReader: repository,
		hostDBRepository:     hostDBRepository,
	}
}

func (hs *StoreHostUseCase) StoreHosts() error {
	activeHosts, err := hs.hostRepositoryReader.GetAllHosts()
	if err != nil {
		return err
	}

	if activeHosts != nil {
		err = hs.hostDBRepository.StoreHosts(activeHosts)
		if err != nil {
			return err
		}
	}
	return nil
}

func (hs *StoreHostUseCase) GetHostByIp(ip string) (host.Host, error) {
	current, err := hs.hostDBRepository.GetHostByIp(ip)
	if err != nil {
		return host.Host{}, err
	}

	return current, nil
}
