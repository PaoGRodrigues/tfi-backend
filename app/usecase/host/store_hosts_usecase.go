package host

import (
	"github.com/PaoGRodrigues/tfi-backend/app/domain/host"
	ports "github.com/PaoGRodrigues/tfi-backend/app/ports/host"
)

type StoreHostUseCase struct {
	hostRepositoryReader ports.HostReader
	hostDBRepository     ports.HostDBRepository
}

func NewHostsStorage(repository ports.HostReader, hostDBRepository ports.HostDBRepository) *StoreHostUseCase {
	return &StoreHostUseCase{
		hostRepositoryReader: repository,
		hostDBRepository:     hostDBRepository,
	}
}

func (usecase *StoreHostUseCase) StoreHosts() error {
	activeHosts, err := usecase.hostRepositoryReader.GetAllHosts()
	if err != nil {
		return err
	}

	if activeHosts != nil {
		err = usecase.hostDBRepository.StoreHosts(activeHosts)
		if err != nil {
			return err
		}
	}
	return nil
}

func (usecase *StoreHostUseCase) GetHostByIp(ip string) (host.Host, error) {
	current, err := usecase.hostDBRepository.GetHostByIp(ip)
	if err != nil {
		return host.Host{}, err
	}

	return current, nil
}
