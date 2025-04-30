package traffic

import (
	traffic "github.com/PaoGRodrigues/tfi-backend/app/domain/traffic"
	trafficPorts "github.com/PaoGRodrigues/tfi-backend/app/ports/traffic"
)

type BytesAggregatorParser struct {
	flowsStorage trafficPorts.TrafficDBRepository
}

func NewBytesParser(flowsStorage trafficPorts.TrafficDBRepository) *BytesAggregatorParser {
	return &BytesAggregatorParser{
		flowsStorage: flowsStorage,
	}
}

func (parser *BytesAggregatorParser) GetBytesPerCountry() ([]traffic.BytesPerCountry, error) {
	serversList, err := parser.flowsStorage.GetServers()

	if err != nil {
		return []traffic.BytesPerCountry{}, err
	}

	servers := traffic.FilterPublicServers(serversList)

	flows := []traffic.TrafficFlow{}
	for _, server := range servers {
		flow, err := parser.flowsStorage.GetFlowByKey(server.Key)
		if err != nil {
			return []traffic.BytesPerCountry{}, err
		}
		flow.Server = server
		flows = append(flows, flow)
	}

	parsedBytesCountry := parsePerCountry(flows)
	bytesCountry := sumCountries(parsedBytesCountry)
	return bytesCountry, nil
}

func sumCountries(bpc []traffic.BytesPerCountry) []traffic.BytesPerCountry {
	m := map[string]int{}
	for _, v := range bpc {
		m[v.Country] += v.Bytes
	}

	bsCountry := []traffic.BytesPerCountry{}
	for country, bytes := range m {
		bpc := traffic.BytesPerCountry{
			Country: country,
			Bytes:   bytes,
		}
		bsCountry = append(bsCountry, bpc)
	}
	return bsCountry
}

func parsePerCountry(flows []traffic.TrafficFlow) []traffic.BytesPerCountry {

	bytesCn := []traffic.BytesPerCountry{}

	for _, flow := range flows {
		bpc := traffic.BytesPerCountry{
			Bytes:   flow.Bytes,
			Country: flow.Server.Country,
		}
		bytesCn = append(bytesCn, bpc)
	}

	return bytesCn
}
