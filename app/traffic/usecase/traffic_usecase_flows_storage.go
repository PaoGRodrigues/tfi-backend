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

func (fs *FlowsRepository) GetFlows(attr string) (domains.Server, error) {
	flow, err := fs.trafficRepo.GetServerByAttr(attr)
	if err != nil {
		return domains.Server{}, err
	}
	return flow, nil
}

func (fs *FlowsRepository) GetClientsList() ([]domains.Client, error) {
	clients, err := fs.trafficRepo.GetClients()
	if err != nil {
		return nil, err
	}
	return clients, nil
}

func (fs *FlowsRepository) GetServersList() ([]domains.Server, error) {
	servers, err := fs.trafficRepo.GetServers()
	if err != nil {
		return nil, err
	}
	return servers, nil
}

func (fs *FlowsRepository) GetFlowByKey(key string) (domains.ActiveFlow, error) {
	flow, err := fs.trafficRepo.GetFlowByKey(key)
	if err != nil {
		return domains.ActiveFlow{}, err
	}
	return flow, nil
}
