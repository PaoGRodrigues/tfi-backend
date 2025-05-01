package traffic

import (
	domains "github.com/PaoGRodrigues/tfi-backend/app/domain/traffic"
	hostPorts "github.com/PaoGRodrigues/tfi-backend/app/ports/host"
	trafficPorts "github.com/PaoGRodrigues/tfi-backend/app/ports/traffic"
)

type StoreTrafficFlowsUseCase struct {
	trafficReader       trafficPorts.TrafficReader
	trafficDBRepository trafficPorts.TrafficDBRepository
	hostDBRepository    hostPorts.HostDBRepository
}

func NewStoreTrafficFlowsUseCase(trafficReader trafficPorts.TrafficReader, trafficDBRepository trafficPorts.TrafficDBRepository, hostDBRepository hostPorts.HostDBRepository) *StoreTrafficFlowsUseCase {
	return &StoreTrafficFlowsUseCase{
		trafficReader:       trafficReader,
		trafficDBRepository: trafficDBRepository,
		hostDBRepository:    hostDBRepository,
	}
}

func (usecase *StoreTrafficFlowsUseCase) StoreTrafficFlows() error {
	current, err := usecase.trafficReader.GetTrafficFlows()
	if err != nil {
		return err
	}
	if len(current) != 0 {
		flows, err := usecase.enrichData(current)
		if err != nil {
			return err
		}
		err = usecase.trafficDBRepository.StoreTrafficFlows(flows)
		if err != nil {
			return err
		}
	}
	return err
}

func (usecase *StoreTrafficFlowsUseCase) enrichData(trafficFlows []domains.TrafficFlow) ([]domains.TrafficFlow, error) {
	newFlows := []domains.TrafficFlow{}
	for _, flow := range trafficFlows {
		serv, err := usecase.hostDBRepository.GetHostByIp(flow.Server.IP)
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
