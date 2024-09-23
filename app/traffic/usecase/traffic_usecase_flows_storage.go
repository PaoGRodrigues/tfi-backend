package usecase

import (
	host_domains "github.com/PaoGRodrigues/tfi-backend/app/hosts/domains"
	"github.com/PaoGRodrigues/tfi-backend/app/traffic/domains"
)

type FlowsStorage struct {
	trafficSearcher domains.TrafficUseCase
	trafficRepo     domains.TrafficRepository
	hostFilter      host_domains.HostsFilter
}

func NewFlowsStorage(trafSearcher domains.TrafficUseCase, trafRepo domains.TrafficRepository, hostFilter host_domains.HostsFilter) *FlowsStorage {
	return &FlowsStorage{
		trafficSearcher: trafSearcher,
		trafficRepo:     trafRepo,
		hostFilter:      hostFilter,
	}
}

func (fs *FlowsStorage) StoreFlows() error {
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
	err = fs.trafficRepo.StoreFlows(flows)
	return err
}

func (fs *FlowsStorage) enrichData(activeFlows []domains.ActiveFlow) ([]domains.ActiveFlow, error) {
	newFlows := []domains.ActiveFlow{}
	for _, flow := range activeFlows {
		serv, err := fs.hostFilter.GetHost(flow.Server.IP)
		if err != nil {
			serv, err = fs.hostFilter.GetHost(flow.Server.Name)
			if err != nil {
				return []domains.ActiveFlow{}, err
			}
		}
		flow.Server.Country = serv.Country
		newFlows = append(newFlows, flow)
	}
	return newFlows, nil
}
