package usecase

import (
	host_domains "github.com/PaoGRodrigues/tfi-backend/app/domain/host"
	"github.com/PaoGRodrigues/tfi-backend/app/traffic/domains"
)

type FlowsStorage struct {
	trafficSearcher domains.TrafficUseCase
	trafficRepo     domains.TrafficRepository
	hostStorage     host_domains.HostsStorage
}

func NewFlowsStorage(trafSearcher domains.TrafficUseCase, trafRepo domains.TrafficRepository, hostStorage host_domains.HostsStorage) *FlowsStorage {
	return &FlowsStorage{
		trafficSearcher: trafSearcher,
		trafficRepo:     trafRepo,
		hostStorage:     hostStorage,
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
		serv, err := fs.hostStorage.GetHostByIp(flow.Server.IP)
		if err != nil {
			return []domains.ActiveFlow{}, err
		}
		flow.Server.Country = serv.Country
		if flow.Server.IsBroadcastDomain {
			flow.Server.Name = flow.Server.IP
		}

		newFlows = append(newFlows, flow)
	}
	return newFlows, nil
}
