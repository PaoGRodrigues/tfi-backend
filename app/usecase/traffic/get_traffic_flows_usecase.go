package traffic

import (
	domains "github.com/PaoGRodrigues/tfi-backend/app/domain/traffic"
	ports "github.com/PaoGRodrigues/tfi-backend/app/ports/traffic"
)

type GetTrafficFlowsUseCase struct {
	repository  ports.TrafficReader
	activeFlows []domains.TrafficFlow
}

func NewTrafficFlowsUseCase(trafSrv ports.TrafficReader) *GetTrafficFlowsUseCase {
	return &GetTrafficFlowsUseCase{
		repository: trafSrv,
	}
}

func (usecase *GetTrafficFlowsUseCase) GetTrafficFlows() ([]domains.TrafficFlow, error) {
	res, err := usecase.repository.GetTrafficFlows()
	if err != nil {
		return []domains.TrafficFlow{}, err
	}
	usecase.activeFlows = res
	return res, nil
}
