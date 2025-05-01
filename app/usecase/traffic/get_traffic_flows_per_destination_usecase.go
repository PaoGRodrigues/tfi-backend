package traffic

import (
	traffic "github.com/PaoGRodrigues/tfi-backend/app/domain/traffic"
	trafficPorts "github.com/PaoGRodrigues/tfi-backend/app/ports/traffic"
)

type BytesPerDestination struct {
	Bytes       int
	Destination string
}

type GetTrafficFlowsPerDestinationUseCase struct {
	trafficDBRepository trafficPorts.TrafficDBRepository
}

func NewGetTrafficFlowsPerDestinationUseCase(trafficDBRepository trafficPorts.TrafficDBRepository) *GetTrafficFlowsPerDestinationUseCase {
	return &GetTrafficFlowsPerDestinationUseCase{
		trafficDBRepository: trafficDBRepository,
	}
}

func (usecase *GetTrafficFlowsPerDestinationUseCase) GetTrafficFlowsPerDestinations() ([]BytesPerDestination, error) {
	serversList, err := usecase.trafficDBRepository.GetServers()

	if err != nil {
		return []BytesPerDestination{}, err
	}

	servers := traffic.FilterPublicServers(serversList)

	flows := []traffic.TrafficFlow{}
	for _, server := range servers {
		flow, err := usecase.trafficDBRepository.GetFlowByKey(server.Key)
		if err != nil {
			return []BytesPerDestination{}, err
		}
		flow.Server = server
		flows = append(flows, flow)
	}

	parsedBytesDst := parsePerDest(flows)
	bytesDst := sumBytes(parsedBytesDst)
	return bytesDst, nil
}

func parsePerDest(flows []traffic.TrafficFlow) []BytesPerDestination {

	bytesDst := []BytesPerDestination{}

	for _, flow := range flows {
		var serverName string
		if flow.Server.Name != "" {
			serverName = flow.Server.Name
		} else {
			serverName = flow.Server.IP
		}
		bpd := BytesPerDestination{
			Bytes:       flow.Bytes,
			Destination: serverName,
		}
		bytesDst = append(bytesDst, bpd)
	}
	return bytesDst
}

func sumBytes(bpd []BytesPerDestination) []BytesPerDestination {
	m := map[string]int{}
	for _, v := range bpd {
		m[v.Destination] += v.Bytes
	}

	newBpd := []BytesPerDestination{}
	for dest, bytes := range m {
		new := BytesPerDestination{
			Destination: dest,
			Bytes:       bytes,
		}
		newBpd = append(newBpd, new)
	}

	return newBpd
}
