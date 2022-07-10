package usecase

import "github.com/PaoGRodrigues/tfi-backend/app/traffic/domains"

type FlowsRepository struct {
	trafficSearcher domains.TrafficUseCase
	trafficRepo     domains.TrafficRepository
}

func NewFlowsStorage(trafSearcher domains.TrafficUseCase, trafRepo domains.TrafficRepository) *FlowsRepository {
	return &FlowsRepository{
		trafficSearcher: trafSearcher,
		trafficRepo:     trafRepo,
	}
}

func (fs *FlowsRepository) StoreFlows() error {
	activeFlows := fs.trafficSearcher.GetActiveFlows()
	if len(activeFlows) == 0 {
		current, err := fs.trafficSearcher.GetAllActiveTraffic()
		if err != nil {
			return err
		}
		activeFlows = current
	}
	err := fs.trafficRepo.AddActiveFlows(activeFlows)
	return err
}
