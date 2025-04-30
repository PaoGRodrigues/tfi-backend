package usecase

import (
	"net"

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

func (parser *BytesAggregatorParser) GetBytesPerDestination() ([]traffic.BytesPerDestination, error) {
	serversList, err := parser.flowsStorage.GetServers()

	if err != nil {
		return []traffic.BytesPerDestination{}, err
	}

	servers := filterPublicServers(serversList)

	flows := []traffic.TrafficFlow{}
	for _, server := range servers {
		flow, err := parser.flowsStorage.GetFlowByKey(server.Key)
		if err != nil {
			return []traffic.BytesPerDestination{}, err
		}
		flow.Server = server
		flows = append(flows, flow)
	}

	parsedBytesDst := parsePerDest(flows)
	bytesDst := sumBytes(parsedBytesDst)
	return bytesDst, nil
}

func filterPublicServers(flows []traffic.Server) []traffic.Server {
	servers := []traffic.Server{}

	for _, srv := range flows {
		ip := net.ParseIP(srv.IP)
		if !ip.IsPrivate() {
			servers = append(servers, srv)
		}
	}
	return servers
}

func parsePerDest(flows []traffic.TrafficFlow) []traffic.BytesPerDestination {

	bytesDst := []traffic.BytesPerDestination{}

	for _, flow := range flows {
		var serverName string
		if flow.Server.Name != "" {
			serverName = flow.Server.Name
		} else {
			serverName = flow.Server.IP
		}
		bpd := traffic.BytesPerDestination{
			Bytes:       flow.Bytes,
			Destination: serverName,
		}
		bytesDst = append(bytesDst, bpd)
	}
	return bytesDst
}

func sumBytes(bpd []traffic.BytesPerDestination) []traffic.BytesPerDestination {
	m := map[string]int{}
	for _, v := range bpd {
		m[v.Destination] += v.Bytes
	}

	newBpd := []traffic.BytesPerDestination{}
	for dest, bytes := range m {
		new := traffic.BytesPerDestination{
			Destination: dest,
			Bytes:       bytes,
		}
		newBpd = append(newBpd, new)
	}

	return newBpd
}

func (parser *BytesAggregatorParser) GetBytesPerCountry() ([]traffic.BytesPerCountry, error) {
	serversList, err := parser.flowsStorage.GetServers()

	if err != nil {
		return []traffic.BytesPerCountry{}, err
	}

	servers := filterPublicServers(serversList)

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
