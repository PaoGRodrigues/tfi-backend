package usecase

import (
	"net"

	domains "github.com/PaoGRodrigues/tfi-backend/app/domain/traffic"
)

type BytesAggregatorParser struct {
	flowsStorage domains.TrafficRepository
}

func NewBytesParser(flowsStorage domains.TrafficRepository) *BytesAggregatorParser {
	return &BytesAggregatorParser{
		flowsStorage: flowsStorage,
	}
}

func (parser *BytesAggregatorParser) GetBytesPerDestination() ([]domains.BytesPerDestination, error) {
	serversList, err := parser.flowsStorage.GetServers()

	if err != nil {
		return []domains.BytesPerDestination{}, err
	}

	servers := filterPublicServers(serversList)

	flows := []domains.TrafficFlow{}
	for _, server := range servers {
		flow, err := parser.flowsStorage.GetFlowByKey(server.Key)
		if err != nil {
			return []domains.BytesPerDestination{}, err
		}
		flow.Server = server
		flows = append(flows, flow)
	}

	parsedBytesDst := parsePerDest(flows)
	bytesDst := sumBytes(parsedBytesDst)
	return bytesDst, nil
}

func filterPublicServers(flows []domains.Server) []domains.Server {
	servers := []domains.Server{}

	for _, srv := range flows {
		ip := net.ParseIP(srv.IP)
		if !ip.IsPrivate() {
			servers = append(servers, srv)
		}
	}
	return servers
}

func parsePerDest(flows []domains.TrafficFlow) []domains.BytesPerDestination {

	bytesDst := []domains.BytesPerDestination{}

	for _, flow := range flows {
		var serverName string
		if flow.Server.Name != "" {
			serverName = flow.Server.Name
		} else {
			serverName = flow.Server.IP
		}
		bpd := domains.BytesPerDestination{
			Bytes:       flow.Bytes,
			Destination: serverName,
		}
		bytesDst = append(bytesDst, bpd)
	}
	return bytesDst
}

func sumBytes(bpd []domains.BytesPerDestination) []domains.BytesPerDestination {
	m := map[string]int{}
	for _, v := range bpd {
		m[v.Destination] += v.Bytes
	}

	newBpd := []domains.BytesPerDestination{}
	for dest, bytes := range m {
		new := domains.BytesPerDestination{
			Destination: dest,
			Bytes:       bytes,
		}
		newBpd = append(newBpd, new)
	}

	return newBpd
}

func (parser *BytesAggregatorParser) GetBytesPerCountry() ([]domains.BytesPerCountry, error) {
	serversList, err := parser.flowsStorage.GetServers()

	if err != nil {
		return []domains.BytesPerCountry{}, err
	}

	servers := filterPublicServers(serversList)

	flows := []domains.TrafficFlow{}
	for _, server := range servers {
		flow, err := parser.flowsStorage.GetFlowByKey(server.Key)
		if err != nil {
			return []domains.BytesPerCountry{}, err
		}
		flow.Server = server
		flows = append(flows, flow)
	}

	parsedBytesCountry := parsePerCountry(flows)
	bytesCountry := sumCountries(parsedBytesCountry)
	return bytesCountry, nil
}

func sumCountries(bpc []domains.BytesPerCountry) []domains.BytesPerCountry {
	m := map[string]int{}
	for _, v := range bpc {
		m[v.Country] += v.Bytes
	}

	bsCountry := []domains.BytesPerCountry{}
	for country, bytes := range m {
		bpc := domains.BytesPerCountry{
			Country: country,
			Bytes:   bytes,
		}
		bsCountry = append(bsCountry, bpc)
	}
	return bsCountry
}

func parsePerCountry(flows []domains.TrafficFlow) []domains.BytesPerCountry {

	bytesCn := []domains.BytesPerCountry{}

	for _, flow := range flows {
		bpc := domains.BytesPerCountry{
			Bytes:   flow.Bytes,
			Country: flow.Server.Country,
		}
		bytesCn = append(bytesCn, bpc)
	}

	return bytesCn
}
