package host

import (
	"github.com/PaoGRodrigues/tfi-backend/app/domain/host"
	ports "github.com/PaoGRodrigues/tfi-backend/app/ports/host"
)

type GetLocalhostsUseCase struct {
	repository ports.HostReader
}

func NewGetLocalhostsUseCase(repository ports.HostReader) *GetLocalhostsUseCase {
	return &GetLocalhostsUseCase{
		repository: repository,
	}
}

func (usecase *GetLocalhostsUseCase) GetLocalHosts() ([]host.Host, error) {
	current, err := usecase.repository.GetAllHosts()
	if err != nil {
		return []host.Host{}, err
	}
	localHosts := host.GetLocalHosts(current)
	return localHosts, nil
}
