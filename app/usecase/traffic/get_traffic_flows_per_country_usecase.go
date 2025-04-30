package traffic

import (
	traffic "github.com/PaoGRodrigues/tfi-backend/app/domain/traffic"
	trafficPorts "github.com/PaoGRodrigues/tfi-backend/app/ports/traffic"
)

type BytesPerCountry struct {
	Bytes   int    `json:"bytes"`
	Country string `json:"country"`
}

type GetTrafficFlowsPerCountryUseCase struct {
	trafficDBRepository trafficPorts.TrafficDBRepository
}

func NewGetTrafficFlowsPerCountryUseCase(trafficDBRepository trafficPorts.TrafficDBRepository) *GetTrafficFlowsPerCountryUseCase {
	return &GetTrafficFlowsPerCountryUseCase{
		trafficDBRepository: trafficDBRepository,
	}
}

func (parser *GetTrafficFlowsPerCountryUseCase) GetBytesPerCountry() ([]BytesPerCountry, error) {
	serversList, err := parser.trafficDBRepository.GetServers()

	if err != nil {
		return []BytesPerCountry{}, err
	}

	servers := traffic.FilterPublicServers(serversList)

	flows := []traffic.TrafficFlow{}
	for _, server := range servers {
		flow, err := parser.trafficDBRepository.GetFlowByKey(server.Key)
		if err != nil {
			return []BytesPerCountry{}, err
		}
		flow.Server = server
		flows = append(flows, flow)
	}

	parsedBytesCountry := parsePerCountry(flows)
	bytesCountry := sumCountries(parsedBytesCountry)
	return bytesCountry, nil
}

func sumCountries(bpc []BytesPerCountry) []BytesPerCountry {
	m := map[string]int{}
	for _, v := range bpc {
		m[v.Country] += v.Bytes
	}

	bsCountry := []BytesPerCountry{}
	for country, bytes := range m {
		bpc := BytesPerCountry{
			Country: country,
			Bytes:   bytes,
		}
		bsCountry = append(bsCountry, bpc)
	}
	return bsCountry
}

func parsePerCountry(flows []traffic.TrafficFlow) []BytesPerCountry {

	bytesCn := []BytesPerCountry{}

	for _, flow := range flows {
		bpc := BytesPerCountry{
			Bytes:   flow.Bytes,
			Country: flow.Server.Country,
		}
		bytesCn = append(bytesCn, bpc)
	}

	return bytesCn
}
