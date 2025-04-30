package traffic

import (
	domains "github.com/PaoGRodrigues/tfi-backend/app/domain/traffic"
	hostPorts "github.com/PaoGRodrigues/tfi-backend/app/ports/host"
	trafficPorts "github.com/PaoGRodrigues/tfi-backend/app/ports/traffic"
)

type FlowsStorage struct {
	trafficReader trafficPorts.TrafficReader
	trafficRepo   domains.TrafficRepository
	hostStorage   hostPorts.HostDBRepository
}

func NewFlowsStorage(trafSearcher trafficPorts.TrafficReader, trafRepo domains.TrafficRepository, hostStorage hostPorts.HostDBRepository) *FlowsStorage {
	return &FlowsStorage{
		trafficReader: trafSearcher,
		trafficRepo:   trafRepo,
		hostStorage:   hostStorage,
	}
}

func (fs *FlowsStorage) StoreFlows() error {
	current, err := fs.trafficReader.GetTrafficFlows()
	if err != nil {
		return err
	}
	if len(current) != 0 {
		flows, err := fs.enrichData(current)
		if err != nil {
			return err
		}
		err = fs.trafficRepo.StoreFlows(flows)
		if err != nil {
			return err
		}
	}
	return err
}

func (fs *FlowsStorage) enrichData(activeFlows []domains.TrafficFlow) ([]domains.TrafficFlow, error) {
	newFlows := []domains.TrafficFlow{}
	for _, flow := range activeFlows {
		serv, err := fs.hostStorage.GetHostByIp(flow.Server.IP)
		if err != nil {
			return []domains.TrafficFlow{}, err
		}
		flow.Server.Country = serv.Country
		if flow.Server.IsBroadcastDomain {
			flow.Server.Name = flow.Server.IP
		}

		newFlows = append(newFlows, flow)
	}
	return newFlows, nil
}
