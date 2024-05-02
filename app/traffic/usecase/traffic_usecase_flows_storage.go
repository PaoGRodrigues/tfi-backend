package usecase

import (
	host_domains "github.com/PaoGRodrigues/tfi-backend/app/hosts/domains"
	"github.com/PaoGRodrigues/tfi-backend/app/traffic/domains"
)

type FlowsRepository struct {
	trafficSearcher domains.TrafficUseCase
	trafficRepo     domains.TrafficRepository
	hostFilter      host_domains.HostsFilter
}

func NewFlowsStorage(trafSearcher domains.TrafficUseCase, trafRepo domains.TrafficRepository, hostFilter host_domains.HostsFilter) *FlowsRepository {
	return &FlowsRepository{
		trafficSearcher: trafSearcher,
		trafficRepo:     trafRepo,
		hostFilter:      hostFilter,
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

	flows, err := fs.enrichData(activeFlows)
	if err != nil {
		return err
	}
	err = fs.trafficRepo.AddActiveFlows(flows)
	return err
}

func (fs *FlowsRepository) enrichData(activeFlows []domains.ActiveFlow) ([]domains.ActiveFlow, error) {
	newFlows := []domains.ActiveFlow{}
	for _, flow := range activeFlows {
		serv, err := fs.hostFilter.GetHost(flow.Server.IP)
		if err != nil {
			return []domains.ActiveFlow{}, err
		}
		flow.Server.Country = serv.Country
		newFlows = append(newFlows, flow)
	}
	return newFlows, nil
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
