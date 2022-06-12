package usecase

import "github.com/PaoGRodrigues/tfi-backend/app/traffic/domains"

type FlowsStorage struct {
	trafficSearcher domains.TrafficUseCase
	trafficStorage  domains.TrafficRepoStore
}

func NewFlowsStorage(trafSearcher domains.TrafficUseCase, trafStorage domains.TrafficRepoStore) *FlowsStorage {
	return &FlowsStorage{
		trafficSearcher: trafSearcher,
		trafficStorage:  trafStorage,
	}
}

func (fs *FlowsStorage) Store() error {
	activeFlows := fs.trafficSearcher.GetActiveFlows()
	if len(activeFlows) == 0 {
		current, err := fs.trafficSearcher.GetAllActiveTraffic()
		if err != nil {
			return err
		}
		activeFlows = current
	}
	err := fs.trafficStorage.StoreActiveFlows(activeFlows)
	return err
}
