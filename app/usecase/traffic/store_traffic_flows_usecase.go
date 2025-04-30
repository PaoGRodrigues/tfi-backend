package traffic

import (
	domains "github.com/PaoGRodrigues/tfi-backend/app/domain/traffic"
	hostPorts "github.com/PaoGRodrigues/tfi-backend/app/ports/host"
	trafficPorts "github.com/PaoGRodrigues/tfi-backend/app/ports/traffic"
)

type StoreTrafficFlowsUseCase struct {
	trafficReader       trafficPorts.TrafficReader
	trafficDBRepository trafficPorts.TrafficDBRepository
	hostStorage         hostPorts.HostDBRepository
}

func NewStoreTrafficFlowsUseCase(trafficReader trafficPorts.TrafficReader, trafficDBRepository trafficPorts.TrafficDBRepository, hostStorage hostPorts.HostDBRepository) *StoreTrafficFlowsUseCase {
	return &StoreTrafficFlowsUseCase{
		trafficReader:       trafficReader,
		trafficDBRepository: trafficDBRepository,
		hostStorage:         hostStorage,
	}
}

func (fs *StoreTrafficFlowsUseCase) StoreFlows() error {
	current, err := fs.trafficReader.GetTrafficFlows()
	if err != nil {
		return err
	}
	if len(current) != 0 {
		flows, err := fs.enrichData(current)
		if err != nil {
			return err
		}
		err = fs.trafficDBRepository.StoreTrafficFlows(flows)
		if err != nil {
			return err
		}
	}
	return err
}

func (fs *StoreTrafficFlowsUseCase) enrichData(activeFlows []domains.TrafficFlow) ([]domains.TrafficFlow, error) {
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
